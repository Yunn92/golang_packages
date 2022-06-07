package tempconv

func CToF(c Celsius) Farenheit{
	return Farenheit(c*9/5+32)
}
func CToK(c Celsius) Kelvin{
	return Kelvin(c - 273.15)
}
func KToC(k Kelvin) Celsius{
	return Celsius(k + 273.15)
}
func KToF(k Kelvin) Farenheit{
	return Farenheit(k*1.8 - 459.67)
}
func FToK(f Farenheit) Kelvin{
	return Kelvin((f+459.67)/1.8)
}
func FToC(f Farenheit) Celsius{
	return Celsius((f-32)*5/9)
}

