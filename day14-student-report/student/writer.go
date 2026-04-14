package student

import (
	"os"
)

func WriteReport(filename string, report string) error {
	return os.WriteFile(filename, []byte(report), 0644)
}
