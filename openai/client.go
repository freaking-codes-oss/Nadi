package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct { Endpoint, APIKey string; HTTPClient *http.Client }
type Message struct { Role string `json:"role"`; Content string `json:"content,omitempty"` }
type Request struct { Model string `json:"model"`; Messages []Message `json:"messages"`; Stream bool `json:"stream,omitempty"` }
type Response struct { Choices []struct { Message Message `json:"message"` } `json:"choices"` }
func (c Client) Chat(ctx context.Context, model string, messages []Message) (Response, error) {
	body, err := json.Marshal(Request{Model:model, Messages:messages}); if err != nil { return Response{}, err }
	endpoint := strings.TrimRight(c.Endpoint, "/") + "/chat/completions"
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body)); if err != nil { return Response{}, err }
		req.Header.Set("Content-Type", "application/json"); if c.APIKey != "" { req.Header.Set("Authorization", "Bearer "+c.APIKey) }
		client := c.HTTPClient; if client == nil { client = &http.Client{Timeout: 60*time.Second} }
		resp, err := client.Do(req); if err != nil { if ctx.Err()!=nil { return Response{}, ctx.Err() }; if attempt==2{return Response{},err}; time.Sleep(time.Duration(attempt+1)*200*time.Millisecond); continue }
		data, readErr := io.ReadAll(io.LimitReader(resp.Body, 4<<20)); resp.Body.Close()
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 { if attempt<2 { time.Sleep(time.Duration(attempt+1)*200*time.Millisecond); continue } }
		if resp.StatusCode < 200 || resp.StatusCode >= 300 { return Response{}, fmt.Errorf("provider returned HTTP %d", resp.StatusCode) }
		if readErr != nil { return Response{}, readErr }; var out Response; if err=json.Unmarshal(data,&out); err!=nil{return Response{},err}; return out,nil
	}
	return Response{}, fmt.Errorf("provider request failed")
}
