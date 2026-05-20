# \AnalysisAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCloudEval**](AnalysisAPI.md#ApiCloudEval) | **Get** /api/cloud-eval | Get cloud evaluation of a position.



## ApiCloudEval

> ApiCloudEval200Response ApiCloudEval(ctx).Fen(fen).MultiPv(multiPv).Variant(variant).Execute()

Get cloud evaluation of a position.



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
	fen := "r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 3 3" // string | X-FEN of the position
	multiPv := int32(56) // int32 | Number of variations (optional) (default to 1)
	variant := "standard" // string | Variant (optional) (default to "standard")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.ApiCloudEval(context.Background()).Fen(fen).MultiPv(multiPv).Variant(variant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.ApiCloudEval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCloudEval`: ApiCloudEval200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.ApiCloudEval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCloudEvalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | **string** | X-FEN of the position | 
 **multiPv** | **int32** | Number of variations | [default to 1]
 **variant** | **string** | Variant | [default to &quot;standard&quot;]

### Return type

[**ApiCloudEval200Response**](ApiCloudEval200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

