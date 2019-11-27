package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

//Person -
type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{"Alex", 10}
    pbytes, _ := json.Marshal(person)
    buff := bytes.NewBuffer(pbytes)
    resp, err := http.Post("http://httpbin.org/post", "application/json", buff)

    //...(elided)...
}  