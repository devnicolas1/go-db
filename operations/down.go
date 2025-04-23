package operations

import (
	"fmt"
	"go-db/types"
)

func Down(currentLevel *types.Level) *types.Level {
	newLevel := types.NewLevel(currentLevel, currentLevel.Depth+1)
	fmt.Printf("Moving down to level %d\n", newLevel.Depth)
	return newLevel
}
