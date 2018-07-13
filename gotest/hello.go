package main

import (
  "strings"
  "net/url"
  "net/http"
  "io/ioutil"
  "bytes"
  "fmt"
)

func splitName(fullname string) (string, string) {
  firstname := strings.Split(fullname, " ")[0]
  lastname := strings.Split(fullname, " ")[1]

  return firstname, lastname
}

func main() {
  baseUrl, err := url.Parse("https://jsonplaceholder.typicode.com/posts/1")
  if err != nil {
    return
  }

  resp, err := http.Get(baseUrl.String())
  if err != nil {
    return
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  buf := bytes.NewBuffer(body)
  html :=  buf.String()
  fmt.Println(html)
}
