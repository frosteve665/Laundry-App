package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println()
	return scanner.Text()
}
