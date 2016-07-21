package runc

import (
  "errors"
  "os"
  "os/exec"
)

func Events(containerID *string) error {

  if *containerID == "" {
    return errors.New("Need an ID")
  }

  cmd := exec.Command("runc", "events", *containerID)
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr

  return cmd.Run()
}
