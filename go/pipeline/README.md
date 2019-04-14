

- definition: a pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function.


                              -->[ func1 channel1 ] --
                              |                      |  
upstream --> inbound channels -->[ func1 channel2 ] --->  outbound channels --> downstream
                              |                      |  
                              -->[ func1 channel3 ] -- 

There is a pattern to our pipeline functions:

- stages close their outbound channels when all the send operations are done.
- stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.


- A buffer chan
- Explicit cancellation: do this by closing a channel

