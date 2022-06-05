package tempconv

func CToF(c Celsius){
	return Farenheit(c*9/5+32)
}
func CToK(c Celsius){
	return Kelvin(c - 273.15)
}
func KToC(k Kelvin){
	return Celsius(k + 273.15)
}
func KToF(k Kelvin){
	return Farenheit(k*1.8 - 459.67)
}
func FToK(f Farenheit){
	return Kelvin((f+459.67)/1.8)
}
func FToC(f Farenheit) Celsius{
	return Celsius((f-32)*5/9)
}

