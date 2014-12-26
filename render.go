package amber

import (
	"github.com/eknkc/amber"
	"github.com/rocwong/neko"
	"html/template"
)

type (
	Options struct {
		// BaseDir represents a base directory of the amber templates.
		BaseDir string
		// Extension represents an extension of files.
		Extension string
		// Setting if pretty printing is enabled.
		// Pretty printing ensures that the output html is properly indented and in human readable form.
		// If disabled, produced HTML is compact. This might be more suitable in production environments.
		// Default: true
		PrettyPrint bool
	}
	amberRenderer struct {
		context *neko.Context
	}
)

var amberCache map[string]*template.Template

func Renderer(options ...*Options) neko.HandlerFunc {
	var opt *Options
	if options != nil {
		opt = options[0]
	} else {
		opt = &Options{BaseDir: "views", Extension: ".amber", PrettyPrint: true}
	}
	amberCache, _ = amber.CompileDir(
		opt.BaseDir,
		amber.DirOptions{Ext: opt.Extension, Recursive: true},
		amber.Options{PrettyPrint: opt.PrettyPrint, LineNumbers: false},
	)

	return func(ctx *neko.Context) {
		ctx.HtmlEngine = &amberRenderer{context: ctx}
	}
}

func (c *amberRenderer) Render(view string, context interface{}, status ...int) (err error) {
	code := 200
	if status != nil {
		code = status[0]
	}
	c.context.Writer.WriteHeader(code)
	err = amberCache[view].Execute(c.context.Writer, context)
	return
}
