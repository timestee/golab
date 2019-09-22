Given a nil channel c:
 - <-c receiving from c blocks forever
 - c <- v sending into c blocks forever
 - close(c) closing c panics
