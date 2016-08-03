package main
import (
	"fmt"
	"encoding/json"
	//"os"
)

type Response1 struct  {
	Page int
	Fruits []string
}

type Response2 struct  {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main()  {

	bolB, _ := json.Marshal(true)
	fmt.Println("bol:",string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("int:",string(intB))

	fltB, _ := json.Marshal(2.35)
	fmt.Println("flt:", string(fltB))

	strB, _ := json.Marshal("string1")
	fmt.Println("str:", string(strB))

	strC := []string{"one", "two", "three"}
	strR, _ := json.Marshal(strC)
	fmt.Println("strC:", string(strR))

	mapS := map[string]int{"one":1, "two":2, "three":3, "four":4}
	mar, _ := json.Marshal(mapS)
	fmt.Println("map:", string(mar))

	resD := &Response1{
		Page : 1,
		Fruits: []string{"apple", "banana", "orange"}}
	resB, _ := json.Marshal(resD)
	fmt.Println("resD:", string(resB))

	redD1 := &Response2{Page:1, Fruits:[]string{"apple", "banana", "orange"}}
	resB1, _ := json.Marshal(redD1)
	fmt.Println("resD1:", string(resB1))

	byt := []byte(`{"num":6.13, "fields":["a", "b"]}`)

	var dat map[string] interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println("num:", num)

	fields := dat["fields"].([]interface{})
	fmt.Println("fields0:", fields[1].(string))

}
