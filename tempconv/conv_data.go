package tempconv
import "fmt"

type Celsius float64
type Farenheit float64
type Kelvin float64

const(
		AbsoluteZeroC Celsius = -273.15
		FreezingC Celsius = 0
		BoilingC Celsius = 100
		test int = 5
	)	

func (c Celsius) String() string{
	return fmt.Sprintf("%gC",c)
}
func (f Farenheit) String() string{
	return fmt.Sprintf("%gF",f)
}
func (k Kelvin) String() string{
	return fmt.Sprintf("%gK",k)
}
