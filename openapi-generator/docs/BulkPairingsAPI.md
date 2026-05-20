# \BulkPairingsAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPairingCreate**](BulkPairingsAPI.md#BulkPairingCreate) | **Post** /api/bulk-pairing | Create a bulk pairing
[**BulkPairingDelete**](BulkPairingsAPI.md#BulkPairingDelete) | **Delete** /api/bulk-pairing/{id} | Cancel a bulk pairing
[**BulkPairingGet**](BulkPairingsAPI.md#BulkPairingGet) | **Get** /api/bulk-pairing/{id} | Show a bulk pairing
[**BulkPairingIdGamesGet**](BulkPairingsAPI.md#BulkPairingIdGamesGet) | **Get** /api/bulk-pairing/{id}/games | Export games of a bulk pairing
[**BulkPairingList**](BulkPairingsAPI.md#BulkPairingList) | **Get** /api/bulk-pairing | View your bulk pairings
[**BulkPairingStartClocks**](BulkPairingsAPI.md#BulkPairingStartClocks) | **Post** /api/bulk-pairing/{id}/start-clocks | Manually start clocks



## BulkPairingCreate

> BulkPairingList200ResponseInner BulkPairingCreate(ctx).Players(players).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).PairAt(pairAt).StartClocksAt(startClocksAt).Rated(rated).Variant(variant).Fen(fen).Message(message).Rules(rules).Execute()

Create a bulk pairing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {
	players := "players_example" // string | OAuth tokens of all the players to pair, with the syntax `tokenOfWhitePlayerInGame1:tokenOfBlackPlayerInGame1,tokenOfWhitePlayerInGame2:tokenOfBlackPlayerInGame2,...`. The 2 tokens of the players of a game are separated with `:`. The first token gets the white pieces. Games are separated with `,`. Up to 1000 tokens can be sent, for a max of 500 games. Each token must be included at most once. Example: `token1:token2,token3:token4,token5:token6`  (optional)
	clockLimit := int32(56) // int32 | Clock initial time in seconds. Example: `600`  (optional)
	clockIncrement := int32(56) // int32 | Clock increment in seconds. Example: `2`  (optional)
	days := int32(56) // int32 | Days per turn. For correspondence games only. (optional)
	pairAt := int64(789) // int64 | Date at which the games will be created as a Unix timestamp in milliseconds. Up to 7 days in the future. Omit, or set to current date and time, to start the games immediately. Example: `1612289869919`  (optional)
	startClocksAt := int64(789) // int64 | Date at which the clocks will be automatically started as a Unix timestamp in milliseconds. Up to 7 days in the future. Note that the clocks can start earlier than specified, if players start making moves in the game. If omitted, the clocks will not start automatically. Example: `1612289869919`  (optional)
	rated := true // bool | Game is rated and impacts players ratings (optional) (default to false)
	variant := "variant_example" // string |  (optional) (default to "standard")
	fen := "fen_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	message := "message_example" // string | Message that will be sent to each player, when the game is created.  It is sent from your user account. `{opponent}` and `{game}` are placeholders that will be replaced with the opponent and the game URLs. You can omit this field to send the default message, but if you set your own message, it must at least contain the `{game}` placeholder.  (optional) (default to "Your game with {opponent} is ready: {game}.")
	rules := "rules_example" // string | Extra game rules separated by commas. Example: `noAbort,noRematch`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingCreate(context.Background()).Players(players).ClockLimit(clockLimit).ClockIncrement(clockIncrement).Days(days).PairAt(pairAt).StartClocksAt(startClocksAt).Rated(rated).Variant(variant).Fen(fen).Message(message).Rules(rules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingCreate`: BulkPairingList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **players** | **string** | OAuth tokens of all the players to pair, with the syntax &#x60;tokenOfWhitePlayerInGame1:tokenOfBlackPlayerInGame1,tokenOfWhitePlayerInGame2:tokenOfBlackPlayerInGame2,...&#x60;. The 2 tokens of the players of a game are separated with &#x60;:&#x60;. The first token gets the white pieces. Games are separated with &#x60;,&#x60;. Up to 1000 tokens can be sent, for a max of 500 games. Each token must be included at most once. Example: &#x60;token1:token2,token3:token4,token5:token6&#x60;  | 
 **clockLimit** | **int32** | Clock initial time in seconds. Example: &#x60;600&#x60;  | 
 **clockIncrement** | **int32** | Clock increment in seconds. Example: &#x60;2&#x60;  | 
 **days** | **int32** | Days per turn. For correspondence games only. | 
 **pairAt** | **int64** | Date at which the games will be created as a Unix timestamp in milliseconds. Up to 7 days in the future. Omit, or set to current date and time, to start the games immediately. Example: &#x60;1612289869919&#x60;  | 
 **startClocksAt** | **int64** | Date at which the clocks will be automatically started as a Unix timestamp in milliseconds. Up to 7 days in the future. Note that the clocks can start earlier than specified, if players start making moves in the game. If omitted, the clocks will not start automatically. Example: &#x60;1612289869919&#x60;  | 
 **rated** | **bool** | Game is rated and impacts players ratings | [default to false]
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **fen** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **message** | **string** | Message that will be sent to each player, when the game is created.  It is sent from your user account. &#x60;{opponent}&#x60; and &#x60;{game}&#x60; are placeholders that will be replaced with the opponent and the game URLs. You can omit this field to send the default message, but if you set your own message, it must at least contain the &#x60;{game}&#x60; placeholder.  | [default to &quot;Your game with {opponent} is ready: {game}.&quot;]
 **rules** | **string** | Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60;  | 

### Return type

[**BulkPairingList200ResponseInner**](BulkPairingList200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingDelete

> AccountKidPost200Response BulkPairingDelete(ctx, id).Execute()

Cancel a bulk pairing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {
	id := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingDelete`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingDeleteRequest struct via the builder pattern


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


## BulkPairingGet

> BulkPairingList200ResponseInner BulkPairingGet(ctx, id).Execute()

Show a bulk pairing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {
	id := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingGet`: BulkPairingList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkPairingList200ResponseInner**](BulkPairingList200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingIdGamesGet

> string BulkPairingIdGamesGet(ctx, id).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()

Export games of a bulk pairing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {
	id := "5IrD6Gzz" // string | 
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to false)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to false)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to false)
	literate := true // bool | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingIdGamesGet(context.Background(), id).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingIdGamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingIdGamesGet`: string
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingIdGamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingIdGamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to false]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to false]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to false]
 **literate** | **bool** | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingList

> []BulkPairingList200ResponseInner BulkPairingList(ctx).Execute()

View your bulk pairings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingList`: []BulkPairingList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingListRequest struct via the builder pattern


### Return type

[**[]BulkPairingList200ResponseInner**](BulkPairingList200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPairingStartClocks

> AccountKidPost200Response BulkPairingStartClocks(ctx, id).Execute()

Manually start clocks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapigenerator"
)

func main() {
	id := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkPairingsAPI.BulkPairingStartClocks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkPairingsAPI.BulkPairingStartClocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPairingStartClocks`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BulkPairingsAPI.BulkPairingStartClocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkPairingStartClocksRequest struct via the builder pattern


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

