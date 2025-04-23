package operations

import (
	"fmt"
	"go-db/types"
)

func Discard(currentLevel *types.Level) *types.Level {
	if currentLevel.Depth == 0 {
		fmt.Println("Error: Cannot discard changes at level 0")
		return currentLevel
	}

	parent := currentLevel.Parent
	fmt.Printf("Discarded changes at level %d and moved up to level %d\n",
		currentLevel.Depth, parent.Depth)
	return parent
}
