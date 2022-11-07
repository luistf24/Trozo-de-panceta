package internal

type Appliance struct {
	Name string
	Consumption float32 // expressed in Wh
	Duration int // expressed in minutes
}

func NewAppliance(name string, consumption float32, duration int) Appliance {
	return Appliance {
		Name: name,
		Consumption: consumption,
		Duration: duration,
	}
}
