# korname

한국어 이름을 랜덤으로 작명해 주는 라이브러리인 [korean name generator](https://github.com/agemor/korean-name-generator) 의 Go 언어 포트입니다.

## 설치

```
go get github.com/dfkdream/korname
```

## 사용

```go
package main

import(
    "fmt"
    "github.com/dfkdream/korname"
)

func main(){
    fmt.Println(korname.Generate(korname.Male))   //남자 이름
    fmt.Println(korname.Generate(korname.Female)) //여자 이름
}
```