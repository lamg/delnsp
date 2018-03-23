package main

import (
	"flag"
	"os"
	"path"
	"strings"
)

func main() {
	var dir string
	flag.StringVar(&dir, "d", "",
		"Directory with files to rename")
	flag.Parse()

	f, e := os.Open(dir)
	var sl []string
	if e == nil {
		sl, e = f.Readdirnames(-1)
	}
	for i := 0; e == nil && i != len(sl); i++ {
		nn := newName(sl[i])
		if nn != sl[i] {
			old, nw := path.Join(dir, sl[i]), path.Join(dir, nn)
			e = os.Rename(old, nw)
		}
	}
	if e != nil {
		os.Stderr.WriteString(e.Error() + "\n")
	}
}

func newName(s string) (r string) {
	r = strings.Replace(s, "： ", "：", 1)
	return
}

/*
Graph Theory： 06 Sum of Degrees is ALWAYS Twice the Number of Edges.mp4
Graph Theory：01. Seven Bridges of Konigsberg.mp4
*/
