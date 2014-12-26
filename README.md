#Amber Render
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/neko-contrib/amber)

[Amber](https://github.com/eknkc/amber) template for [Neko](https://github.com/rocwong/neko).

## Usage
~~~go
  package main
  
  import (
      "github.com/neko-contrib/amber"
      "github.com/rocwong/neko"
  )
  app := neko.Classic()
  app.Use(amber.Renderer(&amber.Options{BaseDir: "core/views", Extension: ".amber", PrettyPrint: true}))
  app.Run(":3000")
~~~

####type Options
~~~go
  type Options struct {
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
~~~


