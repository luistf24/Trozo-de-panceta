package internal

type User struct {
	Appliances []Appliance
}

// Function that allow a user to add its appliances
func (u *User) AddNewAppliance(name string, consumption int, duration int) {
	newappliance := NewAppliance(name, consumption, duration)
	u.Appliances = append(u.Appliances, newappliance)
}