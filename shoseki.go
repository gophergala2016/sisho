package sisho

import (
	"log"
	"os"
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
	contents []content
}

type content struct {
	path string
	name string
}

const (
	tmpDir string = ".tmp"
)

func Run() {
	s := NewSisho("github.com/kogai/golip")
	var err error

	err = s.clone()
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
		repoURI:  "git@" + ss[0] + ":" + ss[1] + "/" + ss[2] + ".git",
	}
}
