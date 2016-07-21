package runc

import (
  "errors"
  "os"
  "os/exec"
)

func Exec(containerID *string, args []string) error {

  if *containerID == "" {
    return errors.New("Need an ID")
  }

  cmd := exec.Command("runc", append([]string{"exec", *containerID}, args...)...)
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr

  return cmd.Run()
}
