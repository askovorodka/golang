package main
import "fmt"

const cel1 float64 = 32.0

func main(){
	fmt.Print("Enter fahrenheit degres: ")
	var (
		fahrenheit float64
	)
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - cel1)*(5.0/9.0)
	fmt.Printf("Celsius: %.2f", celsius)
}

