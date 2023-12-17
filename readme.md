Declaring variables in go:

var variableName type

variableName ;= value


GO is  a static type language


Data types:

bool, string, int, float64


Function declaration 


func function_name return_type(){
return “” 
}



Data types:

bool, string, int, float64
array,map

type deck[] string   ⇒ tells go we want to create array of strings and add bunch of functions specially made to work with it.



Reciever functions

In Go, a method with a receiver is referred to as a "receiver function." The (c Circle) part in the method signature func (c Circle) Area() float64 is the receiver. This means that the Area method is associated with the Circle type, and you can call it on instances of the Circle type.

func (d deck) print(){
} 

Any variable of the deck type will get access to the print method   ==> that is any variable of type deck can call this function on itself.
The deck type variable which we will pass to the print function is available as 'd'


VAriable '_'


Slice range syntax

fruit = {'a', 'b', 'c', 'd'}

fruit [startIndexIncluding, endIndexExcluding]
fruit[0:2] ==> a,b


Byte slice

Type conversion

greeting := "Hi There"
fmt.println([]byte(greeting))


STRUCTS - defining, declaring a struct

type ==> zero-value
string ==> ""
int ==> 0
float ==> 0
bool ==> false

Printf

Go is a pass-by-value language==> whenever we pass a variable to a function(as reciever) it creates a separate copy of that variable at another memory location which is referenced by a new variable


Pointers:

&variable ==> give me the memory address of value this variable is pointing at.
*pointer ==> give me the value this memory address is pointing at.


Struct vs slice

Array vs slices


value types ==> int,float,string, bool, structs           ==> Use pointers to change these things in function
reference types ==> slices,maps, channels, pointers, functions   ==> Don't worry about pointers with these

clean 

INTERFACES

There are 2 types ==> concrete type ==> we can define value and interface type (we can't define it's value)

Interfaces are not generic types ==> code reusability increases
They are implict.For declaring a type to be within interface that particular custom type should have the functions defined within the interface ==> for the types to be considered within interface  


- Interfaces are just for reusing code ==> they are just contract to help us manage types


READER INTERFACE



WRITER INTERFACE
any other custom type having the implementation of function Write, implements the writer interface


io.Copy(os.Stdout, resp.Body)
io.copy ==> (something-implements-writer-interface, something-implements-reader-interface)

-  something implements writer interface && 

   os.Stdout ==> value of type file ==> *file ==> file has function Write ==> hence it implements writer interface


-  something that implements reader interface

   resp.Body ==> it has Read function ==> hence it implements reader interface.          


------------------------------------------------------------------------------------------
   GO routines and Channels
------------------------------------------------------------------------------------------


concurrency vs parallelism

Concurrently ==> a program is concurrent if it has ability to load up multiple go routines at the same time. They might be running on one core ==> 
Our program has abiility to run diff things kind of at the same time but not at the same time.
We can scheduler work to be done, we don't have to wait for one go routine to finish to start another

Parallelism:
It is possible with multiple cores. We are doing multiple things at exact same time. When one core is picking up one go routine for execution, another core can pick up another routine for execution at the same time.

Concurrency ==> we can schedule work(routines) and change between them on the fly
Parallelism ==> we can literally do multiple things at the same time.



we use 'go' keyword in front of function call 

Channels:
channels are used to communicate between different go routines



Never everr try to share variables between parent and child go routines
Never try to access same variable from different child routines. We can access var by passing it as arg, or communicating it through channel



-------------------------------------------------------------------------------------------------
GOLANG COMMANDS:
-------------------------------------------------------------------------------------------------



~/bin/exercism download --exercise=list-ops --track=go

-------------------------------------------------------------------------------------------------
MAKE
------------------------------------------------------------------------------------------------------
the make function is used to create and initialize certain composite types, namely slices, maps, and channels



-------------------------------------------------------------------------------------------------
FUNCTIONAL LITERAL
------------------------------------------------------------------------------------------------------

funcitonal literal in go is just like an anonymous function in javascript


------------------------------------------------------------------------------------------------------
Converting int to string and vice sas versa 
------------------------------------------------------------------------------------------------------

strconv package has func to convert string to an integer &vice sa versa

	myString := "42"

	myInt,err := strconv.Atoi(myString)


   strconv.Itoa(myInt)


String literals are immutable in Go, meaning you cannot change the individual characters of a string once it's created. If you need to modify strings, you can convert them to a slice of runes, which is the Unicode character type in Go, and then perform modifications.

EXAMPLE:
str := "Hello, Go!"
runes := []rune(str)                // Convert string to a slice of runes
runes[7] = '!'
modifiedStr := string(runes)




/home/10683796/exercism

TYPE CONVERSION


int(var_name)
float64(var_name)
string(var_name)

are three ways for type conversion in go

we can convert int to float, string
               float to int

Converting string to different types and vice-sa versa


	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	
	s := "2147483647" // biggest int32
	i64, err := strconv.ParseInt(s, 10, 32)

	i := int32(i64)

	s := strconv.FormatBool(true)
	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s := strconv.FormatInt(-42, 16)
	s := strconv.FormatUint(42, 16)


RUNES:

In Go, a rune is a built-in type that represents a Unicode code point. A Unicode code point is an integer value that uniquely identifies a character in the Unicode standard. Unlike some other programming languages, where characters are represented as bytes, Go uses runes to represent characters in a more abstract and Unicode-friendly way.

