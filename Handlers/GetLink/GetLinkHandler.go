package handlers

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"

	models "github.com/91diego/api-rest-challenge/Models"
	"github.com/gocolly/colly/v2"
)

func GetLInkHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)
	fName := "elments.xlsx"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write XSLS header
	writer.Write([]string{"URL"})

	c := colly.NewCollector()

	// Find and all <a> label
	c.OnHTML("a", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.Text,
		})
	})

	// Start scraping on post URL
	c.Visit(params["url"])

	response := models.Response{
		Code:    http.StatusAccepted,
		Status:  "success",
		Message: "XSLX Downloaded! Verify on root project folder",
		Data:    "Tag a added to xlsx from url " + params["url"],
	}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
