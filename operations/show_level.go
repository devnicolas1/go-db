package operations

import (
	"fmt"
	"go-db/types"
	"strings"
)

func ShowLevel(currentLevel *types.Level) {
	fmt.Println("Level structure:")
	levels := []*types.Level{}
	for level := currentLevel; level != nil; level = level.Parent {
		levels = append([]*types.Level{level}, levels...)
	}

	for i, level := range levels {
		prefix := strings.Repeat("  ", i)
		if level == currentLevel {
			fmt.Printf("%sLevel %d (CURRENT)\n", prefix, level.Depth)
		} else {
			fmt.Printf("%sLevel %d\n", prefix, level.Depth)
		}
	}
}
