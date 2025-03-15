package mydict

import (
	"fmt"
)

// Errorf 는 오류 메시지를 만들 때 사용하는 함수
// errors.New 에 format 을 사용할 수 있게 해준다.
// 아래 처럼 할 필요 없음
// errors.New(fmt.Sprintf("not found %s", word))
func makeErrorNotFound(word string) error {
	return fmt.Errorf("not found %s", word)
}

func makeErrorWordExists(word string) error {
	return fmt.Errorf("%s is already exists", word)
}

var errorCanNotUpdate = fmt.Errorf("can not update non-existing word")

var errorCanNotDelete = fmt.Errorf("can not delete non-existing word")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	value, exists := d[word]
  if(exists) {
		return value, nil
	}

	return "", makeErrorNotFound(word)
}

func (d Dictionary) Add(key, value string) error {
  _, existsError := d.Search(key)

  if existsError == nil {
		return makeErrorWordExists(key)
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, existsError := d.Search(key)

	if existsError != nil {
		return errorCanNotUpdate
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, existsError := d.Search(key)

	if existsError != nil {
		return errorCanNotDelete
	}
	
	delete(d, key)
	return nil
}

