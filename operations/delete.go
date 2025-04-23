package operations

import (
	"fmt"
	"go-db/types"
)

func Delete(currentLevel *types.Level, parts []string) {
	if len(parts) >= 2 {
		key := parts[1]
		if _, ok := currentLevel.Get(key); !ok {
			fmt.Printf("Key '%s' not found in database\n", key)
		} else {
			delete(currentLevel.Changes, key)
			fmt.Printf("Deleted entry: %s\n", key)
		}
	} else {
		fmt.Println("Error: DELETE requires a key")
	}
}
