package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Option struct {
    Text string `json:"text"`
    ArcName string `json:"arc"`
}

type Arc struct {
    Title string  `json:"title"`
    Story []string `json:"story"`
    Options []Option `json:"options"`
}

 var data map[string]Arc
 
 

func main() {
	// Open the JSON file
	file, err := os.Open("./gopher.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file at the end

	// Read the contents of the file into a byte slice
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
    // Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

    templateFile := "base.html"
    template, err := template.ParseFiles(templateFile)
	if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}


    // fmt.Println("Test reading file",jsonData)
    http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {
		// Step 1: Get the Plate (http.ResponseWriter)
         query := r.URL.Query()
         arc := query.Get("page")
         fmt.Println("arc is ",arc)

          ad, ok:= data[arc]

	if !ok {
		fmt.Println("Arc not found:", err)
		return
	}
		// Execute the template with the data and write to ResponseWriter
		err = template.Execute(w, ad)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}



		// Step 4: Set the Headers (Optional)
		w.Header().Set("Content-Type", "text/plain")

		// Step 5: Send the Plate (Flush the response)
		w.(http.Flusher).Flush()
	})

	http.ListenAndServe(":8080", nil)
}

