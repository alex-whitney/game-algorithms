# Game Algorithms
I intend for this to be the last part in an approximately one day workshop,
aimed at students aged 13 - 18. The goal is to get students interested in
computer science and to provide some hands-on learning.

This project is a work-in-progress. This repository consists of a number of
simple games that can be played fully via console, or "in real life" against an
AI opponent (so long as a human maniputes the physical board).

The intention is for a small group to describe an algorithm or for building
a simple AI, and for the exercise leader to code the solution. This mob
programming exercise will allow students to see what code looks like, how it is
written, and understand that the code is just a tool to implement an algorithm.

My goal is that students with experience writing code or who have taken computer
programming classes may be able to write a solution on their own, though that
might be a little ambitous.

I've chosen Go because the language is fairly simple, IDE tooling is reasonably
good, and error messages are at least somewhat helpful.

## Running the app
Currently, it's a bit of a kludge.

`go run .` from the root to run what is "configured" - you'll need to muck with
`main.go` to change what game is being played, and what player(s) are involved.
Do note that an AI can play against an AI.

