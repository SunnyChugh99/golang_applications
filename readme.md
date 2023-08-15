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