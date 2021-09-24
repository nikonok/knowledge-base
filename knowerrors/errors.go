package knowerrors

import (
	"golang.org/x/xerrors"
)

var (
	ErrNotImplemented = xerrors.New("not implemented")

	ErrCannotCreateDoc = xerrors.New("connot create doc")

	ErrCannotReadFile = xerrors.New("cannot read file")

	ErrCannotUpdateFile = xerrors.New("cannot update file")

	ErrCannotDeleteFile = xerrors.New("cannot delete file")

	ErrStorageUnknown = xerrors.New("unknown storage error")
)
