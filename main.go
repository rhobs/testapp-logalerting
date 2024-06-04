package main

import (
    "time"
    "log"  
   
)

func main() {

  sum := 0
  for {
	sum++ // repeated forever
        time.Sleep(5 * time.Second)
        log.Println("Testing logging alerting")
  }

}
