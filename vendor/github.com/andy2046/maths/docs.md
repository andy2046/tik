

# maths
`import "github.com/andy2046/maths"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Range(end int) []struct{}](#Range)
* [func Ternary(expr bool, trueVal, falseVal interface{}) interface{}](#Ternary)
* [type Byte](#Byte)
  * [func (Byte) Max(numerics ...byte) byte](#Byte.Max)
  * [func (Byte) Min(numerics ...byte) byte](#Byte.Min)
* [type Float32](#Float32)
  * [func (Float32) Max(numerics ...float32) float32](#Float32.Max)
  * [func (Float32) Min(numerics ...float32) float32](#Float32.Min)
* [type Float64](#Float64)
  * [func (Float64) Max(numerics ...float64) float64](#Float64.Max)
  * [func (Float64) Min(numerics ...float64) float64](#Float64.Min)
* [type Int](#Int)
  * [func (Int) Max(numerics ...int) int](#Int.Max)
  * [func (Int) Min(numerics ...int) int](#Int.Min)
* [type Int16](#Int16)
  * [func (Int16) Max(numerics ...int16) int16](#Int16.Max)
  * [func (Int16) Min(numerics ...int16) int16](#Int16.Min)
* [type Int32](#Int32)
  * [func (Int32) Max(numerics ...int32) int32](#Int32.Max)
  * [func (Int32) Min(numerics ...int32) int32](#Int32.Min)
* [type Int64](#Int64)
  * [func (Int64) Max(numerics ...int64) int64](#Int64.Max)
  * [func (Int64) Min(numerics ...int64) int64](#Int64.Min)
* [type Int8](#Int8)
  * [func (Int8) Max(numerics ...int8) int8](#Int8.Max)
  * [func (Int8) Min(numerics ...int8) int8](#Int8.Min)
* [type Rune](#Rune)
  * [func (Rune) Max(numerics ...rune) rune](#Rune.Max)
  * [func (Rune) Min(numerics ...rune) rune](#Rune.Min)
* [type String](#String)
  * [func (String) Max(numerics ...string) string](#String.Max)
  * [func (String) Min(numerics ...string) string](#String.Min)
* [type Uint](#Uint)
  * [func (Uint) Max(numerics ...uint) uint](#Uint.Max)
  * [func (Uint) Min(numerics ...uint) uint](#Uint.Min)
* [type Uint16](#Uint16)
  * [func (Uint16) Max(numerics ...uint16) uint16](#Uint16.Max)
  * [func (Uint16) Min(numerics ...uint16) uint16](#Uint16.Min)
* [type Uint32](#Uint32)
  * [func (Uint32) Max(numerics ...uint32) uint32](#Uint32.Max)
  * [func (Uint32) Min(numerics ...uint32) uint32](#Uint32.Min)
* [type Uint64](#Uint64)
  * [func (Uint64) Max(numerics ...uint64) uint64](#Uint64.Max)
  * [func (Uint64) Min(numerics ...uint64) uint64](#Uint64.Min)
* [type Uint8](#Uint8)
  * [func (Uint8) Max(numerics ...uint8) uint8](#Uint8.Max)
  * [func (Uint8) Min(numerics ...uint8) uint8](#Uint8.Min)


#### <a name="pkg-files">Package files</a>
[maths.go](/src/github.com/andy2046/maths/maths.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // StringVar represents instance of string.
    StringVar = new(String)
    // IntVar represents instance of int.
    IntVar = new(Int)
    // Int8Var represents instance of int8.
    Int8Var = new(Int8)
    // Int16Var represents instance of int16.
    Int16Var = new(Int16)
    // Int32Var represents instance of int32.
    Int32Var = new(Int32)
    // Int64Var represents instance of int64.
    Int64Var = new(Int64)
    // UintVar represents instance of uint.
    UintVar = new(Uint)
    // Uint8Var represents instance of uint8.
    Uint8Var = new(Uint8)
    // Uint16Var represents instance of uint16.
    Uint16Var = new(Uint16)
    // Uint32Var represents instance of uint32.
    Uint32Var = new(Uint32)
    // Uint64Var represents instance of uint64.
    Uint64Var = new(Uint64)
    // ByteVar represents instance of byte.
    ByteVar = new(Byte)
    // RuneVar represents instance of rune.
    RuneVar = new(Rune)
    // Float32Var represents instance of float32.
    Float32Var = new(Float32)
    // Float64Var represents instance of float64.
    Float64Var = new(Float64)
)
```


## <a name="Range">func</a> [Range](/src/target/maths.go?s=1827:1857#L74)
``` go
func Range(end int) []struct{}
```
Range creates a range progressing from zero up to, but not including end.



## <a name="Ternary">func</a> [Ternary](/src/target/maths.go?s=8017:8083#L401)
``` go
func Ternary(expr bool, trueVal, falseVal interface{}) interface{}
```
Ternary for no ternary operation in Go.




## <a name="Byte">type</a> [Byte](/src/target/maths.go?s=580:589#L31)
``` go
type Byte byte
```
Byte represents type byte.










### <a name="Byte.Max">func</a> (Byte) [Max](/src/target/maths.go?s=3160:3198#L137)
``` go
func (Byte) Max(numerics ...byte) byte
```
Max returns the largest in numerics.




### <a name="Byte.Min">func</a> (Byte) [Min](/src/target/maths.go?s=3360:3398#L148)
``` go
func (Byte) Min(numerics ...byte) byte
```
Min returns the smallest in numerics.




## <a name="Float32">type</a> [Float32](/src/target/maths.go?s=670:685#L35)
``` go
type Float32 float32
```
Float32 represents type float32.










### <a name="Float32.Max">func</a> (Float32) [Max](/src/target/maths.go?s=1933:1980#L79)
``` go
func (Float32) Max(numerics ...float32) float32
```
Max returns the largest in numerics.




### <a name="Float32.Min">func</a> (Float32) [Min](/src/target/maths.go?s=2154:2201#L88)
``` go
func (Float32) Min(numerics ...float32) float32
```
Min returns the smallest in numerics.




## <a name="Float64">type</a> [Float64](/src/target/maths.go?s=724:739#L37)
``` go
type Float64 float64
```
Float64 represents type float64.










### <a name="Float64.Max">func</a> (Float64) [Max](/src/target/maths.go?s=2374:2421#L97)
``` go
func (Float64) Max(numerics ...float64) float64
```
Max returns the largest in numerics.




### <a name="Float64.Min">func</a> (Float64) [Min](/src/target/maths.go?s=2568:2615#L106)
``` go
func (Float64) Min(numerics ...float64) float64
```
Min returns the smallest in numerics.




## <a name="Int">type</a> [Int](/src/target/maths.go?s=122:129#L11)
``` go
type Int int
```
Int represents type int.










### <a name="Int.Max">func</a> (Int) [Max](/src/target/maths.go?s=7210:7245#L357)
``` go
func (Int) Max(numerics ...int) int
```
Max returns the largest in numerics.




### <a name="Int.Min">func</a> (Int) [Min](/src/target/maths.go?s=7407:7442#L368)
``` go
func (Int) Min(numerics ...int) int
```
Min returns the smallest in numerics.




## <a name="Int16">type</a> [Int16](/src/target/maths.go?s=206:217#L15)
``` go
type Int16 int16
```
Int16 represents type int16.










### <a name="Int16.Max">func</a> (Int16) [Max](/src/target/maths.go?s=6406:6447#L313)
``` go
func (Int16) Max(numerics ...int16) int16
```
Max returns the largest in numerics.




### <a name="Int16.Min">func</a> (Int16) [Min](/src/target/maths.go?s=6609:6650#L324)
``` go
func (Int16) Min(numerics ...int16) int16
```
Min returns the smallest in numerics.




## <a name="Int32">type</a> [Int32](/src/target/maths.go?s=252:263#L17)
``` go
type Int32 int32
```
Int32 represents type int32.










### <a name="Int32.Max">func</a> (Int32) [Max](/src/target/maths.go?s=6001:6042#L291)
``` go
func (Int32) Max(numerics ...int32) int32
```
Max returns the largest in numerics.




### <a name="Int32.Min">func</a> (Int32) [Min](/src/target/maths.go?s=6204:6245#L302)
``` go
func (Int32) Min(numerics ...int32) int32
```
Min returns the smallest in numerics.




## <a name="Int64">type</a> [Int64](/src/target/maths.go?s=298:309#L19)
``` go
type Int64 int64
```
Int64 represents type int64.










### <a name="Int64.Max">func</a> (Int64) [Max](/src/target/maths.go?s=5596:5637#L269)
``` go
func (Int64) Max(numerics ...int64) int64
```
Max returns the largest in numerics.




### <a name="Int64.Min">func</a> (Int64) [Min](/src/target/maths.go?s=5799:5840#L280)
``` go
func (Int64) Min(numerics ...int64) int64
```
Min returns the smallest in numerics.




## <a name="Int8">type</a> [Int8](/src/target/maths.go?s=162:171#L13)
``` go
type Int8 int8
```
Int8 represents type int8.










### <a name="Int8.Max">func</a> (Int8) [Max](/src/target/maths.go?s=6811:6849#L335)
``` go
func (Int8) Max(numerics ...int8) int8
```
Max returns the largest in numerics.




### <a name="Int8.Min">func</a> (Int8) [Min](/src/target/maths.go?s=7011:7049#L346)
``` go
func (Int8) Min(numerics ...int8) int8
```
Min returns the smallest in numerics.




## <a name="Rune">type</a> [Rune](/src/target/maths.go?s=622:631#L33)
``` go
type Rune rune
```
Rune represents type rune.










### <a name="Rune.Max">func</a> (Rune) [Max](/src/target/maths.go?s=2761:2799#L115)
``` go
func (Rune) Max(numerics ...rune) rune
```
Max returns the largest in numerics.




### <a name="Rune.Min">func</a> (Rune) [Min](/src/target/maths.go?s=2961:2999#L126)
``` go
func (Rune) Min(numerics ...rune) rune
```
Min returns the smallest in numerics.




## <a name="String">type</a> [String](/src/target/maths.go?s=78:91#L9)
``` go
type String string
```
String represents type string.










### <a name="String.Max">func</a> (String) [Max](/src/target/maths.go?s=7603:7647#L379)
``` go
func (String) Max(numerics ...string) string
```
Max returns the largest in numerics.




### <a name="String.Min">func</a> (String) [Min](/src/target/maths.go?s=7809:7853#L390)
``` go
func (String) Min(numerics ...string) string
```
Min returns the smallest in numerics.




## <a name="Uint">type</a> [Uint](/src/target/maths.go?s=342:351#L21)
``` go
type Uint uint
```
Uint represents type uint.










### <a name="Uint.Max">func</a> (Uint) [Max](/src/target/maths.go?s=5197:5235#L247)
``` go
func (Uint) Max(numerics ...uint) uint
```
Max returns the largest in numerics.




### <a name="Uint.Min">func</a> (Uint) [Min](/src/target/maths.go?s=5397:5435#L258)
``` go
func (Uint) Min(numerics ...uint) uint
```
Min returns the smallest in numerics.




## <a name="Uint16">type</a> [Uint16](/src/target/maths.go?s=434:447#L25)
``` go
type Uint16 uint16
```
Uint16 represents type uint16.










### <a name="Uint16.Max">func</a> (Uint16) [Max](/src/target/maths.go?s=4381:4425#L203)
``` go
func (Uint16) Max(numerics ...uint16) uint16
```
Max returns the largest in numerics.




### <a name="Uint16.Min">func</a> (Uint16) [Min](/src/target/maths.go?s=4587:4631#L214)
``` go
func (Uint16) Min(numerics ...uint16) uint16
```
Min returns the smallest in numerics.




## <a name="Uint32">type</a> [Uint32](/src/target/maths.go?s=484:497#L27)
``` go
type Uint32 uint32
```
Uint32 represents type uint32.










### <a name="Uint32.Max">func</a> (Uint32) [Max](/src/target/maths.go?s=3970:4014#L181)
``` go
func (Uint32) Max(numerics ...uint32) uint32
```
Max returns the largest in numerics.




### <a name="Uint32.Min">func</a> (Uint32) [Min](/src/target/maths.go?s=4176:4220#L192)
``` go
func (Uint32) Min(numerics ...uint32) uint32
```
Min returns the smallest in numerics.




## <a name="Uint64">type</a> [Uint64](/src/target/maths.go?s=534:547#L29)
``` go
type Uint64 uint64
```
Uint64 represents type uint64.










### <a name="Uint64.Max">func</a> (Uint64) [Max](/src/target/maths.go?s=3559:3603#L159)
``` go
func (Uint64) Max(numerics ...uint64) uint64
```
Max returns the largest in numerics.




### <a name="Uint64.Min">func</a> (Uint64) [Min](/src/target/maths.go?s=3765:3809#L170)
``` go
func (Uint64) Min(numerics ...uint64) uint64
```
Min returns the smallest in numerics.




## <a name="Uint8">type</a> [Uint8](/src/target/maths.go?s=386:397#L23)
``` go
type Uint8 uint8
```
Uint8 represents type uint8.










### <a name="Uint8.Max">func</a> (Uint8) [Max](/src/target/maths.go?s=4792:4833#L225)
``` go
func (Uint8) Max(numerics ...uint8) uint8
```
Max returns the largest in numerics.




### <a name="Uint8.Min">func</a> (Uint8) [Min](/src/target/maths.go?s=4995:5036#L236)
``` go
func (Uint8) Min(numerics ...uint8) uint8
```
Min returns the smallest in numerics.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
