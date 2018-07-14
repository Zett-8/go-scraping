package basic

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  url := "http://toshikikamei.com"

  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("error occured")
    panic(err)
  }
  defer resp.Body.Close()

  byteArray, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(byteArray))
}

