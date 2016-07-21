package main

import (
  "flag"
  "fmt"
  "math/rand"
  "os/exec"
  "time"

  "github.com/cloudfoundry-incubator/tartarus/runc"
)

var (
  rootfs      *string
  volumeArg   *string
  cgroupName  *string
  bundleDir   *string
  containerID *string
  cleanup     *bool
  volume      []string
)

func init() {
  rootfs = flag.String("rootfs", "", "the rootfs path")
  cgroupName = flag.String("cgroups", "", "the cgroup name")
  volumeArg = flag.String("volume", "", "the volume to mount")
  bundleDir = flag.String("bundle", "", "the bundle path")
  containerID = flag.String("id", fmt.Sprintf("%d-%d", time.Now().Unix(), rand.Int()), "the container ID")
  cleanup = flag.Bool("keep", true, "keep the rootfs on exit")
}

func main() {
  flag.Parse()
  args := flag.Args()

  must(runc.Create(containerID, rootfs, bundleDir, volumeArg, cgroupName))

  must(runc.Exec(containerID, args))

  if *cleanup {
    cmd := exec.Command("runc", "delete", *containerID)
    must(cmd.Run())
  }
}

func must(err error) {
  if err != nil {
    fmt.Printf("ERROR! %s\n", err.Error())
    panic(err)
  }
}
