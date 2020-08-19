package main

import (
	"fmt"
	"log"
	"net/http"
	"math"
)

func looping(msg string) string {
	var x = 0.00001
	for i:=0; i<=1000000; i++ {
		x += math.Sqrt(x)
	}
	return msg
}

func main()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, looping("Code.education Rocks"))
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}