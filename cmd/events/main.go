package main

import (
  "flag"
  "fmt"

  "github.com/cloudfoundry-incubator/tartarus/runc"
)

var (
  containerID *string
)

func init() {
  containerID = flag.String("id", "", "the container ID")
}

func main() {
  flag.Parse()

  must(runc.Events(containerID))
}

func must(err error) {
  if err != nil {
    fmt.Printf("ERROR! %s\n", err.Error())
    panic(err)
  }
}
