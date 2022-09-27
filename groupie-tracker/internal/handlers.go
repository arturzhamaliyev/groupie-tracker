package internal

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func Server() error {
	if err := Construct(); err != nil {
		return fmt.Errorf("%s", err)
	}

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/artists/", groupHandler)
	http.HandleFunc("/locations/", locationHandler)
	http.HandleFunc("/dates/", dateHandler)
	http.HandleFunc("/relation/", relationHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	if err := http.ListenAndServe("localhost:1337", nil); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func connected() bool {
	if _, err := http.Get("https://groupietrackers.herokuapp.com/api"); err != nil {
		return false
	}

	return true
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if !connected() {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, Artists)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	if !connected() {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	reg := regexp.MustCompile(`\d+`)
	str := reg.FindString(r.URL.Path)
	id, err := strconv.Atoi(str)

	if r.URL.Path != "/artists/"+str || !isValid(id) || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/html/group.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, Artists[id-1])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func locationHandler(w http.ResponseWriter, r *http.Request) {
	if !connected() {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	reg := regexp.MustCompile(`\d+`)
	str := reg.FindString(r.URL.Path)
	id, err := strconv.Atoi(str)

	if r.URL.Path != "/locations/"+str || !isValid(id) || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/html/locations.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, Locations.Info[id-1])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	if !connected() {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	reg := regexp.MustCompile(`\d+`)
	str := reg.FindString(r.URL.Path)
	id, err := strconv.Atoi(str)

	if r.URL.Path != "/dates/"+str || !isValid(id) || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/html/dates.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, Dates.Info[id-1])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func relationHandler(w http.ResponseWriter, r *http.Request) {
	if !connected() {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	reg := regexp.MustCompile(`\d+`)
	str := reg.FindString(r.URL.Path)
	id, err := strconv.Atoi(str)

	if r.URL.Path != "/relation/"+str || !isValid(id) || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/html/relation.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, Relations.Info[id-1])
	if err != nil {
		fmt.Printf("template error: %v\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func isValid(id int) bool {
	if id > 0 && id < len(Artists) {
		return true
	}
	return false
}
