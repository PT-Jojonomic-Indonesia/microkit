// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/jojocoders/microkit/tracer"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	url := os.Getenv("JAEGER_ENDPOINT")

	tracer.InitOtel(url, "Sum Caller 2", "v1.0.0", "development")

	ctx := context.TODO()
	_, span := tracer.Start(ctx, "Initiate Connection", "", "")
	span.End()

	ctx = context.TODO()
	ctx2, span2 := tracer.Start(ctx, "Call Sum", "", "")

	payload := struct {
		Number1 float64 `json:"number1"`
		Number2 float64 `json:"number2"`
	}{
		Number1: 23.3,
		Number2: 21.4,
	}
	j, _ := json.Marshal(payload)

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Make sure you pass the context to the request to avoid broken traces.
	req, err := http.NewRequestWithContext(ctx2, "POST", "http://localhost:8181/sum?trace_id="+span2.SpanContext().TraceID().String()+"&span_id="+span2.SpanContext().SpanID().String(), bytes.NewBuffer(j))
	if err != nil {
		panic(err)
	}

	// All requests made with this client will create spans.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	resp.Body.Close()
	log.Println(string(body))

	span2.End()

	time.Sleep(10 * time.Second)

}
