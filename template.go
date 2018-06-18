package main 

import (
        "fmt"
        "net/http"
        "html/template"
        "io/ioutil"
        "encoding/xml"
)

type Sitemapindex struct {
    Locations []string `xml:"sitemap>loc"`
}

type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

type News struct {
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
}

type NewsMap struct {
    Keyword string
    Location string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
    var s Sitemapindex
    var n News

    news_map := make(map[string]NewsMap)

    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &s)
    resp.Body.Close()

    for _, Location := range s.Locations {
        resp, _ := http.Get(Location)
        bytes, _ := ioutil.ReadAll(resp.Body)
        xml.Unmarshal(bytes, &n)
        resp.Body.Close()

        // fmt.Println(n.Title)
        for idx, _ := range n.Titles {
            news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
        }
    }

    p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
    t, _ := template.ParseFiles("newsaggtemplate.html")  // _
    // t.Execute(w, p)
    fmt.Println(t.Execute(w, p)) // print out err if err.
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", newsAggHandler)
    http.ListenAndServe(":8000", nil)
}