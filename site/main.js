let hostname = location.hostname;
if (!hostname) {
  hostname = "localhost";
}

let baseUrl = "http://" + hostname + ":8083/v1";
let instanceUrl = baseUrl + "/instance";
let k8sControlPlanesUrl = baseUrl + "/k8s/control-plane";
let emUrl = baseUrl + "/elastic-metal";

let instanceImpactUrl = baseUrl + "/impact/instance";
let emImpactUrl = baseUrl + "/impact/elastic-metal";
let k8sImpactUrl = baseUrl + "/impact/k8s";

let poolCount = 0;

console.log("API URL: " + baseUrl);

function sortList(listIn) {
  listIn.sort(function(a, b) {
    if (a.type > b.type) return 1;
    if (a.type < b.type) return -1;
    return 0
  })

  return listIn;
}

function loadEM() {
  // Load elastic metal list
  $.ajax({
    url: emUrl,
    type: 'GET',
    dataType: 'json',
    success: function(data) {
      emTypes = data.elasticMetals;
      emTypes = sortList(emTypes);

      $.each(emTypes, function(i, emType) {
        $('#em-select').append(
          $('<option>').text(emType.description).attr('value', emType.type)
        );
      });
    },
    error: function(data) {
      var errorMsg = data.responseJSON.message;
      console.log("Error loading elastic metal list: " + errorMsg);
      alert("Error loading elastic metal list:\n" + errorMsg);
    }
  });
}

function loadInstances() {
  // Load instance list
  $.ajax({
    url: instanceUrl,
    type: 'GET',
    dataType: 'json',
    success: function(data) {
      instances = data.instances;
      instances = sortList(instances);

      $.each(instances, function(i, instance) {
        $('#instance-select').append(
          $('<option>').text(instance.description).attr('value', instance.type)
        );

        $('.k8s-pool-instance-type').append(
          $('<option>').text(instance.description).attr('value', instance.type)
        );
      });
    },
    error: function(data) {
      var errorMsg = data.responseJSON.message;
      console.log("Error loading instance list: " + errorMsg);
      alert("Error loading instance list:\n" + errorMsg);
    }
  });
}

function loadK8sControlPlanes() {
  // Load control planes list
  $.ajax({
    url: k8sControlPlanesUrl,
    type: 'GET',
    dataType: 'json',
    success: function(data) {
      controlPlanes = data.controlPlanes;
      controlPlanes = sortList(controlPlanes);

      $.each(controlPlanes, function(i, controlPlane) {
        $('#k8s-cp-select').append(
          $('<option>').text(controlPlane.description).attr('value', controlPlane.type)
        );
      });
    },
    error: function(data) {
      var errorMsg = data.responseJSON.message;
      console.log("Error loading control plane list: " + errorMsg);
      alert("Error loading control plane list:\n" + errorMsg);
    }
  });
}

function addK8sPool() {
  poolCount += 1;
  console.log("Adding pool " + poolCount);

  // Clone default pool
  var clonedPool = $("#k8s-pool-default").clone();

  clonedPool.prop("id", "k8s-pool-" + poolCount);
  clonedPool.find("h6").text("Pool " + poolCount);
  clonedPool.appendTo("#k8s-pool-list");
}

function setUpFormSubmit() {
  // Set up K8s pool button
  $("#k8s-add-pool-btn").on("click", function(e) {
    e.preventDefault();

    addK8sPool();
  });

  // Form submit
  $("#usage-form").on("submit", function(e) {
    // Cancel event
    e.preventDefault();

    var data = {
      "usage": {
        "timeSeconds": getTimeUsage(),
        "region": getUsageRegion(),
        "count": getCount(),
        "loadPercentage": getLoadPercentage()
      }
    };

    var url = null;

    var id = $('.tab-content .active').attr('id');
    if (id == "pills-instance") {
      url = instanceImpactUrl;

      var instanceType = $('#instance-select').val();
      var instanceCount = $('#instance-count').val();

      data["usage"]["count"] = instanceCount;
      data["instance"] = {
        "type": instanceType
      };
    }
    else if (id == "pills-em") {
      url = emImpactUrl;

      var emType = $('#em-select').val();
      var emCount = $('#em-count').val();

      data["usage"]["count"] = emCount;
      data["elasticMetal"] = {
        "type": emType
      };
    }
    else if (id == "pills-k8s") {
      url = k8sImpactUrl;

      // Control plane
      var cpType = $('#k8s-cp-select').val();
      data["controlPlane"] = {
        "type": cpType
      };

      // Build list of pools
      data["pools"] = [];

      $('.k8s-pool').each(function(i, pool) {
        var instanceType = $(this).find('.k8s-pool-instance-type').val();
        var instanceCount = $(this).find('.k8s-pool-instance-count').val();

        console.log("Found pool " + instanceType + " " + instanceCount);
        data["pools"].push({
          "instance": {
            "type": instanceType,
          },
          "count": instanceCount
        });
      });
    }
    else {
      alert("No product selected");
      return;
    }

    submitUsage(url, data);
  });
}

function getCount() {
  return $('#count-input').val();
}

function getLoadPercentage() {
  return $('#load-input').val();
}

function getUsageRegion() {
  return $('#region-select').val();
}

function getTimeUsage() {
  var timeSeconds =
    ($('#time-years').val() * 365 * 24 * 60 * 60) +
    ($('#time-days').val() * 24 * 60 * 60) +
    ($('#time-hours').val() * 60 * 60);

  return timeSeconds;
}

function roundNumber(amount) {
  var textAmount = null;
  if (amount < 10) {
    textAmount = amount.toPrecision(2);
  } else {
    textAmount = Math.round(amount);
  }

  return textAmount;
}

function submitUsage(url, usageData) {
  $.ajax({
    type: 'POST',
    url: url,
    data: JSON.stringify(usageData),
    dataType: 'json',
    contentType: "application/json",
    success: function(data) {
      populateImpacts(data);
    },
    error: function(data) {
      var errorMsg = data.responseJSON.message;
      console.log("Error getting results: " + errorMsg);
      alert("Error getting results:\n" + errorMsg);
    }
  });
}

function formatUseImpact(label, data) {
  return label + roundNumber(data.use) + " " + data.unit;
}

function formatManufactureImpact(label, data) {
  return label + roundNumber(data.manufacture) + " " + data.unit;
}

function populateImpacts(data) {
  var gwpLabel = "Global warming potential: "
  var adpLabel = "Abiotic depletion potential: "
  var peLabel = "Primary energy: "

  // Results
  $('#gwp-use').text(formatUseImpact(gwpLabel, data.impacts.gwp));
  $('#gwp-manufacture').text(formatManufactureImpact(gwpLabel, data.impacts.gwp));

  $('#adp-use').text(formatUseImpact(adpLabel, data.impacts.adp));
  $('#adp-manufacture').text(formatManufactureImpact(adpLabel, data.impacts.adp));

  $('#pe-use').text(formatUseImpact(peLabel, data.impacts.pe));
  $('#pe-manufacture').text(formatManufactureImpact(peLabel, data.impacts.pe));

  // Equivalents
  $('#equivalents-use').empty();
  $('#equivalents-manufacture').empty();

  $.each(data.equivalentsUse, function(i, equivalent) {
    var amount = Number.parseFloat(equivalent.amount);

    var textAmount = roundNumber(amount);
    $('#equivalents-use').append(
      $('<li>').text(textAmount + " " + equivalent.thing)
    );
  });

  $.each(data.equivalentsManufacture, function(i, equivalent) {
    var amount = Number.parseFloat(equivalent.amount);

    var textAmount = roundNumber(amount);
    $('#equivalents-manufacture').append(
      $('<li>').text(textAmount + " " + equivalent.thing)
    );
  });
}

$(document).ready(function() {
  // Load data
  loadInstances();
  loadEM();
  loadK8sControlPlanes();

  // Set up form
  setUpFormSubmit();
});
