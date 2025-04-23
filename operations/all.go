package operations

import (
	"fmt"
	"go-db/types"
)

func All(currentLevel *types.Level) {
	if len(currentLevel.GetAllCombined()) == 0 {
		fmt.Println("Database is empty")
	} else {
		fmt.Println("Database contents:")
		for key, value := range currentLevel.GetAllCombined() {
			fmt.Printf("  %s = %s\n", key, value)
		}
	}
}
