package utiles

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
	func ParseBody(r *http.Request, x interface{}) {
		if body, err := ioutil.ReadAll(r.Body); err != nil {
			if err := json.Unmarshal([]byte(body), x); err != nil {
				return
			}
		}
	}
*/
func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
