# Goal

After becoming comfortable with Golang's syntax, this service will be the actual implementation of the idea of "campaign-gateway"

# Architectural & Codebase Considerations/Limitations

- Did not add a "flight system" which schedules batches of user data as it would be campaign dependent and has it's own complexities (for example what if a campaigns runs every 15 minutes and the volume of data is huge - would have to consider data duplicacy, number of messages received by a customer, etc. ) since it is more of a product design problem first. Given that this is a personal project, I didn't feel the need for such a mechanism to exist.
- Did not integrate any SQL migrations.
- Adding to the flight system limitation, there are multiple ways that campaigns can be segregated based on the volume as well. For now, we are following a very basic approach of running campaigns regardless of it's volume or it's frequency.
- Still not sure if this is the right way to write Golang (given that this is my first 'real' attempt).
- Some functions might not be implemented or might require changes. They have been marked with "TODO: " comments.


# Final Thoughts

- I had to fight the urge to abstract my code and organize it "neatly". I think the approach of writing "bare metal" code has now grown on me, especially for personal projects. I did miss the neat decorators and types present in frameworks like NestJS (for example request body validation) however my development speed was not affected.
- The standard library is more than enough for majority of my use cases.
- Go routines are very easy to understand and integrate.