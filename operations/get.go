package operations

import "fmt"

func Get(database map[string]string, parts []string) {
	if len(parts) >= 2 {
		key := parts[1]
		value, ok := database[key]
		if !ok {
			fmt.Printf("Key '%s' not found in database\n", key)
		} else {
			fmt.Printf("Value for key '%s': %s\n", key, value)
		}
	} else {
		fmt.Println("Error: GET requires a key")
	}
}
