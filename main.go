package main

import (
	// "bytes"
	// "encoding/xml"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	xmlFilePath := "https://github.com/Sneha-Jayakumar123/FirstRepo/blob/main/mule.xml" // Replace with your XML file path

	response, err := http.Get(xmlFilePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    defer response.Body.Close()
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// Create a buffer to read the file content
	buf := make([]byte, 1024) // You can adjust the buffer size as needed

	var xmlString string

	for {
		n, err := xmlFile.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				// Reached the end of the file, break the loop
				break
			}
			fmt.Println("Error reading XML file:", err)
			break
		}

		// Print the content read in this iteration
		fmt.Print(string(buf[:n]))
		xmlString = string(buf[:n])
	}

	// Your XML string
	// xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

	// Define a struct for parsing the XML
	type Root struct {
		XMLName xml.Name 
	}

	var data Root
	err = xml.Unmarshal([]byte(xmlString), &data)
	if err != nil {
		fmt.Printf("Error unmarshaling XML: %v\n", err)
		return
	}

	// Print the XML tags
	printXMLTags(data.XMLName, xmlString)
	//fmt.Println(`<?xml profile><test>` + Xml(`test '123'`) + `</test>`)
}

// func Xml(in string) string {
// 	var b bytes.Buffer
// 	xml.EscapeText(&b, []byte(in))
// 	return b.String()
// }

// ------------------------------------------------------------------------------------------------------
// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Your XML string (the structure is unknown)
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Unmarshal the XML into an interface{} to handle dynamic content
// 	// var data interface{}
// 	type Root struct {
// 		XMLName xml.Name `xml:"root"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Marshal the interface into JSON
// 	_, err = json.Marshal(data)
// 	if err != nil {
// 		fmt.Printf("Error marshaling JSON: %v\n", err)
// 		return
// 	}

// 	// Print the JSON string
// 	// fmt.Println(string(jsonData))
// 	printXMLTags(data.XMLName, xmlString)
// }

//-------------------------------------------------------------------------------------------------
// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// func main() {
// 	// Your XML string
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Define a struct that matches the structure of the XML
// 	type Root struct {
// 		Element1 string `xml:"element1"`
// 		Element2 string `xml:"element2"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Now you can access the values based on the tags
// 	fmt.Printf("Element1: %s\n", data.Element1)
// 	fmt.Printf("Element2: %s\n", data.Element2)
// }
//--------------------------------------------------------------------------------------------

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Your XML string
// 	xmlString := `<root><element1>Value1</element1><element2>Value2</element2></root>`

// 	// Define a struct for parsing the XML
// 	type Root struct {
// 		XMLName xml.Name `xml:"root"`
// 	}

// 	var data Root
// 	err := xml.Unmarshal([]byte(xmlString), &data)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling XML: %v\n", err)
// 		return
// 	}

// 	// Print the XML tags
// 	printXMLTags(data.XMLName, xmlString)
// }

func printXMLTags(name xml.Name, xmlString string) {
	fmt.Println("Root Element:", name.Local)
	startArray :=[]string{}
	endArray:=[]string{}
	decoder := xml.NewDecoder(strings.NewReader(xmlString))
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:{
			startArray=append(startArray,t.Name.Local)
			fmt.Println("Start Tag:",t.Name.Local )
		}
		case xml.EndElement:
			endArray=append(endArray,t.Name.Local)
			fmt.Println("End Tag:",t.Name.Local)
		}
	}
	fmt.Printf("Start Tag:%+v \n", startArray)
	fmt.Printf("End Tag:%+v ",endArray)
}
