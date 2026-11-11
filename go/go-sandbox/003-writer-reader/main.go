package main

import (
	"errors"
	"io"
)

func ReadAll(r io.Reader) ([]byte, error) {
	const CHUNK_SIZE = 4

	chunk := make([]byte, CHUNK_SIZE)
	res := make([]byte, 0)

	for {
		n, err := r.Read(chunk)

		res = append(res, chunk[:n]...)

		if err != nil {
			if errors.Is(err, io.EOF) {
				return res, nil
			}
			return res, err
		}

		if n == 0 {
			return res, nil
		}
	}
}

type LimitedWriter interface {
	io.Writer
}

type limitedWriter struct {
	w                io.Writer
	limit            int
	writenBytesCount int
}

func NewLimitedWriter(w io.Writer, limit int) LimitedWriter {
	return &limitedWriter{
		w:     w,
		limit: limit,
	}
}

var WriteLimitError = errors.New("write limit achived")

func (lw *limitedWriter) Write(p []byte) (n int, err error) {
	pLen := len(p)

	if lw.writenBytesCount+pLen > lw.limit {
		bytesToWrite := lw.limit - lw.writenBytesCount
		n, err := lw.w.Write(p[:bytesToWrite])
		lw.writenBytesCount += n

		if err != nil {
			return n, err
		}

		return n, WriteLimitError
	}

	n, err = lw.w.Write(p)
	lw.writenBytesCount += n

	return n, err
}
