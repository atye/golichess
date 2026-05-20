# \BroadcastsAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastAllRoundsPgn**](BroadcastsAPI.md#BroadcastAllRoundsPgn) | **Get** /api/broadcast/{broadcastTournamentId}.pgn | Export all rounds as PGN
[**BroadcastMyRoundsGet**](BroadcastsAPI.md#BroadcastMyRoundsGet) | **Get** /api/broadcast/my-rounds | Get your broadcast rounds
[**BroadcastPlayerGet**](BroadcastsAPI.md#BroadcastPlayerGet) | **Get** /broadcast/{broadcastTournamentId}/players/{playerId} | Get a player of a broadcast
[**BroadcastPlayersGet**](BroadcastsAPI.md#BroadcastPlayersGet) | **Get** /broadcast/{broadcastTournamentId}/players | Get players of a broadcast
[**BroadcastPush**](BroadcastsAPI.md#BroadcastPush) | **Post** /api/broadcast/round/{broadcastRoundId}/push | Push PGN to a broadcast round
[**BroadcastRoundCreate**](BroadcastsAPI.md#BroadcastRoundCreate) | **Post** /broadcast/{broadcastTournamentId}/new | Create a broadcast round
[**BroadcastRoundGet**](BroadcastsAPI.md#BroadcastRoundGet) | **Get** /api/broadcast/{broadcastTournamentSlug}/{broadcastRoundSlug}/{broadcastRoundId} | Get a broadcast round
[**BroadcastRoundPgn**](BroadcastsAPI.md#BroadcastRoundPgn) | **Get** /api/broadcast/round/{broadcastRoundId}.pgn | Export one round as PGN
[**BroadcastRoundReset**](BroadcastsAPI.md#BroadcastRoundReset) | **Post** /api/broadcast/round/{broadcastRoundId}/reset | Reset a broadcast round
[**BroadcastRoundUpdate**](BroadcastsAPI.md#BroadcastRoundUpdate) | **Post** /broadcast/round/{broadcastRoundId}/edit | Update a broadcast round
[**BroadcastStreamRoundPgn**](BroadcastsAPI.md#BroadcastStreamRoundPgn) | **Get** /api/stream/broadcast/round/{broadcastRoundId}.pgn | Stream an ongoing broadcast round as PGN
[**BroadcastTeamLeaderboardGet**](BroadcastsAPI.md#BroadcastTeamLeaderboardGet) | **Get** /broadcast/{broadcastTournamentId}/teams/standings | Get the team leaderboard of a broadcast
[**BroadcastTourCreate**](BroadcastsAPI.md#BroadcastTourCreate) | **Post** /broadcast/new | Create a broadcast tournament
[**BroadcastTourGet**](BroadcastsAPI.md#BroadcastTourGet) | **Get** /api/broadcast/{broadcastTournamentId} | Get a broadcast tournament
[**BroadcastTourUpdate**](BroadcastsAPI.md#BroadcastTourUpdate) | **Post** /broadcast/{broadcastTournamentId}/edit | Update your broadcast tournament
[**BroadcastsByUser**](BroadcastsAPI.md#BroadcastsByUser) | **Get** /api/broadcast/by/{username} | Get broadcasts created by a user
[**BroadcastsOfficial**](BroadcastsAPI.md#BroadcastsOfficial) | **Get** /api/broadcast | Get official broadcasts
[**BroadcastsSearch**](BroadcastsAPI.md#BroadcastsSearch) | **Get** /api/broadcast/search | Search broadcasts
[**BroadcastsTop**](BroadcastsAPI.md#BroadcastsTop) | **Get** /api/broadcast/top | Get paginated top broadcast previews



## BroadcastAllRoundsPgn

> string BroadcastAllRoundsPgn(ctx, broadcastTournamentId).Clocks(clocks).Comments(comments).Execute()

Export all rounds as PGN



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] }`  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastAllRoundsPgn(context.Background(), broadcastTournamentId).Clocks(clocks).Comments(comments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastAllRoundsPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastAllRoundsPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastAllRoundsPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastAllRoundsPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60;  | [default to true]

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


## BroadcastMyRoundsGet

> BroadcastRoundCreate200Response BroadcastMyRoundsGet(ctx).Nb(nb).Execute()

Get your broadcast rounds



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
	nb := int32(20) // int32 | How many rounds to get (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastMyRoundsGet(context.Background()).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastMyRoundsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastMyRoundsGet`: BroadcastRoundCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastMyRoundsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastMyRoundsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | **int32** | How many rounds to get | 

### Return type

[**BroadcastRoundCreate200Response**](BroadcastRoundCreate200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastPlayerGet

> BroadcastPlayerGet200Response BroadcastPlayerGet(ctx, broadcastTournamentId, playerId).Execute()

Get a player of a broadcast



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID
	playerId := "playerId_example" // string | The unique player ID within the broadcast. This is usually their fideId.  If the player does not have a fideId, it is their name. Consult the [list of players for the broadcast](#tag/broadcasts/GET/broadcast/{broadcastTournamentId}/players) for which ID to use. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastPlayerGet(context.Background(), broadcastTournamentId, playerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastPlayerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastPlayerGet`: BroadcastPlayerGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastPlayerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 
**playerId** | **string** | The unique player ID within the broadcast. This is usually their fideId.  If the player does not have a fideId, it is their name. Consult the [list of players for the broadcast](#tag/broadcasts/GET/broadcast/{broadcastTournamentId}/players) for which ID to use.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastPlayerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BroadcastPlayerGet200Response**](BroadcastPlayerGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastPlayersGet

> []BroadcastPlayersGet200ResponseInner BroadcastPlayersGet(ctx, broadcastTournamentId).Execute()

Get players of a broadcast



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastPlayersGet(context.Background(), broadcastTournamentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastPlayersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastPlayersGet`: []BroadcastPlayersGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastPlayersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastPlayersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BroadcastPlayersGet200ResponseInner**](BroadcastPlayersGet200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastPush

> BroadcastPush200Response BroadcastPush(ctx, broadcastRoundId).Body(body).Execute()

Push PGN to a broadcast round



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
	broadcastRoundId := "broadcastRoundId_example" // string | The broadcast round ID
	body := "body_example" // string | The PGN. It can contain up to 100 games, separated by a double new line.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastPush(context.Background(), broadcastRoundId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastPush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastPush`: BroadcastPush200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastPush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | **string** | The broadcast round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The PGN. It can contain up to 100 games, separated by a double new line. | 

### Return type

[**BroadcastPush200Response**](BroadcastPush200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastRoundCreate

> BroadcastRoundCreate200Response BroadcastRoundCreate(ctx, broadcastTournamentId).Name(name).SyncUrl(syncUrl).SyncUrls(syncUrls).SyncIds(syncIds).SyncUsers(syncUsers).OnlyRound(onlyRound).Slices(slices).SyncSource(syncSource).StartsAt(startsAt).StartsAfterPrevious(startsAfterPrevious).Delay(delay).Status(status).Rated(rated).CustomScoringWhiteWin(customScoringWhiteWin).CustomScoringWhiteDraw(customScoringWhiteDraw).CustomScoringBlackWin(customScoringBlackWin).CustomScoringBlackDraw(customScoringBlackDraw).Period(period).Execute()

Create a broadcast round



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID
	name := "name_example" // string | Name of the broadcast round. Example: `Round 1` 
	syncUrl := "syncUrl_example" // string | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: ```txt https://myserver.org/myevent/round-10/games.pgn ``` 
	syncUrls := "syncUrls_example" // string | URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: ```txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn ``` 
	syncIds := "syncIds_example" // string | Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces. 
	syncUsers := "syncUsers_example" // string | Up to 100 Lichess usernames, separated by spaces 
	onlyRound := int32(56) // int32 | Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: ```txt [Round \\\"3\\\"] [Round \\\"3.1\\\"] ``` If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose `syncUrl` or `syncUrls` as the source.  (optional)
	slices := "slices_example" // string | Select slices of the games  Optional. Select games based on their position in the source. ```txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 ``` Slicing is done after filtering by round number.  It only works if you chose `syncUrl` or `syncUrls` as the source.  (optional)
	syncSource := "syncSource_example" // string | Where the games come from.  (optional) (default to "push")
	startsAt := int64(789) // int64 | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: `1356998400070`  (optional)
	startsAfterPrevious := true // bool | The start date is unknown, and the round will start automatically when the previous round completes.  (optional) (default to false)
	delay := int32(56) // int32 | Delay in seconds for movements to appear on the broadcast. Leave it empty if you don't need it. Example: `900` (15 min)  (optional)
	status := "status_example" // string | Lichess can usually detect the round status, but you can also set it manually if needed.  (optional) (default to "new")
	rated := true // bool | Whether the round is used when calculating players' rating changes. (optional) (default to true)
	customScoringWhiteWin := float32(8.14) // float32 |  (optional)
	customScoringWhiteDraw := float32(8.14) // float32 |  (optional)
	customScoringBlackWin := float32(8.14) // float32 |  (optional)
	customScoringBlackDraw := float32(8.14) // float32 |  (optional)
	period := int32(56) // int32 | (Only for Admins) Waiting time for each poll.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastRoundCreate(context.Background(), broadcastTournamentId).Name(name).SyncUrl(syncUrl).SyncUrls(syncUrls).SyncIds(syncIds).SyncUsers(syncUsers).OnlyRound(onlyRound).Slices(slices).SyncSource(syncSource).StartsAt(startsAt).StartsAfterPrevious(startsAfterPrevious).Delay(delay).Status(status).Rated(rated).CustomScoringWhiteWin(customScoringWhiteWin).CustomScoringWhiteDraw(customScoringWhiteDraw).CustomScoringBlackWin(customScoringBlackWin).CustomScoringBlackDraw(customScoringBlackDraw).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastRoundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastRoundCreate`: BroadcastRoundCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastRoundCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the broadcast round. Example: &#x60;Round 1&#x60;  | 
 **syncUrl** | **string** | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/games.pgn &#x60;&#x60;&#x60;  | 
 **syncUrls** | **string** | URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn &#x60;&#x60;&#x60;  | 
 **syncIds** | **string** | Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces.  | 
 **syncUsers** | **string** | Up to 100 Lichess usernames, separated by spaces  | 
 **onlyRound** | **int32** | Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: &#x60;&#x60;&#x60;txt [Round \\\&quot;3\\\&quot;] [Round \\\&quot;3.1\\\&quot;] &#x60;&#x60;&#x60; If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source.  | 
 **slices** | **string** | Select slices of the games  Optional. Select games based on their position in the source. &#x60;&#x60;&#x60;txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 &#x60;&#x60;&#x60; Slicing is done after filtering by round number.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source.  | 
 **syncSource** | **string** | Where the games come from.  | [default to &quot;push&quot;]
 **startsAt** | **int64** | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: &#x60;1356998400070&#x60;  | 
 **startsAfterPrevious** | **bool** | The start date is unknown, and the round will start automatically when the previous round completes.  | [default to false]
 **delay** | **int32** | Delay in seconds for movements to appear on the broadcast. Leave it empty if you don&#39;t need it. Example: &#x60;900&#x60; (15 min)  | 
 **status** | **string** | Lichess can usually detect the round status, but you can also set it manually if needed.  | [default to &quot;new&quot;]
 **rated** | **bool** | Whether the round is used when calculating players&#39; rating changes. | [default to true]
 **customScoringWhiteWin** | **float32** |  | 
 **customScoringWhiteDraw** | **float32** |  | 
 **customScoringBlackWin** | **float32** |  | 
 **customScoringBlackDraw** | **float32** |  | 
 **period** | **int32** | (Only for Admins) Waiting time for each poll.  | 

### Return type

[**BroadcastRoundCreate200Response**](BroadcastRoundCreate200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastRoundGet

> BroadcastRoundGet200Response BroadcastRoundGet(ctx, broadcastTournamentSlug, broadcastRoundSlug, broadcastRoundId).Execute()

Get a broadcast round



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
	broadcastTournamentSlug := "broadcastTournamentSlug_example" // string | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
	broadcastRoundSlug := "broadcastRoundSlug_example" // string | The broadcast round slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
	broadcastRoundId := "broadcastRoundId_example" // string | The broadcast Round ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastRoundGet(context.Background(), broadcastTournamentSlug, broadcastRoundSlug, broadcastRoundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastRoundGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastRoundGet`: BroadcastRoundGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastRoundGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentSlug** | **string** | The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by &#x60;-&#x60;. Only the &#x60;broadcastRoundId&#x60; is actually used. | 
**broadcastRoundSlug** | **string** | The broadcast round slug. Only used for SEO, the slug can be safely replaced by &#x60;-&#x60;. Only the &#x60;broadcastRoundId&#x60; is actually used. | 
**broadcastRoundId** | **string** | The broadcast Round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BroadcastRoundGet200Response**](BroadcastRoundGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastRoundPgn

> string BroadcastRoundPgn(ctx, broadcastRoundId).Clocks(clocks).Comments(comments).Execute()

Export one round as PGN



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
	broadcastRoundId := "broadcastRoundId_example" // string | The round ID
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] }`  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastRoundPgn(context.Background(), broadcastRoundId).Clocks(clocks).Comments(comments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastRoundPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastRoundPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastRoundPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | **string** | The round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60;  | [default to true]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastRoundReset

> AccountKidPost200Response BroadcastRoundReset(ctx, broadcastRoundId).Execute()

Reset a broadcast round



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
	broadcastRoundId := "broadcastRoundId_example" // string | The broadcast round ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastRoundReset(context.Background(), broadcastRoundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastRoundReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastRoundReset`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastRoundReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | **string** | The broadcast round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundResetRequest struct via the builder pattern


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


## BroadcastRoundUpdate

> BroadcastRoundUpdate200Response BroadcastRoundUpdate(ctx, broadcastRoundId).Name(name).SyncUrl(syncUrl).SyncUrls(syncUrls).SyncIds(syncIds).SyncUsers(syncUsers).Patch(patch).OnlyRound(onlyRound).Slices(slices).SyncSource(syncSource).StartsAt(startsAt).StartsAfterPrevious(startsAfterPrevious).Delay(delay).Status(status).Rated(rated).CustomScoringWhiteWin(customScoringWhiteWin).CustomScoringWhiteDraw(customScoringWhiteDraw).CustomScoringBlackWin(customScoringBlackWin).CustomScoringBlackDraw(customScoringBlackDraw).Period(period).Execute()

Update a broadcast round



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
	broadcastRoundId := "broadcastRoundId_example" // string | The broadcast round ID
	name := "name_example" // string | Name of the broadcast round. Example: `Round 1` 
	syncUrl := "syncUrl_example" // string | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: ```txt https://myserver.org/myevent/round-10/games.pgn ``` 
	syncUrls := "syncUrls_example" // string | URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: ```txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn ``` 
	syncIds := "syncIds_example" // string | Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces. 
	syncUsers := "syncUsers_example" // string | Up to 100 Lichess usernames, separated by spaces 
	patch := true // bool | Only update the provided fields, leaving others unchanged (optional)
	onlyRound := int32(56) // int32 | Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: ```txt [Round \\\"3\\\"] [Round \\\"3.1\\\"] ``` If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose `syncUrl` or `syncUrls` as the source.  (optional)
	slices := "slices_example" // string | Select slices of the games  Optional. Select games based on their position in the source. ```txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 ``` Slicing is done after filtering by round number.  It only works if you chose `syncUrl` or `syncUrls` as the source.  (optional)
	syncSource := "syncSource_example" // string | Where the games come from.  (optional) (default to "push")
	startsAt := int64(789) // int64 | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: `1356998400070`  (optional)
	startsAfterPrevious := true // bool | The start date is unknown, and the round will start automatically when the previous round completes.  (optional) (default to false)
	delay := int32(56) // int32 | Delay in seconds for movements to appear on the broadcast. Leave it empty if you don't need it. Example: `900` (15 min)  (optional)
	status := "status_example" // string | Lichess can usually detect the round status, but you can also set it manually if needed.  (optional) (default to "new")
	rated := true // bool | Whether the round is used when calculating players' rating changes. (optional) (default to true)
	customScoringWhiteWin := float32(8.14) // float32 |  (optional)
	customScoringWhiteDraw := float32(8.14) // float32 |  (optional)
	customScoringBlackWin := float32(8.14) // float32 |  (optional)
	customScoringBlackDraw := float32(8.14) // float32 |  (optional)
	period := int32(56) // int32 | (Only for Admins) Waiting time for each poll.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastRoundUpdate(context.Background(), broadcastRoundId).Name(name).SyncUrl(syncUrl).SyncUrls(syncUrls).SyncIds(syncIds).SyncUsers(syncUsers).Patch(patch).OnlyRound(onlyRound).Slices(slices).SyncSource(syncSource).StartsAt(startsAt).StartsAfterPrevious(startsAfterPrevious).Delay(delay).Status(status).Rated(rated).CustomScoringWhiteWin(customScoringWhiteWin).CustomScoringWhiteDraw(customScoringWhiteDraw).CustomScoringBlackWin(customScoringBlackWin).CustomScoringBlackDraw(customScoringBlackDraw).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastRoundUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastRoundUpdate`: BroadcastRoundUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastRoundUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | **string** | The broadcast round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastRoundUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the broadcast round. Example: &#x60;Round 1&#x60;  | 
 **syncUrl** | **string** | URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/games.pgn &#x60;&#x60;&#x60;  | 
 **syncUrls** | **string** | URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn &#x60;&#x60;&#x60;  | 
 **syncIds** | **string** | Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces.  | 
 **syncUsers** | **string** | Up to 100 Lichess usernames, separated by spaces  | 
 **patch** | **bool** | Only update the provided fields, leaving others unchanged | 
 **onlyRound** | **int32** | Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: &#x60;&#x60;&#x60;txt [Round \\\&quot;3\\\&quot;] [Round \\\&quot;3.1\\\&quot;] &#x60;&#x60;&#x60; If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source.  | 
 **slices** | **string** | Select slices of the games  Optional. Select games based on their position in the source. &#x60;&#x60;&#x60;txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 &#x60;&#x60;&#x60; Slicing is done after filtering by round number.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source.  | 
 **syncSource** | **string** | Where the games come from.  | [default to &quot;push&quot;]
 **startsAt** | **int64** | Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: &#x60;1356998400070&#x60;  | 
 **startsAfterPrevious** | **bool** | The start date is unknown, and the round will start automatically when the previous round completes.  | [default to false]
 **delay** | **int32** | Delay in seconds for movements to appear on the broadcast. Leave it empty if you don&#39;t need it. Example: &#x60;900&#x60; (15 min)  | 
 **status** | **string** | Lichess can usually detect the round status, but you can also set it manually if needed.  | [default to &quot;new&quot;]
 **rated** | **bool** | Whether the round is used when calculating players&#39; rating changes. | [default to true]
 **customScoringWhiteWin** | **float32** |  | 
 **customScoringWhiteDraw** | **float32** |  | 
 **customScoringBlackWin** | **float32** |  | 
 **customScoringBlackDraw** | **float32** |  | 
 **period** | **int32** | (Only for Admins) Waiting time for each poll.  | 

### Return type

[**BroadcastRoundUpdate200Response**](BroadcastRoundUpdate200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastStreamRoundPgn

> string BroadcastStreamRoundPgn(ctx, broadcastRoundId).Clocks(clocks).Comments(comments).Execute()

Stream an ongoing broadcast round as PGN



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
	broadcastRoundId := "broadcastRoundId_example" // string | The broadcast round ID
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] }`  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastStreamRoundPgn(context.Background(), broadcastRoundId).Clocks(clocks).Comments(comments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastStreamRoundPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastStreamRoundPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastStreamRoundPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastRoundId** | **string** | The broadcast round ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastStreamRoundPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60;  | [default to true]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-chess-pgn

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTeamLeaderboardGet

> []BroadcastTeamLeaderboardGet200ResponseInner BroadcastTeamLeaderboardGet(ctx, broadcastTournamentId).Execute()

Get the team leaderboard of a broadcast



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastTeamLeaderboardGet(context.Background(), broadcastTournamentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastTeamLeaderboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastTeamLeaderboardGet`: []BroadcastTeamLeaderboardGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastTeamLeaderboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTeamLeaderboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BroadcastTeamLeaderboardGet200ResponseInner**](BroadcastTeamLeaderboardGet200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTourCreate

> BroadcastsOfficial200Response BroadcastTourCreate(ctx).Name(name).InfoFormat(infoFormat).InfoLocation(infoLocation).InfoTc(infoTc).InfoFideTC(infoFideTC).InfoTimeZone(infoTimeZone).InfoPlayers(infoPlayers).InfoWebsite(infoWebsite).InfoStandings(infoStandings).Markdown(markdown).ShowScores(showScores).ShowRatingDiffs(showRatingDiffs).TeamTable(teamTable).Visibility(visibility).Players(players).Teams(teams).Tier(tier).Tiebreaks(tiebreaks).Execute()

Create a broadcast tournament



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
	name := "name_example" // string | Name of the broadcast tournament.  Example: `Sinquefield Cup` 
	infoFormat := "infoFormat_example" // string | Tournament format. Example: `\\\"8-player round-robin\\\" or \\\"5-round Swiss\\\"`  (optional)
	infoLocation := "infoLocation_example" // string | Tournament Location  (optional)
	infoTc := "infoTc_example" // string | Time control. Example: `\\\"Classical\\\" or \\\"Rapid\\\" or \\\"Rapid & Blitz\\\"`  (optional)
	infoFideTC := "infoFideTC_example" // string | FIDE rating category  (optional)
	infoTimeZone := "infoTimeZone_example" // string | Timezone of the tournament. Example: `America/New_York`. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more.  (optional)
	infoPlayers := "infoPlayers_example" // string | Mention up to 4 of the best players participating.  (optional)
	infoWebsite := "infoWebsite_example" // string | Official website. External website URL  (optional)
	infoStandings := "infoStandings_example" // string | Official Standings. External website URL, e.g. chess-results.com, info64.org  (optional)
	markdown := "markdown_example" // string | Optional long description of the broadcast. Markdown is supported.  (optional)
	showScores := true // bool | Show players scores based on game results  (optional) (default to false)
	showRatingDiffs := true // bool | Show player's rating diffs  (optional) (default to false)
	teamTable := true // bool | Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags.  (optional) (default to false)
	visibility := "visibility_example" // string | Who can view the broadcast. * `public`: Default. Anyone can view the broadcast * `unlisted`: Only people with the link can view the broadcast * `private`: Only the broadcast owner(s) can view the broadcast  (optional) (default to "public")
	players := "players_example" // string | Optional replace player names, ratings and titles.  One line per player, formatted as such:  ```txt player name / FIDE ID ```  Example:  ```txt Magnus Carlsen / 1503014 ```  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\"Jorge Rick Vito\\\" will match \\\"Jorge Rick\\\", \\\"jorge vito\\\", \\\"Rick, Vito\\\", etc.  If the player is NM or WNM, you can:  ```txt player name / FIDE ID / title ```  Alternatively, you may set tags manually, like so:  ```txt player name / rating / title / new name ```  All values are optional. Example: ```txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt ```  (optional)
	teams := "teams_example" // string | Optional: assign players to teams  One line per player, formatted as such: ```txt Team name; Fide Id or Player name ```  Example: ```txt Team Cats ; 3408230 Team Dogs ; Scooby Doo ```  By default the PGN tags WhiteTeam and BlackTeam are used.  (optional)
	tier := int32(56) // int32 | Optional, for Lichess admins only, used to feature on /broadcast.  * `3` for Official: normal tier * `4` for Official: high tier * `5` for Official: best tier  (optional)
	tiebreaks := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastTourCreate(context.Background()).Name(name).InfoFormat(infoFormat).InfoLocation(infoLocation).InfoTc(infoTc).InfoFideTC(infoFideTC).InfoTimeZone(infoTimeZone).InfoPlayers(infoPlayers).InfoWebsite(infoWebsite).InfoStandings(infoStandings).Markdown(markdown).ShowScores(showScores).ShowRatingDiffs(showRatingDiffs).TeamTable(teamTable).Visibility(visibility).Players(players).Teams(teams).Tier(tier).Tiebreaks(tiebreaks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastTourCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastTourCreate`: BroadcastsOfficial200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastTourCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the broadcast tournament.  Example: &#x60;Sinquefield Cup&#x60;  | 
 **infoFormat** | **string** | Tournament format. Example: &#x60;\\\&quot;8-player round-robin\\\&quot; or \\\&quot;5-round Swiss\\\&quot;&#x60;  | 
 **infoLocation** | **string** | Tournament Location  | 
 **infoTc** | **string** | Time control. Example: &#x60;\\\&quot;Classical\\\&quot; or \\\&quot;Rapid\\\&quot; or \\\&quot;Rapid &amp; Blitz\\\&quot;&#x60;  | 
 **infoFideTC** | **string** | FIDE rating category  | 
 **infoTimeZone** | **string** | Timezone of the tournament. Example: &#x60;America/New_York&#x60;. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more.  | 
 **infoPlayers** | **string** | Mention up to 4 of the best players participating.  | 
 **infoWebsite** | **string** | Official website. External website URL  | 
 **infoStandings** | **string** | Official Standings. External website URL, e.g. chess-results.com, info64.org  | 
 **markdown** | **string** | Optional long description of the broadcast. Markdown is supported.  | 
 **showScores** | **bool** | Show players scores based on game results  | [default to false]
 **showRatingDiffs** | **bool** | Show player&#39;s rating diffs  | [default to false]
 **teamTable** | **bool** | Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags.  | [default to false]
 **visibility** | **string** | Who can view the broadcast. * &#x60;public&#x60;: Default. Anyone can view the broadcast * &#x60;unlisted&#x60;: Only people with the link can view the broadcast * &#x60;private&#x60;: Only the broadcast owner(s) can view the broadcast  | [default to &quot;public&quot;]
 **players** | **string** | Optional replace player names, ratings and titles.  One line per player, formatted as such:  &#x60;&#x60;&#x60;txt player name / FIDE ID &#x60;&#x60;&#x60;  Example:  &#x60;&#x60;&#x60;txt Magnus Carlsen / 1503014 &#x60;&#x60;&#x60;  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\&quot;Jorge Rick Vito\\\&quot; will match \\\&quot;Jorge Rick\\\&quot;, \\\&quot;jorge vito\\\&quot;, \\\&quot;Rick, Vito\\\&quot;, etc.  If the player is NM or WNM, you can:  &#x60;&#x60;&#x60;txt player name / FIDE ID / title &#x60;&#x60;&#x60;  Alternatively, you may set tags manually, like so:  &#x60;&#x60;&#x60;txt player name / rating / title / new name &#x60;&#x60;&#x60;  All values are optional. Example: &#x60;&#x60;&#x60;txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt &#x60;&#x60;&#x60;  | 
 **teams** | **string** | Optional: assign players to teams  One line per player, formatted as such: &#x60;&#x60;&#x60;txt Team name; Fide Id or Player name &#x60;&#x60;&#x60;  Example: &#x60;&#x60;&#x60;txt Team Cats ; 3408230 Team Dogs ; Scooby Doo &#x60;&#x60;&#x60;  By default the PGN tags WhiteTeam and BlackTeam are used.  | 
 **tier** | **int32** | Optional, for Lichess admins only, used to feature on /broadcast.  * &#x60;3&#x60; for Official: normal tier * &#x60;4&#x60; for Official: high tier * &#x60;5&#x60; for Official: best tier  | 
 **tiebreaks** | **[]string** |  | 

### Return type

[**BroadcastsOfficial200Response**](BroadcastsOfficial200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTourGet

> BroadcastTourGet200Response BroadcastTourGet(ctx, broadcastTournamentId).Execute()

Get a broadcast tournament



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast tournament ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastTourGet(context.Background(), broadcastTournamentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastTourGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastTourGet`: BroadcastTourGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastTourGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BroadcastTourGet200Response**](BroadcastTourGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastTourUpdate

> AccountKidPost200Response BroadcastTourUpdate(ctx, broadcastTournamentId).Name(name).InfoFormat(infoFormat).InfoLocation(infoLocation).InfoTc(infoTc).InfoFideTC(infoFideTC).InfoTimeZone(infoTimeZone).InfoPlayers(infoPlayers).InfoWebsite(infoWebsite).InfoStandings(infoStandings).Markdown(markdown).ShowScores(showScores).ShowRatingDiffs(showRatingDiffs).TeamTable(teamTable).Visibility(visibility).Players(players).Teams(teams).Tier(tier).Tiebreaks(tiebreaks).Execute()

Update your broadcast tournament



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
	broadcastTournamentId := "broadcastTournamentId_example" // string | The broadcast ID
	name := "name_example" // string | Name of the broadcast tournament.  Example: `Sinquefield Cup` 
	infoFormat := "infoFormat_example" // string | Tournament format. Example: `\\\"8-player round-robin\\\" or \\\"5-round Swiss\\\"`  (optional)
	infoLocation := "infoLocation_example" // string | Tournament Location  (optional)
	infoTc := "infoTc_example" // string | Time control. Example: `\\\"Classical\\\" or \\\"Rapid\\\" or \\\"Rapid & Blitz\\\"`  (optional)
	infoFideTC := "infoFideTC_example" // string | FIDE rating category  (optional)
	infoTimeZone := "infoTimeZone_example" // string | Timezone of the tournament. Example: `America/New_York`. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more.  (optional)
	infoPlayers := "infoPlayers_example" // string | Mention up to 4 of the best players participating.  (optional)
	infoWebsite := "infoWebsite_example" // string | Official website. External website URL  (optional)
	infoStandings := "infoStandings_example" // string | Official Standings. External website URL, e.g. chess-results.com, info64.org  (optional)
	markdown := "markdown_example" // string | Optional long description of the broadcast. Markdown is supported.  (optional)
	showScores := true // bool | Show players scores based on game results  (optional) (default to false)
	showRatingDiffs := true // bool | Show player's rating diffs  (optional) (default to false)
	teamTable := true // bool | Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags.  (optional) (default to false)
	visibility := "visibility_example" // string | Who can view the broadcast. * `public`: Default. Anyone can view the broadcast * `unlisted`: Only people with the link can view the broadcast * `private`: Only the broadcast owner(s) can view the broadcast  (optional) (default to "public")
	players := "players_example" // string | Optional replace player names, ratings and titles.  One line per player, formatted as such:  ```txt player name / FIDE ID ```  Example:  ```txt Magnus Carlsen / 1503014 ```  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\"Jorge Rick Vito\\\" will match \\\"Jorge Rick\\\", \\\"jorge vito\\\", \\\"Rick, Vito\\\", etc.  If the player is NM or WNM, you can:  ```txt player name / FIDE ID / title ```  Alternatively, you may set tags manually, like so:  ```txt player name / rating / title / new name ```  All values are optional. Example: ```txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt ```  (optional)
	teams := "teams_example" // string | Optional: assign players to teams  One line per player, formatted as such: ```txt Team name; Fide Id or Player name ```  Example: ```txt Team Cats ; 3408230 Team Dogs ; Scooby Doo ```  By default the PGN tags WhiteTeam and BlackTeam are used.  (optional)
	tier := int32(56) // int32 | Optional, for Lichess admins only, used to feature on /broadcast.  * `3` for Official: normal tier * `4` for Official: high tier * `5` for Official: best tier  (optional)
	tiebreaks := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastTourUpdate(context.Background(), broadcastTournamentId).Name(name).InfoFormat(infoFormat).InfoLocation(infoLocation).InfoTc(infoTc).InfoFideTC(infoFideTC).InfoTimeZone(infoTimeZone).InfoPlayers(infoPlayers).InfoWebsite(infoWebsite).InfoStandings(infoStandings).Markdown(markdown).ShowScores(showScores).ShowRatingDiffs(showRatingDiffs).TeamTable(teamTable).Visibility(visibility).Players(players).Teams(teams).Tier(tier).Tiebreaks(tiebreaks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastTourUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastTourUpdate`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastTourUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**broadcastTournamentId** | **string** | The broadcast ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastTourUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the broadcast tournament.  Example: &#x60;Sinquefield Cup&#x60;  | 
 **infoFormat** | **string** | Tournament format. Example: &#x60;\\\&quot;8-player round-robin\\\&quot; or \\\&quot;5-round Swiss\\\&quot;&#x60;  | 
 **infoLocation** | **string** | Tournament Location  | 
 **infoTc** | **string** | Time control. Example: &#x60;\\\&quot;Classical\\\&quot; or \\\&quot;Rapid\\\&quot; or \\\&quot;Rapid &amp; Blitz\\\&quot;&#x60;  | 
 **infoFideTC** | **string** | FIDE rating category  | 
 **infoTimeZone** | **string** | Timezone of the tournament. Example: &#x60;America/New_York&#x60;. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more.  | 
 **infoPlayers** | **string** | Mention up to 4 of the best players participating.  | 
 **infoWebsite** | **string** | Official website. External website URL  | 
 **infoStandings** | **string** | Official Standings. External website URL, e.g. chess-results.com, info64.org  | 
 **markdown** | **string** | Optional long description of the broadcast. Markdown is supported.  | 
 **showScores** | **bool** | Show players scores based on game results  | [default to false]
 **showRatingDiffs** | **bool** | Show player&#39;s rating diffs  | [default to false]
 **teamTable** | **bool** | Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags.  | [default to false]
 **visibility** | **string** | Who can view the broadcast. * &#x60;public&#x60;: Default. Anyone can view the broadcast * &#x60;unlisted&#x60;: Only people with the link can view the broadcast * &#x60;private&#x60;: Only the broadcast owner(s) can view the broadcast  | [default to &quot;public&quot;]
 **players** | **string** | Optional replace player names, ratings and titles.  One line per player, formatted as such:  &#x60;&#x60;&#x60;txt player name / FIDE ID &#x60;&#x60;&#x60;  Example:  &#x60;&#x60;&#x60;txt Magnus Carlsen / 1503014 &#x60;&#x60;&#x60;  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\&quot;Jorge Rick Vito\\\&quot; will match \\\&quot;Jorge Rick\\\&quot;, \\\&quot;jorge vito\\\&quot;, \\\&quot;Rick, Vito\\\&quot;, etc.  If the player is NM or WNM, you can:  &#x60;&#x60;&#x60;txt player name / FIDE ID / title &#x60;&#x60;&#x60;  Alternatively, you may set tags manually, like so:  &#x60;&#x60;&#x60;txt player name / rating / title / new name &#x60;&#x60;&#x60;  All values are optional. Example: &#x60;&#x60;&#x60;txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt &#x60;&#x60;&#x60;  | 
 **teams** | **string** | Optional: assign players to teams  One line per player, formatted as such: &#x60;&#x60;&#x60;txt Team name; Fide Id or Player name &#x60;&#x60;&#x60;  Example: &#x60;&#x60;&#x60;txt Team Cats ; 3408230 Team Dogs ; Scooby Doo &#x60;&#x60;&#x60;  By default the PGN tags WhiteTeam and BlackTeam are used.  | 
 **tier** | **int32** | Optional, for Lichess admins only, used to feature on /broadcast.  * &#x60;3&#x60; for Official: normal tier * &#x60;4&#x60; for Official: high tier * &#x60;5&#x60; for Official: best tier  | 
 **tiebreaks** | **[]string** |  | 

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


## BroadcastsByUser

> BroadcastsByUser200Response BroadcastsByUser(ctx, username).Page(page).Html(html).Execute()

Get broadcasts created by a user



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
	username := "username_example" // string | 
	page := int32(1) // int32 |  (optional) (default to 1)
	html := true // bool | Convert the \"description\" field from markdown to HTML (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastsByUser(context.Background(), username).Page(page).Html(html).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastsByUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastsByUser`: BroadcastsByUser200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastsByUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastsByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **html** | **bool** | Convert the \&quot;description\&quot; field from markdown to HTML | 

### Return type

[**BroadcastsByUser200Response**](BroadcastsByUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastsOfficial

> BroadcastsOfficial200Response BroadcastsOfficial(ctx).Nb(nb).Html(html).Live(live).Execute()

Get official broadcasts



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
	nb := int32(56) // int32 | Max number of broadcasts to fetch (optional) (default to 20)
	html := true // bool | Convert the \"description\" field from markdown to HTML (optional)
	live := true // bool | [Filter] only broadcasts where a round is ongoing, i.e. started and not finished (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastsOfficial(context.Background()).Nb(nb).Html(html).Live(live).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastsOfficial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastsOfficial`: BroadcastsOfficial200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastsOfficial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastsOfficialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nb** | **int32** | Max number of broadcasts to fetch | [default to 20]
 **html** | **bool** | Convert the \&quot;description\&quot; field from markdown to HTML | 
 **live** | **bool** | [Filter] only broadcasts where a round is ongoing, i.e. started and not finished | 

### Return type

[**BroadcastsOfficial200Response**](BroadcastsOfficial200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastsSearch

> BroadcastsSearch200Response BroadcastsSearch(ctx).Page(page).Q(q).Execute()

Search broadcasts



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
	page := int32(56) // int32 | Which page to fetch. (optional) (default to 1)
	q := "q_example" // string | Search term (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastsSearch(context.Background()).Page(page).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastsSearch`: BroadcastsSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Which page to fetch. | [default to 1]
 **q** | **string** | Search term | 

### Return type

[**BroadcastsSearch200Response**](BroadcastsSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BroadcastsTop

> BroadcastsTop200Response BroadcastsTop(ctx).Page(page).Html(html).Execute()

Get paginated top broadcast previews



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
	page := int32(56) // int32 | Which page to fetch. Only page 1 has \"active\" broadcasts. (optional) (default to 1)
	html := true // bool | Convert the \"description\" field from markdown to HTML (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BroadcastsAPI.BroadcastsTop(context.Background()).Page(page).Html(html).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BroadcastsAPI.BroadcastsTop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BroadcastsTop`: BroadcastsTop200Response
	fmt.Fprintf(os.Stdout, "Response from `BroadcastsAPI.BroadcastsTop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBroadcastsTopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Which page to fetch. Only page 1 has \&quot;active\&quot; broadcasts. | [default to 1]
 **html** | **bool** | Convert the \&quot;description\&quot; field from markdown to HTML | 

### Return type

[**BroadcastsTop200Response**](BroadcastsTop200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

