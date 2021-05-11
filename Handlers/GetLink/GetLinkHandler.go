package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLInkHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)
	resp, _ := http.Get(params["url"])
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("HTML:\n\n", string(bytes))
	resp.Body.Close()
}
