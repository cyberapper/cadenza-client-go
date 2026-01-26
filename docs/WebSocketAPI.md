# \WebSocketAPI

All URIs are relative to *https://cadenza-api-uat.algo724.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebsocketConnect**](WebSocketAPI.md#WebsocketConnect) | **Post** /connection/websocket | WebSocket Connection



## WebsocketConnect

> WebsocketConnect(ctx).WsCommand(wsCommand).Execute()

WebSocket Connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cyberapper/cadenza-client-go"
)

func main() {
	wsCommand := *openapiclient.NewWsCommand(int32(1)) // WsCommand | WebSocket command (sent after WebSocket connection is established) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebSocketAPI.WebsocketConnect(context.Background()).WsCommand(wsCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSocketAPI.WebsocketConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebsocketConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wsCommand** | [**WsCommand**](WsCommand.md) | WebSocket command (sent after WebSocket connection is established) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

