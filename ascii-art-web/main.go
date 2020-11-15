package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
)

type Data struct {
	Input  string
	Output string
}

type ErrData struct {
	ErrorMessage string
}

var tpl = template.Must(template.ParseFiles("static/index.html")) // ParseFiles - analyze "index.html"

func Handling(w http.ResponseWriter, r *http.Request) {
	// w - where we write the output data (answer to user)
	// r - info about request (from user)

	if r.URL.Path != "/" { // WriteHeader - output error to terminal
		w.WriteHeader(http.StatusNotFound) // Status 404 - not found
		ErrorHandle(w, "Error 404 - Page Not Found." /*, http.StatusNotFound*/)
		return
	}

	// "text" & "fs" - tags from index.html, with these Names respectively
	text := r.FormValue("text")
	fs := r.FormValue("fs")

	if fs == "" {
		fs = "fontStyle/3d.txt"
	}

	if text == "" {
		text = "Type\r\nSomething"
	}

	if !InputValidation(w, text) { // Status 400 - syntax err
		ErrorHandle(w, "Error 400 - Bad Request. Most probably Syntax error." /*, http.StatusBadRequest*/)
		return
	}

	out, err := exec.Command("./ascii-art-project/ascii-art", text, fs, "black").Output()
	// out - array, which contain ascii num of each element from banner(fs)

	if err != nil {
		ErrorHandle(w, "Error 500 - Internal Server Error.")
		return
		// fmt.Println(err.Error()) // exit status
	}

	data := Data{
		Input:  text,        // our Input text from r
		Output: string(out), // our answer
	}

	err = tpl.Execute(w, data) // Execute - run the data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Status 500 - any ISE
	}
}

func ErrorHandle(w http.ResponseWriter, message string /*, status int*/) {
	t, _ := template.ParseFiles("static/error.html")
	dataErr := ErrData{
		ErrorMessage: message,
	}
	t.Execute(w, dataErr)
	// http.Error(w, message, status)
}

func InputValidation(w http.ResponseWriter, input string) bool {
	char := map[int]bool{}
	for i := 0; i <= 94; i++ {
		char[(i + 32)] = true
	}
	char[(13)] = true
	char[(10)] = true

	for _, key := range input {

		if !char[int(key)] {
			fmt.Print("Not valid character: ", string(key), "\n")
			return false
		}
	}
	return true
}

func main() {

	fmt.Println("PORT: 1616\nServer is listening ...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "1616"
	}

	mux := http.NewServeMux() // create new multiplexor HTTP-requests

	// turn on styles(css)
	fs := http.FileServer(http.Dir("static/css"))
	mux.Handle("/css/", http.StripPrefix("/css/", fs))

	mux.HandleFunc("/", Handling)
	http.ListenAndServe(":"+port, mux) // run the server on port
}

// Handlers generate responses to requests!
