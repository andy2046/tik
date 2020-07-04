

# tik
`import "github.com/andy2046/tik"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package tik implements Hierarchical Timing Wheels.




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Callback](#Callback)
* [type Config](#Config)
* [type DefaultTimer](#DefaultTimer)
  * [func (dt *DefaultTimer) Now() uint64](#DefaultTimer.Now)
  * [func (dt *DefaultTimer) Step() &lt;-chan uint64](#DefaultTimer.Step)
  * [func (dt *DefaultTimer) Stop()](#DefaultTimer.Stop)
* [type Option](#Option)
* [type Ticker](#Ticker)
  * [func New(options ...Option) *Ticker](#New)
  * [func (tk *Ticker) AnyExpired() bool](#Ticker.AnyExpired)
  * [func (tk *Ticker) AnyPending() bool](#Ticker.AnyPending)
  * [func (tk *Ticker) Cancel(to *Timeout)](#Ticker.Cancel)
  * [func (tk *Ticker) Close()](#Ticker.Close)
  * [func (tk *Ticker) IsClosed() bool](#Ticker.IsClosed)
  * [func (tk *Ticker) Schedule(delay uint64, cb Callback) *Timeout](#Ticker.Schedule)
* [type Timeout](#Timeout)
  * [func (to *Timeout) Expired() bool](#Timeout.Expired)
  * [func (to *Timeout) Pending() bool](#Timeout.Pending)
* [type Timer](#Timer)
  * [func NewTimer(interval uint64) Timer](#NewTimer)


#### <a name="pkg-files">Package files</a>
[bit.go](/src/github.com/andy2046/tik/bit.go) [ticker.go](/src/github.com/andy2046/tik/ticker.go) [timeout.go](/src/github.com/andy2046/tik/timeout.go) [timer.go](/src/github.com/andy2046/tik/timer.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // DefaultConfig is the default Ticker Config.
    DefaultConfig = Config{
        WheelBitNum: 6,
        WheelNum:    4,
        Timer:       nil,
    }
)
```



## <a name="Callback">type</a> [Callback](/src/target/timeout.go?s=147:162#L12)
``` go
type Callback func()
```
Callback function to trigger when timeout expires.










## <a name="Config">type</a> [Config](/src/target/ticker.go?s=1406:1587#L65)
``` go
type Config struct {
    // number of value bits mapped in each wheel.
    WheelBitNum uint8
    // number of wheels.
    WheelNum uint8
    // Timer to progress the timing wheel.
    Timer Timer
}
```
Config used to init Ticker.










## <a name="DefaultTimer">type</a> [DefaultTimer](/src/target/timer.go?s=352:484#L22)
``` go
type DefaultTimer struct {
    // contains filtered or unexported fields
}
```
DefaultTimer implements Timer interface.










### <a name="DefaultTimer.Now">func</a> (\*DefaultTimer) [Now](/src/target/timer.go?s=1157:1193#L62)
``` go
func (dt *DefaultTimer) Now() uint64
```
Now returns the absolute time when timer started in millisecond.




### <a name="DefaultTimer.Step">func</a> (\*DefaultTimer) [Step](/src/target/timer.go?s=1263:1307#L67)
``` go
func (dt *DefaultTimer) Step() <-chan uint64
```
Step timing wheel by absolute time in millisecond.




### <a name="DefaultTimer.Stop">func</a> (\*DefaultTimer) [Stop](/src/target/timer.go?s=1359:1389#L72)
``` go
func (dt *DefaultTimer) Stop()
```
Stop the DefaultTimer.




## <a name="Option">type</a> [Option](/src/target/ticker.go?s=1634:1656#L75)
``` go
type Option = func(*Config)
```
Option applies config to Ticker Config.










## <a name="Ticker">type</a> [Ticker](/src/target/ticker.go?s=499:1224#L31)
``` go
type Ticker struct {
    // contains filtered or unexported fields
}
```
Ticker progress the timing wheels.







### <a name="New">func</a> [New](/src/target/ticker.go?s=1832:1867#L88)
``` go
func New(options ...Option) *Ticker
```
New initiates a new Ticker.





### <a name="Ticker.AnyExpired">func</a> (\*Ticker) [AnyExpired](/src/target/ticker.go?s=9546:9581#L432)
``` go
func (tk *Ticker) AnyExpired() bool
```
AnyExpired returns true if expiry queue is not empty, false otherwise.




### <a name="Ticker.AnyPending">func</a> (\*Ticker) [AnyPending](/src/target/ticker.go?s=9687:9722#L437)
``` go
func (tk *Ticker) AnyPending() bool
```
AnyPending returns true if there is task in wheels, false otherwise.




### <a name="Ticker.Cancel">func</a> (\*Ticker) [Cancel](/src/target/ticker.go?s=3981:4018#L193)
``` go
func (tk *Ticker) Cancel(to *Timeout)
```
Cancel the Timeout scheduled if it has not yet expired.




### <a name="Ticker.Close">func</a> (\*Ticker) [Close](/src/target/ticker.go?s=4137:4162#L203)
``` go
func (tk *Ticker) Close()
```
Close stop processing any task,
whether it is pending or expired.




### <a name="Ticker.IsClosed">func</a> (\*Ticker) [IsClosed](/src/target/ticker.go?s=4330:4363#L213)
``` go
func (tk *Ticker) IsClosed() bool
```
IsClosed returns true if closed, false otherwise.




### <a name="Ticker.Schedule">func</a> (\*Ticker) [Schedule](/src/target/ticker.go?s=3704:3766#L181)
``` go
func (tk *Ticker) Schedule(delay uint64, cb Callback) *Timeout
```
Schedule creates a one-shot action that executed after the given delay.
`delay` is the time from now to delay execution,
the time unit of the delay depends on the Timer provided, default is millisecond.
`cb` is the task to execute.
it returns `nil` if Ticker is closed.




## <a name="Timeout">type</a> [Timeout](/src/target/timeout.go?s=208:509#L15)
``` go
type Timeout struct {
    // contains filtered or unexported fields
}
```
Timeout represents user timeout logic.










### <a name="Timeout.Expired">func</a> (\*Timeout) [Expired](/src/target/timeout.go?s=1035:1068#L56)
``` go
func (to *Timeout) Expired() bool
```
Expired returns true if timeout is in expired queue, false otherwise.




### <a name="Timeout.Pending">func</a> (\*Timeout) [Pending](/src/target/timeout.go?s=809:842#L48)
``` go
func (to *Timeout) Pending() bool
```
Pending returns true if timeout is in timing wheel, false otherwise.




## <a name="Timer">type</a> [Timer](/src/target/timer.go?s=102:304#L10)
``` go
type Timer interface {
    // Now returns the absolute time when timer started.
    Now() uint64

    // Step channel to step timing wheel by absolute time.
    Step() <-chan uint64

    // Stop the timer.
    Stop()
}
```
Timer progress the timing wheel by current time.







### <a name="NewTimer">func</a> [NewTimer](/src/target/timer.go?s=551:587#L32)
``` go
func NewTimer(interval uint64) Timer
```
NewTimer returns DefaultTimer with interval in millisecond.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
