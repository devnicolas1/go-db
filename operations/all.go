package operations

import "fmt"

func All(database map[string]string) {
	if len(database) == 0 {
		fmt.Println("Database is empty")
	} else {
		fmt.Println("Database contents:")
		for key, value := range database {
			fmt.Printf("  %s = %s\n", key, value)
		}
	}
}
