package main

import (
	"encoding/xml"
	"fmt"
)

/*
type file struct {
	XMLName xml.Name `xml:"RDF"`
	Rdf     string   `xml:"rdf,attr"`
	dc      string   `xml:"dc,attr"`
	Desc    string   `xml:"Description,attr"`
}
*/

type Obj struct {
	Creator string `xml:"creator"`
	Title   string `xml:"title"`
	Subject string `xml:"subject"`
}

func main() {

	object_string := `
	{
  		"<creator> AAAAH </creator>"
  		"<title>bruh</title>"
  		"<subject> attempt</subject>"
	}`

	item1 := new(Obj)
	xml.Unmarshal([]byte(object_string), item1)
	fmt.Println(item1, "ugh")
	item1Str, _ := xml.Marshal(item1)
	fmt.Printf("%s\n", item1Str)
	fmt.Println(item1.Creator, "is the creator")
	/*
		// opening the file
		xmlFile, err := os.Open("fileTrial.rdf")
		// if there is an err, it's handled here
		if err != nil {
			fmt.Println("err", err)
		}
		// TROUBLESHOOT-- fmt.Println("GREAT SUCCESS, the file was opened")

		// defer so we can parse it
		defer xmlFile.Close()

		// reading it as a byteArray
		fileBytes, _ := ioutil.ReadAll(xmlFile)

		var item file

		xml.Unmarshal(fileBytes, &item)

		fmt.Println(item)
		//fmt.Println(item.dc.title)
	*/
}
