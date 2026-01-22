package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync/atomic"
)

func main()  {

	var counter int64 = 0

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		newValue := atomic.AddInt64(&counter, 1)

		tmpl := template.Must(template.ParseFiles("counter.html"))
		tmpl.Execute(w, map[string]int{"Number": int(newValue)})
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		atomic.StoreInt64(&counter, 0)

		tmpl := template.Must(template.ParseFiles("counter.html"))
		tmpl.Execute(w, map[string]int{"Number": 0 })
	})

	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Listening on Port :8080")

	http.ListenAndServe(":8080", nil)
}
