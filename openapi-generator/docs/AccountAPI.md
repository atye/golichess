# \AccountAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Account**](AccountAPI.md#Account) | **Get** /api/account/preferences | Get my preferences
[**AccountEmail**](AccountAPI.md#AccountEmail) | **Get** /api/account/email | Get my email address
[**AccountKid**](AccountAPI.md#AccountKid) | **Get** /api/account/kid | Get my kid mode status
[**AccountKidPost**](AccountAPI.md#AccountKidPost) | **Post** /api/account/kid | Set my kid mode status
[**AccountMe**](AccountAPI.md#AccountMe) | **Get** /api/account | Get my profile
[**Timeline**](AccountAPI.md#Timeline) | **Get** /api/timeline | Get my timeline



## Account

> Account200Response Account(ctx).Execute()

Get my preferences



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.Account(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Account``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Account`: Account200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Account`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountRequest struct via the builder pattern


### Return type

[**Account200Response**](Account200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountEmail

> AccountEmail200Response AccountEmail(ctx).Execute()

Get my email address



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountEmail`: AccountEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountEmailRequest struct via the builder pattern


### Return type

[**AccountEmail200Response**](AccountEmail200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountKid

> AccountKid200Response AccountKid(ctx).Execute()

Get my kid mode status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountKid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountKid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountKid`: AccountKid200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountKid`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountKidRequest struct via the builder pattern


### Return type

[**AccountKid200Response**](AccountKid200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountKidPost

> Ok AccountKidPost(ctx).V(v).Execute()

Set my kid mode status



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
	v := true // bool | Kid mode status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountKidPost(context.Background()).V(v).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountKidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountKidPost`: Ok
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountKidPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountKidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v** | **bool** | Kid mode status | 

### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountMe

> UserExtended AccountMe(ctx).Execute()

Get my profile



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountMe`: UserExtended
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountMeRequest struct via the builder pattern


### Return type

[**UserExtended**](UserExtended.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Timeline

> Timeline Timeline(ctx).Since(since).Nb(nb).Execute()

Get my timeline



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
	since := int32(56) // int32 | Show events since this timestamp. (optional)
	nb := int32(56) // int32 | Max number of events to fetch. (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.Timeline(context.Background()).Since(since).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.Timeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Timeline`: Timeline
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.Timeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int32** | Show events since this timestamp. | 
 **nb** | **int32** | Max number of events to fetch. | [default to 15]

### Return type

[**Timeline**](Timeline.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

