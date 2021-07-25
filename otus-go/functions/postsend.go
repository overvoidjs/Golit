package main

import (
  "fmt"
  "net/http"
  "bytes"
)

func main(){
  data := []byte(`{"key":"2ltalrUZTp44wKzvdZTEJxgY3PSMmNB7Wq5CzJrsnBRMigL8r7QsVNrNgnZx"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost/api/1c/nomenclature/products/update/", "application/json", r)
	fmt.Printf("%v %v", err, resp)
}
