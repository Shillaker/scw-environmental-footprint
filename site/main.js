let hostname = location.hostname;
if (!hostname) {
  hostname = "localhost";
}

let baseUrl = "http://" + hostname + ":8083/v1";
let instanceUrl = baseUrl + "/instance";
let emUrl = baseUrl + "/elastic-metal";

let instanceImpactUrl = baseUrl + "/impact/instance";
let emImpactUrl = baseUrl + "/impact/elastic-metal";

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
      });
    },
    error: function(data) {
      var errorMsg = data.responseJSON.message;
      console.log("Error loading instance list: " + errorMsg);
      alert("Error loading instance list:\n" + errorMsg);
    }
  });
}

function setUpFormSubmit() {
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
      var instanceType = $('#instance-select').val();
      data["instance"] = {
        "type": instanceType
      };
      url = instanceImpactUrl;
    }
    else if (id == "pills-em") {
      var emType = $('#em-select').val();
      data["elasticMetal"] = {
        "type": emType
      };
      url = emImpactUrl;
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

  // Set up form
  setUpFormSubmit();
});
