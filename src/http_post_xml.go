package main

import (
    "bytes"
    "encoding/xml"
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
    pbytes, _ := xml.Marshal(person)  // xml marshal
    buff := bytes.NewBuffer(pbytes)
    resp, err := http.Post("http://httpbin.org/post", "application/xml", buff)
    //...(elided)...
}   