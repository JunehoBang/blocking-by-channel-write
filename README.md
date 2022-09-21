# blocking-by-channel-write

This repository was created to test golang channel's possibility of causing blocking.

Let us assume that a channel instance has the capacity of `1`. 

In case a thread enqueues an element into the channel,
the thread falls into the blocking status until the value is dequeued.

To avoid the blocking status by channel communication, 
it is required to set the length to be `>2`
