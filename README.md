# [Golang Practices by Examples](https://gobyexample.com)
****

[Go](http://golang.org/) is an open source programming language designed for building simple, fast, and reliable software.
Go by Example is a hands-on introduction to Go using annotated example programs.

- [x] [hello-world](https://gobyexample.com/hello-world)
To run the program, put the code in hello-world.go and use go run.	

```sh
$ go run hello-world.go
hello world
```

Sometimes we’ll want to build our programs into binaries. We can do this using go build.	

```sh
$ go build hello-world.go
$ ls
hello-world	hello-world.go
```

We can then execute the built binary directly.

```sh
$ ./hello-world
hello world
```

- [x] [values](https://gobyexample.com/values)

- [x] [variables](https://gobyexample.com/variables)

- [x] [constants](https://gobyexample.com/constants)

- [x] [for](https://gobyexample.com/for)

- [x] [if/Else](https://gobyexample.com/if-else)

- [ ] [Switch]("ttps://gobyexample.com/switch")

- [ ] [Array]("https://gobyexample.com/arrays")

- [ ] [Slices]("https://gobyexample.com/slices")
        <li><a href="https://gobyexample.com/maps">Maps</a></li>
        <li><a href="https://gobyexample.com/range">Range</a></li>
        <li><a href="https://gobyexample.com/functions">Functions</a></li>
        <li><a href="https://gobyexample.com/multiple-return-values">Multiple Return Values</a></li>
        <li><a href="https://gobyexample.com/variadic-functions">Variadic Functions</a></li>
        <li><a href="https://gobyexample.com/closures">Closures</a></li>
        <li><a href="https://gobyexample.com/recursion">Recursion</a></li>
        <li><a href="https://gobyexample.com/pointers">Pointers</a></li>
        <li><a href="https://gobyexample.com/structs">Structs</a></li>
        <li><a href="https://gobyexample.com/methods">Methods</a></li>
        <li><a href="https://gobyexample.com/interfaces">Interfaces</a></li>
        <li><a href="https://gobyexample.com/errors">Errors</a></li>
        <li><a href="https://gobyexample.com/goroutines">Goroutines</a></li>
        <li><a href="https://gobyexample.com/channels">Channels</a></li>
        <li><a href="https://gobyexample.com/channel-buffering">Channel Buffering</a></li>
        <li><a href="https://gobyexample.com/channel-synchronization">Channel Synchronization</a></li>
        <li><a href="https://gobyexample.com/channel-directions">Channel Directions</a></li>
        <li><a href="https://gobyexample.com/select">Select</a></li>
        <li><a href="https://gobyexample.com/timeouts">Timeouts</a></li>
        <li><a href="https://gobyexample.com/non-blocking-channel-operations">Non-Blocking Channel Operations</a></li>
        <li><a href="https://gobyexample.com/closing-channels">Closing Channels</a></li>
        <li><a href="https://gobyexample.com/range-over-channels">Range over Channels</a></li>
        <li><a href="https://gobyexample.com/timers">Timers</a></li>
        <li><a href="https://gobyexample.com/tickers">Tickers</a></li>
        <li><a href="https://gobyexample.com/worker-pools">Worker Pools</a></li>
        <li><a href="https://gobyexample.com/rate-limiting">Rate Limiting</a></li>
        <li><a href="https://gobyexample.com/atomic-counters">Atomic Counters</a></li>
        <li><a href="https://gobyexample.com/mutexes">Mutexes</a></li>
        <li><a href="https://gobyexample.com/stateful-goroutines">Stateful Goroutines</a></li>
        <li><a href="https://gobyexample.com/sorting">Sorting</a></li>
        <li><a href="https://gobyexample.com/sorting-by-functions">Sorting by Functions</a></li>
        <li><a href="https://gobyexample.com/panic">Panic</a></li>
        <li><a href="https://gobyexample.com/defer">Defer</a></li>
        <li><a href="https://gobyexample.com/collection-functions">Collection Functions</a></li>
        <li><a href="https://gobyexample.com/string-functions">String Functions</a></li>
        <li><a href="https://gobyexample.com/string-formatting">String Formatting</a></li>
        <li><a href="https://gobyexample.com/regular-expressions">Regular Expressions</a></li>
        <li><a href="https://gobyexample.com/json">JSON</a></li>
        <li><a href="https://gobyexample.com/time">Time</a></li>
        <li><a href="https://gobyexample.com/epoch">Epoch</a></li>
        <li><a href="https://gobyexample.com/time-formatting-parsing">Time Formatting / Parsing</a></li>
        <li><a href="https://gobyexample.com/random-numbers">Random Numbers</a></li>
        <li><a href="https://gobyexample.com/number-parsing">Number Parsing</a></li>
        <li><a href="https://gobyexample.com/url-parsing">URL Parsing</a></li>
        <li><a href="https://gobyexample.com/sha1-hashes">SHA1 Hashes</a></li>
        <li><a href="https://gobyexample.com/base64-encoding">Base64 Encoding</a></li>
        <li><a href="https://gobyexample.com/reading-files">Reading Files</a></li>
        <li><a href="https://gobyexample.com/writing-files">Writing Files</a></li>
        <li><a href="https://gobyexample.com/line-filters">Line Filters</a></li>
        <li><a href="https://gobyexample.com/command-line-arguments">Command-Line Arguments</a></li>
        <li><a href="https://gobyexample.com/command-line-flags">Command-Line Flags</a></li>
        <li><a href="https://gobyexample.com/environment-variables">Environment Variables</a></li>
        <li><a href="https://gobyexample.com/spawning-processes">Spawning Processes</a></li>
        <li><a href="https://gobyexample.com/execing-processes">Exec'ing Processes</a></li>
        <li><a href="https://gobyexample.com/signals">Signals</a></li>
        <li><a href="https://gobyexample.com/exit">Exit</a></li>
