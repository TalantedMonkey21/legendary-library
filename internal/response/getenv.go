package response

import (
	"fmt"
	"os"
)


func GetEnv(key, defValue string) string {
	if key == "" {
		fmt.Printf("Not found %v, use default\n", key)
		return defValue
	}
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Not found %v value, use default\n", key)
		return defValue
	}
	return value
}