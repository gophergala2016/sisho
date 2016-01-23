package util

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func PathToHTMLfilename(rawStr string) (string, string) {
	s := strings.Split(rawStr, "/")
	title := strings.Join(s[1:len(s)], "-")
	return title, title + ".html"
}

func TimeToString(t time.Time) string {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	result := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day)
	return result
}

func IsBinary(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".png" || ext == ".gif" || ext == ".jpg"
}
