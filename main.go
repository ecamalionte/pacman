package main

import "fmt"
import "os"
import "bufio"

func main() {

	maze, _ := loadMaze()
	init()
	// initialize game

	// load resources

	// game loop
	for {
		printScreen(maze)
		// update screen

		// process input

		// process movement

		// process collisions

		// check game over

		break

		// repeat
	}
}

func init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

/*
File
- # represents a wall
- . represents a dot
- P represents the player
- G represents the ghosts (enemies)
- X represents the power up pills
*/

func printScreen(maze []string) {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func loadMaze() ([]string, error) {
	f, err := os.Open("./maze01.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var maze []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return maze, nil
}
