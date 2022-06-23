package main

import (
  "encoding/xml"
	"fmt"
  "io"
  "os"
  "github.com/knakk/rdf"
)

type XML struct {
  Obj []Obj `xml:"rdf:Description"`
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
  xmlFile, err := os.Open("fileTrial.rdf")

  if err != nil {
    fmt.Println(err)
  }

  defer xmlFile.Close()

  var reader io.Reader
  reader = xmlFile

  result := readTriples(reader)

  for i := 0; i < len(result); i++ {
    fmt.Println(result)
  }

}

func readTriples(xmlFile io.Reader) []rdf.Triple {
  decoder := rdf.NewTripleDecoder(xmlFile, rdf.RDFXML)
  tr, err := decoder.DecodeAll()
  if err != nil {
    fmt.Print(err)
  }
  return tr
}
