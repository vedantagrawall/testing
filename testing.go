package main
import (
    "log"
    "errors"
    "fmt"
 )

func divide(a float32, b float32) (float32, error) {
  if b == 0 {
    return 0, errors.New("can't divide by zero")
  }

  return a / b, nil
}

func main() {

  var a float32 = 10
  var b float32

  ret, err :=  divide(a,b)

  if err != nil{
    log.Print(err)
  }
  
  fmt.Println(ret)
  
}