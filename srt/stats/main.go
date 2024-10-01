package main

import (
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
)

// func getData(w http.ResponseWriter, r *http.Request) {
// 	// Define the path to the temporary JSON file
// 	tempFilePath := "/tmp/srtstats.json"

// 	// Check if the file exists
// 	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
// 		http.Error(w, "Temporary JSON file not found.", http.StatusNotFound)
// 		return
// 	}

// 	// Read the JSON data from the temporary file
// 	data, err := ioutil.ReadFile(tempFilePath)
// 	if err != nil {
// 		http.Error(w, "Error reading JSON file.", http.StatusInternalServerError)
// 		log.Fatal(err)
// 		return
// 	}

// 	// Validate JSON
// 	var jsonData interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		http.Error(w, "Error decoding JSON data.", http.StatusInternalServerError)
// 		return
// 	}



// 	// Set the Content-Type header to application/json
// 	w.Header().Set("Content-Type", "application/json")
// 	// Write the JSON data to the response
// 	w.Write(data)
// }

func main() {
	// http.HandleFunc("/stats.json", getData)
	http.HandleFunc("/stats.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/tmp/srtstats.json")
	})
	// Start the HTTP server on port 8080
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

