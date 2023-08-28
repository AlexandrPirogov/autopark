package autopark

import (
	"encoding/json"
	"log"
)

func unmarshal[T any](body []byte) (T, error) {
	var res T
	err := json.Unmarshal(body, &res)
	log.Printf("unmarshaled into %v", res)
	if err != nil {
		log.Printf("err while unmarshal reequest body %v", err)
		return res, err
	}

	return res, nil
}
