package operations

import (
	"fmt"
	"go-db/types"
)

func Get(currentLevel *types.Level, parts []string) {
	if len(parts) >= 2 {
		key := parts[1]
		value, ok := currentLevel.Get(key)
		if !ok {
			fmt.Printf("Key '%s' not found in database\n", key)
		} else {
			fmt.Printf("Value for key '%s': %s\n", key, value)
		}
	} else {
		fmt.Println("Error: GET requires a key")
	}
}
