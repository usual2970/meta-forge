package openapi

import (
	"context"
	"errors"
)

type Next func()

type Middelware func(ctx context.Context, next Next)

var mwMap = map[string]Middelware{
	"preWps":                PreWps,
	"preWpsPdfConvertToDoc": PreWpsPdfConvertToDoc,
	"postWps":               PostWps,
	"log":                   Log,
	"logNotice":             LogNotice,
	"postWpsNotices":        PostWpsNotices,
}

type Pipeline struct {
	stack    []Middelware
	preIndex int
}

func NewPipeline(names ...string) *Pipeline {
	middlewarews := make([]Middelware, 0)

	for _, name := range names {
		mw, ok := mwMap[name]
		if !ok {
			continue
		}

		middlewarews = append(middlewarews, mw)
	}

	return &Pipeline{
		stack: middlewarews,
	}
}

func (p *Pipeline) Use(middlewares ...Middelware) {
	p.stack = append(p.stack, middlewares...)
}

func (p *Pipeline) runner(ctx context.Context, index int) error {
	if index >= len(p.stack) {
		return nil
	}
	if index == p.preIndex {
		return errors.New("next() called multiple times")
	}

	p.preIndex = index

	mw := p.stack[index]

	if mw != nil {
		mw(ctx, func() {
			p.runner(ctx, index+1)
		})
	}
	return nil
}

func (p *Pipeline) Execute(ctx context.Context) error {
	if len(p.stack) == 0 {
		return nil
	}
	p.preIndex = -1

	return p.runner(ctx, 0)
}
