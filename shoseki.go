package sisho

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
Steps.
* walk repo and generate constitution of repo.
* generate HTML
* create meta files
	* content.opf
	* mimetype (static)
	* toc.ncx
	* META-INF/container.xml (static)
* epubify all files

* If have a time...
* jump defined some func or class or variable.
* compress image files.
* goroutinize generate HTML step.
*/

type Sisho struct {
	log      *log.Logger
	repoName string
	repoUser string
	repoURI  string
	tmpDir   string
	contents []content
}

type content struct {
	path string
	name string
}

const (
	tmpBaseDir string = ".tmp"
)

func Run() {
	s := NewSisho("github.com/kogai/golip")
	var err error

	err = s.clone()
	s.log.Println(err)

	err = s.walkRepo()
	s.log.Println(err)

	err = s.clean()
	s.log.Println(err)
}

func NewSisho(repoPath string) *Sisho {
	var ss []string = strings.Split(repoPath, "/")

	return &Sisho{
		log:      log.New(os.Stdout, "", log.Lshortfile),
		repoUser: ss[1],
		repoName: ss[2],
		tmpDir:   tmpBaseDir + ss[1] + ss[2],
		repoURI:  "git@" + ss[0] + ":" + ss[1] + "/" + ss[2] + ".git",
	}
}

func (s *Sisho) walkRepo() error {
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
		c := content{
			path: path,
			name: tmp[len(tmp)-1],
		}
		s.contents = append(s.contents, c)
		return nil
	})
	if err == nil {
		return err
	}
	return nil
}
