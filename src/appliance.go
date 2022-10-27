package trozo_de_panceta

type Appliance struct {
	Name string `json:"name"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Consumption int `json:"consumption"`
}

// Search for an appliance in API
func (a *Appliance) ExistAppliance(model string) Appliance {
	// TODO
	return *a
}

// Add an appliance to the user
func (a *Appliance) AddAppliance(name string, brand string, model string, consumption int) Appliance {
	// TODO
	return *a
}