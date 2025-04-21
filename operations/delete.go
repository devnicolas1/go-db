package operations

import "fmt"

func Delete(database map[string]string, parts []string) {
	if len(parts) >= 2 {
		key := parts[1]
		if _, ok := database[key]; !ok {
			fmt.Printf("Key '%s' not found in database\n", key)
		} else {
			delete(database, key)
			fmt.Printf("Deleted entry: %s\n", key)
		}
	} else {
		fmt.Println("Error: DELETE requires a key")
	}
}
