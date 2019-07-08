package main
import (
  "fmt"
  "encoding/csv"
  "io"
  "log"
  "os"
)

func main(){
	csvfile, err := os.Open("Names.csv") // returns file descriptor

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err) // prints and exits the os
	}
  
	reader := csv.NewReader(csvfile) // returing a pointer to the reader struct

	for{
		name, err := reader.Read() // reads one slice according to the reader struct and returns the record and err 
	
		if err == io.EOF{
			break
	  	}  

  		if err != nil || name[1] == "" || name[0] == ""{
			
			f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
			
			if err != nil {
    			log.Fatalf("error opening file: %v", err)
			}

			defer f.Close()

			log.SetOutput(f)
			log.Println("One field not found")
			fmt.Printf("There is an error \n")
			break
  		}

		fmt.Printf("first name: %s last name: %s\n", name[0], name[1])
	}
}

