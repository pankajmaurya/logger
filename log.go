 package logger

 import (
  "fmt"
 )
    
 var Version string = "1.0"
    
 func Log(mess string) {
  fmt.Println("[LOG 0.2] " + mess)
 }
