package zip

import (
	"fmt"
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Reader struct {
	r        *zip.ReadCloser
	password string
}

func OpenReader(source string) (Reader, error) {
	r, err := zip.OpenReader(source)
	if err != nil {
		return Reader{}, err
	}

	z := Reader{
		r: r,
	}
	return z, nil
}

func (r Reader) WithPassword(password string) Reader {
	r.password = password
	return r
}

func (r Reader) Extract(output string) {
	for _, f := range r.r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		if f.IsEncrypted() {
			f.SetPassword(r.password)
		}

		r, err := f.Open()
		if err != nil {
			log.Printf("error opening zipped file: %s\n", err)
			continue
		}

		buf, err := io.ReadAll(r)
		if err != nil {
			log.Printf("error reading zipped file: %s\n", err)
			continue
		}

		path := fmt.Sprintf("%s/%s", output, f.Name)
		_ = os.MkdirAll(filepath.Dir(path), 0)
		if err = os.WriteFile(path, buf, 0); err != nil {
			fmt.Printf("error unzipping file: %s\n", err)
		}
	}
	_ = r.r.Close()
}
