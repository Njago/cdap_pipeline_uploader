package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var host string = "http://localhost:11015/v3"
var namespace string = "default"

//var authToken string = ""

func main() {

	//get a dir and the files that dir holds
	dir := "/Users/bajrambojku/pipelines"
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//loop of those files and get the json data
	//call deply function to deploy each file with its file name
	for _, item := range items {
		items, _ := ioutil.ReadFile(dir + "/" + item.Name())
		if strings.HasSuffix(item.Name(), ".json") {
			deployPipeline(items, item.Name(), host, namespace)
		}
	}
}

func deployPipeline(jsonData []byte, fileName string, host string, namespace string) {
	//exported pipelines come with -cdap-data-pipeline.json suffix we only want the name to call the pipline
	fileName = strings.TrimSuffix(fileName, "-cdap-data-pipeline.json")

	// var bearer = "Bearer" + authToken

	request, _ := http.NewRequest("PUT", host+"/namespaces/"+namespace+"/apps/"+fileName, bytes.NewBuffer(jsonData))
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
