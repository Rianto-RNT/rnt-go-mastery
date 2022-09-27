<h1 align="center">
  <br>
  Go (Golang) Mastery. 
  <br>
</h1>

<h4 align="center"> Learn golang from beginner to advanced <a href="https://go.dev/" target="_blank">Go</a>.</h4>

 <p align="center">
 <a href="#deployed-version">Demo</a> •
  <a href="#api-usage">API Usage</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#development-process">Development Process</a> •
  <a href="#installation">Installation</a> •
  <a href="#build-with">Build With</a> •
  <a href="#packages">Packages</a> •
  <a href="#demonstration">Demonstration</a> •
  <a href="#future-updates">Future Updates</a> • 
  <a href="#known-bugs">Known Bugs</a> • 
  <a href="#acknowledgement">Acknowledgement</a>
</p>

## Deployed Version

Live demo (Feel free to visit) 👉 :

- [rnt-example-point](https://github.com)

## API Usage

Check:

[rnt-example-point](https://github.com)

for more info.

## Deployment

The website is deployed with build into Heroku. Below are the steps taken:

1. rnt-example-point

## Installation

You can fork the app or you can git-clone the app into your local machine. Once done that, please install all the dependencies by running

```sh
$ rnt-example-command
```

## Run App in Development

```sh
$ rnt-example-command
```

## Build With

- [Go](https://go.dev/) - Go is an open source programming language supported by Google

## Packages

- [rnt-example-point](https://github.com/)

## Development Process

### 1) Getting Started with Go

- GOPATH and Code Organization

  - Go requires you to organize your code in a speciﬁc way.
  - Go programmers typically keep all their Go code in a single workspace.
  - A workspace is nothing more than a directory in your ﬁle system whose path is stored in
    the environment variable called GOPATH.
  - On Windows , it's in %USERPROFILE%\go (For Example: c:\Users\rianto\go)
  - On Mac & Linux, it's in: ~/go (For Example: /home/rianto/go)
  - You can print out the values of the environment variables by running the go env command.

- Go Application Structure

```sh
$ go run main.go
```

```sh
/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import "fmt"

// package scoped variables and constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

    // Local Scoped Variables and Constants Declarations, calling functions etc
    var a int = 7
    var b float64 = 3.5
    const c int = 10

    // Println() function prints out a line to stdout
    // It belongs to package fmt
    fmt.Println("Hello Go world!") // => Hello Go world!
    fmt.Println(a, b, c)           // => 7 3.5 10

}
```

- Compiling (go build) and Running Go application (go run)

  go is a tool for managing Go source code.

  I. go run: it compiles and runs the application. It doesn’t produce an executable

  - go run ﬁle.go compiles and immediately runs Go programs.

  II. go build: it just compiles the application. It produces an executable

  - go build ﬁle.go compiles a bunch of Go source ﬁles. It compiles packages and dependencies.
  - If you run go build it will compile the ﬁles in the current directory and will produce an executable ﬁle with the name of the current working directory.

  - go build -o app will produce an executable ﬁle called app

```sh
$ go run  main.go

or

$ go run -x main.go
```

```sh
// build

$ go build

or

$ go build -o <application-name.exe>
```

```sh
// on windows OS

$ GOOS=windows GOARCH=amd64 go build -o winapp.exe
```

```sh
// install

$ go install

or

$ go install <rnt-path-directory-application>
```

- Formating Go Source Code (gofmt)

  Go strongly suggests certain styles.

  - gofmt which comes from golang formatter will format a program's source code in an idiomatic way that is easy to read and understand.
  - This tool, gofmt, is built into the language runtime, and it formats Go code according to a set of stable, well-understood language rules.
  - We can run it manually at the command line or we can conﬁgure our IDE, for example VSCode, to run gofmt each time a ﬁle is saved.
    Example: gofmt -l -w main.go

```sh
$ gofmt -w <rnt-file-name>

// or

$ gofmt -w -l <rnt-directory-name>
```

- [![](rnt-example-imageUrl)](https://github.com/Rianto-RNT)

### 2) Go Basic

- Variables in Go
- Multiple Declarations
- Types and zero values
- Comments
- Naming convensions

  - By convention, packages are given lower case, single-word names;
  - Go doesn't provide automatic support for getters and setters. If you have a ﬁeld in a struct called owner, the getter method should be called Owner (upper case, exported), not GetOwner.

  A setter function, if needed, will likely be called SetOwner.

```sh
owner := obj.Owner()

if owner != user {
obj.SetOwner(user)
}
```

- By convention, one-method interfaces are named by the method name plus an -er suﬃx: Reader, Writer, Formatter, etc.
- Golang fmt package
- Constant in Go
  - In Golang, we use the term constant to represent ﬁxed (unchanging) values.
  - We use constants to avoid possible errors (variables that change when they shouldn’t) or to replace a value only in one place instead of in many places
  - All basic literals (1, 3.4, “hello”, true) are in fact unnamed constants.
  - A constant belongs to compile time and it’s created at compile time. It’s value can not be changed while the program is running.
  - Another advantage of using constants is that Go can not detect runtime errors at compile-time but constants belong to compile time so errors can be detect
- Constant rules
- Constant Expressions, Typed vs Untyped Constants
- IOTA
- Data types

  - A type determines a set of values together with operations and methods speciﬁc to those values.
  - There are predeclared types, introduced types with type declarations and composite types: array, slice, map, struct, pointer, function, interface, and channel types.

    #### Predeclared, Built-in Types

  - Numeric types
    - int8, int16, int32, int64
    - uint8, uint16, uint32, uint64: used to represent unsigned (positive) integers.
    - uint is an alias for uint32 or uint64 based on platform.
    - int is an alias for int32 or int64 based on platform.
    - ﬂoat32, ﬂoat64: zero before the decimal point separator can be omitted ( -.5 -3. -0. 1.4).
    - complex64, complex128.
    - byte (alias for uint8).
    - rune (alias for int32)
  - Bool type
    - pre-deﬁned constants true and false.
  - String type
    - Unicode chars written enclosed by double-quotes.
    - A string value is a (possibly empty) sequence of bytes.
  - Array and Slice Type
    - An array is a numbered sequence of elements of a single type, called the element type.
    - An array has a ﬁxed length (we specify how many items are in the array when we declare it), but a a slice has a dynamic length (it can shrink or grow).
  - Map Type
    - A map is an unordered group of elements of one type, indexed by a set of unique keys of another type.
    - A map in Go is similar to dictionary in Python
  - Struct Type (User deﬁned type)
    - A struct is a sequence of named elements, called ﬁelds, each of which has a name and a type.
    - a structure can be compared to class concept in Object Oriented Programming.

```sh
type Car struct {
brand string
price int
}
```

- Struct Type (User deﬁned type)
  - A struct is a sequence of named elements, called ﬁelds, each of which has a name and a type.
  - a structure can be compared to class concept in Object Oriented Programming.

```sh
type Car struct {
    brand string
    price int
}
```

- Operations on Types Arithmetic and Assignment Operators
  - An operator is a symbol of the programming language which is able to operate on values.
  - In Go language, operators can be categorized based upon their different functionality in these categories:
    - Arithmetic and Bitwise Operators: +, -, \*, /, %, &, |, ^, <<, >>
    - Assignment Operators: +=, -=, \*=, /=, %=
    - Increment and Decrement Statements: ++, -- Comparison
      Operators:
      ==,
      !=, <, >, <=, >=
    - Logical
      Operators:
      &&,
      ||
      , !
    - Operators for Pointers (&) and Channels (<-)
  - Arithmetic operators apply to numeric values and are used to perform common mathematical operations.
  - There are the following arithmetic operators:

```sh
+ (addition)
- (subtraction)
* (multiplication)
/ (division)
% (modulus or simply mod)
```

- Assignment Operators which are used to assign values to variables.
- There are the following assignment operators:

```sh
= (simple assignment)
+= (increment assignment)
-= (decrement assignment)
*= (multiplication assignment)
/= (division assignment)
%= (modulus assignment)
```

- The "++" and "--" statements increment or decrement their operands by the untyped constant 1.
- Comparisson and Logical Operators

  - Comparison operators compare two operands and yield an boolean value.
  - There are the following comparison operators:

```sh
    == (equal)
    != (not equal)
    < (less)
    <= (less or equal)
    > (greater)
    >= (greater or equal)
```

- Logical operators apply to boolean values and yield a result of the same type as the operands (bool).
- There are 3 logical operators in Go:

```sh
 && (conditional and)
 || (conditional or)
 ! (not or logical negation)
```

- Overflows
- Converting Numeric Types
- Converting numbers to strings and strings to number
- Define (named) types
  - A deﬁned type also called a named type is a new type created by the programmer from another existing type which is called the underlying or source type.
  - A new deﬁned type must have a new name and can have its new methods.
  - The underlying type provides the representation, operations and size of the newly deﬁned type.
  - Even though the deﬁned type and the source type share the same representation, operations and size they are different types. A new type it’s not just an alias for an existing type, it’s a completely new type.
  - If we want to perform operations between source and deﬁned types we must convert one type into the other type. A type can be converted to another type if they share the same underlying type.
  - There is no type-hierarchy in Go We can attach methods to newly deﬁned types.
  - Type safety. We must convert one type into another to perform operations with them.
  - Readability. When we deﬁned a new type let’s say type usd ﬂoat64 we know that new type represents the US Dollar, not only ﬂoats.
- Alias Declarations

### 03) Program Control Flow

- If Else Statements
- Command line arguments (os.Args)
- Simple if statements
- For loops
- Where is the while loops in go?
- For and continue statements
- For and break statements
- Label Statements
  - Labels are used in break, continue, and goto statements.
  - It is illegal to deﬁne a label that is never used.
  - In contrast to other identiﬁers, labels are not block scoped and do not conﬂict with identiﬁers that are not labels. They live in another space.
  - The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.
  - Most of the time labels are used to terminate outer enclosing loops
  - Goto
  - Switch statements
  - Scope in Go
  - Scope means visibility.
  - The scope or the lifetime of a variable is the interval of time during which it exists as the program executes.
  - A name cannot be declared again in the same scope (for example a function in the package scope), but it can be declared in another scope.
  - In Go there are 3 scope:
    - File Scope
    - Package Scope
    - Block (Local) Scope

```sh
package main

import "fmt"  //file scope

const done = false  //package scope

func main() {
      x := 10   //local (block) scope
    fmt.Println(x)
}
```

### 04) Arrays in Go

- Intro to arrays
- An array is a composite, indexable type that stores a collection of elements of same type.
- An array has a ﬁxed length.
- Every element in an array (or slice) must be of same type.
- Go stores the elements of the array in contiguous memory locations and this way it’s very eﬃcient.
- The length and the elements type determine the type of an array. The length belongs to array type and it’s determined at compile time.

```sh
accounts := [3]int{50, 60, 70}

// The array called accounts that consists of 3 integers has it’s type [3]int
```

- Declaring Arrays
- Array Operations
- Arrays with Keyed Elements

### 05) Slices in Go

- Slice - introduction
  - Array
    - Has a ﬁxed length deﬁned at compile time;
    - The length of an array is part of its type, deﬁned at compile time and cannot be changed;
    - By default an uninitialized array has all elements equal to zero;
  - Slices
    - Has a dynamic length (it can shrink or grow);
    - The length of a slice is not part of its type and it belongs to runtime;
    - An uninitialized slice is equal to nil (its zero value is nil).
  - Both a slice and an array can contain only the same type of elements;
  - We can create a keyed slice like a keyed array;
- Declaring Slices and basic Slices Operations
- Comparing slices
- Appending to slice. Copying slices
- Slice expressions
- Slice internals_Backing Array and slice Header
  - Slice backing (underlying) Array
    - When creating a slice, behind the scenes Go creates a hidden array called Backing Array
    - The Backing Array store the elements, not the slice
    - Go implements a slice as a data structure called slice header
    - Slice Header contains 3 fields:
      - The Address of the backing array (pointer)
      - The length of the slice. len() returns it.
      - the capacity of the slice. cap() returns it
    - Slice header is he runtime representation of a slice
    - a nil slice doesn't a have backing array
- Append, length and, capacity in-depth

### 06) String, Runes, Bytes and UTF-8 in Go

- Intro to string
- Runes, bytes and unicode code points
  - Go has two additional integer types called byte and rune that are aliases for uint8 and int32 data types. In Go, the byte and rune data types are used to distinguish characters from integer values.
  - Golang doesn’t have a char data type. It uses byte and rune to represent character values.
  - Characters or rune literals are expressed in Go by enclosing them in single quotes, as in 'x' or '\n' . Rune literals such as ‘a’ , ‘b’, ‘c’, ‘x’ or ‘\n’ are represented using Unicode Code Points. A code point is a numeric value that represents a rune literal.
  - The character encoding scheme ASCII which is a Unicode subset, comprises 128 code points.
  - A string is a series of bytes values. A string is a slice of bytes and any byte slice can be encoded in a string value.
  - The Go terminology for code points is runes . A rune represent a single unicode character. Rune 0x61 in hexadecimal represents the rune literal ‘a’.
- Runes and Strings. Decoding strings byte by byte and rune by rune
- String length in bytes and runes
- Slicing Strings
- Strings Package - Contains, ContainsAny, Count, ToLower, ToUpper, EqualFold
- ## Intro to Maps

### 07) Maps in Go

- Intro to Maps
  - A Map is a collection type just like an array or a slice and stores key:value pairs.
  - The main advantage of maps is that add, get and delete operations take constant expected time.
  - All the keys and the values in a Map are statically typed and must have the same type.
  - The keys in a map must be unique, but the values don’t have to be unique.
  - A Map allows us to quickly access a value using a unique key!
  - We can use any comparable type as a key map. A comparable type is that type that supports the comparing operator which is the double equals sign.
  - Even if it’s possible, it’s not recommended to use a ﬂoat as a key. A ﬂoat has some comparable issues.
  - We can not compare a map to another map. We can only compare a map to nil.
  - Maps are unordered data structures in Go.
- Declaring Maps, working with Maps
- Comparing Maps
- Map header. cloning maps

### 08) Working with files in Go

- Open, close, rename, move, and remove files
- Operation on file - Coding Program
- Writing bytes to file - os.Write and ioutil.WriteFile
- Writing to files using a buffered writer(bufio package)
- Reading n bytes from a file. reading a file using a buffered reader
- Reading a file line by line using a scanner
- Scanning for user input. reading from Stdin

### 09) Structs in Go

- Organizing Data with Structs
  - A struct is a sequence of named elements, called ﬁelds. Each of them has a name and a type.
  - If you are familiar with OOP from other languages you can think of a struct as of a class. The struct ﬁelds are like the instance attributes we deﬁne in OOP.
  - Unlike traditional Object-Oriented Programming, Go does not have a class-object architecture. Rather we have structs which hold complex data structures.
  - A structs is nothing more that a schema containing a blueprint of data a structure will hold. This blueprint is ﬁxed at compile time. It’s not allowed to change the name or the type of the ﬁelds at runtime. You can’t add or remove ﬁelds from a struct at runtime.
- Creating struct
- Retriving an upating struct fields
- Anonymous struct and anonymous struct fields
- Embedded structs

### 10) functions in Go

- Intro to functions
  - A function is a small piece of code that is dedicated to a perform a particular task based on some input values.
  - Go recommends writing function names in simple word or camelCase.
  - Within the same package function names must be unique!
  - One of Go's features is that functions and methods can return multiple values.
  - Go doesn’t support function overloading.
  - main() and int() are predeﬁned function names.
  - Syntax:

```sh
func (receiver) name(parameters) (returns) {
//code -> function body here
}
```

- Function parameters, arguments and return values
- Variadic functions
- Defer statement
- Anonymous function

### 11) Pointers in Go

- Intro to Pointer
- Computer Memory and Pointers
  - The Computer Memory (RAM) can be thought of as a sequence of boxes or cells, placed one after another in a line. Each cell is labeled with a unique number, which increments sequentially; this number is the address of the cell or its memory location.
- [![](https://github.com/Rianto-RNT/rnt-go-mastery/blob/development/assets/go-pointers-image.png)](https://github.com/Rianto-RNT)
  - Each cell holds a single value. Everything the CPU does is fetching and storing values into memory cells.
  - What is a variable?
    - A variable is just a convenient, alphanumeric nickname or label for a memory location.
    - When we write var a int = 5 we create a label called a for a memory location where the value 5 of type int will be stored.
    - In a nutshell: memory is just a series of numbered cells, and variables are just nicknames for memory locations assigned by the compiler.
  - What is a pointer?
    - A pointer is a variable that stores the memory address of another variable.
    - The pointer points to memory address of a variable, just as a variable represents the memory address of a value.
    - A pointer value is the address of a variable or nil if it hasn’t been initialized yet.
- [![](https://github.com/Rianto-RNT/rnt-go-mastery/blob/development/assets/go-pointers-image-2.png)](https://github.com/Rianto-RNT)
  - In the above image, variable b has value 201 and is stored at memory address 0xAA01BB10. The variable a holds the address of b. Now a is said to point to b or a is a pointer to b.
  - Pointers have the power to mutate or change the data they are pointing to.
  - Unlike C, Go has no pointer arithmetic.
- Declaring Pointers. Address of and dereferencing operators
- Pointer to pointer. comparing pointers
- Passing and returning pointers to functions fom functions
- Passing pointers to functions. Passing by value vs. passing by pointer

### 12) Methods and Interfaces in Go (OOP)

- Reciever functions(Methods)
- Methods with a pointer reciever
- intro to Interfaces
- Implementing Interfaces
- Interface dynamic type and polymorphism
- Embedded interfaces
- Empty interface

### 13) Concurrency in Go

- Concurrency vs Parallelism
  - Concurrency means loading more goroutines at a time. If one goroutine blocks, another one is picked up and started. On single core CPU you can run ONLY concurrent applications but they are not run in parallel.
  - Parallelism means multiple goroutines executed at the same time. It requires multiple CPUs.
  - Concurrency means independently executing processes or dealing with multiple things at once, while parallelism is the simultaneous execution of processes and require multiple core CPUs.
- Goroutines
  - A goroutine is a lightweight thread of execution; goroutines are key ingredients to achieve concurrency in Go.
  - A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by a function invocation;
  - Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes a ﬁxed size of 1-2Mb.
  - An OS Thread Stack is ﬁxed size but a goroutine stack size shrinks and grows as needed.
  - Scheduling a goroutine is much cheaper than scheduling a thread.
  - OS threads are scheduled by the OS kernel, but goroutines are scheduled by its own Go Scheduler using a technique called m:n scheduling, because it multiplexes (or schedules) m goroutines on n OS threads.
  - Goroutines have no identity. There is no notion of identity that is accessible to the programmer.
- Spawning goroutines. the go keyword
- WaitGroups
- Project URL Checker and Page Downloader
- Project Refactoring using WaitGroups (Url Checkerand Page Downloader)
- Data race
- Go Race detecor

```sh
$ go run -race main.go
```

- Mutexes
- Intro to Channels

## Articles

- [The go playground](https://go.dev/play/)
- [The go packages](https://pkg.go.dev/)
- [Go time packages](https://pkg.go.dev/time)
- [Download Go](https://go.dev/dl/)
- [How to Install Go on Windows](https://golangdocs.com/install-go-windows)
- [Inside the Go Playground](https://go.dev/blog/playground)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Golang fmt package](https://pkg.go.dev/fmt)
- [IOTA](https://go.dev/ref/spec#Iota)
- [Golang math big package](https://pkg.go.dev/math/big)
- [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)
- [Packages fmt](https://pkg.go.dev/fmt#Printf)
- [Go constants by Rob Pike](https://go.dev/blog/constants)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)
- [Rapid table (ASCII)](https://www.rapidtables.com/convert/number/ascii-hex-bin-dec-converter.html)
- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [Go maps in action](https://go.dev/blog/maps)
- [There is no pass-by-reference in Go](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
- [Go release Hhstory](https://go.dev/doc/devel/release)
- [rnt-example-point](https://github.com/)

## Future Updates

- rnt-example-point

## Known Bugs

Feel free to email me at rianto.rnt@gmail.com if you run into any issues or have questions, ideas or concerns.
Please enjoy and feel free to share your opinion, constructive criticism, or comments about my work. Thank you! 🙂

## Route Structure

1. rnt-example-point
   > /

# Acknowledgement

- This project is part of the online course I've taken at Udemy. Thanks to Andrei Dumitrescu - Crystal Mind Academy! for creating this awesome course! Link to the course: [Master Go (Golang) Programming:The Complete Go Bootcamp 2022](https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp/)
