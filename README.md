# GoQueue

_Because it sounds like Goku_

## Package

This is a very lightweight implementation of a queue. There are two main functions, `Pop` and `Push`. Supporting functions are `Len` and `Cap`

## Usage

```go
q := queue.NewQueue()
...
q.Push(data)
...
dat, valid := q.Pop()
```

`Len` returns the number of queue objects currently stored in the queue

`Cap` returns the current capacity of the queue

See the `_example` directory for a more in-depth example
