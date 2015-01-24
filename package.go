package main

import (
	"bufio"
	"io"
	"strconv"
	"time"
)

type Package map[string]string

func NewPackage() Package {
	return make(Package)
}

func (p Package) Name() string {
	return p["NAME"]
}

func (p Package) Description() string {
	return p["DESC"]
}

func (p Package) Architecture() string {
	return p["ARCH"]
}

func (p Package) Version() string {
	return p["VERSION"]
}

func (p Package) BuildDate() time.Time {
	text, available := p["BUILDDATE"]
	if !available {
		return time.Time{}
	}
	stamp, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.Unix(stamp, 0)
}

func (p Package) Packager() string {
	return p["PACKAGER"]
}

func (p Package) Extract(r io.Reader) (err error) {
	scanner := bufio.NewScanner(r)

	var section string

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case len(line) <= 0:
			continue
		case line[0] == '%' && line[len(line)-1] == '%':
			section = line[1 : len(line)-1]
		default:
			p[section] = line
		}
	}

	return scanner.Err()
}
