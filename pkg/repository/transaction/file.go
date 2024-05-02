package transaction

import (
	"atp/payment/pkg/utils/domain"
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"os"
)

func (r repository) LastKey(ctx context.Context) (string, error) {
	file, err := os.OpenFile(r.path, os.O_RDONLY|os.O_RDWR, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastkey string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		var prev domain.File
		err = json.Unmarshal(scanner.Bytes(), &prev)
		if err != nil {
			return "", err
		}
		lastkey = prev.Key
	}

	return lastkey, nil
}

func (r repository) Save(ctx context.Context, key string) error {
	_, err := os.Stat(r.path)
	if err != nil {
		if os.IsNotExist(err) { //File not Found !!
			_, err := os.Create(r.path)
			if err != nil {
				return err
			}
		}
	}

	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}

		var prev domain.File
		err = json.Unmarshal(scanner.Bytes(), &prev)
		if err != nil {
			return err
		}

		if prev.Key == key {
			errN := errors.New("key is exist")
			return errN
		}
	}

	hashkey := domain.File{
		Key: key,
	}
	js, err := json.Marshal(hashkey)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(r.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(js)
	f.Write([]byte("\n"))

	return nil
}
