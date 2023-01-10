package debugging

import (
	"log"
	"net/http"
	"os"

	"moul.io/http2curl"
)

func LogCurl(request *http.Request) {
	isDebugging := os.Getenv("DEBUG") == "true"
	if !isDebugging {
		return
	}

	curl, _ := http2curl.GetCurlCommand(request)
	log.Println(curl.String())
}
