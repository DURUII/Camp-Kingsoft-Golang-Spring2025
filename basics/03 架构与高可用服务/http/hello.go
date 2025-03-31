package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 按照最长匹配原则
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := t.Format("2006-01-02 15:04:05")
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
