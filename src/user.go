package trozo_de_panceta

type User struct {
	Name string `json:"name"`
	Last_Name string `json:"last_name"`
	Appliances []Appliance `json:"appliances"`
}

// Function that allow a user to add its appliances
func (u *User) AddNewAppliance(name string, brand string, model string, consumption int) {
	newappliance := NewAppliance(name, brand, model, consumption)
	u.Appliances = append(u.Appliances, newappliance)
}