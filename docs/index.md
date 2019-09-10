# Goverage

A cli to give -coverpkg=all the desire behavior.

## Why

What desire is? When we use `go test -coverpkg=all ./... -cover` it gives a coverage from all libraries usade not only my projects files.

## Poor Solution, but still better

I will take the file with all libs and remove it all, easy peace japanese.
It's poor I know but still better then make tests for all packages, again and again.

## Usage

```sh
goverage clean $FULL_FILE_NAME -o $NEW_FILE_NAME --remove-origin
```

```sh
goverage --help
```
