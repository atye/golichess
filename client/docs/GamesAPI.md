# \GamesAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountPlaying**](GamesAPI.md#ApiAccountPlaying) | **Get** /api/account/playing | Get my ongoing games
[**ApiExportBookmarks**](GamesAPI.md#ApiExportBookmarks) | **Get** /api/games/export/bookmarks | Export your bookmarked games
[**ApiGamesUser**](GamesAPI.md#ApiGamesUser) | **Get** /api/games/user/{username} | Export games of a user
[**ApiImportedGamesUser**](GamesAPI.md#ApiImportedGamesUser) | **Get** /api/games/export/imports | Export your imported games
[**ApiUserCurrentGame**](GamesAPI.md#ApiUserCurrentGame) | **Get** /api/user/{username}/current-game | Export ongoing game of a user
[**GameChatGet**](GamesAPI.md#GameChatGet) | **Get** /game/{gameId}/chat | Fetch the spectator game chat
[**GameImport**](GamesAPI.md#GameImport) | **Post** /api/import | Import one game
[**GamePgn**](GamesAPI.md#GamePgn) | **Get** /game/export/{gameId} | Export one game
[**GamesByIds**](GamesAPI.md#GamesByIds) | **Post** /api/stream/games/{streamId} | Stream games by IDs
[**GamesByIdsAdd**](GamesAPI.md#GamesByIdsAdd) | **Post** /api/stream/games/{streamId}/add | Add game IDs to stream
[**GamesByUsers**](GamesAPI.md#GamesByUsers) | **Post** /api/stream/games-by-users | Stream games of users
[**GamesExportIds**](GamesAPI.md#GamesExportIds) | **Post** /api/games/export/_ids | Export games by IDs
[**StreamGame**](GamesAPI.md#StreamGame) | **Get** /api/stream/game/{id} | Stream moves of a game



## ApiAccountPlaying

> ApiAccountPlaying200Response ApiAccountPlaying(ctx).Nb(nb).Execute()

Get my ongoing games



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
	nb := int32(56) // int32 | Max number of games to fetch (optional) (default to 9)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.ApiAccountPlaying(context.Background()).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.ApiAccountPlaying``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAccountPlaying`: ApiAccountPlaying200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.ApiAccountPlaying`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountPlayingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | **int32** | Max number of games to fetch | [default to 9]

### Return type

[**ApiAccountPlaying200Response**](ApiAccountPlaying200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExportBookmarks

> ApiUserCurrentGame200Response ApiExportBookmarks(ctx).Accept(accept).Since(since).Until(until).Max(max).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).LastFen(lastFen).Sort(sort).Execute()

Export your bookmarked games



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
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	since := int32(56) // int32 | Download games bookmarked since this timestamp. Defaults to account creation date. (optional)
	until := int32(56) // int32 | Download games bookmarked until this timestamp. Defaults to now. (optional)
	max := int32(56) // int32 | How many bookmarked games to download. Leave empty to download all bookmarked games. (optional)
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. The response type must be set to `application/x-ndjson` by the request `Accept` header. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to false)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to false)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to false)
	literate := true // bool | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
	lastFen := true // bool | Include the X-FEN notation of the last position of the game. The response type must be set to `application/x-ndjson` by the request `Accept` header.  (optional) (default to false)
	sort := "sort_example" // string | Sort order of the bookmarks. (optional) (default to "dateDesc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.ApiExportBookmarks(context.Background()).Accept(accept).Since(since).Until(until).Max(max).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).LastFen(lastFen).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.ApiExportBookmarks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiExportBookmarks`: ApiUserCurrentGame200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.ApiExportBookmarks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiExportBookmarksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **since** | **int32** | Download games bookmarked since this timestamp. Defaults to account creation date. | 
 **until** | **int32** | Download games bookmarked until this timestamp. Defaults to now. | 
 **max** | **int32** | How many bookmarked games to download. Leave empty to download all bookmarked games. | 
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to false]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to false]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to false]
 **literate** | **bool** | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **lastFen** | **bool** | Include the X-FEN notation of the last position of the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.  | [default to false]
 **sort** | **string** | Sort order of the bookmarks. | [default to &quot;dateDesc&quot;]

### Return type

[**ApiUserCurrentGame200Response**](ApiUserCurrentGame200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGamesUser

> ApiUserCurrentGame200Response ApiGamesUser(ctx, username).Accept(accept).Since(since).Until(until).Max(max).Vs(vs).Rated(rated).PerfType(perfType).Color(color).Analysed(analysed).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Ongoing(ongoing).Finished(finished).Literate(literate).LastFen(lastFen).WithBookmarked(withBookmarked).Sort(sort).Execute()

Export games of a user



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
	username := "username_example" // string | The user name.
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	since := int32(56) // int32 | Download games played since this timestamp. Defaults to account creation date. (optional)
	until := int32(56) // int32 | Download games played until this timestamp. Defaults to now. (optional)
	max := int32(56) // int32 | How many games to download. Leave empty to download all games. (optional)
	vs := "vs_example" // string | [Filter] Only games played against this opponent (optional)
	rated := true // bool | [Filter] Only rated (`true`) or casual (`false`) games (optional)
	perfType := "perfType_example" // string | [Filter] Only games in these speeds or variants. Multiple perf types can be specified, separated by a comma. Example: blitz,rapid,classical  (optional)
	color := "color_example" // string | [Filter] Only games played as this color. (optional)
	analysed := true // bool | [Filter] Only games with or without a computer analysis available (optional)
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. The response type must be set to `application/x-ndjson` by the request `Accept` header. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to false)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to false)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to false)
	ongoing := true // bool | Ongoing games are delayed by a few seconds ranging from 3 to 60 depending on the time control, as to prevent cheat bots from using this API. (optional) (default to false)
	finished := true // bool | Include finished games. Set to `false` to only get ongoing games. (optional) (default to true)
	literate := true // bool | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
	lastFen := true // bool | Include the X-FEN notation of the last position of the game. The response type must be set to `application/x-ndjson` by the request `Accept` header.  (optional) (default to false)
	withBookmarked := true // bool | Add a `bookmarked: true` JSON field when the logged in user has bookmarked the game. The response type must be set to `application/x-ndjson` by the request `Accept` header.  (optional) (default to false)
	sort := "sort_example" // string | Sort order of the games. (optional) (default to "dateDesc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.ApiGamesUser(context.Background(), username).Accept(accept).Since(since).Until(until).Max(max).Vs(vs).Rated(rated).PerfType(perfType).Color(color).Analysed(analysed).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Ongoing(ongoing).Finished(finished).Literate(literate).LastFen(lastFen).WithBookmarked(withBookmarked).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.ApiGamesUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiGamesUser`: ApiUserCurrentGame200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.ApiGamesUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGamesUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **since** | **int32** | Download games played since this timestamp. Defaults to account creation date. | 
 **until** | **int32** | Download games played until this timestamp. Defaults to now. | 
 **max** | **int32** | How many games to download. Leave empty to download all games. | 
 **vs** | **string** | [Filter] Only games played against this opponent | 
 **rated** | **bool** | [Filter] Only rated (&#x60;true&#x60;) or casual (&#x60;false&#x60;) games | 
 **perfType** | **string** | [Filter] Only games in these speeds or variants. Multiple perf types can be specified, separated by a comma. Example: blitz,rapid,classical  | 
 **color** | **string** | [Filter] Only games played as this color. | 
 **analysed** | **bool** | [Filter] Only games with or without a computer analysis available | 
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to false]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to false]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to false]
 **ongoing** | **bool** | Ongoing games are delayed by a few seconds ranging from 3 to 60 depending on the time control, as to prevent cheat bots from using this API. | [default to false]
 **finished** | **bool** | Include finished games. Set to &#x60;false&#x60; to only get ongoing games. | [default to true]
 **literate** | **bool** | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **lastFen** | **bool** | Include the X-FEN notation of the last position of the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.  | [default to false]
 **withBookmarked** | **bool** | Add a &#x60;bookmarked: true&#x60; JSON field when the logged in user has bookmarked the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.  | [default to false]
 **sort** | **string** | Sort order of the games. | [default to &quot;dateDesc&quot;]

### Return type

[**ApiUserCurrentGame200Response**](ApiUserCurrentGame200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImportedGamesUser

> string ApiImportedGamesUser(ctx).Execute()

Export your imported games



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
	resp, r, err := apiClient.GamesAPI.ApiImportedGamesUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.ApiImportedGamesUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiImportedGamesUser`: string
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.ApiImportedGamesUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedGamesUserRequest struct via the builder pattern


### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserCurrentGame

> ApiUserCurrentGame200Response ApiUserCurrentGame(ctx, username).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()

Export ongoing game of a user



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
	username := "username_example" // string | 
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/json` to get the games in JSON format.  (optional) (default to "application/x-chess-pgn")
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to true)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to true)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to true)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to false)
	literate := true // bool | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.ApiUserCurrentGame(context.Background(), username).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.ApiUserCurrentGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserCurrentGame`: ApiUserCurrentGame200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.ApiUserCurrentGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserCurrentGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/json&#x60; to get the games in JSON format.  | [default to &quot;application/x-chess-pgn&quot;]
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to true]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to true]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to true]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to false]
 **literate** | **bool** | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]

### Return type

[**ApiUserCurrentGame200Response**](ApiUserCurrentGame200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameChatGet

> []GameChatGet200ResponseInner GameChatGet(ctx, gameId).Execute()

Fetch the spectator game chat



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
	gameId := "5IrD6Gzz" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GameChatGet(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GameChatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameChatGet`: []GameChatGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GameChatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameChatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GameChatGet200ResponseInner**](GameChatGet200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameImport

> GameImport200Response GameImport(ctx).Pgn(pgn).Execute()

Import one game



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
	pgn := "pgn_example" // string | The PGN. It can contain only one game. Most standard tags are supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GameImport(context.Background()).Pgn(pgn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GameImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameImport`: GameImport200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GameImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pgn** | **string** | The PGN. It can contain only one game. Most standard tags are supported. | 

### Return type

[**GameImport200Response**](GameImport200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamePgn

> GamePgn200Response GamePgn(ctx, gameId).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).WithBookmarked(withBookmarked).Execute()

Export one game



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
	gameId := "gameId_example" // string | The game ID
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/json` to get the games in JSON format.  (optional) (default to "application/x-chess-pgn")
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to true)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to true)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to true)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to true)
	literate := true // bool | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: `5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)`  (optional) (default to false)
	withBookmarked := true // bool | Add a `bookmarked: true` JSON field when the logged in user has bookmarked the game. The response type must be set to `application/x-ndjson` by the request `Accept` header.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GamePgn(context.Background(), gameId).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).WithBookmarked(withBookmarked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GamePgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamePgn`: GamePgn200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GamePgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The game ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamePgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/json&#x60; to get the games in JSON format.  | [default to &quot;application/x-chess-pgn&quot;]
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to true]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to true]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to true]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to true]
 **literate** | **bool** | Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60;  | [default to false]
 **withBookmarked** | **bool** | Add a &#x60;bookmarked: true&#x60; JSON field when the logged in user has bookmarked the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.  | [default to false]

### Return type

[**GamePgn200Response**](GamePgn200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByIds

> []GamesByIds200ResponseInner GamesByIds(ctx, streamId).Body(body).Execute()

Stream games by IDs



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
	streamId := "myAppName-someRandomId" // string | 
	body := "body_example" // string | Up to 500 or 1000 game IDs separated by commas. Example: `gameId01,gameId02,gameId03` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GamesByIds(context.Background(), streamId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GamesByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesByIds`: []GamesByIds200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GamesByIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId01,gameId02,gameId03&#x60;  | 

### Return type

[**[]GamesByIds200ResponseInner**](GamesByIds200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByIdsAdd

> AccountKidPost200Response GamesByIdsAdd(ctx, streamId).Body(body).Execute()

Add game IDs to stream



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
	streamId := "myAppName-someRandomId" // string | 
	body := "body_example" // string | Up to 500 or 1000 game IDs separated by commas. Example: `gameId04,gameId05,gameId06` 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GamesByIdsAdd(context.Background(), streamId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GamesByIdsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesByIdsAdd`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GamesByIdsAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByIdsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId04,gameId05,gameId06&#x60;  | 

### Return type

[**AccountKidPost200Response**](AccountKidPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByUsers

> []GamesByUsers200ResponseInner GamesByUsers(ctx).Body(body).WithCurrentGames(withCurrentGames).Execute()

Stream games of users



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
	body := "body_example" // string | Up to 300 user IDs separated by commas. Example: `thibault,maia1,maia5` 
	withCurrentGames := true // bool | Include the already started games at the beginning of the stream. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GamesByUsers(context.Background()).Body(body).WithCurrentGames(withCurrentGames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GamesByUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesByUsers`: []GamesByUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GamesByUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGamesByUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Up to 300 user IDs separated by commas. Example: &#x60;thibault,maia1,maia5&#x60;  | 
 **withCurrentGames** | **bool** | Include the already started games at the beginning of the stream. | [default to false]

### Return type

[**[]GamesByUsers200ResponseInner**](GamesByUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesExportIds

> ApiUserCurrentGame200Response GamesExportIds(ctx).Body(body).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()

Export games by IDs



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
	body := "TJxUmbWK,4OtIh2oh,ILwozzRZ" // string | Game IDs separated by commas. Up to 300.
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
	resp, r, err := apiClient.GamesAPI.GamesExportIds(context.Background()).Body(body).Accept(accept).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Literate(literate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GamesExportIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesExportIds`: ApiUserCurrentGame200Response
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GamesExportIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGamesExportIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Game IDs separated by commas. Up to 300. | 
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

[**ApiUserCurrentGame200Response**](ApiUserCurrentGame200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamGame

> []StreamGame200ResponseInner StreamGame(ctx, id).Execute()

Stream moves of a game



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
	id := "LuGQwhBb" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.StreamGame(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.StreamGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamGame`: []StreamGame200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.StreamGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StreamGame200ResponseInner**](StreamGame200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

