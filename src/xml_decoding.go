func main() {
    // 테스트용 데이타
    xmlBytes, _ := xml.Marshal(Member{"Tim", 1, true})

    // XML 디코딩
    var mem Member
    err := xml.Unmarshal(xmlBytes, &mem)
    if err != nil {
        panic(err)
    }
    // mem 구조체 필드 엑세스
    fmt.Println(mem.Name, mem.Age, mem.Active)
}