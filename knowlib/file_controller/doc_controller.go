package filecontroller

import (
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"knowledge-base/knowerrors"
	"knowledge-base/knowlib"
)

const (
	markdown = ".md"
)

type fileDocController struct {
	path   string
	nextId uint64
}

func findNextId(path string) uint64 {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return uint64(len(files))
}

func NewDocController(path string) knowlib.DocController {
	return &fileDocController{
		path:   path,
		nextId: findNextId(path),
	}
}

func (fc *fileDocController) getFilePath(id uint64) string {
	return fc.path + "/" + strconv.FormatUint(id, 10) + markdown
}

func (fc *fileDocController) CreateDoc(doc knowlib.Doc) (uint64, error) {
	log.Info("Creating file with id = ", fc.nextId)

	err := ioutil.WriteFile(fc.getFilePath(fc.nextId), []byte(doc.Body), 0644)
	if err != nil {
		log.Error(err)
		return 0, knowerrors.ErrCannotCreateDoc
	}
	fc.nextId++
	return fc.nextId - 1, nil
}

func (fc *fileDocController) GetDoc(id uint64) (*knowlib.Doc, error) {
	log.Info("Getting file with id = ", id)

	file, err := ioutil.ReadFile(fc.getFilePath(id))
	if err != nil {
		log.Error(err)
		return nil, knowerrors.ErrCannotReadFile
	}
	return &knowlib.Doc{
		ID:   id,
		Body: string(file),
	}, nil
}

func (fc *fileDocController) UpdateDoc(doc knowlib.Doc) error {
	log.Info("Updating file with id = ", doc.ID)

	err := ioutil.WriteFile(fc.getFilePath(doc.ID), []byte(doc.Body), 0644)
	if err != nil {
		log.Error(err)
		return knowerrors.ErrCannotUpdateFile
	}
	return nil
}

func (fc *fileDocController) DeleteDoc(id uint64) error {
	log.Info("Deleting file with id = ", id)

	err := os.Remove(fc.getFilePath(id))
	if err != nil {
		log.Error(err)
		return knowerrors.ErrCannotDeleteFile
	}
	return nil
}
