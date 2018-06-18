package main 

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml"
        )

/*var washPostXML = []byte(`
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
http://www.washingtonpost.com/news-politics-sitemap.xml
</loc>
</sitemap>
<sitemap>
<loc>
http://www.washingtonpost.com/news-blogs-politics-sitemap.xml
</loc>
</sitemap>
<sitemap>
<loc>
http://www.washingtonpost.com/news-opinions-sitemap.xml
</loc>
</sitemap>
</sitemapindex>`)*/

/*func index_handler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "<h1>Hey there</h1>")
    // fmt.Fprintf(w, "<p>Go is fast!</p>")
    // fmt.Fprintf(w, "<p>...and simple!</p>")
    // fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variable</strong>")

    fmt.Fprintf(w, `
<h1>Hey there</h1>"
<p>Go is fast!</p>"
<p>...and simple!</p>"
<p>You %s even add %s</p>"
`, "can", "<strong>variable</strong>")
}*/

type SitemapIndex struct {
    Locations []Location `xml:"sitemap"`
}
// [5 5] int == array
// [] int == slice

type Location struct {
    Loc string `xml:"loc"`
}

func (l Location) String() string {
    return fmt.Sprintf(l.Loc)
}

func main() {
/*    http.HandleFunc("/", index_handler)
    http.ListenAndServe(":8000", nil)*/

    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes, _ := ioutil.ReadAll(resp.Body)
    // string_body := string(bytes)
    // fmt.Println(string_body)
    resp.Body.Close()

    var s SitemapIndex
    xml.Unmarshal(bytes, &s)

    fmt.Println(s.Locations)

    // fmt.Printf("Here %s some %s", "are", "variables")
    for _, Location := range s.Locations {
        fmt.Printf("\n%s", Location)
    }
}