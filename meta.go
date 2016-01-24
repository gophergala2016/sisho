package sisho

import (
	"bytes"
	"encoding/xml"
	"github.com/gophergala2016/sisho/util"
	"html/template"
	"io"
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

	opfD, _ := Asset("templates/content.opf.tmpl")
	ncxD, _ := Asset("templates/toc.ncx.tmpl")

	opfT = template.New("opf")
	ncxT = template.New("ncx")

	opfT, err = opfT.Parse(string(opfD))
	if err != nil {
		return err
	}

	ncxT, err = ncxT.Parse(string(ncxD))
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
	_, err = opfF.WriteString(xml.Header)
	if err != nil {
		return err
	}

	ncxF, err = os.Create(s.buildDir + "/toc.ncx")
	if err != nil {
		return err
	}
	_, err = ncxF.WriteString(xml.Header)
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
	os.Mkdir(s.buildDir+"/assets", 0777)
	os.Mkdir(s.buildDir+"/META-INF", 0777)

	err = CopyFile("templates/assets/style.css", s.buildDir+"/assets/style.css")
	if err != nil {
		return err
	}

	err = CopyFile("templates/static/META-INF/container.xml", s.buildDir+"/META-INF/container.xml")
	if err != nil {
		return err
	}

	err = CopyFile("templates/static/mimetype", s.buildDir+"/mimetype")
	if err != nil {
		return err
	}
	return nil
}

func CopyFile(source, dest string) error {
	d, err := Asset(source)
	if err != nil {
		return err
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destfile, bytes.NewReader(d))
	if err != nil {
		return err
	}

	return nil
}
