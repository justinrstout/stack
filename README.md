# Fast stacks for Go programs

This module implements stacks for `int`s as a lightweight wrapper around the equivalent handwritten code. The intended use is for high performance within stack-based algorithms, and not as a general-purpose structure. As such, bounds checking, error handling, and polymorphism are not features of this library. 

There are two varieties of stacks available:

## `stack/ints`
`stack/ints` is backed by a dynamic array with a growth factor of 1.5. Pop(), Peek(), and Empty() are constant-time operations. Push() is an amortized constant-time operation that performs a memory allocation and copy when the underlying array capacity is exceeded.

## `stack/fixed`
`stack/fixed` is backed by an array with a size of 256 elements. Once the stack is created, all operations are constant-time and no further memory allocations are performed.

# Example

```go
package main

import "github.com/justinrstout/stack/ints"

func example() bool {
	st := ints.New()   
	st.Push(199)
	val := st.Peek()
	eq := st.Pop() == val
	return st.Empty() && eq // returns true
}


```

# Benchmarks
This module achieves 1-2 orders of magnitude better performance than the generic standard library list, and comparable performance to the equivalent handwritten code.
```
goos: darwin
goarch: amd64
pkg: github.com/justinrstout/stack
Benchmark_LargePush/container/list-12         	 8767942	       115 ns/op
Benchmark_LargePush/stack/ints-12             	100000000	        13.3 ns/op
Benchmark_LargePush/handwritten-12            	925843028	         1.53 ns/op
Benchmark_SmallPushPop/container/list-12      	 8065232	       148 ns/op
Benchmark_SmallPushPop/stack/ints-12          	327332500	         3.66 ns/op
Benchmark_SmallPushPop/stack/fixed-12         	297872582	         4.11 ns/op
Benchmark_SmallPushPop/handwritten-12         	230282539	         5.31 ns/op
```