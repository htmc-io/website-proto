package markdown

import (
	"testing"
	"util/log"
)

func TestParseIntro(t *testing.T) {
	intro, e := ParseIntro("/ramdisk/microservice.md")
	if e != nil {
		log.E(e)
	}

	log.H("Title: ", intro.Title)
	log.H("Abstract: ", intro.Abstract)
	log.H("Image: ", intro.Image)
}

func TestParseContent(t *testing.T) {
	bytes, e := ParseContent("/ramdisk/microservice.md")
	if e != nil {
		log.E(e)
	}

	log.H(string(bytes))
}
