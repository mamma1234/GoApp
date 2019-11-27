package main

import "io/ioutil"

func main() {
    //파일 읽기
    bytes, err := ioutil.ReadFile(".\\1.txt")
    println(bytes)
    if err != nil {
        panic(err)
    }
    //파일 쓰기
    err = ioutil.WriteFile(".\\2.txt", bytes, 0)
    if err != nil {
        panic(err)
    }
}