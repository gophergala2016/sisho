package sisho

import (
	"bufio"
	"github.com/gophergala2016/sisho/util"
	"html/template"
	"os"
	"sync"
)

func (s *Sisho) generateHTMLs() error {
	// * generate HTMLs
	os.Mkdir(s.buildDir, 0777)

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	var e chan error = make(chan error)

	for _, c := range s.contents {
		wg.Add(1)
		go s.generateHTML(c, wg, e)
	}

	select {
	case err := <-e:
		s.log.Println(err)
		close(e)
		return err
	default:
		break
	}

	wg.Wait()
	return nil
}

func (s *Sisho) generateHTML(c Code, wg *sync.WaitGroup, e chan error) error {
	title, filename := util.PathToHTMLfilename(c.path)
	c.Title = title

	file, err := os.Open(c.path)
	if err != nil {
		e <- err
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		c.TextLines = append(c.TextLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		e <- err
	}

	t, err := template.ParseFiles("templates/main.tmpl")
	if err != nil {
		e <- err
	}

	f, err := os.Create(s.buildDir + "/" + util.NormalizeDotFile(filename))
	if err != nil {
		e <- err
	}

	t.Execute(f, c)
	defer wg.Done()

	return nil
}
