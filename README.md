## AsyncWait
### Wait, don't sleep :)

#### Helps in testing async methods in your Go program
#### If you need testing your async method, wait rather than sleep

## Uses:
- Helps avoid undeterministic tests
- Helps have faster tests, because you wait enough (No more or less :))

## Usage
```
go get github.com/rShetty/asyncwait
or, just add `github.com/rShetty/asyncwait` as dependency and preferably fix a version.
```

### Code:
``` 
        func doWork(workChan chan int, something Something)
          go func() {
            for i := 0; i < 5; i++ {
              workChan <- i
            }
          }

          something.call()
        }
```

#### The above code pushes 5 int elements into a channel and the aim is to test this async job

### Test:
```
        workChan := make(chan int)
        maxWaitInMillis := 100
        pollIntervalInMillis := 20

        doWork(workChan)

        // Check for the length of channel to be 5
        predicate := func() bool {
            return len(workChan) == 5 
        }

        successful := NewAsyncWait(100, 20).Check(predicate)

        if successful != true {
           // You fail
        }

        // Assert call was called on something
```

