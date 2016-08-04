package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"time"
)

type pageData struct {
	Next string `json:"next"`
	HastNext bool `json:"has_next"`
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Tags []TagItem `json:"results"`
}

type TagItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Picture string `json:"picture"`
	Description string `json:"description"`
	AbsoluteUrl string `json:"absolute_url"`
	Content string `json:"content"`
}

func setPages(urlPage string, pages chan <- string, chapter chan <- bool)  {

	page := 1
	for {
		url := "http://rutube.ru/api/tags/?format=json&page=%d"
		url = fmt.Sprintf(url, page)
		pages <- url
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}

		s := new(pageData)
		err2 := json.Unmarshal([]byte(body), &s)
		if err2 != nil {
			fmt.Println(err2)
			panic(err2.Error())
		}

		if s.HastNext == false {
			close(pages)
			break
		}
		page += 1

	}

}

func setTags(pages <- chan string, tags chan <- TagItem, chapter chan <- bool, check chan bool)  {

	for {
		select {
		case url, ok := <- pages:
			if ok {

				res, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					panic(err.Error())
				}
				defer res.Body.Close()

				body, err := ioutil.ReadAll(res.Body)
				if (err != nil) {
					panic(err.Error())
				}

				page := new(pageData)
				json.Unmarshal([]byte(body), &page)

				if page.Tags != nil {
					for _, Item := range page.Tags {
						tags <- Item
						log.Println("setTags:", Item)
					}
				}

			}
		}
	}

}

func getTags(tags chan TagItem, chapter chan <- bool, check chan bool)  {

	for {
		select {
		case tag, ok := <- tags:
			if ok {
				log.Println("getTags:", tag)
			}

		case <-time.After(time.Second * 3):

			c := len(check)
			t := len(tags)

			if t == 0 {
				if c > 0 {
					return
				} else {
					chapter<- true
					check<- true
					close(tags)
					return
				}
			}


		}
	}


}

func main()  {

	pages := make(chan string, 100)
	tags := make(chan TagItem, 100)
	chapter := make(chan bool)
	check := make(chan bool,1)

	file, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	startUrl := "http://rutube.ru/api/tags/?format=json&page=%d"
	page := 1

	url := fmt.Sprintf(startUrl, page)
	go setPages(url, pages, chapter)

	for i:=1; i<=10;i++  {
		go setTags(pages, tags, chapter, check)
	}

	time.Sleep(time.Second * 10)
	for i:=1; i<=10;i++ {
		go getTags(tags, chapter, check)
	}

	<- chapter

}

