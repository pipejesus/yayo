package Lando

import (
	"bufio"
	"errors"
	"github.com/icza/dyno"
	"io"
	"os"
)

type LandoYamlStructure interface{}

func SetYamlArrayItem(landoStruct LandoYamlStructure, value string, path ...interface{}) (LandoYamlStructure, error) {
	err := dyno.Set(landoStruct, value, path...)

	if err != nil {
		return nil, err
	}

	return landoStruct, nil
}

func ReadLandoYaml(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("nie mogę znaleźć pliku .lando.yml")
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return nil, errors.New("nie mogę odczytać pliku .lando.yml")
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		return nil, errors.New("błąd podczas czytania pliku .lando.yml")
	}

	return bs, nil
}

func WriteLandoYaml(outputFilePath string, landoFileContents []byte) error {
	file, err := os.Create(outputFilePath)
	defer file.Close()

	if err != nil {
		return errors.New("nie mogę znaleźć pliku .lando.yml")
	}

	w := bufio.NewWriter(file)
	_, err = w.Write(landoFileContents)

	if err != nil {
		return err
	}

	err = w.Flush()

	return err
}
