# melon
A simple business indicator tool that uses a sliding window to detect
whether the indicator exceeds the threshold

### Usage
```go
//create the metric
//the ring length is 5000, greater is better for the accuracy compute, it can be equal to instance qps
// if 18 bad in latest 20 point, it return false
// if 70 bad in latest 100 point, the same to above.
me := melon.New(5000, OptionAnchor(20, 18), OptionAnchor(100, 70))

//set metric like below:
//if everything ok
me.Feed(true)
//if something wrong
me.Feed(false)

//get metric health stat
if me.OK() {
   //it means that your system is health
} else {
   //todo some repair
}
```

### Benchmark
> on 4core8G machine, windows system.
#### 1.benchmark feed
```go
//function
func BenchmarkMelonFeed(b *testing.B) {
	me := New(5000, OptionAnchor(100, 70))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 == 0 {
			me.Feed(true)
		} else {
			me.Feed(false)
		}
	}
}

//result:
goos: windows
goarch: amd64
pkg: github.com/alwaysthanks/melon
BenchmarkMelonFeed-4   	50000000	        28.6 ns/op
PASS
```

#### 2.benchmark OK
```go
func BenchmarkMelonOK(b *testing.B) {
	me := New(5000, OptionAnchor(100, 70))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		me.OK()
	}
}

//result:
goos: windows
goarch: amd64
pkg: github.com/alwaysthanks/melon
BenchmarkMelonOk-4   	 1000000	      1095 ns/op
PASS
```