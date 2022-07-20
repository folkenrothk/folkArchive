package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cbroglie/mustache"
)

type File struct {
	XMLName xml.Name `xml:"rdf"`
	Rdf     []Obj    `xml:"metadata"`
}

type Obj struct {
	XMLName    xml.Name `xml:"metadata"`
	Identifier string   `xml:"identifier"`
	Title      string   `xml:"title"`
	Date       string   `xml:"date"`

	//Image
	Desc string `xml:"description"`

	Type        string `xml:"type"`
	Format      string `xml:"format"`
	PhysicalMed string `xml:"PhysicalMedium"`
	Location    string `xml:"Location"`

	Subject  string `xml:"subject"`
	Coverage string `xml:"coverage"`
	Spatial  string `xml:"spatial"`
	Temporal string `xml:"temporal"`

	Source      string `xml:"source"`
	Creator     string `xml:"creator"`
	Contributor string `xml:"contributor"`
	Publisher   string `xml:"publisher"`

	Rights   string `xml:"rights"`
	Language string `xml:"language"`

	//Relation string `xml:"relation"`
}

func main() {
	// opening the file
	xmlFile, err := os.Open("AC00001.rdf")
	// if there is an err, it's handled here
	if err != nil {
		fmt.Println("err", err)
	}
	// TROUBLESHOT-- fmt.Println("GREAT SUCCESS, the file was opened")
	// defer so we can parse it
	defer xmlFile.Close()

	// reading it as a byteArray
	fileBytes, _ := ioutil.ReadAll(xmlFile)

	var item File

	xml.Unmarshal(fileBytes, &item)

	//tell me the truth, doc
	//for i := 0; i < len(item.Rdf); i++ {
	//	fmt.Println("Creator: " + item.Rdf[i].Creator)
	//	fmt.Println("Item Title: " + item.Rdf[i].Title)
	//	fmt.Println("Subject: " + item.Rdf[i].Subject)
	//	fmt.Println("Identifier: " + item.Rdf[i].Identifier)
	//}

	//!can we have an if statement running that if rdf cant be read dont make a page
	//! another> how rdf to read multiple files and make respective pages?
	stachio(item.Rdf[0])

	/*

		//var fileArray []string

			filepath.Walk("./items", VisitFiles)

			for i := 0; i < len(fileArray); i++ {
				//fmt.Println(fileArray[i])
				// opening the file
				//xmlFile, err := os.Open(string(fileArray[i]))
				xmlFile, err := os.Open("AC00001.rdf")
				// if there is an err, it's handled here
				if err != nil {
					fmt.Println("err", err)
				}
				// defer so we can parse it
				defer xmlFile.Close()

				fileBytes, _ := ioutil.ReadAll(xmlFile)
				var item File
				xml.Unmarshal(fileBytes, &item)

				fmt.Println(item.Rdf[0].Date)
				for m := 0; m < len(item.Rdf); m++ {
					fmt.Println("Title: " + item.Rdf[m].Title)
					//stachio(item.Rdf[m])
				}
			}
	*/
}

func stachio(entry Obj) {
	//template, _ := mustache.ParseFile("item.html.mustache")
	//rendered, _ := mustache.RenderFile("item.html.mustache", entry)
	rendered, _ := mustache.RenderFileInLayout("item.html.mustache", "layout.html.mustache", entry)
	ioutil.WriteFile("item.html", []byte(rendered), 0644)
}

/*
func VisitFiles(path string, info os.FileInfo, err error) error {
	// looking through items folder for each item's file
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if info.IsDir() {
		return nil // not a file.  ignore.
	}

	matched, err := filepath.Match("*.rdf", info.Name())
	if err != nil {
		fmt.Println(err) // malformed pattern
		return err       // this is fatal.
	}
	if matched {
		fileArray = append(fileArray, path)
	}
	return nil
}
*/
