# AI Topics - Search

This project is part of the AI Topics series.

It implements informed and uninformed searches algorithms to solve mazes.

Here are implemented 4 algorithms:

__Depth-first search__: An algorithm that does not use external data in the analysis. It is implemented using the "stack" data structure. Faced with a fork, it exhausts the entire chosen path before returning to the branching point.

__Breadth-first search__: An algorithm that does not use external data in the analysis. It is implemented using the "queue" data structure. Faced with a fork, it divides by alternating the search paths.

__Greedy-best-first__: An algorithm that considers the goal's position in its analysis. It implements a heuristic function that measures the distance to the target using the "Manhattan distance."

__A*__: An algorithm that considers the goal's position in its analysis. It implements a heuristic function that measures the distance to the target using the "Manhattan distance" but also takes into account the cost of the already covered distance.

This code is inspired by [CS50's IA](https://cs50.harvard.edu/ai/2024/)

## Run the application:

bash

    go run main.go


Usage

    -m: Specify the maze filename (default is maze1.txt).
    -f: Choose the frontier option (depth or breadth or greedy or a-star, default is queue).
    -s: Display the solution path in the image (default is true).
    -e: Display the explored path in the image (default is false).


## Details

The greedy algorithm option uses Manhatan Distance as the heuristic function.

## Results

The program will output the cost path and generate an image of the maze with the path highlighted.

The image will be saved as {maze filename}-{frontier}.png in the current directory.

### Image example:

![maze solution image example](example.jpg)