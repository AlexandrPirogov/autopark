package autopark

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func unmarshal[T any](body []byte) (T, error) {
	var res T
	err := json.Unmarshal(body, &res)
	log.Warn().Msgf("unmarshaled into %v", res)
	if err != nil {
		log.Warn().Msgf("err while unmarshal reequest body %v", err)
		return res, err
	}

	return res, nil
}
