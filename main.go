package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler function to convert JSON to XML
func HandleConvert(w http.ResponseWriter, r *http.Request) {
	var jsonData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	xmlData := ConvertJSONToXML(jsonData)

	// Write the XML response
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, xmlData)
}

// ConvertJSONToXML is a recursive function to convert JSON to XML
func ConvertJSONToXML(node interface{}) string {
	switch v := node.(type) {
	case map[string]interface{}:
		xmlStr := ""
		for key, val := range v {
			xmlStr += ConvertJSONToXML(val)
			xmlStr = fmt.Sprintf("<%s>%s</%s>", key, xmlStr, key)
		}
		return xmlStr
	case []interface{}:
		xmlStr := ""
		for _, val := range v {
			xmlStr += ConvertJSONToXML(val)
		}
		return xmlStr
	default:
		return fmt.Sprintf("%v", v)
	}
}

func main() {
	// Create a new router
	http.HandleFunc("/convert", HandleConvert)

	// Start the local web server
	http.ListenAndServe(":8080", nil)
}
