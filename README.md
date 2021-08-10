# prodcon
Simple application to generate, read, and show, tokens.

## Solution
We can see this problem as a set of processes. You must first generate tokens,
you must then read them, and save them into database. I decided to tackle this
challenge by breaking it into two applications: producer and consumer.

### Producer
This is the simplest one, a token generator that creates a file with 10 million
(default) random tokens, one per line, each consisting of seven lowercase letters
a-z, and save them to storage/tokens (default).

The process of generating tokens is entirely based on random values. There is a
charset, consisting of all the lowercase letters (\[a-z\]), and, for each position
in a string of size 7 one letter is randomly picked from the charset.

#### **Parameters**
The default values for the following parameters were set to what was asked in the
challenge description, you can change them by overwriting its default values
(--amount=1000000, to set amount to 1M, for example)
* *amount*: the number of tokens to be generated (default: 10000000)
* *length*: the length of the generated token (default: 7)
* *path*: the file location to save the generated tokens (default: storage/tokens)

### Consumer
This is where the tokens are read, and inserted into the database. Since finding
duplicates were also part of the problem, I decided to use an auxiliary data
structure. All the entries are read from the file and saved into a **hash map**,
maping token to total, which is the amount of times this token appears.

It is here, also, that the frequency of non-unique tokens is printed out
concurrently while saving the data into the database.

#### **Parameters**
For the consumer, the available parameters were added to make it easier to test
different configurations for batch size and number of workers. The default values
here worked for my hardware, but might not be the best set for other machines.
* *batch*: the number of tokens to be present in a single insert (default: 1000)
* *workers*: the number of goroutines used to access database (default: 100)
* *path*: the file location to save the generated tokens (default: storage/tokens)
