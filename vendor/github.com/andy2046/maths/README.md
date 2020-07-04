# maths

## makes `Max` `Min` easier
Go `math` pkg provides `Max` `Min` only for `float64`, this lib adds on for other types.

[Why does Go not have feature X?](https://golang.org/doc/faq#Why_doesnt_Go_have_feature_X)

### Install

```
go get github.com/andy2046/maths
```

### Usage

```go
func main() {
    iSlice := []int{1, 2, 3}
    i := IntVar.Max(iSlice...) // 3
    i = IntVar.Min(iSlice...) // 1
}
```
