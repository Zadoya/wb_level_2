package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

type Wget struct {
	colly   *colly.Collector
	Dir     string
	Domain  *url.URL
	isVisit map[string]bool
}

func NewWget() *Wget {
	return &Wget{
		isVisit: make(map[string]bool),
	}
}

func main() {
	wget := NewWget()
	wget.Init()
	wget.MakeDir()
	wget.VisitPage()
	wget.SavePage()
	wget.isVisit[wget.Domain.String()] = true
	err := wget.colly.Visit(wget.Domain.String())
	if err != nil {
		log.Fatal(err)
	}
	wget.colly.Wait()
	fmt.Printf("Save [%v] complete\n", wget.Domain.String())
}

// Init - инициализация с базовыми настройками
func (w *Wget) Init() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	w.Domain, err = url.ParseRequestURI(strings.TrimSuffix(input, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	w.Dir = "test_sites/"
	w.colly = colly.NewCollector(colly.AllowedDomains(w.Domain.Host))
}

// MakeDir - создание базовой директории
func (w *Wget) MakeDir() {
	domain := w.Domain.Hostname()
	if err := os.MkdirAll(w.Dir+domain, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	err := os.Chdir(w.Dir + domain)
	if err != nil {
		log.Fatal(err)
	}

}

// VisitPage - посещение страниц по ссылкам типа <a>
func (w *Wget) VisitPage() {
	w.colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absLink := e.Request.AbsoluteURL(link)
		if !w.isVisit[absLink] {
			err := w.colly.Visit(absLink)
			if err != nil {
				fmt.Printf("error page [%v]: %v\n", absLink, err)
			}
			w.isVisit[absLink] = true
		}

	})
}

// GetPath - получение пути для схранения файлов
func (w *Wget) GetPath(currentPath string) (pathDir string, fullPath string) {
	fullPath = w.Dir + w.Domain.Host + currentPath
	pathDir = fullPath

	if path.Ext(currentPath) == "" { // если путь без расширения, то добавляем '/'
		if fullPath[len(fullPath)-1] != '/' {
			fullPath += "/"
		}
		fullPath += "index.html"
	}
	if index := strings.LastIndex(fullPath, "/"); index != -1 {
		pathDir = fullPath[:index]
	}
	return
}

// SavePage - сохранение страницы в нужные пути
func (w *Wget) SavePage() {
	w.colly.OnResponse(func(r *colly.Response) {
		pathDir, fullPath := w.GetPath(r.Request.URL.Path)
		fmt.Printf("save: [%v] to [%v]\n", r.Request.URL.Path, pathDir)
		//if _, err := os.Stat(pathDir); err != nil {
		if err := os.MkdirAll(pathDir, os.ModePerm); err != nil {
			log.Fatalln(err, 1)
		}
		//}
		err := r.Save(fullPath)
		if err != nil {
			return
		}
	})
}