package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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

	// map is collection

	// slice, err := getSliceOfStringFromFile()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(slice)

	// mapExample()

	// 8. structs

	// structExample()

	// 9. defined types

	// typeExample()

	// method and function
	// methods are similar to function the only difference being it is has a receiver parameter before the funtion name
	// so the receiver specify which type will receive this method
	// myMethod()

	// myType := MyType(32)
	// myType.myMethod()

	// cannot define method for someone else's type (this includes standard library)
	// also it cannot be defined outside package

	// receiver parameter is alterantive to self (python) and this (java)
	// convention for naming for type parameter is use first letter of type name.

	// in case of receiver as well pointer concept applies

	// experimentType := ExperimentType(32)

	// experimentType.receiverPointerExample()
	// fmt.Println(experimentType)
	// experimentType.receiverPointerExampleWithPointer()
	// fmt.Println(experimentType)

	// if we there is no difference in calling the method with pointers

	// go automatically converts reviever to pointer

	// we can also implicitly

	// cannot get address of a value , we will get error.
	// cannot call method on value

	// When calling a method that takes a pointer receiver,
	// Go can automatically convert a value to a pointer to a receiver if it’s stored in a variable.
	// If it’s not, you’ll get an error.

	// 10. Encapsulation

	// d := date{year: 2021, month: 5, day: 27}

	// fmt.Println(d)

	// d.setDay(25)
	// d.setMonth(12)
	// d.setYear(2021)

	// fmt.Println(d)

	// err := d.setYear(-200)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// d2 := date{}

	// d2.setDay(2)
	// d2.year = -56
	// d2.month = -200

	// fmt.Println(d2)

	//even after having validation the behaviour is wrong
	// move the field to separate package if you dont want the fields to be accessible , so we take advantage of unexported variable with package

	// getter method  can be writtent to get the values back similar to setter written below
	// gettter method name should be similar to variable name

	// if we dont need encapsulation it is generally ok to access it directly and use it -- general convention in go

	// Unexported fields dont get promoted.

	//11 . An interface can be satisfied by more than one classes.

	// if we not implement event method defined for interface we cannot use interface type for variable
	// var myVar myInterface = interfaceAbidingType{myFloat: 23.4}
	// in case if we wamt to use pointer method then we need to assign address , go will be handle that

	// var myVar myInterface = &interfaceAbidingType{myFloat: 23.4}
	// myVar.myMethod()
	// myVar.myMethodWithArgument(23)
	// myVar.myMethodWithPointer()
	// fmt.Println(myVar.myMethodWithReturnType())
	// myVar.myMethod2WithReturnType()// cannot call a method that is not part of interface
	// fmt.Println(myVar)

	// methodThatCallsInterface(myVar)

	// rules are same for naming of interface.

	// we need to revisit lesson related to package and interface , sometime again.

	// type assertions

	// myVar2, ok := myVar.(*interfaceAbidingType) // type assertion expose second optional value which can be used to see if type assertion succeded or failed.
	// type assertion dont panic until and unless a method is called on type where it doesnt exist

	// if ok {
	// 	fmt.Println(myVar2.myMethod2WithReturnType())
	// } else {
	// 	fmt.Print(ok)
	// }

	// how to read the above
	// “I know this variable uses the interface type myInterface, but I’m pretty sure this myInterface is actually a interfaceAbidingType.”

	// error value is any value with named error and returns string // error is an interface
	// error interface is give below

	// 	type error interface {
	// 		Error() string
	//  }

	//error is predeclared identifier like int string -- its a universal block

	// we can use anything in universe block encompasses all packages blocks.
	// without importing.

	// fmt package has Stringer interface

	// type Stringer interface{
	// 		string() string
	// }

	// with type interface , you can take in any type as it is an empty interface
	// empty interface can be defined in below way

	// type EmptyInterface interface{}

	// it is recommended not use interface because for each method call we would eventually require type assertions.

	// 12. Panic

	//defer keyword , makes the code run late, as in until all the things in method are done
	// err := deferredResultExample()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//only function and method calls can be deferred

	// how to read directory io/iotuil package there is ReadDir function

	// directoryReadingExample()

	// panicExample()

	// panicDeferredExample()

	// when to call panic ??

	//simplifies code but crashes program.

	// inaccessible file , network failures and bad user -> normal -> errors should be able to help here
	// for impossible situation we use panic like some bug and not users mistake

	// recover

	// recover returns nil in context of normal executing program
	// recoverExample()

	// panic accepts interfaces and recover returns empty interfaces

	// if panic is not intended to be recovered just panic again

	// panic and recover is not be used like exception -- golang maintainers are against it.
	// panic and recover is intentionally clumsy so that people wont use them often

	//13. gocoroutine and channels
	// go coroutine allows different task to be done parallely
	// routine can coordinate work using channels.
	// channels help for passing the data and synchronize.

	// fmt.Println("Executing via uni")
	// uniTaskExample()

	// in go concurrent tasks are called coroutines

	// compared to threads : requires less memory and lesser startup and stop time . so it means we can run more number of gocoroutine
	// we just need to use go keyword for multi tasking
	// main is also a go routine
	// if we run the below program without channels nothing will be printed as in similar to case of join because all of them are executing in parallel

	// fmt.Println("Executing via multi")
	// multiTaskExample()

	// gocoroutine function cannot return value.
	// to communicate values to other function we can use channels.

	// channels are declared with keyword chan

	//below code results in deadlock
	//reason :: main coroutine is waiting for other coroutine to provide response on channel but there is no other coroutine

	// var myChannel chan int
	// myChannel = make(chan int)

	// myChannel2 := make(chan int)

	// myChannel <- 123
	// myChannel2 <- 345
	// fmt.Println(<-myChannel)
	// fmt.Println(<-myChannel2)

	returnChannel := make(chan int)

	multiTaskWithChannel(returnChannel)
	fmt.Println(<-returnChannel)
	fmt.Println(<-returnChannel)
	fmt.Println(<-returnChannel)

	// channel values are conveyed as and when they are provided above one by one
	// until then main thread is blocked

	// if a go coroutine is in sleep then the sending coroutine will block until receving coroutine gets it.

	// we can also send struct over channel
}

func multiTaskWithChannel(returnChannel chan int) {
	go getExampleWithChannel("http://example.com", returnChannel)
	go getExampleWithChannel("http://google.com", returnChannel)
	go getExampleWithChannel("http://facebook.com", returnChannel)

}

func getExampleWithChannel(url string, returnChannel chan int) {
	response, err := http.Get(url) // returns an http response
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() // to close the network connection

	body, err := ioutil.ReadAll(response.Body) // io utils are used to read the contents of response body

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body)) // slice of byte value is provided to string function
	// fmt.Println(len(body))
	returnChannel <- len(body)
}

func multiTaskExample() {
	go getExample("http://example.com")
	go getExample("http://google.com")
	go getExample("http://facebook.com")
	time.Sleep(5. * time.Second)
}

func uniTaskExample() {
	getExample("http://example.com")
	getExample("http://google.com")
	getExample("http://facebook.com")
}

func getExample(url string) {
	response, err := http.Get(url) // returns an http response
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() // to close the network connection

	body, err := ioutil.ReadAll(response.Body) // io utils are used to read the contents of response body

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body)) // slice of byte value is provided to string function
	fmt.Println(len(body))
}

func recoverExample() {
	panickingfunction()
}

func executeRecover() {
	fmt.Println(recover()) // in case of non panic it will return nil
	fmt.Println("Recovering after panic")
}

func panickingfunction() {
	defer executeRecover()
	fmt.Println("Something bad is going to happen") //this statement is not executed :D
	panic("Shit!! bad happened")

}

func panicDeferredExample() {
	pf1()
}

func pf1() {
	defer fmt.Println("Hello boy")
	pf2()
}
func pf2() {
	defer fmt.Println("What's up")
	panicExample()
}

func panicExample() {
	panic("Oh No I fucked up")
}

func directoryReadingExample() {
	readOutCompleteDirectoryContent("D:\\Hustle\\programming\\learning_go\\GoLangLearn\\src\\directoryStruct")
}
func readOutCompleteDirectoryContent(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
			readOutCompleteDirectoryContent(path + "\\" + file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}

func deferredResultExample() error {
	// here the deferred example shows that
	// with defer all the methods below it are completed.
	defer fmt.Println("Hello")
	fmt.Println("Hello2")

	defer fmt.Println("Hello3")
	defer fmt.Println("Hello4")
	if 1 == 1 {
		return fmt.Errorf("this is an error")
	}
	defer fmt.Println("Hello4")

	return nil
}

func methodThatCallsInterface(intface myInterface) {
	intface.myMethod()
}

type myInterface interface {
	// contains set of methods

	myMethod()
	myMethodWithArgument(float64)
	myMethodWithReturnType() string
	myMethodWithPointer()
}

type interfaceAbidingType struct {
	myFloat float64
}

func (i *interfaceAbidingType) myMethodWithPointer() {
	fmt.Println("call from pointer", i)
}

func (i *interfaceAbidingType) myMethod() {
	fmt.Println("Hello")
}

func (i *interfaceAbidingType) myMethodWithArgument(myVal float64) {
	fmt.Println(myVal)
}

func (i *interfaceAbidingType) myMethodWithReturnType() string {
	return "Hello"
}

func (i *interfaceAbidingType) myMethod2WithReturnType() int {
	return 23
}

func (d *date) setYear(year int) error {

	if year < 1 {
		return errors.New("invalid year provided")
	}
	d.year = year
	return nil
}

func (d *date) setMonth(month int) {
	d.month = month
}

func (d *date) setDay(day int) {
	d.day = day
}

type date struct {
	year  int
	month int
	day   int
}

func (e ExperimentType) receiverPointerExample() {
	e = 5
}

func (e *ExperimentType) receiverPointerExampleWithPointer() {
	*e = 7
}

type ExperimentType int

type MyType int

func (myType MyType) myMethod() {
	fmt.Println("Dude dude")

	// we can also print the content of the receiver param

	fmt.Println(myType)
}

type Liters float64
type Gallons float64
type MilliLiters float64

func toGallones(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

// below we cannot redeclare a block, it gives following error

// toGallones redeclared in this blockcompilerDuplicateDecl

// func toGallones(l MilliLiters) MilliLiters {
// 	return MilliLiters(l / 2)
// }

func typeExample() {
	var carFuelInIndia Liters
	var carFuelInUS Gallons

	carFuelInIndia = 24.5
	carFuelInUS = 240.50

	fmt.Println(carFuelInIndia, carFuelInUS)

	carFuelInIndia = Liters(30.9)
	carFuelInUS = Gallons(300.10)

	fmt.Println(carFuelInIndia, carFuelInUS)

	carFuelInUS = Gallons(Liters(30.9))

	fmt.Println(carFuelInUS)

	// cannot assign value from liters var to gallons var but can pass liters for construction of gallons type.

	// but 	carFuelInUS = Gallons(Liters(30.9)) is wrong because gallons are not equal to liter there should be some restriction to this

	// all logical and relational operators by default applies to the what applies to underlying type.

	fmt.Println(toGallones(carFuelInIndia), toLiters(carFuelInUS))

	//++++++++OVERLOADING OF FUNCTIONS IS NOT SUPPORTED ??!! +++++++
	// it doesnt have support because go lang people thinks it is confusing and fragile.

}

func structExample() {
	// building storage structure
	var myStruct struct {
		f1 int
		f2 bool
	}
	fmt.Println(myStruct)

	//  we print with verb for struct prints out value in literal forat

	fmt.Printf("%#v\n", myStruct)

	myStruct.f1 = 1
	myStruct.f2 = true

	fmt.Println(myStruct, myStruct.f1)

	// to create type we use the type keyword.

	type customStruct struct {
		f1 int
		f2 bool
		f3 string
	}

	var myCustomStruct customStruct
	myCustomStruct.f1 = 1
	myCustomStruct.f2 = true
	myCustomStruct.f3 = "custom struct"
	fmt.Println(myCustomStruct)

	var myDiscountMan discount
	myDiscountMan.name = "lol"
	myDiscountMan.rate = 5.0

	fmt.Println("Discount example")
	fmt.Println(myDiscountMan)
	changeRates(myDiscountMan)
	fmt.Println(myDiscountMan)

	//pass by value apples to struc as well that is why the above example doesnt work
	// because we received copy of struct

	changeActualRates(&myDiscountMan)
	fmt.Println(myDiscountMan)

	// always use pointer to pass large structure.
	// struct field names must also be capitalized so that they can be exported.

	myStruct2 := discount{name: "123", rate: 23.45}

	fmt.Println(myStruct2)

	// we may also omit the field while using assignment of struct
	myStruct2 = discount{rate: 46.45}
	fmt.Println(myStruct2)

	var nestedStruct struct {
		name string
		dis  discount
	}
	nestedStruct.name = "faltu example"
	nestedStruct.dis = myDiscountMan

	fmt.Println(nestedStruct.dis.name)

	// anonymous field struct to avoid calling nested field and directly access field

	var nustaDhur myStruct1
	nustaDhur.struct1Name = "bablloo"
	nustaDhur.sturct2Name = "gabloo"
	nustaDhur.internalStruct.sturct3Name = "wabloo"

	fmt.Println(nustaDhur)

	// the anonymous struct variables or are promoted to outer struct
}

type myStruct1 struct {
	struct1Name string
	myStruct2
	internalStruct myStruct3
}

type myStruct2 struct {
	sturct2Name string
}

type myStruct3 struct {
	sturct3Name string
}

func changeActualRates(myDiscountMan *discount) {
	myDiscountMan.rate = 6.7

	// why we did not had to use *operator while assigning value like we do in case of most of the pointer example
	// reason :: dot annotation works on both struct pointer and struct to access value. so we dont need to use * operator in this case

	//  accessing like *myDiscountMan.rate <-- here go would think that rate field in struct is actually a pointer which is not the case
}

func changeRates(myDiscountMan discount) {
	myDiscountMan.rate = 6.7
}

type discount struct {
	name string
	rate float64
}

func mapExample() {
	var countMap map[string]int
	countMap = make(map[string]int)

	slice, err := getSliceOfStringFromFile()
	if err != nil {
		log.Fatal()
	}

	for _, name := range slice {
		countMap[name] = countMap[name] + 1
	}

	for _, name := range slice {
		fmt.Println(countMap[name])
	}

	mapLiteralEx := map[string]string{"hello": "bye", "ohBoy": "ohGirl"}

	fmt.Println(mapLiteralEx)

	mapLiteralEmptyMap := map[int]bool{}
	fmt.Println(mapLiteralEmptyMap)

	mapLiteralEmptyMap[0] = true
	fmt.Println(mapLiteralEmptyMap)

	// zero value within map
	// if we try to access a key that doesnt exist the we get the zero value

	// zero value for map variable is nil

	// play with zero value
	// check if it is a assigned value or real 0 value

	mapZeroValue := make(map[string]int)

	mapZeroValue["non zero"] = 0

	mapZeroValue["value"] = 5

	value, isZeroValue := mapZeroValue["zero value"]
	fmt.Println(value, isZeroValue)

	value, isZeroValue = mapZeroValue["non zero"]
	fmt.Println(value, isZeroValue)

	fmt.Println("Demonstration for delete value")
	fmt.Println(mapZeroValue["value"])
	delete(mapZeroValue, "value")
	fmt.Println(mapZeroValue["value"])

	// while slice and list have range where we get index and value
	// map do also have support for for range
	// but we get key value in this case

	for key, value := range mapZeroValue {
		fmt.Println(key, value)
	}

	// we can ignore uwanted return variable from map for key and value while using map
	// in case if we want to ignore value, we can either not consider writting it or use _

	// map printing is unordered.

	// ordering needs to be taken care via programming
}

func getSliceOfStringFromFile() ([]string, error) {
	osArgs := os.Args
	var slice []string

	file, err := os.Open(osArgs[1])

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return slice, nil
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
