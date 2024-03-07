# AI Topics - Minesweeper

This project is part of the AI Topics series.

This Minesweeper AI program is implemented in the Go programming language. The AI is designed to play the classic Minesweeper game intelligently by making informed moves based on revealed cells and the number of adjacent mines.

Further details: https://cs50.harvard.edu/ai/2024/projects/1/minesweeper/

__This code is inspired by [CS50's IA](https://cs50.harvard.edu/ai/2024/)__

## Features

### Sentence Representation

The AI represents its knowledge using the concept of sentences, where a sentence contains a list of cells and the count of mines around those cells.

### Knowledge Management

The AI dynamically updates its knowledge by marking cells as mines or safe based on the revealed information. It maintains a list of sentences, each representing a set of cells and their mine count.

### Logical Inference

The AI uses logical inference to deduce new safe or mine cells. It iteratively applies rules to update its knowledge until no further deductions can be made.

### Random and Safe Moves

The AI provides methods to make safe moves and random moves. Safe moves are chosen based on the information inferred, while random moves are fallbacks when no safe moves are apparent.

## Run the application

There are some puzzles as examle in the code. To solve them execute:


```bash
    go run main.go
```
