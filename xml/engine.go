package xml

import (
	"embed"
	"io"
	"io/fs"
	"sync"
	"text/template"
)

// Engine motor de templates XML
type Engine struct {
	templates *template.Template
	cache     map[string]*template.Template
	mu        sync.RWMutex
}

// NewEngine crea un nuevo engine desde un filesystem embebido
func NewEngine(fsys embed.FS, root string) (*Engine, error) {
	tmpl := template.New("").Funcs(TemplateFunctions())

	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && isTemplateFile(path) {
			content, err := fs.ReadFile(fsys, path)
			if err != nil {
				return err
			}

			_, err = tmpl.New(d.Name()).Parse(string(content))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Engine{
		templates: tmpl,
		cache:     make(map[string]*template.Template),
	}, nil
}

// NewEngineFromString crea un engine desde strings (para testing)
func NewEngineFromString(templates map[string]string) (*Engine, error) {
	tmpl := template.New("").Funcs(TemplateFunctions())

	for name, content := range templates {
		_, err := tmpl.New(name).Parse(content)
		if err != nil {
			return nil, err
		}
	}

	return &Engine{
		templates: tmpl,
		cache:     make(map[string]*template.Template),
	}, nil
}

// Render renderiza un template
func (e *Engine) Render(w io.Writer, name string, data interface{}) error {
	tmpl := e.templates.Lookup(name)
	if tmpl == nil {
		return ErrTemplateNotFound
	}

	return tmpl.Execute(w, data)
}

// RenderString renderiza un template a string
func (e *Engine) RenderString(name string, data interface{}) (string, error) {
	var buf []byte
	w := &bytesWriter{buf: &buf}

	err := e.Render(w, name, data)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// isTemplateFile verifica si un archivo es un template
func isTemplateFile(path string) bool {
	return len(path) > 5 && path[len(path)-5:] == ".tmpl"
}

// bytesWriter wrapper para escribir a []byte
type bytesWriter struct {
	buf *[]byte
}

func (w *bytesWriter) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}
