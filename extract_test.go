package ziply

import (
	"github.com/restartfu/ziply/rar"
	"github.com/restartfu/ziply/zip"
	"log"
	"testing"
)

func TestZipExtract(t *testing.T) {
	z, err := zip.OpenReader("assets/extract_test.zip")
	if err != nil {
		log.Fatalln(err)
	}

	z = z.WithPassword("test")
	z.Extract("assets/extract_result_zip")
}

func TestRarExtract(t *testing.T) {
	r, err := rar.OpenReader("assets/extract_test.rar", "test")
	if err != nil {
		log.Fatalln(err)
	}

	r.Extract("assets/extract_result_rar")
}
