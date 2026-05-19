# \TournamentsArenaAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamArena**](TournamentsArenaAPI.md#ApiTeamArena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
[**ApiTournament**](TournamentsArenaAPI.md#ApiTournament) | **Get** /api/tournament | Get current tournaments
[**ApiTournamentJoin**](TournamentsArenaAPI.md#ApiTournamentJoin) | **Post** /api/tournament/{id}/join | Join an Arena tournament
[**ApiTournamentPost**](TournamentsArenaAPI.md#ApiTournamentPost) | **Post** /api/tournament | Create a new Arena tournament
[**ApiTournamentTeamBattlePost**](TournamentsArenaAPI.md#ApiTournamentTeamBattlePost) | **Post** /api/tournament/team-battle/{id} | Update a team battle
[**ApiTournamentTerminate**](TournamentsArenaAPI.md#ApiTournamentTerminate) | **Post** /api/tournament/{id}/terminate | Terminate an Arena tournament
[**ApiTournamentUpdate**](TournamentsArenaAPI.md#ApiTournamentUpdate) | **Post** /api/tournament/{id} | Update an Arena tournament
[**ApiTournamentWithdraw**](TournamentsArenaAPI.md#ApiTournamentWithdraw) | **Post** /api/tournament/{id}/withdraw | Pause or leave an Arena tournament
[**ApiUserNameTournamentCreated**](TournamentsArenaAPI.md#ApiUserNameTournamentCreated) | **Get** /api/user/{username}/tournament/created | Get tournaments created by a user
[**ApiUserNameTournamentPlayed**](TournamentsArenaAPI.md#ApiUserNameTournamentPlayed) | **Get** /api/user/{username}/tournament/played | Get tournaments played by a user
[**GamesByTournament**](TournamentsArenaAPI.md#GamesByTournament) | **Get** /api/tournament/{id}/games | Export games of an Arena tournament
[**ResultsByTournament**](TournamentsArenaAPI.md#ResultsByTournament) | **Get** /api/tournament/{id}/results | Get results of an Arena tournament
[**TeamsByTournament**](TournamentsArenaAPI.md#TeamsByTournament) | **Get** /api/tournament/{id}/teams | Get team standing of a team battle
[**Tournament**](TournamentsArenaAPI.md#Tournament) | **Get** /api/tournament/{id} | Get info about an Arena tournament



## ApiTeamArena

> ApiTournament200ResponseCreatedInner ApiTeamArena(ctx, teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()

Get team Arena tournaments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "teamId_example" // string | ID of the team
	max := int32(56) // int32 | How many tournaments to download. (optional) (default to 100)
	status := "status_example" // string | [Filter] Only arena tournaments in this current state.  (optional)
	createdBy := "createdBy_example" // string | [Filter] Only arena tournaments created by a given user.  (optional)
	name := "name_example" // string | [Filter] Only arena tournaments with a given name.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTeamArena(context.Background(), teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTeamArena``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTeamArena`: ApiTournament200ResponseCreatedInner
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTeamArena`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTeamArenaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int32** | How many tournaments to download. | [default to 100]
 **status** | **string** | [Filter] Only arena tournaments in this current state.  | 
 **createdBy** | **string** | [Filter] Only arena tournaments created by a given user.  | 
 **name** | **string** | [Filter] Only arena tournaments with a given name.  | 

### Return type

[**ApiTournament200ResponseCreatedInner**](ApiTournament200ResponseCreatedInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournament

> ApiTournament200Response ApiTournament(ctx).Execute()

Get current tournaments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournament(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournament`: ApiTournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournament`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentRequest struct via the builder pattern


### Return type

[**ApiTournament200Response**](ApiTournament200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournamentJoin

> AccountKidPost200Response ApiTournamentJoin(ctx, id).Password(password).Team(team).PairMeAsap(pairMeAsap).Execute()

Join an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "hL7vMrFQ" // string | The tournament ID.
	password := "password_example" // string | The tournament password, if one is required. Can also be a [user-specific entry code](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) generated and shared by the organizer.  (optional)
	team := "team_example" // string | The team to join the tournament with, for team battle tournaments (optional)
	pairMeAsap := true // bool | If the tournament is started, attempt to pair the user, even if they are not connected to the tournament page. This expires after one minute, to avoid pairing a user who is long gone. You may call \\\"join\\\" again to extend the waiting.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentJoin(context.Background(), id).Password(password).Team(team).PairMeAsap(pairMeAsap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentJoin`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **password** | **string** | The tournament password, if one is required. Can also be a [user-specific entry code](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) generated and shared by the organizer.  | 
 **team** | **string** | The team to join the tournament with, for team battle tournaments | 
 **pairMeAsap** | **bool** | If the tournament is started, attempt to pair the user, even if they are not connected to the tournament page. This expires after one minute, to avoid pairing a user who is long gone. You may call \\\&quot;join\\\&quot; again to extend the waiting.  | [default to false]

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


## ApiTournamentPost

> ApiTournamentPost200Response ApiTournamentPost(ctx).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).TeamBattleByTeam(teamBattleByTeam).ConditionsTeamMemberTeamId(conditionsTeamMemberTeamId).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).ConditionsBots(conditionsBots).ConditionsAccountAge(conditionsAccountAge).Execute()

Create a new Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clockTime := float32(8.14) // float32 | Clock initial time in minutes
	clockIncrement := int32(56) // int32 | Clock increment in seconds
	minutes := int32(56) // int32 | How long the tournament lasts, in minutes
	name := "name_example" // string | The tournament name. Leave empty to get a random Grandmaster name (optional)
	waitMinutes := int32(56) // int32 | How long to wait before starting the tournament, from now, in minutes (optional) (default to 5)
	startDate := int64(789) // int64 | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the `waitMinutes` setting (optional)
	variant := "variant_example" // string |  (optional) (default to "standard")
	rated := true // bool | Games are rated and impact players ratings (optional) (default to true)
	position := "position_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	berserkable := true // bool | Whether the players can use berserk. Only allowed if clockIncrement <= clockTime * 2 (optional) (default to true)
	streakable := true // bool | After 2 wins, consecutive wins grant 4 points instead of 2. (optional) (default to true)
	hasChat := true // bool | Whether the players can discuss in a chat (optional) (default to true)
	description := "description_example" // string | Anything you want to tell players about the tournament (optional)
	password := "password_example" // string | Make the tournament private, and restrict access with a password. You can also [generate user-specific entry codes](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) based on this password.  (optional)
	teamBattleByTeam := "teamBattleByTeam_example" // string | Set the ID of a team you lead to create a team battle. The other teams can be added using the [team battle edit endpoint](#tag/arena-tournaments/POST/api/tournament/team-battle/{id}).  (optional)
	conditionsTeamMemberTeamId := "conditionsTeamMemberTeamId_example" // string | Restrict entry to members of a team. The teamId is the last part of a team URL, e.g. `https://lichess.org/team/coders` has teamId = `coders`. Leave empty to let everyone join the tournament. Do not use this to create team battles, use `teamBattleByTeam` instead.  (optional)
	conditionsMinRatingRating := int32(56) // int32 | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
	conditionsMaxRatingRating := int32(56) // int32 | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
	conditionsNbRatedGameNb := int32(56) // int32 | Minimum number of rated games required to join. (optional)
	conditionsAllowList := "conditionsAllowList_example" // string | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)
	conditionsBots := true // bool | Whether bots are allowed to join the tournament. (optional) (default to false)
	conditionsAccountAge := int32(56) // int32 | Minium account age in days required to join. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentPost(context.Background()).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).TeamBattleByTeam(teamBattleByTeam).ConditionsTeamMemberTeamId(conditionsTeamMemberTeamId).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).ConditionsBots(conditionsBots).ConditionsAccountAge(conditionsAccountAge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentPost`: ApiTournamentPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clockTime** | **float32** | Clock initial time in minutes | 
 **clockIncrement** | **int32** | Clock increment in seconds | 
 **minutes** | **int32** | How long the tournament lasts, in minutes | 
 **name** | **string** | The tournament name. Leave empty to get a random Grandmaster name | 
 **waitMinutes** | **int32** | How long to wait before starting the tournament, from now, in minutes | [default to 5]
 **startDate** | **int64** | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting | 
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **rated** | **bool** | Games are rated and impact players ratings | [default to true]
 **position** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **berserkable** | **bool** | Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2 | [default to true]
 **streakable** | **bool** | After 2 wins, consecutive wins grant 4 points instead of 2. | [default to true]
 **hasChat** | **bool** | Whether the players can discuss in a chat | [default to true]
 **description** | **string** | Anything you want to tell players about the tournament | 
 **password** | **string** | Make the tournament private, and restrict access with a password. You can also [generate user-specific entry codes](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) based on this password.  | 
 **teamBattleByTeam** | **string** | Set the ID of a team you lead to create a team battle. The other teams can be added using the [team battle edit endpoint](#tag/arena-tournaments/POST/api/tournament/team-battle/{id}).  | 
 **conditionsTeamMemberTeamId** | **string** | Restrict entry to members of a team. The teamId is the last part of a team URL, e.g. &#x60;https://lichess.org/team/coders&#x60; has teamId &#x3D; &#x60;coders&#x60;. Leave empty to let everyone join the tournament. Do not use this to create team battles, use &#x60;teamBattleByTeam&#x60; instead.  | 
 **conditionsMinRatingRating** | **int32** | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | **int32** | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | **int32** | Minimum number of rated games required to join. | 
 **conditionsAllowList** | **string** | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 
 **conditionsBots** | **bool** | Whether bots are allowed to join the tournament. | [default to false]
 **conditionsAccountAge** | **int32** | Minium account age in days required to join. | 

### Return type

[**ApiTournamentPost200Response**](ApiTournamentPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournamentTeamBattlePost

> Tournament200Response ApiTournamentTeamBattlePost(ctx, id).Teams(teams).NbLeaders(nbLeaders).Execute()

Update a team battle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID
	teams := "teams_example" // string | All team IDs of the team battle, separated by commas. Make sure to always send the full list. Teams that are not in the list will be removed from the team battle. Example: `coders,zhigalko_sergei-fan-club,hhSwTKZv` 
	nbLeaders := int32(56) // int32 | Number team leaders per team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentTeamBattlePost(context.Background(), id).Teams(teams).NbLeaders(nbLeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentTeamBattlePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentTeamBattlePost`: Tournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentTeamBattlePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentTeamBattlePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teams** | **string** | All team IDs of the team battle, separated by commas. Make sure to always send the full list. Teams that are not in the list will be removed from the team battle. Example: &#x60;coders,zhigalko_sergei-fan-club,hhSwTKZv&#x60;  | 
 **nbLeaders** | **int32** | Number team leaders per team. | 

### Return type

[**Tournament200Response**](Tournament200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournamentTerminate

> AccountKidPost200Response ApiTournamentTerminate(ctx, id).Execute()

Terminate an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "hL7vMrFQ" // string | The tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentTerminate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentTerminate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentTerminate`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentTerminate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentTerminateRequest struct via the builder pattern


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


## ApiTournamentUpdate

> Tournament200Response ApiTournamentUpdate(ctx, id).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).ConditionsBots(conditionsBots).ConditionsAccountAge(conditionsAccountAge).Execute()

Update an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID.
	clockTime := float32(8.14) // float32 | Clock initial time in minutes
	clockIncrement := int32(56) // int32 | Clock increment in seconds
	minutes := int32(56) // int32 | How long the tournament lasts, in minutes
	name := "name_example" // string | The tournament name. Leave empty to get a random Grandmaster name (optional)
	waitMinutes := int32(56) // int32 | How long to wait before starting the tournament, from now, in minutes (optional) (default to 5)
	startDate := int64(789) // int64 | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the `waitMinutes` setting (optional)
	variant := "variant_example" // string |  (optional) (default to "standard")
	rated := true // bool | Games are rated and impact players ratings (optional) (default to true)
	position := "position_example" // string | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. (optional) (default to "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	berserkable := true // bool | Whether the players can use berserk. Only allowed if clockIncrement <= clockTime * 2 (optional) (default to true)
	streakable := true // bool | After 2 wins, consecutive wins grant 4 points instead of 2. (optional) (default to true)
	hasChat := true // bool | Whether the players can discuss in a chat (optional) (default to true)
	description := "description_example" // string | Anything you want to tell players about the tournament (optional)
	password := "password_example" // string | Make the tournament private, and restrict access with a password (optional)
	conditionsMinRatingRating := int32(56) // int32 | Minimum rating to join. Leave empty to let everyone join the tournament. (optional)
	conditionsMaxRatingRating := int32(56) // int32 | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. (optional)
	conditionsNbRatedGameNb := int32(56) // int32 | Minimum number of rated games required to join. (optional)
	conditionsAllowList := "conditionsAllowList_example" // string | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding `%titled` to the list additionally allows any titled player to join. Example: `thibault,german11,%titled`  (optional)
	conditionsBots := true // bool | Whether bots are allowed to join the tournament. (optional) (default to false)
	conditionsAccountAge := int32(56) // int32 | Minium account age in days required to join. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentUpdate(context.Background(), id).ClockTime(clockTime).ClockIncrement(clockIncrement).Minutes(minutes).Name(name).WaitMinutes(waitMinutes).StartDate(startDate).Variant(variant).Rated(rated).Position(position).Berserkable(berserkable).Streakable(streakable).HasChat(hasChat).Description(description).Password(password).ConditionsMinRatingRating(conditionsMinRatingRating).ConditionsMaxRatingRating(conditionsMaxRatingRating).ConditionsNbRatedGameNb(conditionsNbRatedGameNb).ConditionsAllowList(conditionsAllowList).ConditionsBots(conditionsBots).ConditionsAccountAge(conditionsAccountAge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentUpdate`: Tournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clockTime** | **float32** | Clock initial time in minutes | 
 **clockIncrement** | **int32** | Clock increment in seconds | 
 **minutes** | **int32** | How long the tournament lasts, in minutes | 
 **name** | **string** | The tournament name. Leave empty to get a random Grandmaster name | 
 **waitMinutes** | **int32** | How long to wait before starting the tournament, from now, in minutes | [default to 5]
 **startDate** | **int64** | Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting | 
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **rated** | **bool** | Games are rated and impact players ratings | [default to true]
 **position** | **string** | Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated. | [default to &quot;rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1&quot;]
 **berserkable** | **bool** | Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2 | [default to true]
 **streakable** | **bool** | After 2 wins, consecutive wins grant 4 points instead of 2. | [default to true]
 **hasChat** | **bool** | Whether the players can discuss in a chat | [default to true]
 **description** | **string** | Anything you want to tell players about the tournament | 
 **password** | **string** | Make the tournament private, and restrict access with a password | 
 **conditionsMinRatingRating** | **int32** | Minimum rating to join. Leave empty to let everyone join the tournament. | 
 **conditionsMaxRatingRating** | **int32** | Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament. | 
 **conditionsNbRatedGameNb** | **int32** | Minimum number of rated games required to join. | 
 **conditionsAllowList** | **string** | Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60;  | 
 **conditionsBots** | **bool** | Whether bots are allowed to join the tournament. | [default to false]
 **conditionsAccountAge** | **int32** | Minium account age in days required to join. | 

### Return type

[**Tournament200Response**](Tournament200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTournamentWithdraw

> AccountKidPost200Response ApiTournamentWithdraw(ctx, id).Execute()

Pause or leave an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "hL7vMrFQ" // string | The tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiTournamentWithdraw(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiTournamentWithdraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTournamentWithdraw`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiTournamentWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTournamentWithdrawRequest struct via the builder pattern


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


## ApiUserNameTournamentCreated

> ApiTournament200ResponseCreatedInner ApiUserNameTournamentCreated(ctx, username).Nb(nb).Status(status).Execute()

Get tournaments created by a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	username := "username_example" // string | The user whose created tournaments to fetch
	nb := int32(56) // int32 | Max number of tournaments to fetch (optional)
	status := int32(56) // int32 | Include tournaments in the given status: \"Created\" (10), \"Started\" (20), \"Finished\" (30) You can add this parameter more than once to include tournaments in different statuses. Example: `?status=10&status=20`  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiUserNameTournamentCreated(context.Background(), username).Nb(nb).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiUserNameTournamentCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserNameTournamentCreated`: ApiTournament200ResponseCreatedInner
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiUserNameTournamentCreated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user whose created tournaments to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserNameTournamentCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | **int32** | Max number of tournaments to fetch | 
 **status** | **int32** | Include tournaments in the given status: \&quot;Created\&quot; (10), \&quot;Started\&quot; (20), \&quot;Finished\&quot; (30) You can add this parameter more than once to include tournaments in different statuses. Example: &#x60;?status&#x3D;10&amp;status&#x3D;20&#x60;  | 

### Return type

[**ApiTournament200ResponseCreatedInner**](ApiTournament200ResponseCreatedInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserNameTournamentPlayed

> ApiUserNameTournamentPlayed200Response ApiUserNameTournamentPlayed(ctx, username).Nb(nb).Performance(performance).Execute()

Get tournaments played by a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	username := "username_example" // string | The user whose played tournaments to fetch
	nb := int32(56) // int32 | Max number of tournaments to fetch (optional)
	performance := true // bool | Include the player performance rating in the response, at some cost for the server.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ApiUserNameTournamentPlayed(context.Background(), username).Nb(nb).Performance(performance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ApiUserNameTournamentPlayed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserNameTournamentPlayed`: ApiUserNameTournamentPlayed200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ApiUserNameTournamentPlayed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user whose played tournaments to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserNameTournamentPlayedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | **int32** | Max number of tournaments to fetch | 
 **performance** | **bool** | Include the player performance rating in the response, at some cost for the server.  | 

### Return type

[**ApiUserNameTournamentPlayed200Response**](ApiUserNameTournamentPlayed200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GamesByTournament

> string GamesByTournament(ctx, id).Accept(accept).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Execute()

Export games of an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID.
	accept := "accept_example" // string | Specify the desired response format. Use `application/x-chess-pgn` to get the games in PGN format. Use `application/x-ndjson` to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  (optional) (default to "application/x-chess-pgn")
	player := "player_example" // string | Only games of a particular player. Leave empty to fetch games of all players. (optional)
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
	resp, r, err := apiClient.TournamentsArenaAPI.GamesByTournament(context.Background(), id).Accept(accept).Player(player).Moves(moves).PgnInJson(pgnInJson).Tags(tags).Clocks(clocks).Evals(evals).Accuracy(accuracy).Opening(opening).Division(division).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.GamesByTournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GamesByTournament`: string
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.GamesByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGamesByTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript.  | [default to &quot;application/x-chess-pgn&quot;]
 **player** | **string** | Only games of a particular player. Leave empty to fetch games of all players. | 
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


## ResultsByTournament

> ResultsByTournament200Response ResultsByTournament(ctx, id).Nb(nb).Sheet(sheet).Execute()

Get results of an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID.
	nb := int32(56) // int32 | Max number of players to fetch (optional)
	sheet := true // bool | Add a `sheet` field to the player document. It's an expensive server computation that slows down the stream.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.ResultsByTournament(context.Background(), id).Nb(nb).Sheet(sheet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.ResultsByTournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResultsByTournament`: ResultsByTournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.ResultsByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResultsByTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nb** | **int32** | Max number of players to fetch | 
 **sheet** | **bool** | Add a &#x60;sheet&#x60; field to the player document. It&#39;s an expensive server computation that slows down the stream.  | [default to false]

### Return type

[**ResultsByTournament200Response**](ResultsByTournament200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsByTournament

> TeamsByTournament200Response TeamsByTournament(ctx, id).Execute()

Get team standing of a team battle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.TeamsByTournament(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.TeamsByTournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamsByTournament`: TeamsByTournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.TeamsByTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsByTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamsByTournament200Response**](TeamsByTournament200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tournament

> Tournament200Response Tournament(ctx, id).Page(page).Execute()

Get info about an Arena tournament



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | The tournament ID.
	page := int32(1) // int32 | Specify which page of player standings to view. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsArenaAPI.Tournament(context.Background(), id).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsArenaAPI.Tournament``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Tournament`: Tournament200Response
	fmt.Fprintf(os.Stdout, "Response from `TournamentsArenaAPI.Tournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The tournament ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specify which page of player standings to view. | [default to 1]

### Return type

[**Tournament200Response**](Tournament200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

