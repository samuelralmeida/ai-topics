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
	// mazeStackAI.PrintSolve(frontier.NewStackFrontier())
	mazeStackAI.ImageSolve("stack.png", frontier.NewStackFrontier(), true, true)

	mazeQueueAI, err := maze.NewMaze("maze3.txt", "B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("queue solution")
	// mazeQueueAI.PrintSolve(frontier.NewQueueFrontier())
	mazeQueueAI.ImageSolve("queue.png", frontier.NewQueueFrontier(), true, true)
}
