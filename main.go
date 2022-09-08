package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//Добавляем к пути файлы html
	filePath := path.Join("static", "index.html")
	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	//Добавляем к пути файлы html
	filePath := path.Join("static", "data.html")
	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
	//Получаем id
	id, ok := r.URL.Query()["Id"]
	fmt.Fprint(w, id)
	if !ok {
		return
	}
}

func main() {
	http.HandleFunc("/getInfo", getDataHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
