package sisho

import (
	"os"
	"os/exec"
)

func (s *Sisho) clone() error {
	// * clone repository
	_, err := exec.Command("git", "clone", s.repoURI, s.tmpDir).Output()
	if err != nil {
		return err
	}
	return nil
}

func (s *Sisho) clean() error {
	// * clean .tmp directory
	err := os.RemoveAll(s.tmpDir)
	if err != nil {
		return err
	}
	return nil
}
