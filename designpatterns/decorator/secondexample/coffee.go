package secondexample

func GetCoffee() string {
	return "normal coffee"
}

func AddOns(base func() string, addon string) func() string {
	return func() string {
		return base() + " + " + addon
	}
}
