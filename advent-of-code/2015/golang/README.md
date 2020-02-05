# Advent of Code 2015 - Golang

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [01](#01)
  - [Lessons](#lessons)
- [02](#02)
  - [Lessons](#lessons-1)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## 01

```bash
# part 1
$ nodemon -w ./01/part-1/main.go --exec 'go run ./01/part-1/main.go ../01.txt || exit 1'

# part 2
$ nodemon -w ./01/part-2/main.go --exec 'go run ./01/part-2/main.go ../01.txt || exit 1'
```

### Lessons

- to slice a slice or array, use `xs[start:end]`.
- one gets arguments passed to the process via `os.Args` - this is like
    `process.env` in Node
- to get the current working directory, i.e. the directory in which the
    executable lies, we use `os.getWd()`
- to join paths one uses `path.Join(path1, path2, [...])`. This is like
    `require(path.resolve(path1, path2))` in Node
- `io/ioutil` is useful for reading in file data
- `ioutil.ReadFile` returns a byte array - attempting to write it straight out
    may not be meaningful. `ioutil.ReadFile` is similar to `fs.readFile` in Node
- `string(myByteArray)` will format the byte array as a string
- `fmt.Pintf("%s", myByteArray)` will format the byte array as a string without
    the additional `string` call
- the go docs seeem to favour using `log.Fatal(err)` for logging errors -
    `panic` is used in Go By Example
    - `log.Fatal` is equivalent to `fmt.Println(msg); os.Exit(1)`
    - `panic(msg)` will exit once all deferred invocations have completed
        executing


## 02

```bash
# part 1
$ nodemon -w ./02/part-1/main.go --exec 'go run ./02/part-1/main.go ../02.txt || exit 1'

# part 2
$ nodemon -w ./02/part-2/main.go --exec 'go run ./02/part-2/main.go ../02.txt || exit 1'
```

### Lessons

- `backticks` denote string literals, as they do in Javascript. There doesn't
    seem to be string interpolation, however
- "double-quotes" denote interpreted string literals. Special characters, such
    as backslashes, need to be escaped
- the `regexp` package allows for one to define regular expressions
    - start by defining your pattern using `regexp.Compile` or
        `regexp.MustCompile`, followed by using the resulting regular expression
        to perform a match
        - this is like `new RegExp(/pattern/)` in Javascript
    - to find all occurrences in a string:

        ```golang
        newlineRx := regexp.MustCompile(`\n`)

        // or
        newlineRx := regexp.MustCompile("\\n")

        matches := newlineRx.FindAll(str)
        ```
- the `strings` package allows for text manipulation, such as splitting a string
    by a specific character:

    ```golang
    words := strings.Split(str)
    ```
      
    - `strings.Split` is like `String.prototype.split` in Javascript
- use `strconv.Atoi` to parse integers from alphanumeric strings
    - `strconv.Atoi` is similar to `parseInt` in Javascript
- the `sort` package is useful for sorting slices of different types:

    ```golang
    xs := []int{3,2,1}
    // sorts and mutates xs
    sort.Ints(xs)
    ```

    - `sort` is similar to `Array.prototype.sort` in Javascript, but with
        explicitly typed functions
