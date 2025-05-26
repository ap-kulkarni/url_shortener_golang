package url_shortner

import (
	"io"
	"net/http"
)

func ShortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "https://exm.pl/abc123")
}

func LongUrlHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "https://google.com/search?q=fluffy+cat")
}

func MetricHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "{\"youtube\" : 1}")
}
