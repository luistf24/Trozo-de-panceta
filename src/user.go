package trozo_de_panceta

type User struct {
	Name string `json:"name"`
	Last_Name string `json:"last_name"`
	Appliances []Appliance `json:"appliances"`
}