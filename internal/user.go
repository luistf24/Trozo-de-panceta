package internal

type User struct {
	Appliances []Appliance // list of appliances that the user has
}

// Function that allow a user to add its appliances
func (u *User) AddNewAppliance(name string, consumption float32, duration int) {
	newappliance := NewAppliance(name, consumption, duration)
	u.Appliances = append(u.Appliances, newappliance)
}