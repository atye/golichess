# \TablebaseAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AntichessAtomic**](TablebaseAPI.md#AntichessAtomic) | **Get** /antichess | Tablebase lookup for Antichess
[**TablebaseAtomic**](TablebaseAPI.md#TablebaseAtomic) | **Get** /atomic | Tablebase lookup for Atomic chess
[**TablebaseStandard**](TablebaseAPI.md#TablebaseStandard) | **Get** /standard | Tablebase lookup



## AntichessAtomic

> TablebaseStandard200Response AntichessAtomic(ctx).Fen(fen).Execute()

Tablebase lookup for Antichess



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
	fen := "fen_example" // string | X-FEN of the position. Underscores allowed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablebaseAPI.AntichessAtomic(context.Background()).Fen(fen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablebaseAPI.AntichessAtomic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AntichessAtomic`: TablebaseStandard200Response
	fmt.Fprintf(os.Stdout, "Response from `TablebaseAPI.AntichessAtomic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAntichessAtomicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | **string** | X-FEN of the position. Underscores allowed. | 

### Return type

[**TablebaseStandard200Response**](TablebaseStandard200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TablebaseAtomic

> TablebaseStandard200Response TablebaseAtomic(ctx).Fen(fen).Execute()

Tablebase lookup for Atomic chess



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
	fen := "fen_example" // string | X-FEN of the position. Underscores allowed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablebaseAPI.TablebaseAtomic(context.Background()).Fen(fen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablebaseAPI.TablebaseAtomic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TablebaseAtomic`: TablebaseStandard200Response
	fmt.Fprintf(os.Stdout, "Response from `TablebaseAPI.TablebaseAtomic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTablebaseAtomicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | **string** | X-FEN of the position. Underscores allowed. | 

### Return type

[**TablebaseStandard200Response**](TablebaseStandard200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TablebaseStandard

> TablebaseStandard200Response TablebaseStandard(ctx).Fen(fen).Dtc(dtc).Execute()

Tablebase lookup



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
	fen := "fen_example" // string | X-FEN of the position. Underscores allowed.
	dtc := "dtc_example" // string | When to query the tablebase for `dtc` values. The default is `auxiliary`, i.e., only when the position is not covered by one of the other tablebases.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablebaseAPI.TablebaseStandard(context.Background()).Fen(fen).Dtc(dtc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablebaseAPI.TablebaseStandard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TablebaseStandard`: TablebaseStandard200Response
	fmt.Fprintf(os.Stdout, "Response from `TablebaseAPI.TablebaseStandard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTablebaseStandardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | **string** | X-FEN of the position. Underscores allowed. | 
 **dtc** | **string** | When to query the tablebase for &#x60;dtc&#x60; values. The default is &#x60;auxiliary&#x60;, i.e., only when the position is not covered by one of the other tablebases.  | 

### Return type

[**TablebaseStandard200Response**](TablebaseStandard200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

