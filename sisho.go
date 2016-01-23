package sisho

import (
	"log"
	"os"
	"strings"
)

/*
Steps.
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
	buildDir string
	contents []Content
}

type Content struct {
	path      string
	name      string
	Title     string
	TextLines []string
}

const (
	tmpBaseDir string = ".tmp"
)

func Run() {
	s := NewSisho("github.com/kogai/golip")
	var err error

	err = s.walkRepo()
	s.log.Println(err)

	err = s.clone()
	s.log.Println(err)

	err = s.generateHTMLs()
	s.log.Println(err)

	// err = s.clean()
	// s.log.Println(err)
}

func NewSisho(repoPath string) *Sisho {
	var ss []string = strings.Split(repoPath, "/")

	return &Sisho{
		log:      log.New(os.Stdout, "", log.Lshortfile),
		repoUser: ss[1],
		repoName: ss[2],
		tmpDir:   tmpBaseDir + ss[1] + ss[2],
		buildDir: tmpBaseDir + ss[1] + ss[2] + "/.build",
		repoURI:  "git@" + ss[0] + ":" + ss[1] + "/" + ss[2] + ".git",
	}
}
