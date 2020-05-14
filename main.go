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

	//get a dir and the files that dir holds
	items, err := ioutil.ReadDir("/Users/bajrambojku/pipelines")
	if err != nil {
		log.Fatal(err)
	}

	//loop of those files and get the json data
	//call deply function to deploy each file with its file name
	for _, item := range items {
		items, _ := ioutil.ReadFile("/Users/bajrambojku/pipelines/" + item.Name())
		if strings.HasSuffix(item.Name(), ".json") {
			deployPipeline(items, item.Name())
		}
	}
}

func deployPipeline(jsonData []byte, fileName string) {

	fileName = strings.TrimSuffix(fileName, "-cdap-data-pipeline.json")

	// var bearer = "Bearer" + authToken

	request, _ := http.NewRequest("PUT", "http://localhost:11015/v3/namespaces/default/apps/"+fileName, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	// request.Header.Add("Authorization", bearer)
	client := &http.Client{}
	response, err := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)

	//Error handing for err messages
	if err != nil {
		fmt.Printf("Http request failed with error %s\n", err.Error())
	} else {
		fmt.Println(string(data))
	}

}
