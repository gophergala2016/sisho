package sisho

import (
	"github.com/gophergala2016/sisho/util"
	"github.com/jhoonb/archivex"
	"log"
	"os"
	"strings"
)

/*
Steps.

* If have a time...
* jump defined some func or class or variable.
* compress image files.
* [x] goroutinize generate HTML step.
*/

type Sisho struct {
	log      *log.Logger
	repoName string
	repoUser string
	repoURI  string
	tmpDir   string
	buildDir string
	contents []Code
	assets   []Asset
}

type Content struct {
	path         string
	relativePath string
	name         string
	contentType  string
}

type Code struct {
	Content
	Title     string
	TextLines []string
}

type Asset struct {
	Content
	ext string
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

	err = s.generateHTMLs()
	s.log.Println(err)

	err = s.gerenateMeta()
	s.log.Println(err)

	err = s.epubify()
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

func (s *Sisho) epubify() error {
	// * epubify all files
	z := new(archivex.ZipFile)

	z.Create(s.repoName)
	z.AddFile(s.buildDir + "/mimetype")
	z.AddFile(s.buildDir + "/META-INF/container.xml")
	z.AddFile(s.buildDir + "/index.html")

	// z.AddAll(s.buildDir, false)
	defer z.Close()

	var zip, epub string
	zip = s.repoName + ".zip"
	epub = s.repoName + ".epub"

	err := util.CopyFile(zip, epub)
	if err != nil {
		return err
	}
	return nil
}
