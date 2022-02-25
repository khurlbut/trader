package api

import (
    "fmt"
    "time"
    "strconv"
)

func Timestamp() string {
    m := time.Now().Unix()
    return fmt.Sprintf("%s000", strconv.FormatInt(m, 10))
}