package main
import "fmt"

func main(){
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:4]
	y := arr[0:]
	a := arr[:len(arr)]
	c := arr[:]	
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(c)
}
