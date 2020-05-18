package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	var host string = "http://localhost:11015/v3"
	var authToken string = ""
	var namespace string = "default"
	var trim string = "-cdap-data-pipeline.json"

	//get a dir and the files that dir holds
	dir := "/Users/bajrambojku/pipelines"

	items, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//channel for go routine
	c := make(chan string)

	//loop of those files and get the json data
	//call deply function to deploy each file with its file name
	for _, item := range items {
		items, _ := ioutil.ReadFile(dir + "/" + item.Name())
		if strings.HasSuffix(item.Name(), ".json") {
			go deployPipeline(items, item.Name(), host, namespace, trim, c, authToken)
		}
	}

	//for loop for go routine channels
	for i := 0; i < len(items)-1; i++ {
		fmt.Println(<-c)
	}
}

func deployPipeline(jsonData []byte, fileName string, host string, namespace string, trim string, c chan string, authToken string) {

	//exported pipelines come with -cdap-data-pipeline.json suffix we only want the name to call the pipline
	fileName = strings.TrimSuffix(fileName, trim)

	var bearer = "Bearer" + authToken

	request, _ := http.NewRequest("PUT", host+"/namespaces/"+namespace+"/apps/"+fileName, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)

	c <- "Response Recieved"

	//Error handing for err messages
	if err != nil {
		fmt.Printf("Http request failed with error %s\n", err.Error())
	} else {
		fmt.Println(string(data))
	}

	defer response.Body.Close()
}
