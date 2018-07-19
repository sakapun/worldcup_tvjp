package main

import (
  "net/http"
  "log"
  "./dynamo"
  "encoding/json"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  var results []dynamo.AwesomeLink
  db := dynamo.Connect()
  table := db.Table("AwesomeLinks")
  table.Scan().All(&results)
  json, _ := json.Marshal(results)

  w.WriteHeader(200)
  w.Header().Set("Content-Type", "application/json; charset=utf8")
  w.Write(json)
}

func main()  {
  http.HandleFunc("/", rootHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
