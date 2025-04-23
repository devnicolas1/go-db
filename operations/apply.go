package operations

import (
	"fmt"
	"go-db/types"
)

func Apply(currentLevel *types.Level) *types.Level {
	if currentLevel.Depth == 0 {
		fmt.Println("Error: Cannot apply changes at level 0")
		return currentLevel
	}

	parent := currentLevel.Parent
	for k, v := range currentLevel.Changes {
		parent.Changes[k] = v
	}

	fmt.Printf("Applied changes to level %d and moved up\n", parent.Depth)
	return parent
}
