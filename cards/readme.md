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