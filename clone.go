package sisho

import (
	"fmt"
	"os/exec"
	"time"
)

func (s *Sisho) clone() error {
	// Clone git repository
	fmt.Printf("Cloning git repository")

	var c chan bool = make(chan bool)
	go func(c chan bool) {
		for {
			select {
			case ch := <-c:
				if ch {
					close(c)
					return
				}
				break
			case <-time.After(1 * time.Second):
				fmt.Printf(".")
				break
			}
		}
	}(c)

	_, err := exec.Command("git", "clone", s.repoURI, s.tmpDir).Output()
	if err != nil {
		return err
	}

	c <- true
	fmt.Println("OK\nGit repository is cloned.")

	return nil
}
