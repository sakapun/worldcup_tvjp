package main

import (
  "golang.org/x/oauth2"
  "os"
  "fmt"
  "context"
  "github.com/shurcooL/githubv4"
  "./dynamo"
  "./sub"
  "github.com/aws/aws-lambda-go/lambda"
)


type PullRequest struct {
  Title  githubv4.String
  Merged githubv4.Boolean
  MergedAt githubv4.DateTime
  Url   githubv4.URI
  Id    string
}

type Query struct {
  Repository struct {
    PullRequests struct {
      Nodes []struct {
        PullRequest `graphql:"... on PullRequest"`
      }
      PageInfo struct {
        StartCursor githubv4.String
        HasPreviousPage githubv4.Boolean
      }
    } `graphql:"pullRequests(last: 20,states: [MERGED], orderBy: {field: UPDATED_AT,direction: ASC}, before: $startCursor)"`
  } `graphql:"repository(owner: $owner, name: $name)"`
}

func getPullReq(owner string, name string, startCursor string) (Query, string, bool) {
  src := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
  )
  httpClient := oauth2.NewClient(context.Background(), src)

  client := githubv4.NewClient(httpClient)


  var query2 = Query{}


  variables2 := map[string]interface{}{
    "owner": githubv4.String(owner),
    "name":  githubv4.String(name),
    "startCursor":  (*githubv4.String)(nil),
  }
  if startCursor != "" {
    variables2["startCursor"] = githubv4.String(startCursor)
  }
  fmt.Println(variables2)


  err2 := client.Query(context.Background(), &query2, variables2)
  if err2 != nil {
    fmt.Println(err2)
  }

  for _, pr := range query2.Repository.PullRequests.Nodes {
   fmt.Println("---------")
   fmt.Println(pr.Title)
   fmt.Println(pr.Merged)
   fmt.Println(pr.MergedAt)
   fmt.Println(pr.Url)
  }

  return query2, string(query2.Repository.PullRequests.PageInfo.StartCursor),bool(query2.Repository.PullRequests.PageInfo.HasPreviousPage)
}

func AddLoop (pullRequests Query) {
  fmt.Print(pullRequests.Repository.PullRequests.Nodes[0].Url)

  db := dynamo.Connect()
  table := db.Table("AwesomeVueLinks")
  for _, pr := range pullRequests.Repository.PullRequests.Nodes {
    diffs := sub.ExampleScrape(pr.Url.String() + "/files")
    if len(diffs) > 0 {
      url := sub.GetUrl(diffs[0])
      title := sub.GetTitle(diffs[0])

      awesomeLink := dynamo.AwesomeLink{
        Id: pr.Id,
        MergedAt: pr.MergedAt.UTC(),
        Title: title,
        Url: url,
        Type: "link",
      }

      err := table.Put(awesomeLink).Run()
      if err != nil {

      }

      fmt.Print(awesomeLink)
    }
  }
}

func batch()  {
  pullRequests, _, _ := getPullReq("vuejs", "awesome-vue", "")
  fmt.Print(pullRequests)
  AddLoop(pullRequests)
}

func main () {
  lambda.Start(batch)
}
