package dynamo

import (
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/guregu/dynamo"
  "fmt"
)


// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type widget struct {
  UserID int       // Hash key, a.k.a. partition key
  Time   time.Time // Range key, a.k.a. sort key

  Msg       string              `dynamo:"Message"`
  Count     int                 `dynamo:",omitempty"`
  Friends   []string            `dynamo:",set"` // Sets
  Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
  SecretKey string              `dynamo:"-"`    // Ignored
}


func DbRegist() {
  db := Connect()
  table := db.Table("Widgets")

  // put item
  w := widget{UserID: 613, Time: time.Now(), Msg: "hello", Count: 3}
  err := table.Put(w).Run()
  if err != nil {

  }

  // get the same item
  var result widget
  err = table.Get("UserID", w.UserID).
    //Range("Time", dynamo.Equal, w.Time).
    Filter("'Count' = ? AND $ = ?", w.Count, "Message", w.Msg). // placeholders in expressions
    One(&result)
  fmt.Println(result)

  // get all items
  var results []widget
  err = table.Scan().All(&results)
  fmt.Print(results)
}

func Connect () *dynamo.DB {
  return dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
}

type AwesomeLink struct {
  Id         string       // Hash key, a.k.a. partition key
  MergedAt   time.Time // Range key, a.k.a. sort key

  Title      string
  Url        string


  //
  //Msg       string              `dynamo:"Message"`
  //Count     int                 `dynamo:",omitempty"`
  //Friends   []string            `dynamo:",set"` // Sets
  //Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
  //SecretKey string              `dynamo:"-"`    // Ignored
}
