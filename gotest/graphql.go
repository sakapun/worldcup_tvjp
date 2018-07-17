package main

import (
  "golang.org/x/oauth2"
  "os"
  "fmt"
  "context"
  "github.com/shurcooL/githubv4"
  "./dynamo"
)

/*
# Type queries into this side of the screen, and you will
# see intelligent typeaheads aware of the current GraphQL type schema,
# live syntax, and validation errors highlighted within the text.

# We'll get you started with a simple query showing your username!
query {
  repository(owner: "vuejs", name: "awesome-vue") {
    pullRequests(
      states: [MERGED],
      last: 20,
      orderBy: {
        field: UPDATED_AT,
        direction: ASC
      }
    ) {
      nodes {
        title,
        merged,
        mergedAt,
        url,
        id
      }
    }
  }
}
*/

type PullRequest struct {
  Title  githubv4.String
  Merged githubv4.Boolean
  MergedAt githubv4.DateTime
  Url   githubv4.URI
  Id    githubv4.ID
}

type Query struct {
  Repository struct {
    PullRequests struct {
      Nodes []struct {
        PullRequest `graphql:"... on PullRequest"`
      }
    } `graphql:"pullRequests(last: 20,states: [MERGED], orderBy: {field: UPDATED_AT,direction: ASC})"`
  } `graphql:"repository(owner: $owner, name: $name)"`
}

func getPullReq(owner string, name string) Query {
  src := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
  )
  httpClient := oauth2.NewClient(context.Background(), src)

  client := githubv4.NewClient(httpClient)


  var query2 = Query{}


  variables2 := map[string]interface{}{
    "owner": githubv4.String(owner),
    "name":  githubv4.String(name),
  }

  err2 := client.Query(context.Background(), &query2, variables2)
  if err2 != nil {
    fmt.Println(err2)
  }

  //for _, pr := range query2.Repository.PullRequests.Nodes {
  //  fmt.Println("---------")
  //  fmt.Println(pr.Title)
  //  fmt.Println(pr.Merged)
  //  fmt.Println(pr.MergedAt)
  //  fmt.Println(pr.Url)
  //}

  return query2
}

func main() {
  pullRequests := getPullReq("vuejs", "awesome-vue")
  fmt.Print(pullRequests.Repository.PullRequests.Nodes[0].Url)

  //for _, pr := range pullRequests.Repository.PullRequests.Nodes {
  //  diffs := sub.ExampleScrape(pr.Url.String() + "/files")
  //  url := sub.GetUrl(diffs[0])
  //  title := sub.GetTitle(diffs[0])
  //  fmt.Printf("Url: %s, Title: %s\n",url, title)
  //}

  dynamo.DbRegist()

}
