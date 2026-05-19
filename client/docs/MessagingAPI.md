# \MessagingAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InboxUsername**](MessagingAPI.md#InboxUsername) | **Post** /inbox/{username} | Send a private message



## InboxUsername

> AccountKidPost200Response InboxUsername(ctx, username).Text(text).Execute()

Send a private message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess"
)

func main() {
	username := "someplayer" // string | 
	text := "text_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessagingAPI.InboxUsername(context.Background(), username).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessagingAPI.InboxUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InboxUsername`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MessagingAPI.InboxUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInboxUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **text** | **string** |  | 

### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

