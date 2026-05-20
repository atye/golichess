# \BoardAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBoardSeek**](BoardAPI.md#ApiBoardSeek) | **Post** /api/board/seek | Create a seek
[**ApiStreamEvent**](BoardAPI.md#ApiStreamEvent) | **Get** /api/stream/event | Stream incoming events
[**BoardGameAbort**](BoardAPI.md#BoardGameAbort) | **Post** /api/board/game/{gameId}/abort | Abort a game
[**BoardGameBerserk**](BoardAPI.md#BoardGameBerserk) | **Post** /api/board/game/{gameId}/berserk | Berserk a tournament game
[**BoardGameChatGet**](BoardAPI.md#BoardGameChatGet) | **Get** /api/board/game/{gameId}/chat | Fetch the player chat
[**BoardGameChatPost**](BoardAPI.md#BoardGameChatPost) | **Post** /api/board/game/{gameId}/chat | Write in the chat
[**BoardGameClaimDraw**](BoardAPI.md#BoardGameClaimDraw) | **Post** /api/board/game/{gameId}/claim-draw | Claim draw of a game
[**BoardGameClaimVictory**](BoardAPI.md#BoardGameClaimVictory) | **Post** /api/board/game/{gameId}/claim-victory | Claim victory of a game
[**BoardGameDraw**](BoardAPI.md#BoardGameDraw) | **Post** /api/board/game/{gameId}/draw/{accept} | Handle draw offers
[**BoardGameMove**](BoardAPI.md#BoardGameMove) | **Post** /api/board/game/{gameId}/move/{move} | Make a Board move
[**BoardGameResign**](BoardAPI.md#BoardGameResign) | **Post** /api/board/game/{gameId}/resign | Resign a game
[**BoardGameStream**](BoardAPI.md#BoardGameStream) | **Get** /api/board/game/stream/{gameId} | Stream Board game state
[**BoardGameTakeback**](BoardAPI.md#BoardGameTakeback) | **Post** /api/board/game/{gameId}/takeback/{accept} | Handle takeback offers



## ApiBoardSeek

> ApiBoardSeek200Response ApiBoardSeek(ctx).Time(time).Increment(increment).Days(days).Rated(rated).Variant(variant).RatingRange(ratingRange).Color(color).Execute()

Create a seek



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
	time := float32(8.14) // float32 | Clock initial time in minutes. Required for real-time seeks.
	increment := int32(56) // int32 | Clock increment in seconds. Required for real-time seeks.
	days := int32(56) // int32 | Days per turn. Required for correspondence seeks.
	rated := true // bool | Whether the game is rated and impacts players ratings. (optional) (default to false)
	variant := "variant_example" // string |  (optional) (default to "standard")
	ratingRange := "ratingRange_example" // string | The rating range of potential opponents. Better left empty. Example: 1500-1800  (optional)
	color := "color_example" // string | Which color you get to play (optional) (default to "random")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.ApiBoardSeek(context.Background()).Time(time).Increment(increment).Days(days).Rated(rated).Variant(variant).RatingRange(ratingRange).Color(color).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.ApiBoardSeek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBoardSeek`: ApiBoardSeek200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.ApiBoardSeek`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBoardSeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **float32** | Clock initial time in minutes. Required for real-time seeks. | 
 **increment** | **int32** | Clock increment in seconds. Required for real-time seeks. | 
 **days** | **int32** | Days per turn. Required for correspondence seeks. | 
 **rated** | **bool** | Whether the game is rated and impacts players ratings. | [default to false]
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **ratingRange** | **string** | The rating range of potential opponents. Better left empty. Example: 1500-1800  | 
 **color** | **string** | Which color you get to play | [default to &quot;random&quot;]

### Return type

[**ApiBoardSeek200Response**](ApiBoardSeek200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStreamEvent

> ApiStreamEvent200Response ApiStreamEvent(ctx).Execute()

Stream incoming events



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
	resp, r, err := apiClient.BoardAPI.ApiStreamEvent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.ApiStreamEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStreamEvent`: ApiStreamEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.ApiStreamEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiStreamEventRequest struct via the builder pattern


### Return type

[**ApiStreamEvent200Response**](ApiStreamEvent200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameAbort

> AccountKidPost200Response BoardGameAbort(ctx, gameId).Execute()

Abort a game



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameAbort(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameAbort`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameBerserk

> AccountKidPost200Response BoardGameBerserk(ctx, gameId).Execute()

Berserk a tournament game



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameBerserk(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameBerserk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameBerserk`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameBerserk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameBerserkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameChatGet

> []GameChatGet200ResponseInner BoardGameChatGet(ctx, gameId).Execute()

Fetch the player chat



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameChatGet(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameChatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameChatGet`: []GameChatGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameChatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameChatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GameChatGet200ResponseInner**](GameChatGet200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameChatPost

> AccountKidPost200Response BoardGameChatPost(ctx, gameId).Room(room).Text(text).Execute()

Write in the chat



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
	gameId := "5IrD6Gzz" // string | 
	room := "room_example" // string | 
	text := "text_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameChatPost(context.Background(), gameId).Room(room).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameChatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameChatPost`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameChatPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameChatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **room** | **string** |  | 
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


## BoardGameClaimDraw

> AccountKidPost200Response BoardGameClaimDraw(ctx, gameId).Execute()

Claim draw of a game



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameClaimDraw(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameClaimDraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameClaimDraw`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameClaimDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameClaimDrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameClaimVictory

> AccountKidPost200Response BoardGameClaimVictory(ctx, gameId).Execute()

Claim victory of a game



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameClaimVictory(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameClaimVictory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameClaimVictory`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameClaimVictory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameClaimVictoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameDraw

> AccountKidPost200Response BoardGameDraw(ctx, gameId, accept).Execute()

Handle draw offers



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
	gameId := "5IrD6Gzz" // string | 
	accept := *openapiclient.NewBoardGameDrawAcceptParameter() // BoardGameDrawAcceptParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameDraw(context.Background(), gameId, accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameDraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameDraw`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**accept** | [**BoardGameDrawAcceptParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameDrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameMove

> AccountKidPost200Response BoardGameMove(ctx, gameId, move).OfferingDraw(offeringDraw).Execute()

Make a Board move



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
	gameId := "5IrD6Gzz" // string | 
	move := "e2e4" // string | The move to play, in UCI format
	offeringDraw := true // bool | Whether to offer (or agree to) a draw (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameMove(context.Background(), gameId, move).OfferingDraw(offeringDraw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameMove`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**move** | **string** | The move to play, in UCI format | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offeringDraw** | **bool** | Whether to offer (or agree to) a draw | 

### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameResign

> AccountKidPost200Response BoardGameResign(ctx, gameId).Execute()

Resign a game



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameResign(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameResign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameResign`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameResign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameResignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameStream

> BoardGameStream200Response BoardGameStream(ctx, gameId).Execute()

Stream Board game state



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameStream(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameStream`: BoardGameStream200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BoardGameStream200Response**](BoardGameStream200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardGameTakeback

> AccountKidPost200Response BoardGameTakeback(ctx, gameId, accept).Execute()

Handle takeback offers



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
	gameId := "5IrD6Gzz" // string | 
	accept := *openapiclient.NewBoardGameDrawAcceptParameter() // BoardGameDrawAcceptParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.BoardGameTakeback(context.Background(), gameId, accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.BoardGameTakeback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BoardGameTakeback`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.BoardGameTakeback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**accept** | [**BoardGameDrawAcceptParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBoardGameTakebackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

