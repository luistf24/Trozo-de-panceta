package internal

type User struct {
	Appliances []Appliance
}

// Function that allow a user to add its appliances
func (u *User) AddNewAppliance(brand string, model string, consumption int, category string, duration int) {
	newappliance := NewAppliance(brand, model, consumption, category, duration)
	u.Appliances = append(u.Appliances, newappliance)
}