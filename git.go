package main
 
import (
  "os"
  "github.com/libgit2/git2go"
)

type gitRepo struct {
  repo *git.Repository
  walker *git.RevWalk
}

var gr = gitRepo{}

func gitStartup() {
  repo, err := git.OpenRepository("/Users/tonymarklove/Web/shaded-website")
  if err != nil {
    os.Exit(1)
  }

  gr.repo = repo

  walker, err := gr.repo.Walk()
  if err != nil {
    os.Exit(2)
  }

  gr.walker = walker

  err = gr.walker.PushHead()
  if err != nil {
    os.Exit(3)
  }
}

// {
//   iterator := func(commit *git.Commit) bool {
//     fmt.Printf("%s", commit.Message())
//     return true
//   }

//   walker.Iterate(iterator)
// }
