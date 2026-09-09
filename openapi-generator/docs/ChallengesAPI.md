# \ChallengesAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminChallengeTokens**](ChallengesAPI.md#AdminChallengeTokens) | **Post** /api/token/admin-challenge | Admin challenge tokens
[**ChallengeAccept**](ChallengesAPI.md#ChallengeAccept) | **Post** /api/challenge/{challengeId}/accept | Accept a challenge
[**ChallengeAi**](ChallengesAPI.md#ChallengeAi) | **Post** /api/challenge/ai | Challenge the AI
[**ChallengeCancel**](ChallengesAPI.md#ChallengeCancel) | **Post** /api/challenge/{challengeId}/cancel | Cancel a challenge
[**ChallengeCreate**](ChallengesAPI.md#ChallengeCreate) | **Post** /api/challenge/{username} | Create a challenge
[**ChallengeDecline**](ChallengesAPI.md#ChallengeDecline) | **Post** /api/challenge/{challengeId}/decline | Decline a challenge
[**ChallengeList**](ChallengesAPI.md#ChallengeList) | **Get** /api/challenge | List your challenges
[**ChallengeOpen**](ChallengesAPI.md#ChallengeOpen) | **Post** /api/challenge/open | Open-ended challenge
[**ChallengeShow**](ChallengesAPI.md#ChallengeShow) | **Get** /api/challenge/{challengeId}/show | Show one challenge
[**ChallengeStartClocks**](ChallengesAPI.md#ChallengeStartClocks) | **Post** /api/challenge/{gameId}/start-clocks | Start clocks of a game
[**RoundAddTime**](ChallengesAPI.md#RoundAddTime) | **Post** /api/round/{gameId}/add-time/{seconds} | Add time to the opponent clock



## AdminChallengeTokens

> map[string]string AdminChallengeTokens(ctx).Users(users).Description(description).Execute()

Admin challenge tokens



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
	users := "users_example" // string | Usernames separated with commas
	description := "description_example" // string | User visible description of the token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.AdminChallengeTokens(context.Background()).Users(users).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.AdminChallengeTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminChallengeTokens`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.AdminChallengeTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminChallengeTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | **string** | Usernames separated with commas | 
 **description** | **string** | User visible description of the token | 

### Return type

**map[string]string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeAccept

> Ok ChallengeAccept(ctx, challengeId).Color(color).Execute()

Accept a challenge



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
	challengeId := "5IrD6Gzz" // string | 
	color := "color_example" // string | Accept challenge as this color (only valid if this is an [open challenge](#challenge/open)) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeAccept(context.Background(), challengeId).Color(color).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeAccept`: Ok
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **color** | **string** | Accept challenge as this color (only valid if this is an [open challenge](#challenge/open)) | 

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


## ChallengeAi

> ChallengeAi201Response ChallengeAi(ctx).Level(level).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).Execute()

Challenge the AI



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
	level := int32(56) // int32 | AI strength
	clockLimit := int32(56) // int32 | Clock initial time in seconds. If empty, a correspondence game is created. (optional)
	clockIncrement := int32(56) // int32 | Clock increment in seconds. If empty, a correspondence game is created. (optional)
	days := int32(56) // int32 | Days per move, for correspondence games. Clock settings must be omitted. (optional)
	color := openapiclient.ChallengeColor("white") // ChallengeColor | Which color you get to play (optional) (default to "random")
	variant := openapiclient.VariantKey("standard") // VariantKey |  (optional) (default to "standard")
	fen := "fen_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeAi(context.Background()).Level(level).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Color(color).Variant(variant).Fen(fen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeAi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeAi`: ChallengeAi201Response
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeAi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChallengeAiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **level** | **int32** | AI strength | 
 **clockLimit** | **int32** | Clock initial time in seconds. If empty, a correspondence game is created. | 
 **clockIncrement** | **int32** | Clock increment in seconds. If empty, a correspondence game is created. | 
 **days** | **int32** | Days per move, for correspondence games. Clock settings must be omitted. | 
 **color** | [**ChallengeColor**](ChallengeColor.md) | Which color you get to play | [default to &quot;random&quot;]
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to &quot;standard&quot;]
 **fen** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]

### Return type

[**ChallengeAi201Response**](ChallengeAi201Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeCancel

> Ok ChallengeCancel(ctx, challengeId).OpponentToken(opponentToken).Execute()

Cancel a challenge



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
	challengeId := "5IrD6Gzz" // string | 
	opponentToken := "opponentToken_example" // string | Optional `challenge:write` token of the opponent. If set, the game can be canceled even if both players have moved. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeCancel(context.Background(), challengeId).OpponentToken(opponentToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeCancel`: Ok
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opponentToken** | **string** | Optional &#x60;challenge:write&#x60; token of the opponent. If set, the game can be canceled even if both players have moved. | 

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


## ChallengeCreate

> ChallengeJson ChallengeCreate(ctx, username).Days(days).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Rated(rated).Color(color).Variant(variant).Fen(fen).KeepAliveStream(keepAliveStream).Rules(rules).Execute()

Create a challenge



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
	username := "LeelaChess" // string | 
	days := int32(56) // int32 | Days per turn. Required for correspondence seeks.
	clockLimit := int32(56) // int32 | Clock initial time in seconds. If empty, a correspondence game is created. Valid values are 0, 15, 30, 45, 60, 90, and any multiple of 60 up to 10800 (3 hours). (optional)
	clockIncrement := int32(56) // int32 | Clock increment in seconds. If empty, a correspondence game is created. (optional)
	rated := true // bool | Game is rated and impacts players ratings (optional) (default to false)
	color := openapiclient.ChallengeColor("white") // ChallengeColor | Which color you get to play (optional) (default to "random")
	variant := openapiclient.VariantKey("standard") // VariantKey |  (optional) (default to "standard")
	fen := "fen_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	keepAliveStream := true // bool | If set, the response is streamed as [ndjson](#description/streaming-with-nd-json). The challenge is kept alive until the connection is closed by the client. When the challenge is accepted, declined or canceled, a message of the form `{\\\"done\\\":\\\"accepted\\\"}` is sent, then the connection is closed by the server. If not set, the response is not streamed, and the challenge expires after 20s if not accepted.  (optional)
	rules := "rules_example" // string | Extra game rules separated by commas. Example: `noAbort,noRematch`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeCreate(context.Background(), username).Days(days).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Rated(rated).Color(color).Variant(variant).Fen(fen).KeepAliveStream(keepAliveStream).Rules(rules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeCreate`: ChallengeJson
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Days per turn. Required for correspondence seeks. | 
 **clockLimit** | **int32** | Clock initial time in seconds. If empty, a correspondence game is created. Valid values are 0, 15, 30, 45, 60, 90, and any multiple of 60 up to 10800 (3 hours). | 
 **clockIncrement** | **int32** | Clock increment in seconds. If empty, a correspondence game is created. | 
 **rated** | **bool** | Game is rated and impacts players ratings | [default to false]
 **color** | [**ChallengeColor**](ChallengeColor.md) | Which color you get to play | [default to &quot;random&quot;]
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to &quot;standard&quot;]
 **fen** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **keepAliveStream** | **bool** | If set, the response is streamed as [ndjson](#description/streaming-with-nd-json). The challenge is kept alive until the connection is closed by the client. When the challenge is accepted, declined or canceled, a message of the form &#x60;{\\\&quot;done\\\&quot;:\\\&quot;accepted\\\&quot;}&#x60; is sent, then the connection is closed by the server. If not set, the response is not streamed, and the challenge expires after 20s if not accepted.  | 
 **rules** | **string** | Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60;  | 

### Return type

[**ChallengeJson**](ChallengeJson.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeDecline

> Ok ChallengeDecline(ctx, challengeId).Reason(reason).Execute()

Decline a challenge



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
	challengeId := "5IrD6Gzz" // string | 
	reason := "reason_example" // string | Reason the challenge was declined. Only the values listed below are accepted; any other value falls back to `generic`. The reason will be translated to the player's language. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeDecline(context.Background(), challengeId).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeDecline`: Ok
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | Reason the challenge was declined. Only the values listed below are accepted; any other value falls back to &#x60;generic&#x60;. The reason will be translated to the player&#39;s language. | 

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


## ChallengeList

> ChallengeList200Response ChallengeList(ctx).Execute()

List your challenges



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
	resp, r, err := apiClient.ChallengesAPI.ChallengeList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeList`: ChallengeList200Response
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeListRequest struct via the builder pattern


### Return type

[**ChallengeList200Response**](ChallengeList200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeOpen

> ChallengeOpenJson ChallengeOpen(ctx).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Variant(variant).Fen(fen).Name(name).Rules(rules).Users(users).ExpiresAt(expiresAt).Execute()

Open-ended challenge



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
	rated := true // bool | Game is rated and impacts players ratings (optional) (default to false)
	clockLimit := int32(56) // int32 | Clock initial time in seconds. If empty, a correspondence game is created. (optional)
	clockIncrement := int32(56) // int32 | Clock increment in seconds. If empty, a correspondence game is created. (optional)
	days := int32(56) // int32 | Days per turn. For correspondence challenges. (optional)
	variant := openapiclient.VariantKey("standard") // VariantKey |  (optional) (default to "standard")
	fen := "fen_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	name := "name_example" // string | Optional name for the challenge, that players will see on the challenge page. (optional)
	rules := "rules_example" // string | Extra game rules separated by commas. Example: `noRematch,noGiveTime` The `noAbort` rule is available for Lichess admins only  (optional)
	users := "users_example" // string | Optional pair of usernames, separated by a comma. If set, only these users will be allowed to join the game. The first username gets the white pieces. Example: `Username1,Username2`  (optional)
	expiresAt := int64(789) // int64 | Timestamp in milliseconds to expire the challenge. Defaults to 24h after creation. Can't be more than 2 weeks after creation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeOpen(context.Background()).Rated(rated).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).Variant(variant).Fen(fen).Name(name).Rules(rules).Users(users).ExpiresAt(expiresAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeOpen`: ChallengeOpenJson
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeOpen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChallengeOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rated** | **bool** | Game is rated and impacts players ratings | [default to false]
 **clockLimit** | **int32** | Clock initial time in seconds. If empty, a correspondence game is created. | 
 **clockIncrement** | **int32** | Clock increment in seconds. If empty, a correspondence game is created. | 
 **days** | **int32** | Days per turn. For correspondence challenges. | 
 **variant** | [**VariantKey**](VariantKey.md) |  | [default to &quot;standard&quot;]
 **fen** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **name** | **string** | Optional name for the challenge, that players will see on the challenge page. | 
 **rules** | **string** | Extra game rules separated by commas. Example: &#x60;noRematch,noGiveTime&#x60; The &#x60;noAbort&#x60; rule is available for Lichess admins only  | 
 **users** | **string** | Optional pair of usernames, separated by a comma. If set, only these users will be allowed to join the game. The first username gets the white pieces. Example: &#x60;Username1,Username2&#x60;  | 
 **expiresAt** | **int64** | Timestamp in milliseconds to expire the challenge. Defaults to 24h after creation. Can&#39;t be more than 2 weeks after creation. | 

### Return type

[**ChallengeOpenJson**](ChallengeOpenJson.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeShow

> ChallengeJson ChallengeShow(ctx, challengeId).Execute()

Show one challenge



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
	challengeId := "challengeId_example" // string | The challenge ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeShow(context.Background(), challengeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeShow`: ChallengeJson
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**challengeId** | **string** | The challenge ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChallengeJson**](ChallengeJson.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChallengeStartClocks

> Ok ChallengeStartClocks(ctx, gameId).Token1(token1).Token2(token2).Execute()

Start clocks of a game



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
	gameId := "gameId_example" // string | 
	token1 := "token1_example" // string | OAuth token of a player
	token2 := "token2_example" // string | OAuth token of the other player. Omit for AI games that have only one player. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.ChallengeStartClocks(context.Background(), gameId).Token1(token1).Token2(token2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.ChallengeStartClocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChallengeStartClocks`: Ok
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.ChallengeStartClocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChallengeStartClocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token1** | **string** | OAuth token of a player | 
 **token2** | **string** | OAuth token of the other player. Omit for AI games that have only one player. | 

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


## RoundAddTime

> Ok RoundAddTime(ctx, gameId, seconds).Execute()

Add time to the opponent clock



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
	gameId := "gameId_example" // string | 
	seconds := int32(56) // int32 | How many seconds to give

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChallengesAPI.RoundAddTime(context.Background(), gameId, seconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChallengesAPI.RoundAddTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoundAddTime`: Ok
	fmt.Fprintf(os.Stdout, "Response from `ChallengesAPI.RoundAddTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 
**seconds** | **int32** | How many seconds to give | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoundAddTimeRequest struct via the builder pattern


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

