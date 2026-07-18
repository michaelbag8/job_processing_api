package main

import (
	"context"
	"fmt"
	"strings"
)

type Processor interface{
	Process (ctx context.Context, payload string) (string, error)
}

type ReverseProcessor struct{}

func (r ReverseProcessor) Process(ctx context.Context, payload string) (string,error){
	runes := []rune(payload)
	for r, l := 0, len(runes)-1; r < l; r, l = r+1, l-1{
		runes[r], runes[l] = runes[l], runes[r]
	}
	return string(runes), nil
}

type WordCountProcessor struct{}

func (w WordCountProcessor) Process(ctx context.Context, payload string) (string,error){
	words := strings.Fields(payload)
	return fmt.Sprintf("%d", len(words)),nil
}