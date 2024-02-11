package main

import (
	"fmt"
	"log"

	"github.com/samuelralmeida/ai-topics/search/frontier"
	"github.com/samuelralmeida/ai-topics/search/maze"
)

func main() {
	fmt.Println("AI TOPICS - SEARCH")

	maze, err := maze.NewMaze("maze3.txt", "B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("stack solution")
	maze.PrintSolve(frontier.NewStackFrontier())

	fmt.Println("")
	fmt.Println("queue solution")
	maze.PrintSolve(frontier.NewQueueFrontier())
}
