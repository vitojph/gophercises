package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

// JSONStory parses a JSON file containing a story
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story contains our stories
type Story map[string]Chapter

// Chapter is the main item of a story
type Chapter struct {
	Title string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

// Option contains the availabe choices
type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}

// NewHandler manages the story
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r  *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func init() {	
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

var defaultHandlerTemplate = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
    </body>
</html>`
