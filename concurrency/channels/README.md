* Channels are a typed conduit through wich you can send and receive values
with the channel operator <-

```
ch <- v // Send v to channel ch
v := <- ch // Receive from ch, and assign value to v
```
* Channels must be created before use `ch := make(chan int)`

* By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
* Channels can be buffered. Provide the buffer length as the second argument to `make` to initialize a buffered channel: `ch := make(chan int, 100)`
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
* A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: 
after
`v, ok := <-ch`
ok is false if there are no more values to receive and the channel is closed.

* Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

* Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

## Channels acsiom
1) Channel - FIFO queue. The data is read same order in which it was sent.

2) Sending is blocked until there is a recipient (for unbuffered).
* `c <- x` pauses until someone does `<- c`. This provides handshake synchronization.

3) Receiving is blocked until data is avalible.`<-c` waits until something is sent.

4) The buffer changes the behavior.
* `make(chan T, N)` allows sending up to N elements without waiting.
* When the buffer is full, sending blocks again.
* When the buffer is empty, receiving blocks.

5) Only the sender should close the channel.
* `close(c)` - signal "no more new data".
* If recipient closes the channel: a panic may occur: `send on closed channel`.

6) Reading from a closed channel is always successful.
* If data is still in buffer - it is output
* When the buffer is empty - a zero value of type + "channel closed" flag (separated by a comma) is returned `v, ok := <-c //ok == false -> channel is closed and empty`.

7) Iteration over the channel (range c) is read until closed.
* The cycle will automatically end when the channel is closed and empty.

8) Channel can be passed as a parameter(in both directions or limited)
* `chan T` - two-way
* `chan <-T` - sender only
* `<-chan T` - receiver only

9) Order is not guaranted between goroutines.
* If two goroutines write to same channel - the order depends on the scheduler.
* Within a single goroutine the order is preserved.

10) Channel = data stream + completion signal.
`chan struct{}` is often used as a "signal" channel without data.

