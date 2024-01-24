# Arg Parser

简化 flag 的调用，批量注册参数，方便统一管理。

## 使用

```shell
go get github.com/my-dev-lib/arg-parser-go
```

```go
import (
    argparser "github.com/my-dev-lib/arg-parser-go"
)
```

## 示例

提供 6 种类型的参数。

```go
// 声明参数
argParser := argparser.NewArgParser([][]any{
    // 参数名、参数类型、帮助文本、参数默认值
    {"testInt", argparser.TypeInt, "TypeInt help", 320},
    {"testBool", argparser.TypeBool, "TypeBool help", false},
    {"testUint", argparser.TypeUint, "TypeUint help", uint(32)},
    {"testString", argparser.TypeString, "TypeString help", "stringV"},
    {"testUint64", argparser.TypeUint64, "TypeUint64 help", uint64(64)},
    {"testFloat64", argparser.TypeFloat64, "TypeFloat64 help", 1234.0},
})


// 解析参数
ret, err := argParser.Parse()
if err != nil {
    fmt.Printf("err: %v", err)
    return
}

// 获取参数值
fmt.Println("testInt: " + fmt.Sprint(ret["testInt"].(int)))
fmt.Println("testBool: " + fmt.Sprint(ret["testBool"].(bool)))
fmt.Println("testUint: " + fmt.Sprint(ret["testUint"].(uint)))
fmt.Println("testString: " + ret["testString"].(string))
fmt.Println("testUint64: " + fmt.Sprint(ret["testUint64"].(uint64)))
fmt.Println("testFloat64: " + fmt.Sprint(ret["testFloat64"].(float64)))
```

## 测试

```shell
go run test.go -testFloat64 12 -testInt 34 -testString sss -testUint 56 -testUint64 78 -testBool
```