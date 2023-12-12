package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	fmt.Println("type conversion program\n")
	
	// int => float64, string
	fmt.Println("converting int to float and string") 
	myInt := 32

	myString := strconv.Itoa(myInt)
	myFloat := float64(myInt)
	fmt.Println(reflect.TypeOf(myString), myString)
	fmt.Println(reflect.TypeOf(myFloat), myFloat)


	// float64 => int, string
	fmt.Println("converting flaot to int and string")
	myFloat2 := 32.5
	myInt2 := int(myFloat2)
	fmt.Println(reflect.TypeOf(myInt2), myInt2)
	
	myString2 := strconv.FormatFloat(myFloat2, 'g', 5, 64)
	fmt.Println(myString2 + " myString2 is of type %T", myString2)


	// string => int, float
	myString3 := "42"

	myInt3,err := strconv.Atoi(myString3)
	if err!= nil{
		fmt.Println("error")
	}
	fmt.Println(reflect.TypeOf(myInt3), myInt3)


	// b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415", 64)
	// i, err := strconv.ParseInt("-42", 10, 64)
	// u, err := strconv.ParseUint("42", 10, 64)
	
	// s := "2147483647" // biggest int32
	// i64, err := strconv.ParseInt(s, 10, 32)

	// i := int32(i64)

	// s := strconv.FormatBool(true)
	// s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// s := strconv.FormatInt(-42, 16)
	// s := strconv.FormatUint(42, 16)

	fmt.Println("runes usage")
	str := "Hello, Go!"
	runes := []rune(str)  
	// Convert string to a slice of runes
	fmt.Println(runes)
	runes[7] = '!'
	modifiedStr := string(runes)

	fmt.Println(modifiedStr)
	// str11 := "Hello, 世界"
	for i:=0; i<len(runes)-1; i++{
		runes[i] = runes[i+1]
	}

	modifiedStr2 := string(runes)
	fmt.Println("final str")

	fmt.Println(modifiedStr2)

	fmt.Println("byte conv")
	str12 := "Hello, 世界"

	byteString := []byte(str12)
	fmt.Println(byteString)

	fmt.Println("TYPE ASSERTION IN GO")

	map1 := map[string]interface{}{
		"cricket": []string{"Sunny", "Manish", "Raj"},
		"tennis": []string{"Shankar", "Mayur"},
	}
	fmt.Println(map1)
	for sport, names := range map1{

		fmt.Println("Players selected for sport- " + sport)
		names2 := names.([]string)
		for _,name := range names2{
			fmt.Println(name)
		}
	} 
	
	map2 := map[string]interface{}{
		"name": []string{"Sunny", "Manish", "Raj"},
		"roll_no": []int{1,2,3},
		"sports": "cricket",
	}

	namesList := map2["name"].([]string)
	fmt.Println(namesList)
	fmt.Println(len(namesList))

	rollNoList := map2["roll_no"].([]int)
	fmt.Println(rollNoList)

	sports := map2["sports"].(string)
	fmt.Println(sports)
	fmt.Println(reflect.TypeOf(sports))


}