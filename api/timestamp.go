package api

import (
    "fmt"
    "time"
    "strconv"
)

func Timestamp()  {
    // m := time.Now().UnixMilli()
    m := time.Now().Unix()
    fmt.Printf("Type of m: %T\n", m)
    fmt.Printf("Val of m: %s\n", strconv.FormatInt(m, 10))
    // fmt.Printf("Val of m: %s\n", strconv.Itoa(m))


    return 
    // out, err := exec.Command("date", "+v%s000").Output()
    // fmt.Printf("out = %T\n", out)

    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(string(out))
    // return fmt.Sprintln("%s000", out')
}