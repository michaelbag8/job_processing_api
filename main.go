package main

import (
	"context"
)

type Processor interface{
	Process (ctx context.Context, payload string) (string, error)
}

type ReverseProcessor struct{}

func (p ReverseProcessor) Process(ctx context.Context, payload string) (string,error){
	runes := []rune(payload)
	for r, l := 0, len(runes)-1; r < l; r, l = r+1, l-1{
		runes[r], runes[l] = runes[l], runes[r]
	}
	return string(runes), nil
}