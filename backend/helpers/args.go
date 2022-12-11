package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
)

func GetArgs() map[string]string {
	args := make(map[string]string)
	for _, arg := range os.Args[1:] {
		key, val, err := normalizeArgs(arg)
		if err != nil {
			log.Fatalf("Error reading arguments: %s", err)
		}
		args[key] = val
	}
	return args
}

func normalizeArgs(arg string) (string, string, error) {
	if strings.Contains(arg, "=") {
		strParts := strings.Split(arg, "=")
		key := strParts[0]
		value := strParts[1]
		return key, value, nil
	}
	return "", "", errors.New("arguments not in right format")
}
