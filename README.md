This is a quick test run of go for the Game of Life.  
It uses go routines for threaded calculation of each space of a generation.
Currently this program starts with the "Glider" pattern in the middle of a 25x25 grid.  
A new generation is generated every 5 seconds.


Once checked out, you'll need to do the standard go setup of:

```
$ go install

```

To execute the code:

```
$ $GOPATH/bin/golife
```

