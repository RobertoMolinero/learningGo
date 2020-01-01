# Go Documentation

## Retrieving the documentation of standard libraries with 'go doc'

The documentation of a function from a package can be retrieved in the terminal with 'go doc PACKAGE FUNCTION'.

Using the example of the function 'Println' from the package 'fmt': 

```
$ go doc fmt Println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```

## HTTP

Locally, all comments can also be viewed as HTML pages in the browser. For this purpose 'godoc' starts an HTTP server and generates the corresponding pages.

```
godoc -http :6060
```

## HTTP with external sources

If you look at the page closely you will notice that not only the standard libraries are listed but also all 
external sources that you have downloaded sooner or later with 'go get'.

Suppose you want to use the driver go-sql-driver.

```
$ go get -u github.com/go-sql-driver/mysql
```

Then the code with all the comments it contains ends up in the local repository.

```
$GOPATH/src/github.com/go-sql-driver/mysql
```

And all comments within this repository are now also considered on the 'godoc' page.

```
...
* fmt               Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
* github.com	
    * go-sql-driver	
        * mysql     Package mysql provides a MySQL driver for Go's database/sql package.
...
```

## Own documentation

### Maintaining own comments

Every custom package as well as every externally visible function should have a meaningful comment.

A simple example of a package:

```
// Package myMath provides ACME inc math solutions.
package myMath
```

A simple example of an externally visible function:

```
// Sum adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
```

### Generation of own comments

If you want to see the comments of your own packages and functions on the local page of 'godoc' you have to put these packages in the folder '$GOPATH/src/' as well.

One way is to check in your code to github and get it into your repository with 'go get'.

You can also simply copy the folder with your code to try it out.

For this simple test case:

```
$ cp -r myMath/ $GOPATH/src
$ godoc -http=:6060
```
