package plain_http_client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestDoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"hoge\":\"fuga\"}")
	}))
	defer ts.Close()

	ctx := context.Background()
	var ret map[string]interface{}
	err := DoRequest(ctx, "GET", ts.URL, nil, &ret)
	assert.NoError(t, err)
	assert.Equal(t, "fuga", ret["hoge"])
}
