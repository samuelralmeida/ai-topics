package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/samuelralmeida/ai-topics/search/frontier"
	"github.com/samuelralmeida/ai-topics/search/maze"
)

func main() {
	fmt.Println("AI TOPICS - SEARCH")

	flagMazeFile := flag.String("m", "maze1.txt", "maze filename")
	flagFrontier := flag.String("f", "queue", "frontier option: stack | queue | greedy")
	flagShowSolution := flag.Bool("s", true, "image shows solution path")
	flagShowExplored := flag.Bool("e", false, "image shows explored path")

	flag.Parse()

	m, err := maze.NewMaze(*flagMazeFile, "B")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Goal)

	var f maze.Frontier

	f = frontier.NewQueueFrontier()

	switch *flagFrontier {
	case "stack":
		f = frontier.NewStackFrontier()
	case "greedy":
		f = frontier.NewGreedyFrontier(m.Goal)
	}

	solution, err := m.Solve(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s solution\n", *flagFrontier)
	m.GenerateImage(
		strings.Replace(fmt.Sprintf("%s-%s.png", *flagMazeFile, *flagFrontier), ".txt", "", 1),
		solution, *flagShowSolution, *flagShowExplored,
	)

	fmt.Println("Custo:", solution.NumExplored)
}
