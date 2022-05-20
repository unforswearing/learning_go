package main

/*
 	dued.go
  ========
  dued.go days hours
  eg. dued.go 4 10
  eg. dued.go -8 0

  no arguments prints todays date
  eg. dued.go 
*/

import (
	"time"
	"fmt"
	"os"
	"strconv"
)

func main() {
  var timeFmt string = "2006-01-02 15:04"
  now := time.Now()
	
  args := os.Args[1:]
  
  if len(args) == 0 {
    fmt.Println(now.Format(timeFmt))
    os.Exit(0)
  }

  days, _ := strconv.Atoi(args[0])
  hours, _ := strconv.Atoi(args[1])

  /*if hours == nil {
    hours := 0
  }*/

  hoursCalc := (time.Hour * time.Duration(hours))
  
  generatedTime := now.Add(hoursCalc).AddDate(0, 0, days)

  fmt.Println(generatedTime.Format(timeFmt))
}
