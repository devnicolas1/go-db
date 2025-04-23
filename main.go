package main

import (
	"bufio"
	"fmt"
	"go-db/operations"
	"go-db/types"
	"os"
	"strings"
)

func main() {
	rootLevel := types.NewLevel(nil, 0)
	currentLevel := rootLevel
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple in-memory database")
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToUpper(input) == "EXIT" {
			break
		}
		parts := strings.Fields(input)
		if len(parts) > 0 {
			command := strings.ToUpper(parts[0])
			switch command {
			case "HELP":
				fmt.Println("Commands:")
				fmt.Println("  ALL: Display all entries in the database")
				fmt.Println("  GET <key>: Get the value of a key")
				fmt.Println("  CREATE <key> <value>: Create a new entry in the database")
				fmt.Println("  DELETE <key>: Delete an entry from the database")
				fmt.Println("  LEVEL: Display the current level structure")
				fmt.Println("  DOWN: Move down to the next level")
				fmt.Println("  APPLY: Apply changes to the parent level")
				fmt.Println("  DISCARD: Discard changes and move up to the parent level")
				fmt.Println("  EXIT: Exit the database")
			case "ALL":
				operations.All(currentLevel)
			case "GET":
				operations.Get(currentLevel, parts)
			case "CREATE":
				operations.Create(currentLevel, parts)
			case "DELETE":
				operations.Delete(currentLevel, parts)
			case "LEVEL":
				operations.ShowLevel(currentLevel)
			case "DOWN":
				currentLevel = operations.Down(currentLevel)
			case "APPLY":
				currentLevel = operations.Apply(currentLevel)
			case "DISCARD":
				currentLevel = operations.Discard(currentLevel)
			default:
				fmt.Println(input)
			}
		}
		fmt.Print("> ")
	}
	fmt.Println("Goodbye!")
}
