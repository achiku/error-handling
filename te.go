package main

import "github.com/pkg/errors"

func doSomethingWithNil() error {
	return errors.Wrap(nil, "hogehoge")
}

func handleError() error {
	err := doSomethingWithNil()
	if err != nil {
		return err
	}
	return nil
}

func handleNormal() error {
	return errors.New("erererere")
}
