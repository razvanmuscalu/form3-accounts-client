# form3-accounts-client
A client library in Go to access the fake Form3 account API service.

# A Few Notes

- Please take into consideration that I do not have much experience with Golang (but that I do like it and want to work with it at production level). I only self-taught myself some Golang when I spent 2 weeks making and deploying an AWS Lambda written in Golang. That was in June 2018. Since then, I only touched it during the weekend for this exercise. So I expect that this code could be improved from a Golang best-practices point of view.
- To cover for my small level of experience with the language, I focused instead more heavily on: 
  - writing readable code
  - making use of interface and struct to write the API client as a service
  - extensive testing for as much as I could think of (currently code coverage is 97%)
  - something I wanted to integrate but didn't succeed is to use `https://godoc.org/github.com/go-playground/validator` instead of `validateAccount` function.
- Regarding testing, I chose `goconvey` in order to be able to write tests in BDD style (because instructions mention that Form3 engineers tend to favour BDD - yay :) ).
  - However, I maintained a Golang idiomatic style and separated the BDD tests/features into separate Golang test functions. I thought this is good as I can run individual more granular tests while developing as opposed to having all the create-related tests in one Golang test for example.
  - I had a quick look online and was in-between: `goconvey`, `ginkgo`, `goblin`, `godog`. I eliminated `ginkgo` after reading some negative opinions, and `godog` as it would be too far off from the Golang idiomatic style (though I use cucumber heavily and like it very much, also though I can achieve similar outcomes with the other mentioned frameworks). I went with `goconvey` over `goblin` in the end because it has more stars on GitHub (more popularity, better documentation, better support).
- I ended up building 2 docker integrations:
  - one runs the BDDs on top of the fake Accounts API at command line via `go test -v`. I tried making the image smaller here by using alpine, but still it's still 390MB (359MB are coming from golang-alpine however). I see your fake Accounts API image has only 63.4MB though so I must be doing something wrong.
  - one starts the `goconvey` web application on `localhost:8081` (this web application comes by default with `goconvey` so thought I'd make use of it). It's pretty cool actually as it detects changes to code and re-runs the tests in realtime. Had to use the full blown golang docker image as this web application requires `gcc`.
  
# How To Run My Code

- as required, `docker-compose up` will run the tests on command line and also spin up the `goconvey` web app on `localhost:8081`

# Instructions

# Form3 Take Home Exercise

## Instructions

This exercise has been designed to be completed in 4-8 hours. The goal of this exercise is to write a client library 
in Go to access our fake [account API](http://api-docs.form3.tech/api.html#organisation-accounts) service. 

### Should
- Client library should be written in Go
- Document your technical decisions
- Implement the `Create`, `Fetch`, `List` and `Delete` operations on the `accounts` resource. Note that filtering of the List operation is not required, but you should support paging
- Focus on writing full-stack tests that cover the full range of expected and unexpected use-cases
 - Tests can be written in Go idiomatic style or in BDD style. Form3 engineers tend to favour BDD. Make sure tests are easy to read
 - If you encounter any problems running the fake accountapi we would encourage you to do some debugging first, 
before reaching out for help

#### Docker-compose

 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service 

### Please don't
- Use a code generator to write the client library
- Implement an authentication scheme

## How to submit your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, copy the `docker-compose` from this repository
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 and @form3tech-interviewer-2 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team
- Usernames of the developers reviewing your code will then be provided for you to grant them access to your private repository
- Put your name in the README