package mock

import (
	"fmt"
)

type DbMock struct {
}

func (*DbMock) Create(i interface{}) error {
	fmt.Println("Created")
	return nil
}

func (*DbMock) Find(i interface{}) {
	fmt.Println("Find")
}

func (*DbMock) Update(i interface{}, id string) error {
	return nil
}

func (*DbMock) Delete(i interface{}, id string) error {
	return nil
}
