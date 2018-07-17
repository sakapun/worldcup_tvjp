package main

import (
  "golang.org/x/oauth2"
  "os"
  "fmt"
  "context"
  "github.com/shurcooL/githubv4"
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
        url
      }
    }
  }
}
*/

func getPullReq(owner string, name string) {
  src := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
  )
  httpClient := oauth2.NewClient(context.Background(), src)

  client := githubv4.NewClient(httpClient)

  type PullRequest struct {
    Title  githubv4.String
    Merged githubv4.Boolean
    MergedAt githubv4.DateTime
    Url   githubv4.URI
  }

  var query2 struct {
    Repository struct {
      PullRequests struct {
        Nodes []struct {
          PullRequest `graphql:"... on PullRequest"`
        }
      } `graphql:"pullRequests(last: 20,states: [MERGED], orderBy: {field: UPDATED_AT,direction: ASC})"`
    } `graphql:"repository(owner: $owner, name: $name)"`
  }


  variables2 := map[string]interface{}{
    "owner": githubv4.String(owner),
    "name":  githubv4.String(name),
  }

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



}

func main() {
  getPullReq("vuejs", "awesome-vue")
}
