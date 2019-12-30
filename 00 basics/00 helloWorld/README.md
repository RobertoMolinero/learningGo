# Hello World

The classic 'Hello World' inspired by the legendary Dennis Ritchie.

To start this program type in the console:

```
go run heyHoLetsGo.go
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

To be able to use the program system-wide in every terminal it must be installed. For this purpose, there is the command 'install'.

```
go install
```

This command builds the application and moves the result to the $GOBIN folder. The same result can also be achieved by manual copying using the compiler option '-o'. This procedure also gives you the possibility to name the result of the build directly. 

```
go build -o $GOBIN/heyHoLetsGo
```

For problems with the path variables please refer to the installation instructions in the parent path.