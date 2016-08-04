## Slice - Length and Capacity

Slice is a dynamic array which grows or shrinks in size.

https://candysmurf.github.io/GoReflectionTalk

## Links

https://blog.golang.org/slices  
https://blog.golang.org/go-slices-usage-and-internals  
https://www.goinggo.net/2013/08/understanding-slices-in-go-programming.html


## Code Review

### Length and Capacity

Example shows that appending items to the pre-allocated slice.  
[Problem](code/problem/problem.go) ([Go Playground](https://play.golang.org/p/N--hzBDMtI))    

Example shows slice's length and capacity at its declared state.  
[Init](code/init/init.go) ([Go Playground](https://play.golang.org/p/fhHed9UBwL))  

Example shows that a slice is accessible when its index is between 0 to len(s)-1.  
[Accessible](code/accessible/accessible.go) ([Go Playground](https://play.golang.org/p/J6rdgyjBlL))  

Example shows that the valid index is in the range of x <= y <= cap(s) for slice s[x:y].  
[Sliceable](code/sliceable/sliceable.go) ([Go Playground](https://play.golang.org/p/_z-XSU9oEq))  

Example shows that append can append arbitrary length of data.  
[Appendable](code/appendable/appendable.go) ([Go Playground](https://play.golang.org/p/wRV3ac6tRr))    

Example shows that the initial cap allocation varies based on the data type. Afterward, it doubles.  
[CapInitialGrow](code/capInitialGrow/capInitialGrow.go) ([Go Playground](https://play.golang.org/p/4Y1HldIabG))  
  
___
All material is licensed under the [Apache License Version 2.0, August 2016](http://www.apache.org/licenses/LICENSE-2.0).
