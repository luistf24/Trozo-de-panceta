package trozo_de_panceta

type Appliance struct {
	Name string `json:"name"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Consumption int `json:"consumption"`
}

// Create a new Appliance
func NewAppliance(name string, brand string, model string, consumption int) Appliance {
	return Appliance {
		Name: name,
		Brand: brand,
		Model: model,
		Consumption: consumption,
	}
}