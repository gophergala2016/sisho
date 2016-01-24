package sisho

import (
	"log"
	"os"
	"strings"
)

type Sisho struct {
	log      *log.Logger
	repoName string
	repoUser string
	repoURI  string
	tmpDir   string
	buildDir string
	contents []Code
	assets   []ContentAsset
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

type ContentAsset struct {
	Content
	ext string
}

const (
	tmpBaseDir string = ".tmp"
)

func Run(pathToRepository string) error {
	s := NewSisho(pathToRepository)
	var err error

	if err = s.clone(); err != nil {
		return err
	}

	if err = s.walkRepo(); err != nil {
		return err
	}

	if err = s.generateHTMLs(); err != nil {
		return err
	}

	if err = s.gerenateMeta(); err != nil {
		return err
	}

	if err = s.epubify(); err != nil {
		return err
	}

	if err = s.clean(); err != nil {
		return err
	}

	return nil
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
