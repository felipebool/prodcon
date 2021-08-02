# prodcon
Simple application to generate, read, and show, tokens.

## Producer
A token generator that creates a file with 10 million (default) random tokens,
one per line, each consisting of seven lowercase letters a-z.

## Consumer
A token reader that reads the previously created file and stores the tokens in
a relational DB. It also deals with duplicates, not saving them to DB.

## Visualizer
A visualizer to show a list of all non-unique tokens and their frequencies.
