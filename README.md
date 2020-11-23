# merge-intervals
Parse and merge numerical intervals. The `merge.go` file contains the implementation for the `Intervals.Merge()` method,
as well as the associated types and helper functions. `merge_test.go` contains unit tests.

The module has no external dependencies except the `testify` test suite library. This module was developed and tested
with Go version `1.14`, but should work with older and newer versions as well.

## Unit tests
Run the tests for the module with the following command:
```
go test -v .
```

## Run on command line
The sub-package `merge-intervals` contains a command line tool for parsing and merging integer intervals.
The tool can be directly invoked with the following command:
```
go run ./merge-intervals 25 30 2 19 14 23 4 8 
```

Alternatively, install the tool into your `GOPATH` and run it from there:
```
go install ./merge-intervals
merge-intervals 25 30 2 19 14 23 4 8
```
