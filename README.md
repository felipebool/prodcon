# prodcon
Simple application to generate, read, and show, tokens.

## Producer
A token generator that creates a file with 10 million (default) random tokens,
one per line, each consisting of seven lowercase letters a-z, and save them to
storage/tokens (ignored).
```bash
go run cmd/producer/main.go --amount=1000 --length=10
```
### Parameters
* *amount*: the number of tokens to be generated (default: 10000000)
* *length*: the length of the generated token (default: 7)

## Consumer
A token reader that reads the previously created file and stores the tokens in
a relational DB. It also deals with duplicates, not saving them to DB.

## Visualizer
A visualizer to show a list of all non-unique tokens and their frequencies.
