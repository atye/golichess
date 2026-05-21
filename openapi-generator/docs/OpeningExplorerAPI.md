# \OpeningExplorerAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpeningExplorerLichess**](OpeningExplorerAPI.md#OpeningExplorerLichess) | **Get** /lichess | Lichess games
[**OpeningExplorerMaster**](OpeningExplorerAPI.md#OpeningExplorerMaster) | **Get** /masters | Masters database
[**OpeningExplorerMasterGame**](OpeningExplorerAPI.md#OpeningExplorerMasterGame) | **Get** /masters/pgn/{gameId} | OTB master game
[**OpeningExplorerPlayer**](OpeningExplorerAPI.md#OpeningExplorerPlayer) | **Get** /player | Player games



## OpeningExplorerLichess

> OpeningExplorerLichess OpeningExplorerLichess(ctx).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Moves(moves).TopGames(topGames).RecentGames(recentGames).History(history).Execute()

Lichess games



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
	variant := openapiclient.VariantKey("standard") // VariantKey | Variant (optional) (default to "standard")
	fen := "rnbqkbnr/ppp1pppp/8/3pP3/8/8/PPPP1PPP/RNBQKBNR b KQkq - 0 2" // string | X-FEN or EPD of the root position (optional)
	play := "d2d4,d7d5,c2c4,c7c6,c4d5" // string | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to "")
	speeds := []openapiclient.Speed{openapiclient.Speed("ultraBullet")} // []Speed | Comma separated list of game speeds to filter by (optional)
	ratings := []int32{int32(123)} // []int32 | Comma separated list of ratings groups to filter by. Each group ranges from its value to the next higher group in the enum (`0` from 0 to 999, `1000` from 1000 to 1199, ..., `2500` from 2500 to any rating above).  (optional)
	since := "since_example" // string | Include only games from this month or later (optional) (default to "1952-01")
	until := "until_example" // string | Include only games from this month or earlier (optional) (default to "3000-12")
	moves := int32(56) // int32 | Number of most common moves to display (optional) (default to 12)
	topGames := int32(56) // int32 | Maximum number of top games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected.  (optional) (default to 4)
	recentGames := int32(56) // int32 | Maximum number of recent games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected.  (optional) (default to 4)
	history := true // bool | Optionally retrieve history (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpeningExplorerAPI.OpeningExplorerLichess(context.Background()).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Ratings(ratings).Since(since).Until(until).Moves(moves).TopGames(topGames).RecentGames(recentGames).History(history).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerAPI.OpeningExplorerLichess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpeningExplorerLichess`: OpeningExplorerLichess
	fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerAPI.OpeningExplorerLichess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerLichessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to &quot;standard&quot;]
 **fen** | **string** | X-FEN or EPD of the root position | 
 **play** | **string** | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to &quot;&quot;]
 **speeds** | [**[]Speed**](Speed.md) | Comma separated list of game speeds to filter by | 
 **ratings** | **[]int32** | Comma separated list of ratings groups to filter by. Each group ranges from its value to the next higher group in the enum (&#x60;0&#x60; from 0 to 999, &#x60;1000&#x60; from 1000 to 1199, ..., &#x60;2500&#x60; from 2500 to any rating above).  | 
 **since** | **string** | Include only games from this month or later | [default to &quot;1952-01&quot;]
 **until** | **string** | Include only games from this month or earlier | [default to &quot;3000-12&quot;]
 **moves** | **int32** | Number of most common moves to display | [default to 12]
 **topGames** | **int32** | Maximum number of top games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected.  | [default to 4]
 **recentGames** | **int32** | Maximum number of recent games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected.  | [default to 4]
 **history** | **bool** | Optionally retrieve history | [default to false]

### Return type

[**OpeningExplorerLichess**](OpeningExplorerLichess.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerMaster

> OpeningExplorerMasters OpeningExplorerMaster(ctx).Fen(fen).Play(play).Since(since).Until(until).Moves(moves).TopGames(topGames).Execute()

Masters database



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
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" // string | X-FEN of the root position (optional)
	play := "d2d4,d7d5,c2c4,c7c6,c4d5" // string | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to "")
	since := int32(56) // int32 | Include only games from this year or later (optional) (default to 1952)
	until := int32(56) // int32 | Include only games from this year or earlier (optional)
	moves := int32(56) // int32 | Number of most common moves to display (optional) (default to 12)
	topGames := int32(56) // int32 | Number of top games to display (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpeningExplorerAPI.OpeningExplorerMaster(context.Background()).Fen(fen).Play(play).Since(since).Until(until).Moves(moves).TopGames(topGames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerAPI.OpeningExplorerMaster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpeningExplorerMaster`: OpeningExplorerMasters
	fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerAPI.OpeningExplorerMaster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerMasterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fen** | **string** | X-FEN of the root position | 
 **play** | **string** | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to &quot;&quot;]
 **since** | **int32** | Include only games from this year or later | [default to 1952]
 **until** | **int32** | Include only games from this year or earlier | 
 **moves** | **int32** | Number of most common moves to display | [default to 12]
 **topGames** | **int32** | Number of top games to display | [default to 15]

### Return type

[**OpeningExplorerMasters**](OpeningExplorerMasters.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpeningExplorerMasterGame

> string OpeningExplorerMasterGame(ctx, gameId).Execute()

OTB master game



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpeningExplorerAPI.OpeningExplorerMasterGame(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerAPI.OpeningExplorerMasterGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpeningExplorerMasterGame`: string
	fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerAPI.OpeningExplorerMasterGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerMasterGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## OpeningExplorerPlayer

> OpeningExplorerPlayer OpeningExplorerPlayer(ctx).Player(player).Color(color).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Modes(modes).Since(since).Until(until).Moves(moves).RecentGames(recentGames).Execute()

Player games



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
	player := "revoof" // string | Username or ID of the player
	color := "white" // string | Look for games with *player* on the given side
	variant := openapiclient.VariantKey("standard") // VariantKey | Variant (optional) (default to "standard")
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" // string | X-FEN of the root position (optional)
	play := "d2d4,d7d5" // string | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from `fen`. Required to find an opening name, if `fen` is not an exact match for a named position.  (optional) (default to "")
	speeds := []openapiclient.Speed{openapiclient.Speed("ultraBullet")} // []Speed | Comma separated list of game speeds to look for (optional)
	modes := []string{"Modes_example"} // []string | Comma separated list of modes (optional)
	since := "since_example" // string | Include only games from this month or later (optional) (default to "1952-01")
	until := "until_example" // string | Include only games from this month or earlier (optional) (default to "3000-12")
	moves := int32(56) // int32 | Number of most common moves to display (optional)
	recentGames := int32(56) // int32 | Number of recent games to display (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpeningExplorerAPI.OpeningExplorerPlayer(context.Background()).Player(player).Color(color).Variant(variant).Fen(fen).Play(play).Speeds(speeds).Modes(modes).Since(since).Until(until).Moves(moves).RecentGames(recentGames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpeningExplorerAPI.OpeningExplorerPlayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpeningExplorerPlayer`: OpeningExplorerPlayer
	fmt.Fprintf(os.Stdout, "Response from `OpeningExplorerAPI.OpeningExplorerPlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpeningExplorerPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **player** | **string** | Username or ID of the player | 
 **color** | **string** | Look for games with *player* on the given side | 
 **variant** | [**VariantKey**](VariantKey.md) | Variant | [default to &quot;standard&quot;]
 **fen** | **string** | X-FEN of the root position | 
 **play** | **string** | Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position.  | [default to &quot;&quot;]
 **speeds** | [**[]Speed**](Speed.md) | Comma separated list of game speeds to look for | 
 **modes** | **[]string** | Comma separated list of modes | 
 **since** | **string** | Include only games from this month or later | [default to &quot;1952-01&quot;]
 **until** | **string** | Include only games from this month or earlier | [default to &quot;3000-12&quot;]
 **moves** | **int32** | Number of most common moves to display | 
 **recentGames** | **int32** | Number of recent games to display | [default to 8]

### Return type

[**OpeningExplorerPlayer**](OpeningExplorerPlayer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

