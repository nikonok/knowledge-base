package filecontroller_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"knowledge-base/knowlib"
	filecontroller "knowledge-base/knowlib/file_controller"
)

func init() {
	os.Mkdir(testPath, os.ModePerm)
}

const (
	testPath = "./test_storage/"
)

func TestCreateFile(t *testing.T) {
	dir, err := os.MkdirTemp(testPath, "example")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	dc := filecontroller.NewDocController(dir)

	newDocBody := "# Hello World \n\nHello!"
	doc := knowlib.Doc{
		Body: newDocBody,
	}

	id, err := dc.CreateDoc(doc)
	assert.NoError(t, err)

	returnedDoc, err := dc.GetDoc(id)
	assert.NoError(t, err)

	assert.Equal(t, newDocBody, returnedDoc.Body)
}
