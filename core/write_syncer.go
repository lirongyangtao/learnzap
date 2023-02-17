package core

import (
	"io"
)

type WriteSyncer interface {
	io.Writer
	Sync() error
}

func AddSyncer(w io.Writer) WriteSyncer {
	switch t := w.(type) {
	case WriteSyncer:
		return t
	}
	return &WrapperSyncer{Writer: w}
}

type WrapperSyncer struct {
	io.Writer
}

func (w *WrapperSyncer) Sync() error {
	return nil
}

type MulWriteSyncer []WriteSyncer

func (ws *MulWriteSyncer) Sync() error {
	var errs []error
	for _, v := range *ws {
		err := v.Sync()
		if err != nil {
			errs = append(errs, err)
			panic(err)
		}
	}
	// TODO merge errrs
	return nil
}

func (ws *MulWriteSyncer) Write(p []byte) (int, error) {
	var errs []error
	nWritten := 0
	for _, w := range *ws {
		n, err := w.Write(p)
		if err != nil {
			panic(err)
			errs = append(errs, err)
		}
		if nWritten == 0 && n != 0 {
			nWritten = n
		} else if n < nWritten {
			nWritten = n
		}
	}
	return nWritten, nil
}
