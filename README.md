[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-fsize)](https://repl.it/github/matt4biz/go-class-fsize)

# Go class: File size example
This example shows two ways to read a file (and figure its size).

In each case, note how errors are handled.

## Part 1
The first example `fsize.go` reads the entire file into a slice and then finds its length. This is simple, but not efficient for large files.

```
$ go run ./fsize test.txt
The file has 490 bytes
```

## Part 2
The second example `flsize.go` reads the file chunk by chunk into a buffer until end-of-file (EOF) is reached. It sums up the total of all the bytes read into the buffer.

```
$ go run ./flsize test.txt
read 100 chars
read 100 chars
read 100 chars
read 100 chars
read 86 chars
The file has 490 bytes
```
