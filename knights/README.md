# AI Topics - Knights

This project is part of the AI Topics series.

It implements Model Checking algorithm to play “Knights and Knaves” logical puzzle created by Raymond Smullyan.

In a Knights and Knaves puzzle, the following information is given: Each character is either a knight or a knave. A knight will always tell the truth: if knight states a sentence, then that sentence is true. Conversely, a knave will always lie: if a knave states a sentence, then that sentence is false.

The objective of the puzzle is, given a set of sentences spoken by each of the characters, determine, for each character, whether that character is a knight or a knave.

Further details: https://cs50.harvard.edu/ai/2024/projects/1/knights/

__This code is inspired by [CS50's IA](https://cs50.harvard.edu/ai/2024/)__

## Algorithm Overview:

The core of the app's functionality lies in a model-checking algorithm, which determines if a given knowledge base logically entails a query. The algorithm operates on logical sentences represented in a symbolic language.

### Logical Sentences:

Logical sentences are constructed using symbols, connectives, and quantifiers. The app supports symbols (representing variables or atomic statements), negation, conjunction (AND), disjunction (OR), implication (=>), and biconditional (<=>) operators.

### Model Checking:

The model-checking algorithm explores different assignments of truth values to symbols within the knowledge base and the query. For each assignment, it evaluates whether the knowledge base logically entails the query. The algorithm systematically checks all possible combinations of truth values, making it an exhaustive method for logical reasoning.

## Usage:

1. Defining Sentences:
- Create logical sentences using symbols and logical connectives.
- Example: A AND (B OR NOT C)

2. Constructing Knowledge Base and Query:
- Define a knowledge base by creating logical sentences that represent the information you have.
- Formulate a query that you want to answer or verify against the knowledge base.

3. Running Model Check:
- Utilize the modelCheck function to check if the knowledge base logically entails the query.
- The algorithm will explore various truth assignments to symbols and determine the logical relationship between the knowledge base and the query.


## Run the application:

There are some puzzles as examle in the code. To solve them execute:


```bash
    go run .
```
