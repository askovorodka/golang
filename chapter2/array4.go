package main
import "fmt"

func main(){
	
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	x := make(map[string]float64)
	x["key"] = 1
	fmt.Println(x["key"])

}

