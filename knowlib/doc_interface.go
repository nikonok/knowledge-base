package knowlib

type Doc struct {
	ID     uint64
	Header string
	Body   string
}

type DocController interface {
	CreateDoc(doc Doc) (uint64, error)
	GetDoc(id uint64) (*Doc, error)
	UpdateDoc(id uint64, doc Doc) error
	DeleteDoc(id uint64, doc Doc) error
}
