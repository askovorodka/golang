package main
import "fmt"
var x [5]int

func main(){
	for i:= 0; i < len(x); i++ {
		x[i] = i * 10 + i
	}
	fmt.Println(x)

	var total float64
	
	for _, value := range x {
		total += float64(value)
	}

	fmt.Printf("total: %.2f\n", total)
	fmt.Printf("division: %.2f\n", total / float64(len(x)) )

	y := [5]float64 {
		98,
		99,
		100,
		101,
		102,
	}

	fmt.Println(y)

}

