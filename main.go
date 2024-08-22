package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	ascii "ascii/ascii"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// content, err := template.ParseFiles("home.html")
	// content, err := os.ReadFile("home.html")
	// if err != nil {
	// 	fmt.Println("error :", err)
	// 	return
	// }

	if r.URL.Path != "/" {
		fmt.Fprint(w, "Page Not Found ", http.StatusNotFound)
		return
	}
	tmp.ExecuteTemplate(w, "home.html", nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	text := strings.ReplaceAll(r.FormValue("input"), "\r\n", "\n")
	banner := r.FormValue("banner")

	if banner != "shadow" && banner != "thinkertoy" && banner != "standard" && banner != "jacky" {
		http.Error(w, "Banner Not Found", http.StatusNotFound)
		return
	}

	splitedText := strings.Split(text, "\n")
	if len(text) > 1000 {
		fmt.Fprint(w, "The length of text bigger than 1000")
		return
	}

	res := ascii.AsciiArrt(splitedText, banner)
	tmp.ExecuteTemplate(w, "home.html", res)
}

func main() {
	// css := http.FileServer(http.Dir("./"))
	// http.Handle("/style.css", css)

	// to read css imges ...
	css := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", css))

	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	// file := http.FileServer(http.Dir("./home.html"))
	fmt.Println("Server started at http://localhost:8080")
	// http.Handle("/home.html", file)

	http.ListenAndServe(":8080", nil)
}
