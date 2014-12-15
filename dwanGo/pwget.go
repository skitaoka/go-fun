// parallel wget
package main

import (
	"net/http"

	"os"
	"io"
	"strings"
	"log"
	"runtime"
	"sync"

	"github.com/codegangsta/cli"
)

// download file from url
func download(url string, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	log.Println(url)
	path := strings.Split(url, "/")

	file, err := os.Create(path[len(path) - 1])
	if err != nil {
		return
	}
	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return
	}

	return
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "pwget"
	app.Usage = "pwget [urls]"
	app.Action = func(c *cli.Context) {
		wg := new(sync.WaitGroup)

		for _, url := range os.Args[1:] {
			wg.Add(1)
			go download(url, wg)
		}
		wg.Wait()
	}

	app.Run(os.Args)
}

