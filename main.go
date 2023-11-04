package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func renderHTML(w http.ResponseWriter, data interface{}) {
	t, err := template.New("index.html").ParseFiles("static/index.html")
	checkErr(err)
	t.Execute(w, data)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", r)
	renderHTML(w, "no data")
}

func createMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("match-name")
		positive := r.FormValue("positive")
		negative := r.FormValue("negative")
		log.Printf("build match name: %s, positive:%s, negative:%s", name, positive, negative)
	}
}

type WaitMatch struct {
	Id       int64
	Name     string
	Positive string
	Negative string
}

func waitMatch(w http.ResponseWriter, r *http.Request) {
	list := make([]WaitMatch, 0)
	match1 := WaitMatch{}
	match1.Id = 1
	match1.Name = "比赛 1"
	match1.Positive = "正方 1"
	match1.Negative = "反方 1"
	list = append(list, match1)
	match2 := WaitMatch{}
	match2.Id = 2
	match2.Name = "比赛 1"
	match2.Positive = "正方 1"
	match2.Negative = "反方 1"
	list = append(list, match2)

	tmpl, err := template.ParseFiles("tmpl/wait-match.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, list)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func startMatch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("tmpl/match-timer.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	match1 := WaitMatch{}
	match1.Id = 1
	match1.Name = "面对生命重的荒谬，我们应该坦然面对/抗争到底"
	match1.Positive = "正方"
	match1.Negative = "反方"
	err = tmpl.Execute(w, match1)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}

}

func goToTest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("tmpl/test.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/test", goToTest)
	http.HandleFunc("/start", startMatch)
	http.HandleFunc("/create-match", createMatch)
	http.HandleFunc("/wait-match", waitMatch)
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
