package internal

type Appliance struct {
	Brand string
	Model string
	Consumption int
	Category string
}

func NewAppliance(brand string, model string, consumption int, category string) Appliance {
	return Appliance {
		Brand: brand,
		Model: model,
		Consumption: consumption,
		Category: category,
	}
}