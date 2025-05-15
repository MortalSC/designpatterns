package functionaloptions

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_HTTPClient(t *testing.T) {
	assert := assert.New(t)

	client := NewHTTPClient(WithTimeout(5 * time.Second))

	resp, err := client.Get("http://example.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	assert.Equal(404, resp.StatusCode, "期望状态码 404")
}
