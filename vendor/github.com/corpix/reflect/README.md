reflect
---------

[![Build Status](https://travis-ci.org/corpix/reflect.svg?branch=master)](https://travis-ci.org/corpix/reflect)

Just a wrapper around `reflect`.

## Type conversions

It could do simple type conversions for you:

``` console
$ go run ./examples/type-conversion/type-conversion.go
Convert []int to []string
([]string) (len=10 cap=10) {
 (string) (len=1) "1",
 (string) (len=1) "2",
 (string) (len=1) "3",
 (string) (len=1) "4",
 (string) (len=1) "5",
 (string) (len=1) "6",
 (string) (len=1) "7",
 (string) (len=1) "8",
 (string) (len=1) "9",
 (string) (len=1) "0"
}
(interface {}) <nil>
Convert map[int]string to map[string]int
(map[string]int) (len=10) {
 (string) (len=1) "8": (int) 8,
 (string) (len=1) "3": (int) 3,
 (string) (len=1) "5": (int) 5,
 (string) (len=1) "6": (int) 6,
 (string) (len=1) "7": (int) 7,
 (string) (len=1) "4": (int) 4,
 (string) (len=1) "9": (int) 9,
 (string) (len=1) "0": (int) 0,
 (string) (len=1) "1": (int) 1,
 (string) (len=1) "2": (int) 2
}
(interface {}) <nil>
```
