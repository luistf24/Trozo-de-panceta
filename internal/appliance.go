package internal

type Appliance struct {
	Name string
	Consumption int
	Duration int
}

func NewAppliance(name string, consumption int, duration int) Appliance {
	return Appliance {
		Name: name,
		Consumption: consumption,
		Duration: duration,
	}
}
