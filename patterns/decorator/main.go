package main

import "fmt"

type DB interface {
	Save(string) error
}

type Store struct{}

func (s *Store) Save(data string) error {
	fmt.Println("saving data: ", data)
	return nil
}

func myExecuteFn(db DB) ExecuteFn {
	fmt.Println("testing func:", db)
	return func(data string) error {
		fmt.Println("before saving data: ", data)
		return db.Save(data)
	}
}

func main() {
	store := &Store{}
	decoratedFn := myExecuteFn(store)
	Execute(decoratedFn)
}

type ExecuteFn func(string) error

func Execute(fn ExecuteFn) {
	err := fn("Hi there")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
