package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type RDF struct {
	XMLName xml.Name
	object  Object
}

type Object struct {
	XMLName string `rdf:"Description"`
}

func main() {
	xmlFile, err := os.Open("itemTrial.rdf")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	fileBytes, _ := ioutil.ReadAll(xmlFile)

	fmt.Print(string(fileBytes))

	item := &RDF{}

	xml.Unmarshal(fileBytes, item)
	fmt.Printf("%#v\n", item, item.XMLName)
}
