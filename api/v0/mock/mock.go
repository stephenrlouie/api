package mock

import (
	"encoding/json"
	"io/ioutil"
)

// simple file based responses for web server
// Will contain JSON that has
// StatusCode: int
// Raw: embedded json to parse

func GetMock(path string, body interface{}) (int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		return 0, err
	}

	// Get the status code
	var mock int
	err = json.Unmarshal(*objmap["StatusCode"], &mock)
	if err != nil {
		return 0, err
	}

	// Get the return object
	if body != nil {
		err = json.Unmarshal(*objmap["Raw"], body)
		if err != nil {
			return 0, err
		}
	}
	return mock, nil
}
