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
