package aws4

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Example_jSONBody() {
	data := strings.NewReader("{}")
	r, _ := http.NewRequest("POST", "https://dynamodb.us-east-1.amazonaws.com/", data)
	r.Header.Set("Content-Type", "application/x-amz-json-1.0")
	r.Header.Set("X-Amz-Target", "DynamoDB_20111205.ListTables")

	resp, err := DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	// Output:
	// 200
}

func Example_formEncodedBody() {
	v := make(url.Values)
	v.Set("Action", "DescribeAutoScalingGroups")

	resp, err := PostForm("https://autoscaling.us-east-1.amazonaws.com/", v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
	// Output:
	// 200
}
