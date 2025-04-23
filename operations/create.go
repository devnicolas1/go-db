package operations

import (
	"fmt"
	"go-db/types"
	"strings"
)

func Create(currentLevel *types.Level, parts []string) {
	if len(parts) >= 3 {
		key := parts[1]
		value := strings.Join(parts[2:], " ")
		currentLevel.Changes[key] = value
		fmt.Printf("Created entry: %s = %s\n", key, value)
	} else {
		fmt.Println("Error: CREATE requires both key and value")
	}
}
