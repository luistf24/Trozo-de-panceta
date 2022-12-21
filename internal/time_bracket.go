package internal

type TimeBracket struct {
	Hour int
	Price float32
	Appliances []Appliance
}

func (t TimeBracket) prueba() int{
	return 1
}
