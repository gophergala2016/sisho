package sisho

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func (s *Sisho) epubify() error {
	newfile, err := os.Create(s.repoName + ".epub")
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	err = addZip(s.buildDir+"/mimetype", s.buildDir, zipWriter)
	if err != nil {
		return err
	}

	err = addAll(s.buildDir, zipWriter)
	if err != nil {
		return err
	}
	return nil
}

func addAll(buildDir string, zipWriter *zip.Writer) error {
	err := filepath.Walk(buildDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if path == buildDir+"/mimetype" {
			return nil
		}

		err = addZip(path, buildDir, zipWriter)
		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func addZip(filename, buildDir string, zipWriter *zip.Writer) error {
	zipfile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	info, err := zipfile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	r := regexp.MustCompile(buildDir + "/")
	header.Name = r.ReplaceAllString(filename, "")
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, zipfile)
	return nil
}
