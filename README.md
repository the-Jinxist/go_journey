## My Go Learning Journey

I will be documenting my jourmey through learning GO in this repository, go help me GOD! Learnt primitive go stuff using the official documentation here: https://www.golangprograms.com/go-language/variables.html. Learning rest api through a short course that probably shouldn't have paid for here: https://www.udemy.com/course/build-a-restful-api-with-golang-go-programming-language/learn/lecture/11760566#overview

### Day 1

Started writing Go. Was unable to import external libraries. After many tries and pulling my hair out, it finally worked.

#### Learned:
- Learnt about variable declaration, different data types and types of variable initialization
- Also learnt about go naming conventions, another thing to look out for.
- Conversion of data types
- Operators in Go
- Learnt conditional and control structures i.e if-else and switch

### Day 2
Learning about strange Go stuff. Somehow exciting

#### Learned:
- For loops.
- Learned functions, anonymous functions, higher-order functions. Also saw that functions that starts with an uppercase letter can be exported and ones with lowercase are not exported.
- Deferred function calls, worth mentioning, make sure a function is called when it's finished
- Panic! at the golang. Learnt about panic and recovery

### Day 3

#### Learned:
- Arrays: Declaring, Initializing.
- Filtering arrays and looping through them.
- Slices: Declarind and Initializing them. They seem like weird stuff to be honest.

### Day 4

#### Learned:
- Maps in Golang
- Structs: Which has to be my favourite part about Golang so far. Syntax is very funky. I love it!

### Day 5

#### Learned:
- Interfaces and self-defined types that are aren't interfaces

### Day 6

#### Learned:
- Learned about Goroutines: Some notes
    - The go keyword makes the function call to return immediately, while the function starts running in the background as a goroutine and the rest of the program continues its execution. The main function of every Golang program is started using a goroutine, so every Golang program runs at least one goroutine.
    - atomic functions should be used to solve race conditions
    - mutext can be used to is used to define critical sections that must must be carried out by one goroutine at a time
- Learned about Channels: Buffered and Unbuffered.
    - In unbuffered channel there is no capacity to hold any value before it's received. In this type of channels both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can complete. If the two goroutines aren't ready at the same instant, the channel makes the goroutine that performs its respective send or receive operation first wait. Synchronization is fundamental in the interaction between the send and receive on the channel. One can't happen without the other.
    - In buffered channel there is a capacity to hold one or more values before they're received. In this types of channels don't force goroutines to be ready at the same instant to perform sends and receives. There are also different conditions for when a send or receive does block. A receive will block only if there's no value in the channel to receive. A send will block only if there's no available buffer to place the value being sent.

### Day 7

#### Learned:
- Started initial work with rest API using the [mux library](github.com/gorilla/mux)
- Created routes for different http methods

### Day 7

#### Learned:
- Implemented the addBooks, updateBooks and deleteBooks endpoint
- Have a basic understanding of routes