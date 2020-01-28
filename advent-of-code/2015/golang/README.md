# Advent of Code 2015 - AWK

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [01](#01)
  - [Lessons](#lessons)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## 01

```bash
# part 1
$ nodemon -w ./01/part-1/main.go --exec 'go run ./01/part-1/main.go ../01.txt || exit 1'

# part 2
$ nodemon -w ./01/part-2/main.go --exec 'go run ./01/part-2/main.go ../01.txt || exit 1'
```

### Lessons

- `io/ioutil` is useful for reading in file data
- `ioutil.ReadFile` returns a byte array - attempting to write it straight out
    may not be meaningful
- `string(myByteArray)` will format the byte array as a string
- `fmt.Pintf("%s", myByteArray)` will format the byte array as a string without
    the additional `string` call
- the go docs seeem to favour using `log.Fatal(err)` for logging errors -
    `panic` is used in Go By Example
    - `log.Fatal` is equivalent to `fmt.Println(msg); os.Exit(1)`
    - `panic(msg)` will exit once all deferred invocations have completed
        executing


