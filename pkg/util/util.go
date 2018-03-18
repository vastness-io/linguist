package util

import (
	"io"
	"os"
)

func GetFileContents(bufferSize int32, filePath string) ([]byte, error) {

	buf := make([]byte, bufferSize)

	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	_, err = f.Read(buf)

	if err = f.Close(); err != nil {
		return nil, err
	}

	if err != io.EOF {
		return nil, err
	}

	return buf, nil
}
