package sisho

import (
	"log"
	"os"
)

/*
Steps.
* clone repository
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
	repoPath string
	contents []content
}

type content struct {
	path string
	name string
}

func Run() {
	s := NewSisho()
	s.log.Println("hello")
}

func NewSisho() *Sisho {
	return &Sisho{
		log: log.New(os.Stdout, "", log.Lshortfile),
	}
}
