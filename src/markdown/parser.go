package markdown

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"util/log"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

/*
* ---
* title: This Is Title
* abstract: Abstract to this article
* image:/img/cover.jpg
* ---
* MD content
 */

// ParseContent is to parse markdown content and return a html byte[]
func ParseContent(path string) ([]byte, error) {
	intro, e := ParseIntro(path)
	if e != nil {
		return nil, e
	}

	_, content, e := splitMD(path)
	if e != nil {
		return nil, e
	}

	title := "# " + intro.Title + "\n\n"

	content = []byte(title + string(content))

	return blackfriday.Run(content), nil
}

func ParseIntro(path string) (*MDDocIntro, error) {
	intro := MDDocIntro{"", "", ""}

	rd, e := NewReader(path)
	if e != nil {
		return nil, e
	}

	var flag bool
	var line string
	for {
		line, e = rd.Readln()
		if e == io.EOF {
			break
		}

		if e != nil {
			return nil, e
		}

		if line == "---" {
			if !flag { // first '---''
				flag = true
				continue
			} else {
				break
			}
		}

		array := split2words(line, ":")
		if len(array) != 2 {
			return nil, errors.New("Error line: " + line)
		}

		name := strings.ToLower(array[0])
		switch name {
		case "title":
			intro.Title = array[1]
		case "abstract":
			intro.Abstract = array[1]
		case "image":
			intro.Image = array[1]
		default:
			log.E("not support property: ", array[0])
		}
	}

	return &intro, nil
}

func split2words(line, seperator string) []string {
	index := strings.Index(line, seperator)
	if index == -1 {
		return nil
	}

	ret := make([]string, 2)
	ret[0], ret[1] = line[:index], line[index+1:]
	ret[1] = strings.Trim(ret[1], " ")
	ret[1] = strings.Trim(ret[1], "\"")
	return ret
}

// ---
// Intro       <--- first return intro byte[]
// ---
// MD content  <--- second return content byte[]
func splitMD(path string) ([]byte, []byte, error) {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, nil, e
	}

	content := string(bytes)
	parts := strings.Split(content, "---\n")

	if len(parts) != 3 {
		return nil, nil, errors.New("Failed in split markdown file: " + path)
	}
	return []byte(parts[1]), []byte(parts[2]), nil
}
