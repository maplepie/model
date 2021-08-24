# model

自定义结构类型，主要用于LeetCode中获取输入和打印输出

## 获取输入/打印输出

`Parse("input string")`即可将LeetCode中的字符串转成相应结构

```go
import "github.com/maplepie/model/list"

func main() {
    l, err := list.Parse([]byte("[1,4,7,11,15]"))
    if err != nil {
    fmt.Println(err)
    }
    result := l
    fmt.Println("result: ", result)
    fmt.Println("value: ", result.GetValue(2))
}
```
