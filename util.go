package sisho

import (
	"strings"
)

func pathToHTMLfilename(rawStr string) (string, string) {
	s := strings.Split(rawStr, "/")
	title := strings.Join(s[1:len(s)], "-")
	return title, title + ".html"
}
