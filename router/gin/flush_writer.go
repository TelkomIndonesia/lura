package gin

import (
	"github.com/gin-gonic/gin"
)

type flushWriter struct {
	gin.ResponseWriter
}

func newFlushWriter(w gin.ResponseWriter) *flushWriter {
	return &flushWriter{w}
}

func (fw flushWriter) Write(p []byte) (int, error) {
	n, err := fw.ResponseWriter.Write(p)
	fw.Flush()
	return n, err
}
