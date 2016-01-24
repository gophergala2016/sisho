package sisho

import (
	"fmt"
	"os"
)

func (s *Sisho) clean() error {
	// * clean .tmp directory
	err := os.RemoveAll(s.tmpDir)
	if err != nil {
		return err
	}
	fmt.Println(s.repoName + ".epub is created!")
	return nil
}
