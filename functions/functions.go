package functions

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	e "ascii-art-web/asci-art"
)
//function for download the output
func Download(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("result")
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Content-Disposition", "attachment ;filename=file.txt")
	w.Header().Add("Content-Length", strconv.Itoa(len(text)))

	http.ServeFile(w, r, "file.txt")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/404.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Indexhandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmp, err := template.ParseFiles("web/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmp.Execute(w, nil)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func Asciiart(w http.ResponseWriter, req *http.Request) {
	myart := ""
	temp, err := template.ParseFiles("web/ascii-art.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := req.FormValue("text")
	input = e.Checkascii(input)
	banner := req.FormValue("font")

	if req.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ban := make(map[string]bool)
	ban["standard"] = true
	ban["thinkertoy"] = true
	ban["shadow"] = true

	if input == "" || !ban[banner] {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if len(input) > 200 {
		http.Error(w, "max length 200", http.StatusBadRequest)
		return
	}

	// ********* new line *********

	word := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")


	for _, b := range word {

		if b != "" {
			myart += e.Art(b, banner)
		} else {
			myart += "\r\n"
		}
		if myart == "" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	e.Output(myart, w)
	temp.Execute(w, myart)
}
