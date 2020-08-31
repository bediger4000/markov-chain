# Daily Coding Problem: Problem #640 [Easy]

This problem was asked by Google.

You are given a starting state start, a list of transition probabilities for a
Markov chain, and a number of steps num_steps. Run the Markov chain starting
from start for num_steps and compute the number of times we visited each state.

For example, given the starting state a, number of steps 5000, and the
following transition probabilities:

```
[
  ('a', 'a', 0.9),
  ('a', 'b', 0.075),
  ('a', 'c', 0.025),
  ('b', 'a', 0.15),
  ('b', 'b', 0.8),
  ('b', 'c', 0.05),
  ('c', 'a', 0.25),
  ('c', 'b', 0.25),
  ('c', 'c', 0.5)
]
```

One instance of running this Markov chain might produce { 'a': 3012, 'b': 1656, 'c': 332 }.

## Building and running

```sh
$ go build .
$ ./markov-chain a 5000 table1
```
This runs the example probabilities.

To check my code, I wrote `table2`, which should end up with 50% of the visits to state `a`,
and 50% to state `b`.

## Analysis

I suspect that the "easy" rating was done by someone who knew what a Markov Chain was,
and had worked this sort of problem a lot, but forgot how the first time they did it felt.

This might be a decent interview question, for a mid- to senior-level position candidate.
There's a lot going on:
getting a floating point probability between 0.0 and 1.0,
finding the next states and their probabilities from the current state,
printing out the number of visits for each state.
There's probably a choice of data structures that would work,
but I chose a go `map[rune][]*TableEntry`,
where `TableEntry` is a struct that has a next state and the probability of transitioning to that next state.
That's a lot of design decisions for a candidate to talk through.

It's also entirely possible that Google, who employs a lot of "data scientists",
wants people who have inside knowledge about this type of problem.
Selecting a next state based on proability seems like the kind of thing
that has a stock solution that only specialists know.
I'm not super happy about my solution.
