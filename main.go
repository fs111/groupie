package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]int)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}
		value, exists := m[s]
		if exists {
			m[s] = value + 1
		} else {
			m[s] = 1
		}
	}
	for key, val := range m {
		fmt.Printf("%d %s\n", val, key)
	}

}
