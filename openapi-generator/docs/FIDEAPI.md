# \FIDEAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FidePlayerGet**](FIDEAPI.md#FidePlayerGet) | **Get** /api/fide/player/{playerId} | Get a FIDE player
[**FidePlayerRatings**](FIDEAPI.md#FidePlayerRatings) | **Get** /api/fide/player/{playerId}/ratings | Get ratings history of a FIDE player
[**FidePlayerSearch**](FIDEAPI.md#FidePlayerSearch) | **Get** /api/fide/player | Search FIDE players



## FidePlayerGet

> FIDEPlayer FidePlayerGet(ctx, playerId).Execute()

Get a FIDE player



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	playerId := int32(56) // int32 | The FIDE player ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDEAPI.FidePlayerGet(context.Background(), playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDEAPI.FidePlayerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FidePlayerGet`: FIDEPlayer
	fmt.Fprintf(os.Stdout, "Response from `FIDEAPI.FidePlayerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **int32** | The FIDE player ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFidePlayerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FIDEPlayer**](FIDEPlayer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FidePlayerRatings

> FIDEPlayerRatings FidePlayerRatings(ctx, playerId).Execute()

Get ratings history of a FIDE player



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	playerId := int32(56) // int32 | The FIDE player ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDEAPI.FidePlayerRatings(context.Background(), playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDEAPI.FidePlayerRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FidePlayerRatings`: FIDEPlayerRatings
	fmt.Fprintf(os.Stdout, "Response from `FIDEAPI.FidePlayerRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **int32** | The FIDE player ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFidePlayerRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FIDEPlayerRatings**](FIDEPlayerRatings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FidePlayerSearch

> []FIDEPlayer FidePlayerSearch(ctx).Q(q).Execute()

Search FIDE players



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	q := "Erigaisi Arjun" // string | The search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDEAPI.FidePlayerSearch(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDEAPI.FidePlayerSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FidePlayerSearch`: []FIDEPlayer
	fmt.Fprintf(os.Stdout, "Response from `FIDEAPI.FidePlayerSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFidePlayerSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The search query. | 

### Return type

[**[]FIDEPlayer**](FIDEPlayer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

