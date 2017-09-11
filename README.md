# csvTokenise
a quick binary for tokenising a column in a csv file

## Usage
```
csvTokenise -column=1 -fname=test.csv
```

## Installation

```
go install
```

## Example
```
$ cat test.csv
abc123, mike dewar
def456, jen smith
$ csvTokenise -column=1 -fname=test.csv
abc123,8mfa 5b7a6
def456,nw8 49hm9
```

## The Hard Stuff

Uses the FF1 format-preserving encryption algorithm, implemented by capitlaone: https://github.com/capitalone/fpe
