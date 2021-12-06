package api

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	l "github.com/cchaiyatad/mss/internal/layout"
	"github.com/cchaiyatad/mss/internal/utils"

	"github.com/go-playground/assert"
)

func getRespAndBodyGivenHandlerFunc(handleFun func(http.ResponseWriter, *http.Request), param string, t *testing.T) (*http.Response, []byte) {
	server := httptest.NewServer(http.HandlerFunc(handleFun))
	defer server.Close()

	url := fmt.Sprintf("%s?%s", server.URL, param)

	resp, err := http.Get(url)
	assert.Equal(t, nil, err)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, nil, err)
	resp.Body.Close()

	return resp, body
}

func TestMain(m *testing.M) {
	os.Setenv("LAYOUT_PATH", "./layout_test/")
	l.InitLayout()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateAPIController(t *testing.T) {
	t.Run("CreateAPIController ", func(t *testing.T) {
		os.Args = []string{"app", "-port", "1234"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		gotController := CreateAPIController()
		expectedPort := "1234"

		assert.Equal(t, expectedPort, gotController.args.Port)
	})
}
func TestGetPort(t *testing.T) {
	t.Run("GetPort ", func(t *testing.T) {
		givenArgs := &utils.Args{Port: "1234"}
		expectedPort := ":1234"

		assert.Equal(t, expectedPort, getPort(givenArgs))
	})
}
func TestLayoutEndpoint(t *testing.T) {
	t.Run("getLayoutsMessage", func(t *testing.T) {
		expected := []byte(`["no-tile"]`)

		assert.Equal(t, expected, getLayoutsOptionJSON())
	})

	t.Run("layout endpoint", func(t *testing.T) {
		gotResp, gotBody := getRespAndBodyGivenHandlerFunc(layout, "", t)

		expectedBody := []byte(`["no-tile"]`)
		expectedStatusCode := http.StatusOK
		fmt.Println(gotBody)

		assert.Equal(t, expectedStatusCode, gotResp.StatusCode)
		assert.Equal(t, expectedBody, gotBody)
	})
}

func TestSolveEndpoint(t *testing.T) {

	t.Run("solve endpoint when everything satisfy", func(t *testing.T) {
		givenParam := "strategy=random&heuristic=random&layout=no-tile"
		gotResp, gotBody := getRespAndBodyGivenHandlerFunc(solve, givenParam, t)

		expectedBody := []byte(`{"board":{"tiles":[],"size":{"x_size":8,"y_size":8,"z_size":2},"layout":"no-tile"}, "order":[], "params":{"strategy":"random","heuristic":"random","layout":"no-tile"}}`)
		expectedStatusCode := http.StatusOK

		assert.Equal(t, expectedStatusCode, gotResp.StatusCode)
		assert.Equal(t, expectedBody, gotBody)
	})

	t.Run("solve endpoint when param is not valid", func(t *testing.T) {
		cases := []struct {
			givenLayout    string
			givenStrategy  string
			givenHeuristic string
		}{
			{"no-tile", "random", "not-valid"},
			{"no-tile", "not-valid", "random"},
			{"not-valid", "random", "random"},
			{"no-tile", "not-valid", "not-valid"},
			{"not-valid", "not-valid", "random"},
			{"not-valid", "random", "not-valid"},
			{"not-valid", "not-valid", "not-valid"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("request when layout:%s, strategy:%s givenHeuristic: %s", tc.givenLayout, tc.givenStrategy, tc.givenHeuristic), func(t *testing.T) {
				t.Parallel()
				given := fmt.Sprintf("strategy=%s&heuristic=%s&layout=%s", tc.givenStrategy, tc.givenHeuristic, tc.givenLayout)

				gotResp, gotBody := getRespAndBodyGivenHandlerFunc(solve, given, t)

				expectedBody := []byte(http.StatusText(http.StatusBadRequest) + "\n")
				expectedStatusCode := http.StatusBadRequest

				assert.Equal(t, expectedStatusCode, gotResp.StatusCode)
				assert.Equal(t, expectedBody, gotBody)
			})
		}
	})
	t.Run("solve endpoint when param is missing", func(t *testing.T) {
		cases := []struct {
			givenParam string
		}{
			{"strategy=random&heuristic=random"},
			{"strategy=random&layout=no-tile"},
			{"heuristic=random&layout=no-tile"},
			{"strategy=random"},
			{"layout=no-tile"},
			{"heuristic=random"},
			{""},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("request when param:%s", tc.givenParam), func(t *testing.T) {
				t.Parallel()
				gotResp, gotBody := getRespAndBodyGivenHandlerFunc(solve, tc.givenParam, t)

				expectedBody := []byte(http.StatusText(http.StatusBadRequest) + "\n")
				expectedStatusCode := http.StatusBadRequest

				assert.Equal(t, expectedStatusCode, gotResp.StatusCode)
				assert.Equal(t, expectedBody, gotBody)
			})
		}
	})

}
