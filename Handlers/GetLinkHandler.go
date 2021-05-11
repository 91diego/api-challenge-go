package handlers

import (
	"fmt"
	"net/http"
)

func GetLInkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Download file")
}
