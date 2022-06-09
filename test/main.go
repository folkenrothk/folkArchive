package main

import (
  "encoding/xml"
	"fmt"
  "io/ioutil"
  "os"
)

type XML struct {
  Obj Obj `xml:"rdf Description"`
}

type Obj struct {
  XMLName    xml.Name
  identifier string       `xml:"dc identifier"`
	title      string       `xml:"dc title"`
	date       string       `xml:"dc date"`
	// Image
	description    string   `xml:"dc description"`
	typeItem       string   `xml:"dc type"`
	format         string   `xml:"dc format"`
	PhysicalMedium string   `xml:"dc PhysicalMedium"`
	Location       string   `xml:"dc Location"`
	subject        string   `xml:"dc subject"`
	coverage       string   `xml:"dc coverage"`
	spatial        string   `xml:"dc spatial"`
	temporal       string   `xml:"dc temporal"`
	source         string   `xml:"dc source"`
	creator        string   `xml:"dc creator"`
	contributor    string   `xml:"dc contributor"`
	publisher      string   `xml:"dc publisher"`
	rights         string   `xml:"dc rights"`
	language       string   `xml:"dc language"`
	// relation

}

func main() {
  xmlFile, err := os.Open("itemTrial.rdf")

  if err != nil {
    fmt.Println(err)
  }

  defer xmlFile.Close()

  fileBytes, _ := ioutil.ReadAll(xmlFile)

  item := &XML{}
  xml.Unmarshal(fileBytes, item)

  fmt.Print(item.Obj)

}
