package main
 
import (
  "os/exec"
  "log"
  "strings"
)

func Git(cmdLine string, moreArgs ...string) (string) {
  args := []string{"--git-dir=/Users/tonymarklove/Web/test-repo/.git","--work-tree=/Users/tonymarklove/Web/test-repo"}

  args = append(args, strings.Fields(cmdLine)...)
  args = append(args, moreArgs...)

  command := exec.Command("/usr/bin/git", args...)

  output, err := command.Output()
  if err != nil {
    log.Print(command.Args)
    log.Print(output)
  }

  return string(output)
}

func GitChangedFiles() ([]string) {
  result := Git("ls-files --modified")
  return strings.Split(strings.TrimSpace(result), "\n")
}
