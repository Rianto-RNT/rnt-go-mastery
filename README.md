<h1 align="center">
  <br>
  Go (Golang) Mastery. 
  <br>
</h1>

<h4 align="center"> Learn golang from beginner to advenced <a href="https://go.dev/" target="_blank">Go</a>.</h4>

 <p align="center">
 <a href="#deployed-version">Demo</a> •
  <a href="#api-usage">API Usage</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#development-process">Development Process</a> •
  <a href="#installation">Installation</a> •
  <a href="#build-with">Build With</a> •
  <a href="#npm-packages">NPM Packages</a> •
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

### number) rnt-example-point

## Articles

- [The go playground](https://go.dev/play/)
- [The go packages](https://pkg.go.dev/)
- [Go time packages](https://pkg.go.dev/time)
- [Download Go](https://go.dev/dl/)
- [How to Install Go on Windows](https://golangdocs.com/install-go-windows)
- [Inside the Go Playground](https://go.dev/blog/playground)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Golang fmt package](https://pkg.go.dev/fmt)
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
