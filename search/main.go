package main

import (
	"fmt"
	"log"

	"github.com/samuelralmeida/ai-topics/search/frontier"
	"github.com/samuelralmeida/ai-topics/search/maze"
)

func main() {
	fmt.Println("AI TOPICS - SEARCH")

	mazeStackAI, err := maze.NewMaze("maze3.txt", "B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("stack solution")
	mazeStackAI.PrintSolve(frontier.NewStackFrontier())

	mazeQueueAI, err := maze.NewMaze("maze3.txt", "B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("queue solution")
	mazeQueueAI.PrintSolve(frontier.NewQueueFrontier())
}
