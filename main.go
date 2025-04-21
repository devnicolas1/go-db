package main

import (
	"bufio"
	"fmt"
	"go-db/operations"
	"os"
	"strings"
)

func main() {
	database := make(map[string]string)
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
				fmt.Println("  EXIT: Exit the program")
			case "ALL":
				operations.All(database)
			case "GET":
				operations.Get(database, parts)
			case "CREATE":
				operations.Create(database, parts)
			case "DELETE":
				operations.Delete(database, parts)
			default:
				fmt.Println(input)
			}
		}
		fmt.Print("> ")
	}
	fmt.Println("Goodbye!")
}
