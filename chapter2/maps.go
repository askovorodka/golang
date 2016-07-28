package main

import (
	"fmt"
)

func main()  {

	var map1 = make(map[string]int)

	map1["k1"] = 7
	map1["k2"] = 9

	fmt.Println("map:", map1)

	v1 := map1["k1"]
	fmt.Println("v1", v1)

	fmt.Println("len:", len(map1))

	delete(map1,"k2")
	fmt.Println("map:",map1)

	_,prs := map1["k1"]
	fmt.Println("prs:",prs)

	n := map[string]int{"one":1, "two":2}
	fmt.Println("n:",n)
}