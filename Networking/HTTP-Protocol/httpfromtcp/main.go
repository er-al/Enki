package main

import (
  "fmt"
  "os"
  "io"
 )

func main() {
   f, err := os.Open("messages.txt")
   if err != nil {
     fmt.Println("Error opening file:", err);
     return 
   }
   
   defer f.Close()

   buf := make([]byte, 8)

   for {
     n , err := f.Read(buf)
     if err != nil {
  	if err == io.EOF {
           break
        }
        fmt.Println("Read error:", err)
        return
      } 

      fmt.Printf("read: %s\n", buf[:n])     

    }
}
