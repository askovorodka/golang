package main
import "fmt"

func main(){

	var str1 string = "hello"
	var str2 string = "Hello"
	str3 := "hello"
	
	fmt.Println(str1 == str2)
	fmt.Println(str2 == str3)
	fmt.Println(str1 == str3)
}
