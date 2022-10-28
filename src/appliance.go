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


// Search for an appliance in API
func ExistAppliance(model string) bool {
	// TODO
	return true;
}

// Add an appliance to the user
func AddAppliance(name string, brand string, model string, consumption int) bool {
	// TODO
	return true;
}