package filecontroller_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"knowledge-base/knowerrors"
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

func TestCRUD(t *testing.T) {
	dir, err := os.MkdirTemp(testPath, "example")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	dc := filecontroller.NewDocController(dir)

	docBody := "# Hello World \n\nHello!"
	doc := knowlib.Doc{
		Body: docBody,
	}

	id, err := dc.CreateDoc(doc)
	assert.NoError(t, err)

	returnedDoc, err := dc.GetDoc(id)
	assert.NoError(t, err)

	assert.Equal(t, docBody, returnedDoc.Body)

	newDocBody := "# NOT THE SAME!"
	newDoc := knowlib.Doc{
		ID:   id,
		Body: newDocBody,
	}

	err = dc.UpdateDoc(newDoc)
	assert.NoError(t, err)

	returnedDoc, err = dc.GetDoc(id)
	assert.NoError(t, err)

	assert.Equal(t, newDocBody, returnedDoc.Body)

	err = dc.DeleteDoc(id)
	assert.NoError(t, err)

	_, err = dc.GetDoc(id)
	assert.Equal(t, knowerrors.ErrCannotReadFile, err)
}
