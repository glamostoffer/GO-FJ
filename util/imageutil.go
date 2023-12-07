package util

import (
	"io"
	"mime/multipart"
	"os"
)

type Service struct {
	Destination string
}

func (s *Service) SaveImages(files []*multipart.FileHeader) ([]string, error) {
	paths := make([]string, 0)
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		dst, err := os.Create(s.Destination + file.Filename)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		paths = append(paths, s.Destination+file.Filename)
	}
	return paths, nil
}
