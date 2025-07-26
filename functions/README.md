## Functions ##
* Functions are first class objects. They may be assign to a variable, passed as an argument to another function, or returned as the value of another function.

```
func funcName(arguments) returnedTypes {
    // body
}
```

## Closures ##
* Closures - a function value that references variables out the body of the function.

```
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }

}
```

## Pointer and value argument passing ## 
* All arguments are passed by value.
* To change the argument being passed, pass a pointer:

```
func update(x *int) {
    *x = 42
}
``` 
* For slices, maps, and channels, a “pointer header” is passed, so they behave like a pointer pass.

## Methods ##
* Method are function with reciever.

```
func (t MyType) DoSomething() // by value
func (t *MyType) ModifySomething // by pointer
```
* Reciever may be passed by value or pointer, depending on the desired behavior.
* Methods can only be defined only for types, declared in the same package.


## Interfaces ##
* An Interface is a set of method signatures.
* A type implements an interface implicity if it has the right methods

```
type Stringer() interface {
    String() string
}
```

## No overloading and functions with a varible number of named parametrs##
* You cannot declare two functions with the same name and different parameters.
* You can use variative parametrs

```
func sum(nums...int) int
```

## Functions can return multiple values ##
* Includes value + error

```
func divide(a, b int) (int, error)
```

## Anonymous functions ##
* Can be defined without a name, useful for closures and one-liners.

```
f := func(x int) int {
    return x * x
}
```

## Defer functions ##
* Executed in reverse order when exiting the function.
* They are used to free up resources:

```
defer file.CLose()
```

