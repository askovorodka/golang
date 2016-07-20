package main

import (
	"fmt"
)

func main()  {
	nums := []int{2,3,4}
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("sum:", sum)

	for key, value := range nums{
		if key == 2 {
			fmt.Println("key:", key, "value:", value)
		}
	}

	kvs := map[string]string{"a" : "apple", "b":"banana","c":"cola","d":"doner"}
	for key, value := range kvs  {
		fmt.Printf("%s -> %s \n", key, value)
	}

	for key, value := range "go"{
		fmt.Println(key,value)
	}

}
