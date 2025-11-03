## Code sample on how to work with iterators in Go lang

I've just practice with `iter` and find more interesting that:
* it's super annoyng in case of naming, really
* `iter.Seq` is something similar to Kotlin `sequence`, so in terms of "pull" iterator it's same as a standard Java iterator pattern implementation
* main work lays on `for .. range` language constraction and somethimes it feels uncomfortable to work with, especially in case with `nil` or `nil-iter`


I've foound that we can build OOPish sequence methods and chaining operation over container\collection. I think It could work in some cases with domain models \ DDD-sh patterns. But I don't have real-world expertice in big project where it would work, so it's just a throughts for future me.

What's I don't like - we have to use `Seq2` instead of `Seq` time to time and it's a bit wired in case of error handling and indroduction local struct like first\second or left\right `Pair` (but it's general question for Go lang and as I've read this question is appears time to time). What I mean, in go errors model we try to avoid throwing or launching panic, so in case when we need to convert something and it have a chance for an errors, then we force to use more transofmation or complex syntax. But it's the language limitation, I've accept it.

And the last: iters are awesome to practicce Golang generics, I worked with go before they announced them and this small excersice was a great as learning this feature too. 
