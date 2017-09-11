# csvTokenise
a quick binary for tokenising a column in a csv file

## Usage
```
csvTokenise -column=1 -fname=test.csv -key=FF4359D8D580AA4F7F036D6F04FC6A94 -tweak=D8E7920AFA330A73
```

## Limitations

* converts everything to lowercase
* tokens are all alphanumeric (radix 36)
* you shouldn't use the default key or tweak

## Installation

```
go install
```

## Help
```
$ ./csvTokenise --help
Usage of ./csvTokenise:
  -column int
        column to tokenise
  -fname string
        input filename
  -key string
        key for the FF1 algorithm (default "FF4359D8D580AA4F7F036D6F04FC6A94")
  -tweak string
        tweak for the FF1 algorithm (default "D8E7920AFA330A73")
```


## Example
```
$ cat test.csv
abc123, mike dewar
def456, jen smith
xyz890, Bobby Fingers
111222, Jen Smith
aaabbb, Jen Fitzroy

$ ./csvTokenise -column=1 -fname=test.csv
abc123,8mfa 5b7a6
def456,nw8 49hm9
xyz890,66n88 7xsje5
111222,nw8 49hm9
aaabbb,nw8 nband0p
```

## The Hard Stuff

Uses the FF1 format-preserving encryption algorithm, implemented by capitlaone: https://github.com/capitalone/fpe

