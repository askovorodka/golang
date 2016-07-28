package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil" )

func getJson(url string, data interface{}) error {
	response, error := http.Get(url)
	if error != nil {
		return error
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(data)
	fmt.Println(data)
	return nil
}

type Result struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Picture string `json:"picture"`
}

type data struct {
	Next string `json:"next"`
	Has_next bool `json:"has_next"`
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	/*Results []struct{
		Id int `json:"id"`
		Name string `json:"name"`
		Picture string `json:"picture"`
		Description string `json:"description"`
	}*/
	Results []Result `json:"results"`
}

func getPage(url string) (*data, error)  {
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s := new(data)
	err2 := json.Unmarshal([]byte(body), &s)
	if err2 != nil {
		panic(err2.Error())
	}

	return s, err2
}

func getResults(body []byte) (*data, error)  {
	s := new(data)
	err := json.Unmarshal(body, &s)
	if err != nil {
		panic(err.Error())
	}
	return s, err
}

func main()  {
	/*
	var foo1 Result
	getJson("http://rutube.ru/api/tags/?format=json", foo1)
	fmt.Println(foo1)*/
	/*var data struct{
		next string
		results map[string]string
	}
	content, err := http.Get("http://rutube.ru/api/tags/?format=json")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}*/
	//defer content.Body.Close()
	//dec := json.NewDecoder(content.Body)
	//dec.Decode(&data)
	/*error := json.Unmarshal(content.Body, &data)
	if error != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Println(data)
	*/

	startUrl := "http://rutube.ru/api/tags/?format=json&page=%d"
	page := 1
	for {
		startUrl = fmt.Sprintf(startUrl, page)
		pageContent, err := getPage(startUrl)
		if err != nil {
			panic(err.Error())
		}

		if pageContent.Has_next == false {
			break
		}
		page += 1

	}


	/*res, error := http.Get("http://rutube.ru/api/tags/?format=json")
	if error != nil {
		panic(error.Error())
	}

	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		panic(error.Error())
	}

	var s data
	err := json.Unmarshal([]byte(body), &s)
	if err != nil {
		panic(err.Error())
	}

	for _, Item := range s.Results {
		fmt.Printf("Id: %d, Name: %s, Picture: %s\n", Item.Id, Item.Name, Item.Picture)
	}*/

}
