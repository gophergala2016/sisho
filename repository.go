package sisho

import (
	"github.com/gophergala2016/sisho/util"
	"mime"
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
		name := tmp[len(tmp)-1]
		ext := filepath.Ext(path)
		r := regexp.MustCompile(s.tmpDir + "/")

		if util.IsBinary(ext) {
			a := Asset{
				Content: Content{
					path:         path,
					relativePath: r.ReplaceAllString(path, ""),
					name:         name,
					contentType:  mime.TypeByExtension(ext),
				},
				ext: ext,
			}
			s.assets = append(s.assets, a)
			return nil
		}

		c := Code{
			Content: Content{
				path:         path,
				relativePath: r.ReplaceAllString(path, ""),
				name:         name,
				contentType:  mime.TypeByExtension(ext),
			},
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
