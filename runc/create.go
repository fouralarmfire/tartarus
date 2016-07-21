package runc

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "strings"

  "github.com/opencontainers/runtime-spec/specs-go"
)

func Create(containerID, rootfs, bundleDir, volumeArg, cgroupName *string) error {
  mountVols := []specs.Mount{
    specs.Mount{
      Destination: "/proc",
      Type:        "proc",
      Source:      "",
    },
  }

  if *volumeArg != "" {
    volume := strings.Split(*volumeArg, ":")
    mountVols = append(mountVols, specs.Mount{
      Destination: volume[1],
      Type:        "bind",
      Source:      volume[0],
      Options:     []string{"rbind", "rw"},
    })
  }

  containerSpec := specs.Spec{
    Version:  "0.2.0",
    Platform: specs.Platform{OS: "linux", Arch: "amd64"},
    Process: specs.Process{
      Args: []string{"/bin/sh"},
      Env:  []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
      Cwd:  "/",
    },
    Root: specs.Root{
      Path: *rootfs,
    },
    Mounts: mountVols,
    Linux: &specs.Linux{
      CgroupsPath: cgroupName,
      Namespaces: []specs.Namespace{
        specs.Namespace{Type: specs.UTSNamespace},
        specs.Namespace{Type: specs.PIDNamespace},
        specs.Namespace{Type: specs.MountNamespace},
      },
    },
  }

  jsonData, err := json.Marshal(containerSpec)
  if err != nil {
    return err
  }

  ioutil.WriteFile(path.Join(*bundleDir, "config.json"), jsonData, 0600)

  cmd := exec.Command("runc", "create", *containerID, "--bundle", *bundleDir)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}
