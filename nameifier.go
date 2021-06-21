package main

import (
	_ "embed"
	"encoding/json"
	"hash/fnv"

	log "github.com/sirupsen/logrus"
)

//go:embed data/nouns.json
var nouns []byte

//go:embed data/adjectives.json
var adjectives []byte

type NameGenerator interface {
	Nameify(seed string) (string, error)
}

func hash(s string, max int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % uint32(max)
}

func loadJson(file []byte, nameifier *Nameifier) {
	err := json.Unmarshal(file, &nameifier)
	if err != nil {
		log.Errorf("Eror parsing json from file: %s", err)
	}
}

type Nameifier struct {
	Nouns      []string
	Adjectives []string
}

func NewNameifier() *Nameifier {
	n := &Nameifier{}
	loadJson(nouns, n)
	loadJson(adjectives, n)
	return n
}

func (n *Nameifier) Nameify(seed string) (string, error) {
	return n.Adjectives[hash(seed, len(n.Adjectives))] + "-" + n.Nouns[hash(seed, len(n.Nouns))], nil
}
