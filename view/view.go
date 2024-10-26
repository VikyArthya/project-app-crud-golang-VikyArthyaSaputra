package view

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(data interface{}) {
    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        fmt.Println("Error formatting JSON:", err)
        return
    }
    fmt.Println(string(jsonData))
}
