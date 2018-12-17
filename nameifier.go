package main

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	NounsFilename  = "data/nouns.json"
	AdjectivesFile = "data/adjectives.json"
)

type NameGenerator interface {
	Nameify(seed string) (string, error)
}

func hash(s string, max int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % uint32(max)
}

func loadJson(filename string, nameifier *Nameifier) {
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err.Error())
		return
	}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("Can't read the nouns file! (%s) %s", filename, err)
	}

	err = json.Unmarshal(raw, &nameifier)
	if err != nil {
		log.Errorf("Eror parsing json from the nouns file '%s': %s", filename, err)
	}
}

type Nameifier struct {
	Nouns      []string
	Adjectives []string
	BasePath   string
}

func NewNameifier(basePath string) *Nameifier {
	n := &Nameifier{
		BasePath: basePath,
	}
	loadJson(n.BasePath+"/"+NounsFilename, n)
	loadJson(n.BasePath+"/"+AdjectivesFile, n)
	return n
}

func (n *Nameifier) Nameify(seed string) (string, error) {
	return n.Adjectives[hash(seed, len(n.Adjectives))] + "-" + n.Nouns[hash(seed, len(n.Nouns))], nil
}
