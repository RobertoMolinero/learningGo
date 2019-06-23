# Hello World

The classic 'Hello World' inspired by the legendary Dennis Ritchie.

To start this program type in the console:

```
go run main.go
```

You can also translate the program first and then run the created binary.

```
go build
./00_helloWorld 
```

The default name of the file is the name of the parent folder. To change the name of the binary use the compiler option '-o'.

```
go build -o heyHoLetsGo
./heyHoLetsGo
```

To make the application available in every terminal export it to 'GOBIN'. 

```
go build -o $GOBIN/heyHoLetsGo
```

If the variable is not yet defined correctly, it can be entered as follows.

```
export GOBIN=$GOPATH/bin
```
