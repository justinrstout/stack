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
Benchmark_LargePush/container/list-12         	10899460	       126 ns/op
Benchmark_LargePush/stack/ints-12             	100000000	        12.4 ns/op
Benchmark_LargePush/stack/fixed-12            	972272236	         1.23 ns/op
Benchmark_LargePush/handwritten-12            	937995576	         1.36 ns/op
Benchmark_SmallPushPop/container/list-12      	 8018734	       146 ns/op
Benchmark_SmallPushPop/stack/ints-12          	327493423	         3.67 ns/op
Benchmark_SmallPushPop/stack/fixed-12         	290399660	         4.08 ns/op
Benchmark_SmallPushPop/handwritten-12         	226686604	         5.27 ns/op
```