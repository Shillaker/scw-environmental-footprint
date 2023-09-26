package model

const (
	CO2EKgFlightLondonNY    = 900
	CO2EKgFlightLondonParis = 122
	CO2EKg1KmDrive          = 0.14
	CO2EKg1KmTrain          = 0.04
	CO2EKg1KgCement         = 0.90
	CO2EKg1gBeef            = 0.013
	CO2EKg1lKettle          = 0.006

	EquivalentFlightLondonNY    = "flights from London to New York"
	EquivalentFlightLondonParis = "flights from London to Paris"
	EquivalentKmDrive           = "kms driven in a petrol car"
	EquivalentKgCement          = "kgs of cement manufactured"
	EquivalentGBeef             = "grams of beef eaten"
	EquivalentLitreBoiled       = "litres of water boiled in a kettle"
)

func CalculateEquivalentCO2E(co2EKg float32) []EquivalentCO2E {
	return []EquivalentCO2E{
		{Thing: EquivalentFlightLondonNY, Amount: co2EKg / CO2EKgFlightLondonNY},
		{Thing: EquivalentFlightLondonParis, Amount: co2EKg / CO2EKgFlightLondonParis},
		{Thing: EquivalentKmDrive, Amount: co2EKg / CO2EKg1KmDrive},
		{Thing: EquivalentKgCement, Amount: co2EKg / CO2EKg1KgCement},
		{Thing: EquivalentGBeef, Amount: co2EKg / CO2EKg1gBeef},
		{Thing: EquivalentLitreBoiled, Amount: co2EKg / CO2EKg1lKettle},
	}
}
