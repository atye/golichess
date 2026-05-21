# \BotAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBotOnline**](BotAPI.md#ApiBotOnline) | **Get** /api/bot/online | Get online bots
[**ApiStreamEvent**](BotAPI.md#ApiStreamEvent) | **Get** /api/stream/event | Stream incoming events
[**BotAccountUpgrade**](BotAPI.md#BotAccountUpgrade) | **Post** /api/bot/account/upgrade | Upgrade to Bot account
[**BotGameAbort**](BotAPI.md#BotGameAbort) | **Post** /api/bot/game/{gameId}/abort | Abort a game
[**BotGameChat**](BotAPI.md#BotGameChat) | **Post** /api/bot/game/{gameId}/chat | Write in the chat
[**BotGameChatGet**](BotAPI.md#BotGameChatGet) | **Get** /api/bot/game/{gameId}/chat | Fetch the game chat
[**BotGameClaimDraw**](BotAPI.md#BotGameClaimDraw) | **Post** /api/bot/game/{gameId}/claim-draw | Claim draw of a game
[**BotGameClaimVictory**](BotAPI.md#BotGameClaimVictory) | **Post** /api/bot/game/{gameId}/claim-victory | Claim victory of a game
[**BotGameDraw**](BotAPI.md#BotGameDraw) | **Post** /api/bot/game/{gameId}/draw/{accept} | Handle draw offers
[**BotGameMove**](BotAPI.md#BotGameMove) | **Post** /api/bot/game/{gameId}/move/{move} | Make a Bot move
[**BotGameResign**](BotAPI.md#BotGameResign) | **Post** /api/bot/game/{gameId}/resign | Resign a game
[**BotGameStream**](BotAPI.md#BotGameStream) | **Get** /api/bot/game/stream/{gameId} | Stream Bot game state
[**BotGameTakeback**](BotAPI.md#BotGameTakeback) | **Post** /api/bot/game/{gameId}/takeback/{accept} | Handle takeback offers



## ApiBotOnline

> User ApiBotOnline(ctx).Nb(nb).Execute()

Get online bots



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
	nb := int32(20) // int32 | How many bot users to fetch (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.ApiBotOnline(context.Background()).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ApiBotOnline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBotOnline`: User
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.ApiBotOnline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBotOnlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | **int32** | How many bot users to fetch | [default to 100]

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

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
	resp, r, err := apiClient.BotAPI.ApiStreamEvent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.ApiStreamEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStreamEvent`: ApiStreamEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.ApiStreamEvent`: %v\n", resp)
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


## BotAccountUpgrade

> Ok BotAccountUpgrade(ctx).Execute()

Upgrade to Bot account



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
	resp, r, err := apiClient.BotAPI.BotAccountUpgrade(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotAccountUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotAccountUpgrade`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotAccountUpgrade`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotAccountUpgradeRequest struct via the builder pattern


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


## BotGameAbort

> Ok BotGameAbort(ctx, gameId).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameAbort(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameAbort`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BotGameChat

> Ok BotGameChat(ctx, gameId).Room(room).Text(text).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameChat(context.Background(), gameId).Room(room).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameChat`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **room** | **string** |  | 
 **text** | **string** |  | 

### Return type

[**Ok**](Ok.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameChatGet

> []SpectatorGameChatInner BotGameChatGet(ctx, gameId).Execute()

Fetch the game chat



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
	resp, r, err := apiClient.BotAPI.BotGameChatGet(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameChatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameChatGet`: []SpectatorGameChatInner
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameChatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameChatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SpectatorGameChatInner**](SpectatorGameChatInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGameClaimDraw

> Ok BotGameClaimDraw(ctx, gameId).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameClaimDraw(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameClaimDraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameClaimDraw`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameClaimDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameClaimDrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BotGameClaimVictory

> Ok BotGameClaimVictory(ctx, gameId).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameClaimVictory(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameClaimVictory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameClaimVictory`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameClaimVictory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameClaimVictoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BotGameDraw

> Ok BotGameDraw(ctx, gameId, accept).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameDraw(context.Background(), gameId, accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameDraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameDraw`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**accept** | [**BoardGameDrawAcceptParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameDrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## BotGameMove

> Ok BotGameMove(ctx, gameId, move).OfferingDraw(offeringDraw).Execute()

Make a Bot move



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
	resp, r, err := apiClient.BotAPI.BotGameMove(context.Background(), gameId, move).OfferingDraw(offeringDraw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameMove`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**move** | **string** | The move to play, in UCI format | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offeringDraw** | **bool** | Whether to offer (or agree to) a draw | 

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


## BotGameResign

> Ok BotGameResign(ctx, gameId).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameResign(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameResign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameResign`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameResign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameResignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BotGameStream

> BoardGameStream200Response BotGameStream(ctx, gameId).Execute()

Stream Bot game state



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
	resp, r, err := apiClient.BotAPI.BotGameStream(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameStream`: BoardGameStream200Response
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameStreamRequest struct via the builder pattern


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


## BotGameTakeback

> Ok BotGameTakeback(ctx, gameId, accept).Execute()

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
	resp, r, err := apiClient.BotAPI.BotGameTakeback(context.Background(), gameId, accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.BotGameTakeback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGameTakeback`: Ok
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.BotGameTakeback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**accept** | [**BoardGameDrawAcceptParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGameTakebackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

