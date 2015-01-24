package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"strings"
)

type Repository struct {
	Packages []Package
}

func OpenRepository(r io.Reader) (db *Repository, err error) {

	var (
		reader *gzip.Reader
		header *tar.Header
	)

	if reader, err = gzip.NewReader(r); err != nil {
		return
	}

	archive := tar.NewReader(reader)

	db = new(Repository)
	db.Packages = make([]Package, 0, 10)

	for {
		if header, err = archive.Next(); err != nil {
			return
		}

		if strings.HasSuffix(header.Name, "/desc") {

			pkg := NewPackage()

			if err = pkg.Extract(archive); err != nil {
				return
			}

			db.Packages = append(db.Packages, pkg)
		}
	}

	return
}
