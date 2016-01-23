package sisho

import (
	"os"
	"os/exec"
)

func (s *Sisho) clone() error {
	// * clone repository
	_, err := exec.Command("git", "clone", s.repoURI, tmpDir+s.repoUser+s.repoName).Output()
	if err != nil {
		return err
	}
	return nil
}

func (s *Sisho) clean() error {
	// * clean .tmp directory
	err := os.RemoveAll(tmpDir + s.repoUser + s.repoName)
	if err != nil {
		return err
	}
	return nil
}
