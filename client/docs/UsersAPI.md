# \UsersAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCrosstable**](UsersAPI.md#ApiCrosstable) | **Get** /api/crosstable/{user1}/{user2} | Get crosstable
[**ApiPlayerAutocomplete**](UsersAPI.md#ApiPlayerAutocomplete) | **Get** /api/player/autocomplete | Autocomplete usernames
[**ApiUser**](UsersAPI.md#ApiUser) | **Get** /api/user/{username} | Get user public data
[**ApiUserActivity**](UsersAPI.md#ApiUserActivity) | **Get** /api/user/{username}/activity | Get user activity
[**ApiUserPerf**](UsersAPI.md#ApiUserPerf) | **Get** /api/user/{username}/perf/{perf} | Get performance statistics of a user
[**ApiUserRatingHistory**](UsersAPI.md#ApiUserRatingHistory) | **Get** /api/user/{username}/rating-history | Get rating history of a user
[**ApiUsers**](UsersAPI.md#ApiUsers) | **Post** /api/users | Get users by ID
[**ApiUsersStatus**](UsersAPI.md#ApiUsersStatus) | **Get** /api/users/status | Get real-time users status
[**Player**](UsersAPI.md#Player) | **Get** /api/player | Get all top 10
[**PlayerTopNbPerfType**](UsersAPI.md#PlayerTopNbPerfType) | **Get** /api/player/top/{nb}/{perfType} | Get one leaderboard
[**ReadNote**](UsersAPI.md#ReadNote) | **Get** /api/user/{username}/note | Get notes for a user
[**StreamerLive**](UsersAPI.md#StreamerLive) | **Get** /api/streamer/live | Get live streamers
[**WriteNote**](UsersAPI.md#WriteNote) | **Post** /api/user/{username}/note | Add a note for a user



## ApiCrosstable

> ApiCrosstable200Response ApiCrosstable(ctx, user1, user2).Matchup(matchup).Execute()

Get crosstable



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
	user1 := "user1_example" // string | 
	user2 := "user2_example" // string | 
	matchup := true // bool | Whether to get the current match data, if any (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiCrosstable(context.Background(), user1, user2).Matchup(matchup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiCrosstable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCrosstable`: ApiCrosstable200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiCrosstable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user1** | **string** |  | 
**user2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCrosstableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **matchup** | **bool** | Whether to get the current match data, if any | 

### Return type

[**ApiCrosstable200Response**](ApiCrosstable200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPlayerAutocomplete

> ApiPlayerAutocomplete200Response ApiPlayerAutocomplete(ctx).Term(term).Exists(exists).Object(object).Names(names).Friend(friend).Team(team).Tour(tour).Swiss(swiss).Teacher(teacher).Execute()

Autocomplete usernames



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
	term := "term_example" // string | The beginning of a username
	exists := true // bool | If `true`, only checks if the user exists.  (optional) (default to false)
	object := true // bool | - `false` returns an array of usernames - `true` returns an object with matching users  (optional) (default to false)
	names := true // bool | - `false` returns an array of usernames - `true` returns an array of usernames with preferred casing  (optional) (default to false)
	friend := true // bool | Returns followed players matching `term` if any, else returns other players. Requires [OAuth](#tag/OAuth).  (optional)
	team := "team_example" // string | Search within a team. Use team ID/slug.  (optional)
	tour := "tour_example" // string | Search within a arena tournament. Use tournament ID.  (optional)
	swiss := "swiss_example" // string | Search within a Swiss tournament.  (optional)
	teacher := true // bool | Only search for players who also have a teacher role.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiPlayerAutocomplete(context.Background()).Term(term).Exists(exists).Object(object).Names(names).Friend(friend).Team(team).Tour(tour).Swiss(swiss).Teacher(teacher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiPlayerAutocomplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPlayerAutocomplete`: ApiPlayerAutocomplete200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiPlayerAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPlayerAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** | The beginning of a username | 
 **exists** | **bool** | If &#x60;true&#x60;, only checks if the user exists.  | [default to false]
 **object** | **bool** | - &#x60;false&#x60; returns an array of usernames - &#x60;true&#x60; returns an object with matching users  | [default to false]
 **names** | **bool** | - &#x60;false&#x60; returns an array of usernames - &#x60;true&#x60; returns an array of usernames with preferred casing  | [default to false]
 **friend** | **bool** | Returns followed players matching &#x60;term&#x60; if any, else returns other players. Requires [OAuth](#tag/OAuth).  | 
 **team** | **string** | Search within a team. Use team ID/slug.  | 
 **tour** | **string** | Search within a arena tournament. Use tournament ID.  | 
 **swiss** | **string** | Search within a Swiss tournament.  | 
 **teacher** | **bool** | Only search for players who also have a teacher role.  | [default to false]

### Return type

[**ApiPlayerAutocomplete200Response**](ApiPlayerAutocomplete200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUser

> ApiUser200Response ApiUser(ctx, username).Trophies(trophies).Profile(profile).Rank(rank).FideId(fideId).Execute()

Get user public data



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
	username := "username_example" // string | 
	trophies := true // bool | Include user trophies (optional) (default to false)
	profile := true // bool | Include user profile data (optional) (default to true)
	rank := true // bool | Include global lichess ranking for each perf (optional) (default to false)
	fideId := true // bool | Include public FIDE ID if any (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUser(context.Background(), username).Trophies(trophies).Profile(profile).Rank(rank).FideId(fideId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUser`: ApiUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trophies** | **bool** | Include user trophies | [default to false]
 **profile** | **bool** | Include user profile data | [default to true]
 **rank** | **bool** | Include global lichess ranking for each perf | [default to false]
 **fideId** | **bool** | Include public FIDE ID if any | [default to false]

### Return type

[**ApiUser200Response**](ApiUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserActivity

> []ApiUserActivity200ResponseInner ApiUserActivity(ctx, username).Execute()

Get user activity



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
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUserActivity(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUserActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserActivity`: []ApiUserActivity200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUserActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiUserActivity200ResponseInner**](ApiUserActivity200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserPerf

> ApiUserPerf200Response ApiUserPerf(ctx, username, perf).Execute()

Get performance statistics of a user



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
	username := "username_example" // string | 
	perf := "perf_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUserPerf(context.Background(), username, perf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUserPerf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserPerf`: ApiUserPerf200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUserPerf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**perf** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserPerfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiUserPerf200Response**](ApiUserPerf200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserRatingHistory

> []ApiUserRatingHistory200ResponseInner ApiUserRatingHistory(ctx, username).Execute()

Get rating history of a user



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
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUserRatingHistory(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUserRatingHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserRatingHistory`: []ApiUserRatingHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUserRatingHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserRatingHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiUserRatingHistory200ResponseInner**](ApiUserRatingHistory200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsers

> []ApiUsers200ResponseInner ApiUsers(ctx).Body(body).Profile(profile).Rank(rank).Execute()

Get users by ID



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
	body := "thibault,maia1,maia5" // string | User IDs separated by commas.
	profile := true // bool | Include user profile data (optional) (default to true)
	rank := true // bool | Include global lichess ranking for each perf (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUsers(context.Background()).Body(body).Profile(profile).Rank(rank).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsers`: []ApiUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | User IDs separated by commas. | 
 **profile** | **bool** | Include user profile data | [default to true]
 **rank** | **bool** | Include global lichess ranking for each perf | [default to false]

### Return type

[**[]ApiUsers200ResponseInner**](ApiUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersStatus

> []ApiUsersStatus200ResponseInner ApiUsersStatus(ctx).Ids(ids).WithSignal(withSignal).WithGameIds(withGameIds).WithGameMetas(withGameMetas).Execute()

Get real-time users status



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
	ids := "thibault,maia1,maia5" // string | User IDs separated by commas. Up to 100 IDs.
	withSignal := true // bool | Also return the network signal of the player, when available. It ranges from 1 (poor connection, lag > 500ms) to 4 (great connection, lag < 150ms) Defaults to `false` to preserve server resources.  (optional)
	withGameIds := true // bool | Also return the ID of the game being played, if any, for each player, in a `playingId` field. Defaults to `false` to preserve server resources.  (optional)
	withGameMetas := true // bool | Also return the id, time control and variant of the game being played, if any, for each player, in a `playing` field. Defaults to `false` to preserve server resources. Disables `withGameIds`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ApiUsersStatus(context.Background()).Ids(ids).WithSignal(withSignal).WithGameIds(withGameIds).WithGameMetas(withGameMetas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ApiUsersStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUsersStatus`: []ApiUsersStatus200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ApiUsersStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | User IDs separated by commas. Up to 100 IDs. | 
 **withSignal** | **bool** | Also return the network signal of the player, when available. It ranges from 1 (poor connection, lag &gt; 500ms) to 4 (great connection, lag &lt; 150ms) Defaults to &#x60;false&#x60; to preserve server resources.  | 
 **withGameIds** | **bool** | Also return the ID of the game being played, if any, for each player, in a &#x60;playingId&#x60; field. Defaults to &#x60;false&#x60; to preserve server resources.  | 
 **withGameMetas** | **bool** | Also return the id, time control and variant of the game being played, if any, for each player, in a &#x60;playing&#x60; field. Defaults to &#x60;false&#x60; to preserve server resources. Disables &#x60;withGameIds&#x60;.  | 

### Return type

[**[]ApiUsersStatus200ResponseInner**](ApiUsersStatus200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Player

> Player200Response Player(ctx).Execute()

Get all top 10



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
	resp, r, err := apiClient.UsersAPI.Player(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.Player``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Player`: Player200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.Player`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerRequest struct via the builder pattern


### Return type

[**Player200Response**](Player200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerTopNbPerfType

> PlayerTopNbPerfType200Response PlayerTopNbPerfType(ctx, nb, perfType).Execute()

Get one leaderboard



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
	nb := int32(100) // int32 | How many users to fetch
	perfType := "bullet" // string | The speed or variant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PlayerTopNbPerfType(context.Background(), nb, perfType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PlayerTopNbPerfType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlayerTopNbPerfType`: PlayerTopNbPerfType200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PlayerTopNbPerfType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nb** | **int32** | How many users to fetch | 
**perfType** | **string** | The speed or variant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerTopNbPerfTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlayerTopNbPerfType200Response**](PlayerTopNbPerfType200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.lichess.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNote

> []ReadNote200ResponseInner ReadNote(ctx, username).Execute()

Get notes for a user



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
	username := "thibault" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ReadNote(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ReadNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNote`: []ReadNote200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ReadNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ReadNote200ResponseInner**](ReadNote200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamerLive

> []StreamerLive200ResponseInner StreamerLive(ctx).Execute()

Get live streamers



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
	resp, r, err := apiClient.UsersAPI.StreamerLive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.StreamerLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamerLive`: []StreamerLive200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.StreamerLive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamerLiveRequest struct via the builder pattern


### Return type

[**[]StreamerLive200ResponseInner**](StreamerLive200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteNote

> AccountKidPost200Response WriteNote(ctx, username).Text(text).Execute()

Add a note for a user



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
	username := "thibault" // string | 
	text := "text_example" // string | The contents of the note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.WriteNote(context.Background(), username).Text(text).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.WriteNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WriteNote`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.WriteNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **text** | **string** | The contents of the note | 

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

