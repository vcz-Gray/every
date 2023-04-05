# every

![GitHub go.mod Go version](https://img.shields.io/badge/Go-1.20-brightgreen)

It will save your time such as implement every functions

## Installation
To install `every` package, you need to install Go and set your Go workspace first.
```sh
$ go get -u github.com/quepark/every
```

## Contents

- [Contents](#contents)
  - [Ints](#Ints)
  - [Floats](#Floats)
  - [Strings](#Strings)
  - [Bytes](#Bytes)
  - [Complexes](#Complexes)
  - [Any](#Any)
    
## Description

#### Ints
```go
func Ints([]int, func(int) bool) bool
```

#### Floats
```go
func Floats([]float64, func(float64) bool) bool
```

#### Strings
```go
func Strings([]float64, func(float64) bool) bool
```

#### Bytes
```go
func Bytes([]byte, func(byte) bool) bool
```

#### Bools
```go
func Bools([]bool, func(bool) bool) bool
```

#### Complexes
```go
func Complexes([]complex128, func(complex128) bool) bool
```

#### Any
```go
func Any(interface{}, func(interface{}) bool) bool
```