# Evans Analysis

## Installation

The project page with all files and instructions can be found here:

https://github.com/ktr0731/evans

## Start

```
$ evans -p 50051 -r
  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client

math.MathService@127.0.0.1:50051>
```

## Usage

To search and navigate at package level:

```
show package
package math
```

To search and navigate at service level:

```
show service
service MathService
```

### Call Unary Function

```
math.MathService@127.0.0.1:50051> call Sum
s::first (TYPE_INT32) => 3
s::second (TYPE_INT32) => 5
{
  "result": 8
}
```

### Call Server/Client Streaming Function

Abort the streaming side: CTRL+D

```
math.MathService@127.0.0.1:50051> call ComputeAverage
number (TYPE_INT32) => 1
number (TYPE_INT32) => 4
number (TYPE_INT32) => 8
number (TYPE_INT32) => 54
number (TYPE_INT32) => 22
number (TYPE_INT32) => 74
number (TYPE_INT32) => 753
number (TYPE_INT32) => 
{
  "average": 130.85714285714286
}
```

### Call Bidirectional Function

...

### Call Unary Function with Error Handling

It is also possible to receive and check error codes from the server.

```
math.MathService@127.0.0.1:50051> call SquareRoot
number (TYPE_INT32) => 33
{
  "root": 5.744562646538029
}

math.MathService@127.0.0.1:50051> call SquareRoot
number (TYPE_INT32) => -42
command call: failed to send a request: rpc error: code = InvalidArgument desc = Received a negative number: -42
```