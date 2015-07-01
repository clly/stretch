// Author: Connor Kelly

package stretch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest(host string, endpoint string) []byte {
	url := fmt.Sprintf("%s/%s", host, endpoint)
	result, _ := http.Get(url)
	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
