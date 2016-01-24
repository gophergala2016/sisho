package sisho

import (
	"bufio"
	"fmt"
	"github.com/gophergala2016/sisho/util"
	"html/template"
	"os"
	"sync"
)

func (s *Sisho) generateHTMLs() error {
	fmt.Printf("Generating page files")

	// Generate HTML files
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
		panic(err)
		break
	default:
		break
	}

	wg.Wait()

	fmt.Println("OK\nPage files are generated.")
	return nil
}

func (s *Sisho) generateHTML(c Code, wg *sync.WaitGroup, e chan error) error {
	title, filename := util.PathToHTMLfilename(c.path)
	c.Title = title

	file, err := os.Open(c.path)
	if err != nil {
		s.log.Println(err)
		e <- err
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		c.TextLines = append(c.TextLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		s.log.Println(err)
		e <- err
	}

	d, _ := Asset("templates/main.tmpl")
	tmpl := template.New("main")

	t, err := tmpl.Parse(string(d))
	if err != nil {
		s.log.Println(err)
		e <- err
	}

	f, err := os.Create(s.buildDir + "/" + util.NormalizeDotFile(filename))
	if err != nil {
		s.log.Println(err)
		e <- err
	}

	t.Execute(f, c)

	fmt.Printf(".")
	defer wg.Done()

	return nil
}
