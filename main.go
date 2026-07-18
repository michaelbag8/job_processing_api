package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Processor interface {
	Process(ctx context.Context, payload string) (string, error)
}

type ReverseProcessor struct{}

func (r ReverseProcessor) Process(ctx context.Context, payload string) (string, error) {
	runes := []rune(payload)
	for r, l := 0, len(runes)-1; r < l; r, l = r+1, l-1 {
		runes[r], runes[l] = runes[l], runes[r]
	}
	return string(runes), nil
}

type WordCountProcessor struct{}

func (w WordCountProcessor) Process(ctx context.Context, payload string) (string, error) {
	words := strings.Fields(payload)
	return fmt.Sprintf("%d", len(words)), nil
}

type HashProcessor struct{}

func (h HashProcessor) Process(ctx context.Context, payload string) (string, error) {
	data := sha256.Sum256([]byte(payload))
	hashstring := hex.EncodeToString(data[:])
	return hashstring, nil
}

var processorRegistry = map[string]Processor{
	"reverse": ReverseProcessor{},
	"wordcount": WordCountProcessor{},
	"hash": HashProcessor{},
}

func GetProcessor(jobType string) (Processor, error) {
	value, ok := processorRegistry[jobType]
	if !ok {
		return nil, fmt.Errorf("%s doesn't exist", jobType)
	}
	
	return value, nil
}
