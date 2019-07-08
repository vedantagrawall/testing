package main
import (
  "fmt"
  "encoding/csv"
  "io"
  "log"
  "os"
)

func reading(reader1 *Reader) (string, string){
  endoffile := "0"
  name, err := reader1.Read() // reads one slice according to the reader struct and returns the record and err 
  if err == io.EOF{
    return endoffile, endoffile
  } 

  if err != nil{
    log.Fatal(err)
  }

 return name[0], name[1]

}

func main(){
  csvfile, err := os.Open("Names.csv") // returns file descriptor
  if err != nil {
    log.Fatalln("Couldn't open the csv file", err) // prints and exits the os
  }
  
  reader := csv.NewReader(csvfile) // returing a pointer to the reader struct

  for{
    firstname, lastname := reading(reader)

    if(firstname == "0" && lastname == "0"){
      break
    }

    fmt.Printf("First name: %s Last name: %s\n", firstname, lastname)
  }
  fmt.Printf("The file is read")
}

