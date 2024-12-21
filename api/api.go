package api

import (
	"fmt"
	"time"
)

func GetDefinition(provider DictionaryProvider, query string) (Word, error) {
	result := make(chan Word, 1)
	error := make(chan error, 1)
	go func() {
		res, err := provider.Fetch(query)
		if err != nil {
			error <- err
		}
		result <- res
		error <- nil
	}()
	select {
	case res := <-result:
		return res, <-error
	case <-time.After(30 * time.Second):
		return <-result, fmt.Errorf("request timeout")
	}
}
