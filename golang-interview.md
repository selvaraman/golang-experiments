Declaring and Initializing variable:
------------------------------------
:= short hand for declaring and initializing

var f string = "apple"
f := "apple"

No ternary operator, no while
-----------------------------
Array:
-----
    var a [5]int
    b := [5]int{1, 2, 3, 4, 5}
    b = [...]int{1, 2, 3, 4, 5}
    var twoD [2][3]int
  twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }

Array vs Slice:
-------------
Array - fixed size
Slice - dynamic size

Array - can compare two arrays with same type
Slice - can not compare two slice

slice := make([]int, 0, 0)
=> create new array in heap
=> copy from old to new array
=> abondon the old array - GC- track, scan, free  -  world-pause - latency spikes
String:
------
'' - rune, char
for index, rune := range str {
strings are equivalent to []byte

range:
------
when ranging map, slice, go knows when the range is done, knows length of the map or slice
you can also range channel.
but ranging channel, go does not know when the channel done being sent. solve this using defer

struct
------
type person struct {
    name string
    age  int
}

map internal:
------------
m := map[int]string {1: "selvam"}
delete(m, 1)
m[1] = "raman"
for key, value := range m {
} // Note: this is unordered.
len(m)

hash(keys) / no-of-buckets
double when it reaches max
1. score, ok = m[key] // ok => true if exists
2. unordered
3. keys 
keys => string, int, float, bool, pointers { make(map[*int]string) }, channels, and interface
can not be keys => Slices, maps, and functions cannot be map keys
go.mod vs go.sum
----------------
go.mod - keeps dependencies required for the project
go.sum - holds checksums(hash) of direct, indirect, transitive dependencies, used to verify downloaded packages are correct, prevents tampering

mod vs package
--------------
Package - Collection of Go source files in a directory [ Used for dependency management and versioning ]
Module - Collection of related packages [ Used for code organization ]

GOROOT vs GOPATH
----------------
GOROOT - where Go lives { Go compiler, Standard library, Go tools (go, gofmt, etc) }
GOPATH - where your code/tools live

go install vs go get
--------------------
go install - build, install executables
go get - add or update dependencies required for the project

methods vs functions
--------------------
function - standalone
method - receiver, pointer receiver with mutation capability

variable length args:
--------------------
func list[T int | float32](args ...T) {

callback function
-----------------
sending a function as args to another function

execute(a int, b int, operation func(int, int)) {
operation(a, b)
}

execute(a, b, add)

Latest Golang
-------------
May 7, 2026
1.26.3

time.NewTimer / time.NewTicker
------------------------------
time.NewTimer => executes only once {timer.Stop}
timer := time.NewTimer(3 * time.Second)
<-timer.C

time.NewTicker => repeated execution {ticker.Stop, ticker.Reset}
ticker := time.NewTicker(1 * time.Second)

Channel:
-------
How to check a chennel is closed -  comma, ok idiom
msg, ok := <-ch

Generics:
---------
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyInt int

type GenType interface {
	~int | ~float32
}

func add[T GenType](a, b T) T {
	c := a + b
	return c
}

func sub[T constraints.Integer](a, b T) T {
	c := a - b
	return c
}

func main() {
	result1 := add[int](5, 6)
	result2 := add[float32](5.5, 6.6)
	result3 := add[MyInt](1, 1)
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
	fmt.Println("Result 3:", result3)
	result4 := sub[int](4, 5)
	fmt.Println("Result 4:", result4)
}
//comparable
//along with int, float, we can use struct too - 
type creature interface {
	dog | cat
}
func info[T creature](t T) string {


errors:
------
type error interface {
    Error() string
}

func (e errorStruct) Error() string {
errors.New
errors.As
errors.Is
Sentinel Error(errors.Is) vs Wrapped Error(errors.As)
var ErrNotFound = errors.New("not found") ==> Sentinel Error
fmt.Errorf("user lookup failed: %w", ErrNotFound) ==> Wrapped Error
fmt.Println(errors.New("host not provided") == errors.New("host not provided")) => pointer comparision returns false. everytime it creates new object.

enum:
----
const (
	A = iota // 0
	B        // 1
	C        // 2
)

const (
	a int = iota // 0
	b            // 1
	c            // 2
)

new vs make (allocates in heap):
-------------------------------
new => int, bool, string, struct => returns pointer
make => map, chan, slice => Ready to use values


Build:
-----
GOOS=windows GOARCH=amd64
GOOS=linux GOARCH=arm64
GOOS=darwin GOARCH=arm64

High Throughput and Low Latency:
-------------------------------
Throughput - How much work a system can complete in a given time.
Latency - How long it takes for a single request to complete.

Synchronous vs Asynchronous:
----------------------------
Synchronous => A task must finish before starting the next task
Asynchronous => A task started, the program does not wait for it to finish

Context:
-------
https://leapcell.medium.com/golang-context-deep-dive-from-zero-to-hero-d9793617e61c

SOLID
-----
https://medium.com/@vishal/understanding-solid-principles-in-golang-a-guide-with-examples-f887172782a3

Concurrency Pattern:
-------------------
https://freedium-mirror.cfd/https://medium.com/@mehul25/concurrency-patterns-in-go-4ffca7b9e295

GC:
---
Stop The World, Trie-Color-Marking - Terminate STW - Reclaim the memory marked white.
https://rugu.dev/en/blog/understanding-go-gc/

Race Condition:
--------------
sync.Mutex => mu.Lock(), mu.Unlock()

sync.RWMutex => mu.RLock(), mu.RUnlock(), mu.Lock(), mu.Unlock() => Read Heavy workload
=> Allows multipel readers
RLock() = people reading books.
Lock() = librarian updating the catalog.
Many readers can read together
While the librarian is updating, nobody can read
While readers are reading, the librarian waits.

sync/atomic => AddInt64, LoadInt64 => shared state access without lock
=> hardware-supported atomic CPU instructions
=> Limited only Integers and Pointers
=> atomic operations are limited. Add/Sub/Store/CompareAndSwap/Swap/Load

Idiomatic Go Guide:
-------------------
https://learning.christofherkost.org/


For Dev:
-------
https://github.com/Melkeydev/go-blueprint
