# Basics

The first part deals with the basic components of the language Go.

Especially the first examples are often kept so small that the code itself is already the best documentation.

First I take a look at the unique selling points of the language.

Afterwards I am interested in the testability and documentability. This aspect is covered in detail in the last two parts.

# Installation

Download latest version of Go.

```
https://golang.org/dl/
```
 
Extract the file.

```
tar -C /usr/local -xzf <FILENAME>
```

Extend the path variable 'PATH' to the bin folder of the Go installation.

```
export PATH=$PATH:/usr/local/go/bin
```

Test the installation and check the currently installed version.

```
roberto@denkBrett:~$ go version
go version go1.13.5 linux/amd64
```

Set the variable GOPATH to the folder which is used as repository for included external libraries.

```
export GOPATH=$HOME/go
```

Set the variable GOBIN to the folder where your self-written programs should be installed. From there they can be used by any terminal.

```
export GOBIN=$HOME/go/bin
```

To make the changes permanent, the changes must be entered into the profile file of the used shell. Name and location of this file differ depending on the shell and Linux variant.

```
# set PATH so it includes go version go1.13.5 linux/amd64
export PATH=$PATH:/usr/local/go/bin

# set GOPATH for external packages
export GOPATH=$HOME/go

# set GOBIN for self-written, system-wide programs
export GOBIN=$HOME/go/bin
```