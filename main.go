package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./urlshort"
)

func main() {

	mux := http.NewServeMux()
	mapHandler := urlshort.MapHandler(map[string]string{}, mux)

	yamlFileName := flag.String("yaml", "conf.yml", "a YAML file describing url routing")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*yamlFileName)
	if err != nil {
		log.Printf("Get err #%v", err)
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	if err != nil {
		log.Printf("Get err #%v", err)
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}
