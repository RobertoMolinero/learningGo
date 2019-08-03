# Test

## Unit Test

The test file for a class m.go is called m_test.go according to the convention.

The command to run all existing tests is as follows:

```
go test
```

As usual with Unix, only the errors are displayed. If there are no errors, only 'PASS' and a kind of totals line are 
output. 

If you want to see all tests, no matter if they were negative, you have to append the option '-v'.

### Construction

The package for tests in Go supports a data-driven test setup.

The naming convention for a test method is Test_<Method Name>. The parameter of the test method is an object of the 
test library and is needed later to control the test procedure. 

```
func Test_mySum(t *testing.T) {
...
}
```

The first thing to do is to build a structure that describes this test data. You can then define test cases in this 
structure.

```
	type test struct {
		name   string
		data   []int
		answer int
	}

	tests := []test{
		{"2+3=5", []int{2, 3}, 5},
		{"4+7=11", []int{4, 7}, 11},
		{"5+9=14", []int{5, 9}, 14},
	}
```

The last step is the execution of the test cases. In a loop, all test data is iterated and the input and output are 
used to test the function.

```
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySum(tt.data...); got != tt.answer {
				t.Errorf("mySum() = %v, want %v", got, tt.answer)
			}
		})
	}
```

## Coverage

To measure the test coverage there is the parameter '-cover'. The number of tests is set in relation to the number of 
non-tests.

```
go test -cover 
```

However, the result of the last command is only a single number. If you want to improve the measured coverage it is 
necessary to know which lines are not yet covered by a test.

There are the following two commands. The first one outputs the measured result to a file. The second command takes 
this file and displays the result in the standard browser. 

The result is the source code file whose code is marked with 3 different colors. The color represents the lines that 
have test coverage. The lines in the color red have no test coverage. And the lines that are displayed in gray cannot 
be tested. This is the case, for example, with import statements and function declarations.

```
go test -coverprofile c.out 
go tool cover -html=c.out
```

## Example Test

...

## Benchmark Test

... (-benchmem)

```
go test -bench .
```
