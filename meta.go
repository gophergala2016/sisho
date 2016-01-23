package sisho

import (
	"github.com/gophergala2016/sisho/util"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Manifest struct {
	Path        string
	ContentType string
}

type TableOfContents struct {
	Title        string
	Author       string
	Year         int
	Date         string
	HTMLContents []Manifest
	Assets       []Manifest
}

func (s *Sisho) gerenateMeta() error {
	// * create meta files
	// 	* content.opf
	// 	* toc.ncx

	var opfT, ncxT *template.Template
	var err error

	opfT, err = template.ParseFiles("templates/content.opf.tmpl")
	ncxT, err = template.ParseFiles("templates/toc.ncx.tmpl")
	if err != nil {
		return err
	}

	now := time.Now()

	var htmlContens []Manifest = []Manifest{}
	r := regexp.MustCompile(s.buildDir + "/")

	err = filepath.Walk(s.buildDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		htmlContens = append(htmlContens, Manifest{r.ReplaceAllString(path, ""), ""})
		return nil
	})

	if err != nil {
		return err
	}

	var assets []Manifest = []Manifest{}
	for _, v := range s.assets {
		assets = append(assets, Manifest{v.relativePath, v.contentType})
	}

	var opfF, ncxF *os.File

	opfF, err = os.Create(s.buildDir + "/content.opf")
	if err != nil {
		return err
	}
	ncxF, err = os.Create(s.buildDir + "/toc.ncx")
	if err != nil {
		return err
	}

	t := &TableOfContents{
		Title:        s.repoName,
		Author:       s.repoUser,
		Year:         now.Year(),
		Date:         util.TimeToString(now),
		HTMLContents: htmlContens,
		Assets:       assets,
	}

	// generate content.opf
	opfT.Execute(opfF, t)

	// generate toc.ncx
	ncxT.Execute(ncxF, t)

	// copy static files
	err = util.CopyDir("templates/static", s.buildDir)
	if err != nil {
		return err
	}
	return nil
}
