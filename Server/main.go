package main

import (
	data "firstservice/Data"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
)

func main(){
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.Split(r.URL.Path, "/")
	if urlPath[1]=="roman_numeral"{
  number, _ := strconv.Atoi(strings.TrimSpace(urlPath[2]))
if number <1 || number > 10 {
w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))

}else{
fmt.Fprintf(w, "%q",
html.EscapeString(data.NumerialData[number]))
}
}else{
 w.WriteHeader(http.StatusNotFound)
 w.Write([]byte("ERROR"))
	}
 })
 s := &http.Server{
Addr: ":9000",
 }
 s.ListenAndServe()
}