package api

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

/***********************************************************************************************
* As of now, there is no way to clear url aggregate data unless we implement clear data api.   *
* Clubbing these tests with other API tests doesn't provide clean state. This makes writing    *
* predictable and reproducible tests difficult. Hence, writing metrics api tests separately.   *
************************************************************************************************/

func TestMetricHandler(t *testing.T) {
	t.Run("metrics with empty state returns no content response", func(t *testing.T) {
		w := executeMetricRequest()
		if w.Code != http.StatusNoContent {
			t.Errorf("Received wrong status code: %v, expected: %v", w.Code, http.StatusNoContent)
		}
	})

	t.Run("metric with single converted domain", func(t *testing.T) {
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://google.com/1"))
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://google.com/2"))
		w := executeMetricRequest()
		if w.Code != http.StatusOK {
			t.Errorf("Received wrong status code: %v, expected: %v", w.Code, http.StatusOK)
			t.FailNow()
		}
		var response map[string]int
		responseBytes, _ := io.ReadAll(w.Body)
		err := json.Unmarshal(responseBytes, &response)
		if err != nil {
			t.Errorf("possibly received malformed response: %v", err)
			t.FailNow()
		}
		got := response["google.com"]
		want := 2
		assertCorrectMessage(t, got, want)
	})

	t.Run("metric with multiple converted domains", func(t *testing.T) {
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://google.com/1"))
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://wiki.com/1"))
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://google.com/2"))
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://wiki.com/2"))
		_ = executeShortenUrlRequest(getShortenUrlRequestBody("http://wiki.com/3"))

		w := executeMetricRequest()
		assertCorrectStatusCode(t, w.Code, http.StatusOK)
		if t.Failed() {
			t.FailNow()
		}
		var response map[string]int
		responseBytes, _ := io.ReadAll(w.Body)
		err := json.Unmarshal(responseBytes, &response)
		if err != nil {
			t.Errorf("possibly received malformed response: %v", err)
			t.FailNow()
		}
		got := response["wiki.com"]
		want := 3
		assertCorrectMessage(t, got, want)
	})
}
