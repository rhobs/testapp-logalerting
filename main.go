package main

import (
    "log"
    "time" 
    //zlog "github.com/rs/zerolog/log"
   
)

func main() {

  sum := 0
  for {
	sum++ // repeated forever
        time.Sleep(5 * time.Second)
        log.Println("Testing logging alerting")
        //zlog.print("Testing logging alerting with a structured line")
  }

}
