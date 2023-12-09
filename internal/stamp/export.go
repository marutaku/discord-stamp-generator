package stamp

import (
	"os"
)

func Export(imgBytes []byte, filePath string) error {
	distFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer distFile.Close()
	_, err = distFile.Write(imgBytes)
	if err != nil {
		return err
	}
	return nil
}
