------------------------------------------------------------------------------------------------------------
GOLANG & IT'S USE
------------------------------------------------------------------------------------------------------------
Go is a statically typed, concurrent, and garbage-collected programming language created at Google in 2009. 

It is designed to be simple, efficient, and easy to learn, making it a popular choice for building scalable network services, web applications, and command-line tools.

Go is known for its support for concurrency, which is the ability to run multiple tasks simultaneously. Concurrency is achieved in Go through the use of Goroutines and Channels, which allow you to write code that can run multiple operations at the same time. This makes Go an ideal choice for building high-performance and scalable network services, as well as for solving complex computational problems.

Another important feature of Go is its garbage collection, which automatically manages memory for you.

WHY GO?

Because Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aims to be modern, with support for networked and multicore computing. 


------------------------------------------------------------------------------------------------------------
VARIABLE DECLARATION & DATA TYPES
------------------------------------------------------------------------------------------------------------
var variable_name data_type = value_of_variable
var card string = "nine of spades"
card := "nine of spades"
variable_name := value_of_variable


------------------------------------------------------------------------------------------------------------
LOOPING 
------------------------------------------------------------------------------------------------------------







------------------------------------------------------------------------------------------------------------
LISTS
------------------------------------------------------------------------------------------------------------






------------------------------------------------------------------------------------------------------------
MAPS
------------------------------------------------------------------------------------------------------------

In Go, a map is a built-in data structure that provides an unordered collection of key-value pairs. It is sometimes referred to as an associative array, hash map, or dictionary in other programming languages. The keys in a map must be of a comparable type, and the values can be of any type.

ages := make(map[string]int)  // creates an empty map with string keys and int values


You can also use a map literal to create and initialize a map:
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Charlie": 22,
}


------------------------------------------------------------------------------------------------------------
STRUCTS
------------------------------------------------------------------------------------------------------------

In Go, a struct is a composite data type that groups together variables (fields) under a single name. It allows you to create a custom data type with multiple properties, similar to a class or a record in other programming languages


------------------------------------------------------------------------------------------------------------
INTERFACES
------------------------------------------------------------------------------------------------------------

In Go, interfaces and polymorphism play a significant role in achieving flexibility and code reusability. Let's explore these concepts:

Interfaces:
An interface in Go is a collection of method signatures. It defines a set of methods that a type must implement to satisfy the interface. A type implicitly satisfies an interface if it implements all the methods declared by that interface. Unlike some other languages, Go uses implicit interface implementation, meaning you don't need to explicitly declare that a type implements an interface.

Interfaces essentially act as placeholders for methods that will have multiple implementations based on what object is using it.

Polymorphism:
Polymorphism refers to the ability of a single piece of code to work with values of different types. In Go, polymorphism is realized through interface types and the ability of different types to implement these interfaces.


Benefits of Polymorphism in Go:

Code Reusability: Polymorphism allows you to write functions and methods that can be reused with different types, as long as they adhere to the specified interfaces.
Flexibility: You can add new types that satisfy existing interfaces, extending the functionality of your code without modifying existing functions.



------------------------------------------------------------------------------------------------------------
POINTERS
------------------------------------------------------------------------------------------------------------




------------------------------------------------------------------------------------------------------------
RECIEVER FUNCTIONS
------------------------------------------------------------------------------------------------------------

In Go, a receiver function is a method associated with a type. A receiver is a special parameter of a method that allows you to call the method on an instance of that type. In other programming languages, this concept is often referred to as a method or member function. Receiver functions are a way to associate behavior with a type in Go.





------------------------------------------------------------------------------------------------------------
DEFER,PANIC,RECOVER
------------------------------------------------------------------------------------------------------------
In Go, defer, panic, and recover are keywords used for handling errors and cleanup operations in a program.

defer: The defer statement is used to ensure that a function call is performed later in a program's execution, usually for cleanup operations. The deferred function call is placed on a stack, and it will be executed after the surrounding function returns.


panic: The panic function is used to terminate the normal flow of control and begin panicking. It's often used to indicate a runtime error. When a function encounters a panic, it immediately stops executing, and the deferred functions from its call stack are executed.


recover: The recover function is used to regain control of a panicking goroutine. It should be called in a deferred function. If a deferred function calls recover and the surrounding function is panicking, recover stops the panicking and returns the value passed to the panic function.





------------------------------------------------------------------------------------------------------------
GO ROUNTINES AND CHANNELS
------------------------------------------------------------------------------------------------------------
Go routines:

Goroutines are lightweight threads managed by the Go runtime. They are used for concurrent programming and can be thought of as functions running concurrently with other functions. They are more efficient than traditional threads, and the Go runtime handles their scheduling.


Channel:
a channel is a communication mechanism that allows one goroutine to send data to another goroutine. It provides a way for two concurrent goroutines to synchronize execution and safely communicate by sending and receiving values.

for channel creation:

c := make(chan string)



------------------------------------------------------------------------------------------------------------
CONCURRENCY VS PARALLELISM
------------------------------------------------------------------------------------------------------------

Concurrently ==> a program is concurrent if it has ability to load up multiple go routines at the same time. They might be running on one core ==> 
Our program has abiility to run diff things kind of at the same time but not at the same time.
We can scheduler work to be done, we don't have to wait for one go routine to finish to start another

Parallelism:
It is possible with multiple cores. We are doing multiple things at exact same time. When one core is picking up one go routine for execution, another core can pick up another routine for execution at the same time.

Concurrency ==> we can schedule work(routines) and change between them on the fly
Parallelism ==> we can literally do multiple things at the same time.




------------------------------------------------------------------------------------------------------------
ERROR HANDLING
-----------------------------------------------------------------------------------------------------------

In Go, error handling is done through the use of the error type, and it's a common practice to return errors as a second return value from functions. 

------------------------------------------------------------------------------------------------------------
GO PACKAGE HANDLING
------------------------------------------------------------------------------------------------------------

a.  go mod init
The go mod init command is used to initialize a new module. A module is a collection of related Go packages that are versioned together as a single unit. This command is typically used when you are starting a new project.
SYNTAX:  go mod init example.com/myproject

b.  go get
The go get command is used to add dependencies to your project. It can be used to download and install packages from version control systems (like Git) and add them to your go.mod file.
SYNTAX: go get <module-path>
Example: go get github.com/example/package

This command fetches the specified package and its dependencies, updates the go.mod file with the new dependency information, and creates or updates the go.sum file, which contains checksums for the downloaded module versions.

c - go.mod File:

The go.mod file is the module definition file for your project. It includes the module path, the Go version compatibility, and the direct dependencies along with their versions.

d go.sum File:

The go.sum file contains checksums for module versions. It helps ensure the integrity of the dependencies by storing cryptographic hashes.
These files help ensure that your project's dependencies are reproducible and that others working on your project get the same versions of dependencies.


-----------------------------------------------------------------------------------------------------------
Boolean operators and precedence
-----------------------------------------------------------------------------------------------------------

&&, ||, !

The three boolean operators each have a different operator precedence. As a consequence, they are evaluated in this order: ! first, && second, and finally ||. If you want to force a different ordering, you can enclose a boolean expression in parentheses (ie. ()), as the parentheses have an even higher operator precedence.


-----------------------------------------------------------------------------------------------------------
Strings:
-----------------------------------------------------------------------------------------------------------



-----------------------------------------------------------------------------------------------------------
Variables access control 
-----------------------------------------------------------------------------------------------------------

Go determines if an item can be called by code in other packages through how it is declared. To make a function, type, variable, constant or struct field externally visible (known as exported) the name must start with a capital letter.

-----------------------------------------------------------------------------------------------------------
String Formatting
-----------------------------------------------------------------------------------------------------------

Go provides an in-built package called fmt (format package) which offers a variety of functions to manipulate the format of input and output. The most commonly used function is Sprintf, which uses verbs like %s to interpolate values into a string and returns that string.


-----------------------------------------------------------------------------------------------------------
Rand Package
-----------------------------------------------------------------------------------------------------------

Package math/rand provides support for generating pseudo-random numbers.

n := rand.Intn(100) // n is a random int, 0 <= n < 100
f := rand.Float64() // f is a random float64, 0.0 <= f < 1.0


Shuffle
x := []string{"a", "b", "c", "d", "e"}

rand.Shuffle(len(x), func(i, j int) {
	x[i], x[j] = x[j], x[i]
})

-----------------------------------------------------------------------------------------------------------
Parameters vs. Arguments
-----------------------------------------------------------------------------------------------------------
Let's quickly cover two terms that are often confused together: parameters and arguments. Function parameters are the names defined in the function's signature, such as greeting and name in the function PrintGreetingName above. Function arguments are the concrete values passed to the function parameters when we invoke the function. For instance, in the example below, "Hello" and "Katrina" are the arguments passed to the greeting and name parameters:

PrintGreetingName("Hello", "Katrina")

-----------------------------------------------------------------------------------------------------------
Named Return Values and Naked Return
-----------------------------------------------------------------------------------------------------------
As well as parameters, return values can optionally be named. If named return values are used, a return statement without arguments will return those values. This is known as a 'naked' return.


Normal function with return value (normal)
func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}

with Named Return Values and Naked Return

func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
    sum, mult = a+b, a*b
    sum -= c
    mult -= c
    return
}