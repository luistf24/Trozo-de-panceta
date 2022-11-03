package internal

type Appliance struct {
	Brand string
	Model string
	Consumption int
	Category string
	Duration int
}

func NewAppliance(brand string, model string, consumption int, category string, duration int) Appliance {
	return Appliance {
		Brand: brand,
		Model: model,
		Consumption: consumption,
		Category: category,
	}
}