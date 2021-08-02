# prodcon
Simple application to generate, read, and show, tokens.

## Producer
A token generator that creates a file with 10 million (default) random tokens,
one per line, each consisting of seven lowercase letters a-z, and save them to
storage/tokens (ignored).
```bash
go run cmd/producer/main.go --amount=1000 --length=10 --path=storage/file_name
```
### Parameters
* *amount*: the number of tokens to be generated (default: 10000000)
* *length*: the length of the generated token (default: 7)
* *path*: the file location to save the generated tokens (default: storage/tokens)

## Consumer
A token reader that reads the previously created file and stores the tokens in
a cache, and then into a relational DB, dealing with with duplicates.
```bash
go run cmd/consumer/main.go --path=storage/file_name
```
### Parameters
* *path*: the file location to read the tokens from (default: storage/tokens)

## Visualizer
A visualizer to show a list of all non-unique tokens and their frequencies.
