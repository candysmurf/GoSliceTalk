## Slice - Length and Capacity

Slice is a dynamic array which grows or shrinks in size.

## Links

https://blog.golang.org/slices
https://blog.golang.org/go-slices-usage-and-internals
https://www.goinggo.net/2013/08/understanding-slices-in-go-programming.html


## Code Review

### Length and Capacity

Example shows that appending the value to the end of a slice instead of initilized slice slots.  
[Problem](code/problem/problem.go) ([Go Playground](https://play.golang.org/p/N--hzBDMtI))    

Example shows slice's length and capacity at initial states.  
[Init](code/init/init.go) ([Go Playground](https://play.golang.org/p/fhHed9UBwL))  

Example shows that a slice slot is accessible when its index < len(s). Otherwise, a panic occurs.  
[Accessible](code/accessible/accessible.go) ([Go Playground](https://play.golang.org/p/XhxE6m4cpE))  

Example shows that slicing s[x:y] has to have indexes in the range of x <= y <= cap(s).  
[Sliceable](code/sliceable/sliceable.go) ([Go Playground](https://play.golang.org/p/_z-XSU9oEq))  

Example shows append works for unknown length data.  
[Appendable](code/appendable/appendable.go) ([Go Playground](https://play.golang.org/p/GC951VMF7U))  

Example shows that the initial cap allocation varies based on the data type. Afterwards it doubles.  
[CapInitialGrow](code/interface/slice/slice.go) ([Go Playground](https://play.golang.org/p/GWYFCHCZdj))  
  
___
All material is licensed under the [Apache License Version 2.0, August 2016](http://www.apache.org/licenses/LICENSE-2.0).
