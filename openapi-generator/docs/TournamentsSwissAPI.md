# \TournamentsSwissAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSwissJoin**](TournamentsSwissAPI.md#ApiSwissJoin) | **Post** /api/swiss/{id}/join | Join a Swiss tournament
[**ApiSwissNew**](TournamentsSwissAPI.md#ApiSwissNew) | **Post** /api/swiss/new/{teamId} | Create a new Swiss tournament
[**ApiSwissScheduleNextRound**](TournamentsSwissAPI.md#ApiSwissScheduleNextRound) | **Post** /api/swiss/{id}/schedule-next-round | Manually schedule the next round
[**ApiSwissTerminate**](TournamentsSwissAPI.md#ApiSwissTerminate) | **Post** /api/swiss/{id}/terminate | Terminate a Swiss tournament
[**ApiSwissUpdate**](TournamentsSwissAPI.md#ApiSwissUpdate) | **Post** /api/swiss/{id}/edit | Update a Swiss tournament
[**ApiSwissWithdraw**](TournamentsSwissAPI.md#ApiSwissWithdraw) | **Post** /api/swiss/{id}/withdraw | Pause or leave a swiss tournament
[**ApiTeamSwiss**](TournamentsSwissAPI.md#ApiTeamSwiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
[**GamesBySwiss**](TournamentsSwissAPI.md#GamesBySwiss) | **Get** /api/swiss/{id}/games | Export games of a Swiss tournament
[**ResultsBySwiss**](TournamentsSwissAPI.md#ResultsBySwiss) | **Get** /api/swiss/{id}/results | Get results of a swiss tournament
[**Swiss**](TournamentsSwissAPI.md#Swiss) | **Get** /api/swiss/{id} | Get info about a Swiss tournament
[**SwissTrf**](TournamentsSwissAPI.md#SwissTrf) | **Get** /swiss/{id}.trf | Export TRF of a Swiss tournament



## ApiSwissJoin

> AccountKidPost200Response ApiSwissJoin(ctx, id).Password(password).Execute()

Join a Swiss tournament



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
	id := "hL7vMrFQ" // string | The tournament ID.
	password := "password_example" // string | The tournament password, if one is required (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiSwissJoin(context.Background(), id).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSwissJoin`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiSwissJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | **string** | The tournament password, if one is required | 

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


## ApiSwissNew

> ApiSwissNew200Response ApiSwissNew(ctx, teamId).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Position(position).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsPlayYourGames(conditionsPlayYourGames).ConditionsAllowList(conditionsAllowList).Execute()

Create a new Swiss tournament



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
	teamId := "teamId_example" // string | ID of the team
	clockLimit := int32(56) // int32 | Clock initial time in seconds
	clockIncrement := int32(56) // int32 | Clock increment in seconds
	nbRounds := int32(56) // int32 | Maximum number of rounds to play
	name := "name_example" // string | The tournament name. Leave empty to get a random Grandmaster name (optional)
	startsAt := int64(789) // int64 | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. (optional)
	roundInterval := int32(56) // int32 | How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI. If empty or -1, a sensible value is picked automatically.  (optional)
	variant := "variant_example" // string |  (optional) (default to "standard")
	position := "position_example" // string | Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	description := "description_example" // string | Anything you want to tell players about the tournament (optional)
	rated := true // bool | Games are rated and impact players ratings (optional) (default to true)
	password := "password_example" // string | Make the tournament private and restrict access with a password. (optional)
	forbiddenPairings := "forbiddenPairings_example" // string | Usernames of players that must not play together. Two usernames per line, separated by a space.  (optional)
	manualPairings := "manualPairings_example" // string | Manual pairings for the next round. Two usernames per line, separated by a space. Example: ``` PlayerA PlayerB PlayerC PlayerD ``` To give a bye (1 point) to a player instead of a pairing, add a line like so: ``` PlayerE 1 ``` Missing players will be considered absent and get zero points.  (optional)
	chatFor := int32(56) // int32 | Who can read and write in the chat. - 0  = No-one - 10 = Only team leaders - 20 = Only team members - 30 = All Lichess players  (optional) (default to 20)
	conditionsMinRatingRating := int32(56) // int32 | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
	conditionsMaxRatingRating := int32(56) // int32 | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
	conditionsNbRatedGameNb := int32(56) // int32 | Minimum number of rated games required to join. (optional)
	conditionsPlayYourGames := true // bool | Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won't be able to enter yours. This results in a better swiss experience for the players who actually show up.  (optional) (default to false)
	conditionsAllowList := "conditionsAllowList_example" // string | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiSwissNew(context.Background(), teamId).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Position(position).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsPlayYourGames(conditionsPlayYourGames).ConditionsAllowList(conditionsAllowList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSwissNew`: ApiSwissNew200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiSwissNew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockLimit** | **int32** | Clock initial time in seconds | 
 **clockIncrement** | **int32** | Clock increment in seconds | 
 **nbRounds** | **int32** | Maximum number of rounds to play | 
 **name** | **string** | The tournament name. Leave empty to get a random Grandmaster name | 
 **startsAt** | **int64** | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. | 
 **roundInterval** | **int32** | How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI. If empty or -1, a sensible value is picked automatically.  | 
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **position** | **string** | Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **description** | **string** | Anything you want to tell players about the tournament | 
 **rated** | **bool** | Games are rated and impact players ratings | [default to true]
 **password** | **string** | Make the tournament private and restrict access with a password. | 
 **forbiddenPairings** | **string** | Usernames of players that must not play together. Two usernames per line, separated by a space.  | 
 **manualPairings** | **string** | Manual pairings for the next round. Two usernames per line, separated by a space. Example: &#x60;&#x60;&#x60; PlayerA PlayerB PlayerC PlayerD &#x60;&#x60;&#x60; To give a bye (1 point) to a player instead of a pairing, add a line like so: &#x60;&#x60;&#x60; PlayerE 1 &#x60;&#x60;&#x60; Missing players will be considered absent and get zero points.  | 
 **chatFor** | **int32** | Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players  | [default to 20]
 **conditionsMinRatingRating** | **int32** | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | **int32** | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | **int32** | Minimum number of rated games required to join. | 
 **conditionsPlayYourGames** | **bool** | Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won&#39;t be able to enter yours. This results in a better swiss experience for the players who actually show up.  | [default to false]
 **conditionsAllowList** | **string** | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 

### Return type

[**ApiSwissNew200Response**](ApiSwissNew200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSwissScheduleNextRound

> ApiSwissScheduleNextRound(ctx, id).Date(date).Execute()

Manually schedule the next round



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
	id := "hL7vMrFQ" // string | The tournament ID.
	date := int64(789) // int64 | Timestamp in milliseconds to start the next round at a given date and time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TournamentsSwissAPI.ApiSwissScheduleNextRound(context.Background(), id).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissScheduleNextRound``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissScheduleNextRoundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **int64** | Timestamp in milliseconds to start the next round at a given date and time. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSwissTerminate

> AccountKidPost200Response ApiSwissTerminate(ctx, id).Execute()

Terminate a Swiss tournament



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
	id := "W5FrxusN" // string | The Swiss tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiSwissTerminate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissTerminate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSwissTerminate`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiSwissTerminate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Swiss tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissTerminateRequest struct via the builder pattern


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


## ApiSwissUpdate

> ApiSwissNew200Response ApiSwissUpdate(ctx, id).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Position(position).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsPlayYourGames(conditionsPlayYourGames).ConditionsAllowList(conditionsAllowList).Execute()

Update a Swiss tournament



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
	id := "hL7vMrFQ" // string | The tournament ID.
	clockLimit := int32(56) // int32 | Clock initial time in seconds
	clockIncrement := int32(56) // int32 | Clock increment in seconds
	nbRounds := int32(56) // int32 | Maximum number of rounds to play
	name := "name_example" // string | The tournament name. Leave empty to get a random Grandmaster name (optional)
	startsAt := int64(789) // int64 | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. (optional)
	roundInterval := int32(56) // int32 | How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI, or [with the API](#tag/tournaments-swiss/POST/api/swiss/{id}/schedule-next-round). If empty or -1, a sensible value is picked automatically.  (optional)
	variant := "variant_example" // string |  (optional) (default to "standard")
	position := "position_example" // string | Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	description := "description_example" // string | Anything you want to tell players about the tournament (optional)
	rated := true // bool | Games are rated and impact players ratings (optional) (default to true)
	password := "password_example" // string | Make the tournament private and restrict access with a password. (optional)
	forbiddenPairings := "forbiddenPairings_example" // string | Usernames of players that must not play together. Two usernames per line, separated by a space.  (optional)
	manualPairings := "manualPairings_example" // string | Manual pairings for the next round. Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  (optional)
	chatFor := int32(56) // int32 | Who can read and write in the chat. - 0  = No-one - 10 = Only team leaders - 20 = Only team members - 30 = All Lichess players  (optional) (default to 20)
	conditionsMinRatingRating := int32(56) // int32 | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
	conditionsMaxRatingRating := int32(56) // int32 | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
	conditionsNbRatedGameNb := int32(56) // int32 | Minimum number of rated games required to join. (optional)
	conditionsPlayYourGames := true // bool | Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won't be able to enter yours. This results in a better swiss experience for the players who actually show up.  (optional) (default to false)
	conditionsAllowList := "conditionsAllowList_example" // string | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiSwissUpdate(context.Background(), id).ClockLimit(clockLimit).ClockIncrement(clockIncrement).NbRounds(nbRounds).Name(name).StartsAt(startsAt).RoundInterval(roundInterval).Variant(variant).Position(position).Description(description).Rated(rated).Password(password).ForbiddenPairings(forbiddenPairings).ManualPairings(manualPairings).ChatFor(chatFor).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsPlayYourGames(conditionsPlayYourGames).ConditionsAllowList(conditionsAllowList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSwissUpdate`: ApiSwissNew200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiSwissUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockLimit** | **int32** | Clock initial time in seconds | 
 **clockIncrement** | **int32** | Clock increment in seconds | 
 **nbRounds** | **int32** | Maximum number of rounds to play | 
 **name** | **string** | The tournament name. Leave empty to get a random Grandmaster name | 
 **startsAt** | **int64** | Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation. | 
 **roundInterval** | **int32** | How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI, or [with the API](#tag/tournaments-swiss/POST/api/swiss/{id}/schedule-next-round). If empty or -1, a sensible value is picked automatically.  | 
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **position** | **string** | Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **description** | **string** | Anything you want to tell players about the tournament | 
 **rated** | **bool** | Games are rated and impact players ratings | [default to true]
 **password** | **string** | Make the tournament private and restrict access with a password. | 
 **forbiddenPairings** | **string** | Usernames of players that must not play together. Two usernames per line, separated by a space.  | 
 **manualPairings** | **string** | Manual pairings for the next round. Two usernames per line, separated by a space. Present players without a valid pairing will be given a bye, which is worth 1 point. Forfeited players will get 0 points.  | 
 **chatFor** | **int32** | Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players  | [default to 20]
 **conditionsMinRatingRating** | **int32** | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | **int32** | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | **int32** | Minimum number of rated games required to join. | 
 **conditionsPlayYourGames** | **bool** | Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won&#39;t be able to enter yours. This results in a better swiss experience for the players who actually show up.  | [default to false]
 **conditionsAllowList** | **string** | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 

### Return type

[**ApiSwissNew200Response**](ApiSwissNew200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSwissWithdraw

> AccountKidPost200Response ApiSwissWithdraw(ctx, id).Execute()

Pause or leave a swiss tournament



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
	id := "hL7vMrFQ" // string | The tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiSwissWithdraw(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiSwissWithdraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSwissWithdraw`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiSwissWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiSwissWithdrawRequest struct via the builder pattern


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


## ApiTeamSwiss

> ApiSwissNew200Response ApiTeamSwiss(ctx, teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()

Get team swiss tournaments



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
	teamId := "coders" // string | 
	max := int32(56) // int32 | How many tournaments to download. (optional) (default to 100)
	status := "status_example" // string | [Filter] Only swiss tournaments in this current state.  (optional)
	createdBy := "createdBy_example" // string | [Filter] Only swiss tournaments created by a given user.  (optional)
	name := "name_example" // string | [Filter] Only swiss tournaments with a given name.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ApiTeamSwiss(context.Background(), teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ApiTeamSwiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTeamSwiss`: ApiSwissNew200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ApiTeamSwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamSwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int32** | How many tournaments to download. | [default to 100]
 **status** | **string** | [Filter] Only swiss tournaments in this current state.  | 
 **createdBy** | **string** | [Filter] Only swiss tournaments created by a given user.  | 
 **name** | **string** | [Filter] Only swiss tournaments with a given name.  | 

### Return type

[**ApiSwissNew200Response**](ApiSwissNew200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesBySwiss

> string GamesBySwiss(ctx, id).Accept(accept).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Execute()

Export games of a Swiss tournament



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
	id := "id_example" // string | The tournament ID.
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	player := "player_example" // string | Only the games played by a given player (optional)
	moves := true // bool | Include the PGN moves. (optional) (default to true)
	pgnInJson := true // bool | Include the full PGN within the JSON response, in a `pgn` field. (optional) (default to false)
	tags := true // bool | Include the PGN tags. (optional) (default to true)
	clocks := true // bool | Include clock status when available. Either as PGN comments: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }` Or in a `clocks` JSON field, as centisecond integers, depending on the response type.  (optional) (default to false)
	evals := true // bool | Include analysis evaluations and comments, when available. Either as PGN comments: `12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }` Or in an `analysis` JSON field, depending on the response type.  (optional) (default to false)
	accuracy := true // bool | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  (optional) (default to false)
	opening := true // bool | Include the opening name. Example: `[Opening \"King's Gambit Accepted, King's Knight Gambit\"]`  (optional) (default to false)
	division := true // bool | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.GamesBySwiss(context.Background(), id).Accept(accept).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.GamesBySwiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesBySwiss`: string
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.GamesBySwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesBySwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **player** | **string** | Only the games played by a given player | 
 **moves** | **bool** | Include the PGN moves. | [default to true]
 **pgnInJson** | **bool** | Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. | [default to false]
 **tags** | **bool** | Include the PGN tags. | [default to true]
 **clocks** | **bool** | Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type.  | [default to false]
 **evals** | **bool** | Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type.  | [default to false]
 **accuracy** | **bool** | Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON.  | [default to false]
 **opening** | **bool** | Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60;  | [default to false]
 **division** | **bool** | Plies which mark the beginning of the middlegame and endgame. Only available in JSON  | [default to false]

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


## ResultsBySwiss

> ResultsBySwiss200Response ResultsBySwiss(ctx, id).Nb(nb).Execute()

Get results of a swiss tournament



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
	id := "id_example" // string | The tournament ID.
	nb := int32(56) // int32 | Max number of players to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.ResultsBySwiss(context.Background(), id).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.ResultsBySwiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResultsBySwiss`: ResultsBySwiss200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.ResultsBySwiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResultsBySwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | **int32** | Max number of players to fetch | 

### Return type

[**ResultsBySwiss200Response**](ResultsBySwiss200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Swiss

> ApiSwissNew200Response Swiss(ctx, id).Execute()

Get info about a Swiss tournament



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
	id := "id_example" // string | The Swiss tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.Swiss(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.Swiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Swiss`: ApiSwissNew200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.Swiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Swiss tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiSwissNew200Response**](ApiSwissNew200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwissTrf

> string SwissTrf(ctx, id).Execute()

Export TRF of a Swiss tournament



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
	id := "id_example" // string | The tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsSwissAPI.SwissTrf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsSwissAPI.SwissTrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwissTrf`: string
	fmt.Fprintf(os.Stdout, "Response from `TournamentsSwissAPI.SwissTrf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwissTrfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

