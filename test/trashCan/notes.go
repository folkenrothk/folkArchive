//Pulling from the web

//fieldTags
package main
 
import (
    "fmt"
    "encoding/json"
)
 
type Employee struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    City string `json:"city"`
}
 
func main() {
    json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`
 
    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string), emp1)
    fmt.Println(emp1)
 
    emp2 := new(Employee)
    emp2.FirstName = "Ramesh"
    emp2.LastName = "Soni"
    emp2.City = "Mumbai"
    jsonStr, _ := json.Marshal(emp2)
    fmt.Printf("%s\n", jsonStr)
}

//using Struct

package main
 
import (
    "fmt"
    "encoding/json"
)
 
type Employee struct {
    FirstName  string `json:"firstname"`
    LastName   string `json:"lastname"`
    City string `json:"city"`
}
 
func main() {
    json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`
 
    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string), emp1)
    fmt.Println(emp1)
 
    emp2 := new(Employee)
    emp2.FirstName = "Ramesh"
    emp2.LastName = "Soni"
    emp2.City = "Mumbai"
    jsonStr, _ := json.Marshal(emp2)
    fmt.Printf("%s\n", jsonStr)
}

//new thoughts 6.23

```
<?xml version='1.0' encoding='UTF-8'?>
<rdf:RDF xmlns:rdf='http://www.w3.org/1999/02/22-rdf-syntax-ns#'>
<rdf:Description rdf:about='about'
  xmlns:et='http://ns.exiftool.org/1.0/' et:toolkit='Image::ExifTool 12.30'
  xmlns:ExifTool='http://ns.exiftool.org/ExifTool/1.0/'
  xmlns:System='http://ns.exiftool.org/File/System/1.0/'
  xmlns:File='http://ns.exiftool.org/File/1.0/'>
 <ExifTool:ExifToolVersion>12.30</ExifTool:ExifToolVersion>
 <System:FileSize>20 MiB</System:FileSize>
</rdf:Description>
</rdf:RDF>
```

package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type MyXml struct {
    XMLName xml.Name `xml:"rdf:RDF"`
    Rdf     string   `xml:"rdf,attr"`
}

func main() {
    // Open our xmlFile
    xmlFile, err := os.Open("./z.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println("err", err)
    }

    fmt.Println("Successfully Opened z.xml")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(xmlFile)

    var data MyXml

    xml.Unmarshal(byteValue, &data)

    fmt.Println(data)

}
