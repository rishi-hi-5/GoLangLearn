package main

import (
	"bufio"
	"errors"
	"fmt"

	// "greetings"
	// "greetings/englishbye"
	// "greetings/spanishbye"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	// "math/rand" // one you want to include something specific from package
)

// package is looked in following folders
// C:\Program Files\Go\src\greetings (from $GOROOT)
// 	C:\Users\rishi\go\src\greetings (

func main() {

	//1. START

	// the variable will contain a zero value if it is only decleared.

	// if we want to define the variable as we declare use the second line

	// var a int = 1
	// b := 2
	// fmt.Println(a, b)

	// BasicDeclaractionAndInputExample()

	// TimeAndReplacerExample()
	/*
		Function call format is as below

		object.Function()
		object <-- here is a value
		Function <-- method inside the value object.

	*/

	// 2. CONDITIONALS

	// InputExample()

	// AppendExample()
	// if something of same name exists then the later will take precendence

	// fmt := "lol" the variable has shadowed the fmt from package
	// fmt.Println("shadow") // this function will result in error as it is overshadowed

	// reader.ReadString <-- returns a character with \n in above grade example . to remove that we can use trim function
	// WhiteSpaceExample()

	// for conversion from string to number and vice versa method Parse<NumberType> can be used

	// StringConversionExample() // please always check the package name when using
	// BlockExample()

	// the scope is as follows
	// {} <- this indicates one scope also package is also a scope
	// ScopeExample()

	// package name vs import paths
	// import package/path

	// RandomNumberExample()

	// GuessGame()

	//3. FUNCTIONS

	// formatting ouput with printf and sprintf
	// FormatExample()

	//for a function the name must begin with letter

	// VERY IMP POINT ABOUT FUNCTIONS
	// functions with captial name can be exported and used outside the package

	// camel case if multiple words.

	// fmt.Println(add(3, 5))

	// about the error values

	// errorExample()

	// fmt.Println(returnsMultipleValue())
	// fmt.Println(returnsMultipleValueWithDocs())

	// if there are no errors then the value could be set to special keyword nil.
	// always check return value for error and it always needs to be handled if an error is given back

	// go is pass by value : a copy of parameter is passed . to demonstrate

	// in pass by reference , you can change the reference to point to different object.
	// in pass by value , you cannot change a reference to point to different object
	// consider following function to clear out on this.
	// var x float64 = 2.3
	// passByValueExample(x)
	// fmt.Println(x)

	//?? we will return here to present a more valid example
	// refer stack overflow ::https://stackoverflow.com/questions/40480/is-java-pass-by-reference-or-pass-by-value :: answer by scott

	//pointers
	// address of variable can be obtained by using & operator

	// pointerExample()

	// to create pointer to int
	// y := 3
	// var x *int = &y
	// fmt.Println(x)
	// fmt.Println(*x)

	// pointer variable can only hold one type of pointers

	// var x float64 = 2.3
	// fmt.Println("before calling a function", x)

	// passByValueExample(x)
	// fmt.Println("Pass by value without pointer", x)

	// x = 2.3
	// passByValueUpdateWithPointer(&x)
	// fmt.Println("Pass by value with pointer", x)

	// it is possible to return pointer from function

	// fmt.Println(returnPointerFromFunction())
	// fmt.Println(*returnPointerFromFunction())

	// 4. packages

	// go tool looks or package code in a special directory called workspace
	// workspace is a directory named go in current user directory

	// workspace directory contain 3 directory bin pkg src

	/*
		1. bin holds compiled binary executable
		2. pkg holds compiled binary pkg files
		3. src holds go source code

		each package goes in its own subdirectory
	*/

	// always set correcct workspace path with GOPATH before creating or using package structure
	/*
		Struture should be following
		GOPATH
			--> src
				--> create modules
					--> in one of the modules specify package as main
	*/
	// fmt.Println(greetings.HelloSir())

	// non capitalized functions package are not available

	// generally package name should match name of directory , but main is exception
	// main package generally denotes executable.

	// when go run command is given for a file with main packaging
	// go based on GOPATH (workspace path) tries to find the a custom package mentioned

	//package names should all be lower case and no camel casing

	// contents inside package can be named anything.

	// fmt.Println(greetings.HelloSuperSmartBoy())

	// to declare constant use const instead of var

	// const helloWorld = "hello world" // constants are declared and assinged at the same time and assignment operator doesnt exist in case of constants
	// const aValue int = 2             // changing constant can result in problem.

	// A constant with Captal letter is exporeted

	// fmt.Println(greetings.DefaultGreeting)

	// Nested package

	// fmt.Println(englishbye.ByeInEnglish())
	// fmt.Println(spanishbye.ByeInSpanish())

	// go can run without any go file once executable is created
	// go build creates go exe file but if we need to create in proper structured manner
	// then we need to install the same  with go install command the exe file is created in bin older

	// the packages can also name like a uri strucutre.

	// go get github.com/accountname/<name> <-- to download package

	// to view doc
	// go doc <specific package>
	// go doc <specific package> <specific function>

	// to add comments to the code
	// they are simple comments with //
	// for package it should start with Package followed by package name and description
	// for function it is function name and description
	// code examples can also be added with intending them

	// to run go doc server type godoc -http=:<port> . this godoc is different from go doc

	// check why godoc command is not working

	// 5. Arrays

	// array holds specific number of elements and it cannot grow or shrink
	// arrayExample()

	// File Example
	// fileExample()

	// go arrays cannot shrink and grow . to address this we can use slices

	// 6. slices
	// slices are like collection that grow to hold additional items.

	// slice and array has one difference you dont mention anything for number of elements in slice
	// sliceExample()

	// how to get command line arguments
	// slice operator can be used with other slices as well

	// osArgsExample()

	//variadic functions

	// functions that takes in variable number of arguments

	// to make a variadic function we use ...(also called elipsis)

	// variadicFunctionExample("Hi", "How", "Are", "You")
	// variadicFunctionExample("m", "fine")

	// non variadic arg if present are compulsory

	// variadicFunctionExample2("2")
	// variadicFunctionExample2("2", 3, 4)

	// math package offers Max limims
	// fmt.Println(math.Inf(-1)) // - sign denotes negative infinity
	// fmt.Println(math.Inf(1))

	// we can also pass slices to variadic function

	// var sampleSlice []string
	// sampleSlice = append(sampleSlice, "12", "13")
	// variadicFunctionExample(sampleSlice...)

	// 7. Labeling data: Maps
}

func variadicFunctionExample2(compulsory string, notCompulsory ...int) {
	fmt.Println(compulsory)
	fmt.Println(notCompulsory)
}

func variadicFunctionExample(strings ...string) { // strings variables holds a slice
	for _, str := range strings {
		fmt.Println(str)
	}
}

func osArgsExample() {

	myArray := os.Args
	fmt.Println(myArray[1:])
}

func sliceExample() {
	var mySlice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	// calling slice would not make the automatically amake the variable

	var mySlice2 []string
	mySlice2 = make([]string, 7)
	fmt.Println(mySlice2)

	mySlice3 := make([]string, 5)
	mySlice3[2] = "1"
	mySlice3[1] = "3"
	fmt.Println(mySlice3)

	// why not use slice instead of array directly ?

	//slices are built on top of array, understanding slices would be difficult if you cannot understand array used for its impl.

	//creating slice from array

	myArray := [5]int{1, 3, 4, 5, 6}

	sliceFromArray := myArray[1:3] // starting is inclusive and ending is exclusive

	fmt.Println(sliceFromArray)

	//slice operator has default for both start and stop indexes for slice
	// start slice :: 0
	// stop slice :: length of array

	// slice is microscopic focusing on array

	// changing underlying array would also result in changing the slice

	// same goes for the changes from slice to array

	// with make you avoid working on the array from which the slice is created and hence avoid confusion.

	// go offers built in function append which can help you get new slice including older elements.

	var extendSlice []string = make([]string, 2)

	extendSlice[0] = "hello"
	extendSlice[1] = "bye"

	extendSlice = append(extendSlice, "oh", "	hi", "new", "slice")

	fmt.Println(extendSlice, len(extendSlice))

	//always assign the value from append to same slice which was extended to avoid inconsistent behaviour

	// internally to slice an array is managed and if the array doesnt have sufficent space for append a new array will be created
	// only store one slice when appending.

	// slice has nil as there zero value instead of being specific to the type of slice as compared to array.

	// len function returns 0 if nil is passed.

	var sampleSlice []string //nil slice

	fmt.Println(sampleSlice, len(sampleSlice))
	sampleSlice = append(sampleSlice, "oh", "what???")

	fmt.Println(sampleSlice)
}

func fileExample() {
	file, err := os.Open("D:\\Hustle\\programming\\learning_go\\GoLangLearn\\sampleTextFile.txt")

	if err != nil {
		log.Fatal("Error while reading file", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func arrayExample() {
	var animes [3]string
	animes[0] = "Jojo"
	animes[1] = "Naruto"
	animes[2] = "Initial D"

	// default values are assigned in absence of value // also called zero values
	fmt.Println(animes[1])

	// if we already know what will contain in array during declaration then the same can be initialized
	// with array literal

	var intArray [3]int = [3]int{1, 3, 4}

	fmt.Println(intArray[2])

	// using assignment operator

	floatArray := [3]float64{0.2, 0.3, 0.4}

	fmt.Println(floatArray[0])

	fmt.Println(animes, intArray, floatArray)

	fmt.Printf("%#v\n", floatArray)

	// loop through elements

	for i := 0; i < len(animes); i++ {
		fmt.Println(animes[i])
	}

	for index, value := range intArray {
		fmt.Println(index, value)
	}

	for _, value := range intArray {
		fmt.Println(value)
	}
}

func returnPointerFromFunction() *int {
	x := 3
	return &x
}

func passByValueUpdateWithPointer(x *float64) {
	*x = 5.2
	// fmt.Println(*x)
}

func pointerExample() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
}

func passByValueExample(x float64) {
	x = 5.2
	// fmt.Println(x)
}

func returnsMultipleValue() (int, bool, float64) { // paraenthesis are required for multiple return value
	return 1, true, 0.2
}

func returnsMultipleValueWithDocs() (sampleInt int, sampleBool bool, sampleFloat float64) { // name can be assigned for easy documentation of what each value signifies in return
	return 1, true, 0.2
}

func errorExample() {
	err := errors.New("New error")
	fmt.Println(err.Error())
	log.Fatal(err)

	// if we need a formatted error below function can be used to do that

	err = fmt.Errorf("%d New error Encountered", 2)
	fmt.Println(err.Error())
}

func add(a int, b int) int {
	return a + b
}

func FormatExample() {
	fmt.Println("decimal value with high precision", 1.0/3.0)
	fmt.Printf("decimal value with fixed precision %0.2f\n", 1.0/3.0)                     // print with format
	formattedString := fmt.Sprintf("decimal value with fixed precision %0.1f\n", 1.0/3.0) //%precisionf is a verb

	/*
		for string it is %s
		for integer it is %d
		for boolean  it is %t
		for any value %v
		for any value formatted as it appears in go code %#v
		for type $T
		for literal % %%
	*/
	fmt.Println(formattedString)

	fmt.Printf("Check this \n %s %d %4d %t %v %#v %T %% \n", "hello", 1, 1, true, "\t", "\t", "Type")

	fmt.Printf("%5.2f\n", 1.0/3.0)     //5 denotes minimum witdh of decimal number ,including decimal point
	fmt.Printf("%1.2f\n", 254.3333333) //if the width is invalid it will just consider precision part and ignore padding.

}

func RandomNumberExample() {
	randomNum := rand.Intn(100) + 1
	fmt.Println("Not a random number as seed is not provided , fixed value", randomNum) // as we have not given a seed , it will generate same number again and again

	seconds := time.Now().Unix() // we get unix time in int which represents number of seconds since january 1 , 1970
	rand.Seed(seconds)
	randomNum = rand.Intn(100) + 1
	println("Random number with help of seed", randomNum)

	// to generate a true random number we need to provide a see on the basis of which it generates random number
	// this can be done by using time package
}

func GuessGame() {
	//generate a random number
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	randomNum := rand.Intn(10) + 1

	turns := 5
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < turns; i++ {
		fmt.Println("Guess the number in game")

		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		guess, _ := strconv.Atoi(str)

		if randomNum > guess {
			fmt.Println("Oops your guess was LOW.")
		} else if randomNum < guess {
			fmt.Println("Oops your guess was HIGH.")
		} else {
			fmt.Println("Good job!, You guessed it")
			return
		}
	}
	fmt.Println("Sorry, You didn't guess my number. It was:[", randomNum, "]")
}

func ScopeExample() {

	scope := "String"
	// scope := "String 2" // this would not work as in the same scope there could be one variable with that name
	fmt.Println(scope)
	scope2, scope := "hello", "hi" // here this is a strange thing, here if we have one variable for declaration then the already declared variable will just get new value
	// unlike the above commented line which errors out
	println("scope 2", scope2, "scope ", scope)
	{
		{
			scope = "see we are not using assignmen here"
			{
				scope3 := "scope 3 var wont be accessible below"
				fmt.Println(scope3)
				fmt.Println(scope)
			}
			scope3 := "scope 3 var wont be accessible from above scope so reusing same var name is not problem"
			fmt.Println(scope3)
			fmt.Println(scope)
		}
		scope3 := "scope again reused"
		fmt.Println(scope3)
		fmt.Println(scope)
	}
	scope3 := "what?? just kidding"

	println(scope3)
}

func BlockExample() {

	// also maybe refered as scoping
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a float value")
	input, _ := reader.ReadString('\n') // dont use down quote, use single quote for character
	number, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	var result string
	if number <= 60 {
		status := "pass" // status is scoped to a block
		result = status
	} else {
		status := "fail"
		result = status
	}

	// fmt.Println("Status is ",status) //doesnt work it is not here the status
	fmt.Println("Status is", result) //works
}

func StringConversionExample() {
	var str string = "23"
	var float, _ = strconv.ParseFloat(str, 64)

	fmt.Println(float * 2)
}

func BasicDeclaractionAndInputExample() {
	fmt.Println("Hello", "Beasty")
	fmt.Println(strings.Title("head First go"))
	fmt.Println(math.Ceil(5.4))

	fmt.Println('A', 'B')
	fmt.Println(123)

	// errors printed are one for each function
	// go is statically typed.

	// check types

	fmt.Println(reflect.TypeOf(123))
	fmt.Println(reflect.TypeOf("String"))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(3.145))

	//variables
	var x, y, z int
	x = 1
	y = 2
	z = 3

	fmt.Println(x, y, z)
	x, y, z = 4, 5, 6
	fmt.Println(x, y, z)

}

func TimeAndReplacerExample() {
	//time package

	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println("The year is", year)

	replacestr := "I am the best"
	replacer := strings.NewReplacer("I am", "Rishi")
	replaced := replacer.Replace(replacestr)

	fmt.Println(replaced)

}

func InputExample() {
	//take input

	fmt.Println("Enter the grade:")
	reader := bufio.NewReader(os.Stdin)   // where does it get input from
	input, err := reader.ReadString('\n') // return everything until the mentioned character is encountered
	if err != nil {
		log.Fatal("Error ocured", err) // here if log .fatal is used it will terminate the program
	}
	fmt.Println(input, err)

	secondInput, _ := reader.ReadString('\n') // return everything until the mentioned character is encountered
	// _ as a blank to ignore the variable that no one wants to use

	fmt.Println(secondInput)
	// if input < 60 {
	// 	fmt.Println("Poor Performance")
	// } else if input > 100 {
	// 	fmt.Println("Good Performance")
	// } else {
	// 	fmt.Println("Extra ordinary")
	// }

}

func AppendExample() {
	// usage of append function

	var lang []string = append([]string{}, "Hello", "Bye", "What ??!!") // [] represents an array
	// append function appends the element the end of provided list

	lang = append(lang, "heard of bhai lang")

	fmt.Println(lang)

}

func WhiteSpaceExample() {
	wow := "\t the whitespace at start end will be gone    \n"
	fmt.Println(wow)
	wow = strings.TrimSpace(wow)
	fmt.Println("Printing again")
	fmt.Println(wow)

}
