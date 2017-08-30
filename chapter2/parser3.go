package main

import (
	"fmt"
	//"time"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"github.com/streadway/amqp"
)

type TVData struct {
	HasNext bool `json:"has_next"`
	Next string `json:"next"`
	Page int `json:"page"`
	Results []TVListItem `json:"results"`
}

type TVListItem struct  {
	Id int `json:"id"`
	Content string `json:"content"`
	Name string `json:"name"`
	Description string `json:"description"`
	Picture string `json:"picture"`
}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func setTvContents(tvContents chan <- string, chapter chan <- bool)  {

	page := 1

	for {
		url := "http://rutube.ru/api/metainfo/tv/?format=json&page=%d"
		url = fmt.Sprintf(url, page)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		s := new(TVData)
		err2 := json.Unmarshal([]byte(body), &s)
		if err2 != nil {
			panic(err2)
		}

		if s.Results != nil {
			for _, item := range s.Results{
				if item.Content != ""{
					tvContents<- string(item.Content)
				}
			}
		}

		log.Println("tv page:", url)
		if s.HasNext == false {
			chapter<- true
			break
		}

		page += 1
	}


}

func main()  {

	conn, err := amqp.Dial("amqp://username:password@localhost:15672/")
	failOnError(err, "fail to connect RabbitMQ")
	defer conn.Close()

	tvContent := make(chan string, 10000)
	chapter := make(chan bool, 1)

	file, err := os.OpenFile("parser3.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	setTvContents(tvContent, chapter)

	<- chapter

}