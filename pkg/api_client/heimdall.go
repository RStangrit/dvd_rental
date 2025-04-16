package apiclient

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

func Heimdall(context *gin.Context) {

	backoffInterval := 2 * time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond
	backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)

	retrier := heimdall.NewRetrier(backoff)
	timeout := 1000 * time.Millisecond

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(1),
	)

	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Printf("Error when requesting external API: %v", err)
		context.JSON(500, gin.H{"error": "Unable to fetch data from external API"})
		return
	}

	body, _ := io.ReadAll(res.Body)
	context.JSON(200, gin.H{"response": string(body)})
}
