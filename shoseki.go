package sisho

import (
	"log"
	"os"
	"os/exec"
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
* clean .tmp directory

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
	contents []content
}

type content struct {
	path string
	name string
}

func Run() {
	s := NewSisho("github.com/kogai/golip")
	err := s.clone()
	s.log.Println(err)
}

func NewSisho(repoPath string) *Sisho {
	var ss []string = strings.Split(repoPath, "/")

	return &Sisho{
		log:      log.New(os.Stdout, "", log.Lshortfile),
		repoUser: ss[1],
		repoName: ss[2],
		repoURI:  "git@" + ss[0] + ":" + ss[1] + "/" + ss[2] + ".git",
	}
}

func (s *Sisho) clone() error {
	// * clone repository
	_, err := exec.Command("git", "clone", s.repoURI, ".tmp").Output()
	if err != nil {
		return err
	}
	return nil
}
