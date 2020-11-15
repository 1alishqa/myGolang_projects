package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/process.html")
	r.ParseForm()
	text := r.FormValue("string")
	file := r.FormValue("banner")
	if !inputValidation(text) {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	banner, err := getCharacters("banners/" + file + ".txt")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	convText := convert(banner, text)
	createFile("public/output.txt")
	writeFile(convText, "public/output.txt")
	t.Execute(w, convText)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, _ := template.ParseFiles("templates/notfound.html")
		t.Execute(w, nil)
		return
	}
	if status == http.StatusBadRequest {
		t, _ := template.ParseFiles("templates/badrequest.html")
		t.Execute(w, nil)
		return
	}
	if status == http.StatusInternalServerError {
		t, _ := template.ParseFiles("templates/servererror.html")
		t.Execute(w, nil)
		return
	}
}
