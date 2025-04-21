package operations

import "fmt"

func Create(database map[string]string, parts []string) {
	if len(parts) >= 3 {
		key := parts[1]
		// Join the rest as the value (in case value has spaces)
		// value := strings.Join(parts[2:], " ")
		// Store in database
		database[key] = parts[2]
		fmt.Printf("Created entry: %s = %s\n", key, parts[2])
	} else {
		fmt.Println("Error: CREATE requires both key and value")
	}
}
