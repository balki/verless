package pipeline

import (
	"github.com/verless/verless/builder"
	"github.com/verless/verless/parser"
	"github.com/verless/verless/plugin"
	"github.com/verless/verless/writer"
)

const (
	parallelism = 4
)

type Context struct {
	files   <-chan string
	parser  parser.Parser
	builder builder.Builder
	writer  writer.Writer
}

type Pipeline interface {
	AddPlugin(plugin plugin.Plugin)
	Run(ctx Context) error
}

func New() Pipeline {
	p := pipeline{
		plugins: make([]plugin.Plugin, 0),
	}
	return &p
}

type pipeline struct {
	plugins []plugin.Plugin
}

func (p *pipeline) AddPlugin(plugin plugin.Plugin) {
	p.plugins = append(p.plugins, plugin)
}

func (p *pipeline) Run(ctx Context) error {
	err := runParallel(func(file string) error {
		return nil
	}, ctx.files, parallelism)

	if err != nil {
		return err
	}

	site, err := ctx.builder.Dispatch()
	if err != nil {
		return err
	}

	for _, _plugin := range p.plugins {
		if err := _plugin.Finalize(); err != nil {
			return err
		}
	}

	if err := ctx.writer.Write(site); err != nil {
		return err
	}

	return nil
}
