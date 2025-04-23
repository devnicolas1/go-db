package types

type Level struct {
	Changes map[string]string
	Parent  *Level
	Depth   int
}

func NewLevel(parent *Level, depth int) *Level {
	return &Level{
		Changes: make(map[string]string),
		Parent:  parent,
		Depth:   depth,
	}
}

func (l *Level) Get(key string) (string, bool) {
	if value, exists := l.Changes[key]; exists {
		return value, true
	}

	if l.Parent != nil {
		return l.Parent.Get(key)
	}

	return "", false
}

func (l *Level) GetAllCombined() map[string]string {
	result := make(map[string]string)

	if l.Parent != nil {
		for k, v := range l.Parent.GetAllCombined() {
			result[k] = v
		}
	}

	for k, v := range l.Changes {
		result[k] = v
	}

	return result
}
