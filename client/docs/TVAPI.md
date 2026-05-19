# \TVAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TvChannelFeed**](TVAPI.md#TvChannelFeed) | **Get** /api/tv/{channel}/feed | Stream current TV game of a TV channel
[**TvChannelGames**](TVAPI.md#TvChannelGames) | **Get** /api/tv/{channel} | Get best ongoing games of a TV channel
[**TvChannels**](TVAPI.md#TvChannels) | **Get** /api/tv/channels | Get current TV games
[**TvFeed**](TVAPI.md#TvFeed) | **Get** /api/tv/feed | Stream current TV game



## TvChannelFeed

> TvChannelFeed200Response TvChannelFeed(ctx, channel).Execute()

Stream current TV game of a TV channel



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
	channel := "channel_example" // string | The name of the channel in camel case.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TVAPI.TvChannelFeed(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TVAPI.TvChannelFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvChannelFeed`: TvChannelFeed200Response
	fmt.Fprintf(os.Stdout, "Response from `TVAPI.TvChannelFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The name of the channel in camel case. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTvChannelFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TvChannelFeed200Response**](TvChannelFeed200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TvChannelGames

> string TvChannelGames(ctx, channel).Accept(accept).Nb(nb).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Opening(opening).Execute()

Get best ongoing games of a TV channel



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
	channel := "channel_example" // string | The name of the channel in camel case.
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	nb := int32(56) // int32 | Number of games to fetch. (optional) (default to 10)
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TVAPI.TvChannelGames(context.Background(), channel).Accept(accept).Nb(nb).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Opening(opening).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TVAPI.TvChannelGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvChannelGames`: string
	fmt.Fprintf(os.Stdout, "Response from `TVAPI.TvChannelGames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The name of the channel in camel case. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTvChannelGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **nb** | **int32** | Number of games to fetch. | [default to 10]
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TvChannels

> TvChannels200Response TvChannels(ctx).Execute()

Get current TV games



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TVAPI.TvChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TVAPI.TvChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvChannels`: TvChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `TVAPI.TvChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTvChannelsRequest struct via the builder pattern


### Return type

[**TvChannels200Response**](TvChannels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TvFeed

> TvFeed200Response TvFeed(ctx).Execute()

Stream current TV game



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TVAPI.TvFeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TVAPI.TvFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvFeed`: TvFeed200Response
	fmt.Fprintf(os.Stdout, "Response from `TVAPI.TvFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTvFeedRequest struct via the builder pattern


### Return type

[**TvFeed200Response**](TvFeed200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

