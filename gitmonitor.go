package gitmonitor

import "context"
import "fmt"

import "github.com/google/go-github/github"

func Go() {
    client := github.NewClient(nil)
    opt := &github.RepositoryListByOrgOptions{Type: "public"}
    ctx := context.Background()
    repos, _, err := client.Repositories.ListByOrg(ctx, "newridetech", opt)

    if err != nil {
        panic(err)
    }

    for _, repo := range repos {
        fmt.Println(*repo.FullName)
    }
}
