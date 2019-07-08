package main
import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "os"
)

type Teams struct{
	Teams []Team `json:"f1 teams"`
}

type Team struct{
	Name string `json:"team"`
	Drivers Drivers `json:"Drivers"`
	Color string `json:"Color"`
	Engine string `json:"Engine"`
}

type Drivers struct{
	One string `json:"1"`
	Two string `json:"2"`
}

func main(){
	inputfile, err := os.Open("f1.json")
	if err != nil{
		fmt.Println(err)
	}
	defer inputfile.Close()

	bytearray, _ := ioutil.ReadAll(inputfile)
	
	fmt.Println("opening is done")

	var team Teams
	json.Unmarshal(bytearray, &team)

	fmt.Println("File has been unmarshaled")
	for i := 0; i<len(team.Teams); i++ {
		fmt.Println("Team name: "+team.Teams[i].Name)
		fmt.Println("Diver 1: "+team.Teams[i].Drivers.One)
		fmt.Println("Driver 2: "+team.Teams[i].Drivers.Two)
		fmt.Println("Car color: " + team.Teams[i].Color)
		fmt.Println("Enginer company: " + team.Teams[i].Engine)
		fmt.Println("-----------------------------------------------")

	} 

}

