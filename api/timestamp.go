package api

import (
    "fmt"
    "time"
    "strconv"
)

func Timestamp() string {
    return fmt.Sprintf("%s000", strconv.FormatInt(time.Now().Unix(), 10))
}