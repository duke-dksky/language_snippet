Curious Channels
====================================

[参考文章](https://dave.cheney.net/2013/04/30/curious-channels)

1. A closed channel never blocks
  * This powerful idiom allows you to use a channel to send a signal to an unknown number of goroutines, 
    without having to know anything about them, or worrying about deadlock.

2. A nil channel always blocks
  * a channel value that has not been initalised, or has been set to nil will always block. 


In conclusion, the simple properties of closed and nil channels are powerful building blocks that can be used to create highly concurrent programs that are simple to reason about.


