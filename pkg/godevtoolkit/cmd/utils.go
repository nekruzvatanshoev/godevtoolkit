package cmd

import (
	"encoding/json"
	"fmt"
)

func AsJSON(i interface{}) string {

	out, err := json.MarshalIndent(i, "", "   ")
	if err != nil {
		return fmt.Sprintf("%#v", i)
	}
	return string(out)

}
