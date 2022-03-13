package api

import (
    "fmt"
    "log"
    "os/exec"
)
	//  	$ docker build  -t hello-world:v1 -f /Users/Ke015t7/.cf-kube/Dockerfile .

func Dk() {
    fmt.Printf("Docker Playground")

    out, err := exec.Command("ls", "-l").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}