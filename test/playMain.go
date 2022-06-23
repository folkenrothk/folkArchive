package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type file struct {
	XMLName xml.Name `xml:"RDF"`
	rdf     struct {
		RDFNS string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns#"`
		SINS  string `xml:"https://www.w3schools.com/rdf/"`
	} `xml:"RDF"`
	obj struct {
		title string `xml:""https://www.w3schools.com/rdf/ title"`
	} `xml:Description`
}

func main() {
	xmlFile, err := os.Open("itemTrial.rdf")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	fileBytes, _ := ioutil.ReadAll(xmlFile)

	var item file

	if err := xml.Unmarshal(fileBytes, &item); err != nil {
		panic(err)
	}

	fmt.Println(item)
	fmt.Println(item.obj.title)

}
