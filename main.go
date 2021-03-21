package main

import (
    "fmt"
    "html/template"
    "net/http"
    "net/url"
    "os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type Source struct {
    ID   interface{} `json:"id"`
    Name       string   `json:"gname"`
    Birthdate  string   `json:"bdate"`
    Gender     string   `json:"gend"`
}

type Result  struct {
    Source      Source    `json:"source"`
    Family     string   `json:"fam"`
    Name       string   `json:"gname"`
    Birthdate  string   `json:"bdate"`
    Gender     string   `json:"gend"`
}

type Search struct {
    SearchKey  string
    NextPage   int
    TotalPages int
    
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
    u, err := url.Parse(r.URL.String())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    params := u.Query()
    searchKey := params.Get("q")
    page := params.Get("page")
    if page == "" {
        page = "1"
    }

    fmt.Println("Search Query is: ", searchKey)
    fmt.Println("Results page is: ", page)
}
func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/search", searchHandler)
    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}
