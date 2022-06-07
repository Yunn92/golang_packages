package main
import (
	"tempconv"
	"fmt"
	"os"
	"strconv"
)

func main(){
	for _, arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil{
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Farenheit(t)
		fmt.Printf("%s = %s\n%s = %s\n%s = %s\n\n", f, tempconv.FToC(f), f, tempconv.FToK(f), tempconv.FToK(f), tempconv.FToC(f))
	}
}