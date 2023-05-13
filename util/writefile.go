package util

import (
	"os"
)

func Writefile(file string, content string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		return err2
	}

	return nil
}
