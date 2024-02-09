package main

import (
"fmt"
"time"
)

func main() {
tn := time.Date(2023, 9, 13, 0,0,0,0, time.Local)
fmt.Println("hellow", tn.Before(time.Now()))
}
