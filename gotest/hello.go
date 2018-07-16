package main

import (
  "github.com/PuerkitoBio/goquery"
  "net/http"
  "log"
  "regexp"
  "strings"
  "fmt"
)

func ExampleScrape() []string {
  // Request the HTML page.
  res, err := http.Get("https://github.com/vuejs/awesome-vue/pull/2272/files/84e1d84d7902c0447dd262150fd0454723c5b0f3")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  return doc.Find(".blob-code-addition .blob-code-inner").Map(func(_ int, s *goquery.Selection) string {
    // For each item found, get the band and title
    return s.Text()
  })
}

func GetUrl(diffStr string) string {
  re1 := regexp.MustCompile(`\((https:.*)\)`)
  diff := re1.FindString(diffStr)
  r := strings.NewReplacer("(", "", ")", "")
  return r.Replace(diff)
}

func GetTitle(diffStr string) string {
  re1 := regexp.MustCompile(`\[.*?\]`)
  diff := re1.FindString(diffStr)
  r := strings.NewReplacer("[", "", "]", "")
  return r.Replace(diff)
}

func main() {
  diffs := ExampleScrape()
  mathcher := diffs[0]
  url := GetUrl(mathcher)
  title := GetTitle(mathcher)
  fmt.Printf("Url: %s, Title: %s",url, title)
}
