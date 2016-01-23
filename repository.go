package sisho

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func (s *Sisho) walkRepo() error {
	// * walk repo and generate constitution of repo.
	err := filepath.Walk(s.tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		isGitDir, isGitDirErr := regexp.MatchString(s.tmpDir+"/.git/", path)
		if isGitDirErr != nil {
			return err
		}
		if isGitDir {
			return nil
		}

		tmp := strings.Split(path, "/")
		c := Content{
			path:      path,
			name:      tmp[len(tmp)-1],
			Title:     "",
			TextLines: []string{},
		}
		s.contents = append(s.contents, c)
		return nil
	})
	if err == nil {
		return err
	}
	return nil
}
