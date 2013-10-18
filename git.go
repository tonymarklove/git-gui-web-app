package main
 
import (
  "os/exec"
  "log"
  "strings"
)

func Git(cmdLine string) (string) {
  args := []string{"--git-dir=/Users/tonymarklove/Web/shaded-website/.git","--work-tree=/Users/tonymarklove/Web/shaded-website"}

  args = append(args, strings.Fields(cmdLine)...)

  command := exec.Command("/usr/bin/git", args...)

  output, err := command.Output()
  if err != nil {
    log.Print(command.Args)
  }

  return string(output)
}
