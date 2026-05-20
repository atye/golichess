# \TeamsAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTeamArena**](TeamsAPI.md#ApiTeamArena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
[**ApiTeamSwiss**](TeamsAPI.md#ApiTeamSwiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
[**TeamAll**](TeamsAPI.md#TeamAll) | **Get** /api/team/all | Get popular teams
[**TeamIdJoin**](TeamsAPI.md#TeamIdJoin) | **Post** /team/{teamId}/join | Join a team
[**TeamIdKickUserId**](TeamsAPI.md#TeamIdKickUserId) | **Post** /api/team/{teamId}/kick/{userId} | Kick a user from your team
[**TeamIdPmAll**](TeamsAPI.md#TeamIdPmAll) | **Post** /team/{teamId}/pm-all | Message all members
[**TeamIdQuit**](TeamsAPI.md#TeamIdQuit) | **Post** /team/{teamId}/quit | Leave a team
[**TeamIdUsers**](TeamsAPI.md#TeamIdUsers) | **Get** /api/team/{teamId}/users | Get members of a team
[**TeamOfUsername**](TeamsAPI.md#TeamOfUsername) | **Get** /api/team/of/{username} | Teams of a player
[**TeamRequestAccept**](TeamsAPI.md#TeamRequestAccept) | **Post** /api/team/{teamId}/request/{userId}/accept | Accept join request
[**TeamRequestDecline**](TeamsAPI.md#TeamRequestDecline) | **Post** /api/team/{teamId}/request/{userId}/decline | Decline join request
[**TeamRequests**](TeamsAPI.md#TeamRequests) | **Get** /api/team/{teamId}/requests | Get join requests
[**TeamSearch**](TeamsAPI.md#TeamSearch) | **Get** /api/team/search | Search teams
[**TeamShow**](TeamsAPI.md#TeamShow) | **Get** /api/team/{teamId} | Get a single team



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
	resp, r, err := apiClient.TeamsAPI.ApiTeamArena(context.Background(), teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.ApiTeamArena``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTeamArena`: ApiTournament200ResponseCreatedInner
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.ApiTeamArena`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	teamId := "coders" // string | 
	max := int32(56) // int32 | How many tournaments to download. (optional) (default to 100)
	status := "status_example" // string | [Filter] Only swiss tournaments in this current state.  (optional)
	createdBy := "createdBy_example" // string | [Filter] Only swiss tournaments created by a given user.  (optional)
	name := "name_example" // string | [Filter] Only swiss tournaments with a given name.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.ApiTeamSwiss(context.Background(), teamId).Max(max).Status(status).CreatedBy(createdBy).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.ApiTeamSwiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTeamSwiss`: ApiSwissNew200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.ApiTeamSwiss`: %v\n", resp)
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


## TeamAll

> TeamAll200Response TeamAll(ctx).Page(page).Execute()

Get popular teams



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
	page := int32(1) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamAll(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamAll`: TeamAll200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]

### Return type

[**TeamAll200Response**](TeamAll200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamIdJoin

> AccountKidPost200Response TeamIdJoin(ctx, teamId).Message(message).Password(password).Execute()

Join a team



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
	teamId := "coders" // string | 
	message := "message_example" // string | Required if team manually reviews admission requests. (optional)
	password := "password_example" // string | Optional password, if the team requires one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamIdJoin(context.Background(), teamId).Message(message).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamIdJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamIdJoin`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamIdJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **string** | Required if team manually reviews admission requests. | 
 **password** | **string** | Optional password, if the team requires one. | 

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


## TeamIdKickUserId

> AccountKidPost200Response TeamIdKickUserId(ctx, teamId, userId).Execute()

Kick a user from your team



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
	teamId := "coders" // string | 
	userId := "neio" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamIdKickUserId(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamIdKickUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamIdKickUserId`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamIdKickUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdKickUserIdRequest struct via the builder pattern


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


## TeamIdPmAll

> AccountKidPost200Response TeamIdPmAll(ctx, teamId).Message(message).Execute()

Message all members



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
	teamId := "coders" // string | 
	message := "message_example" // string | The message to send to all your team members. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamIdPmAll(context.Background(), teamId).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamIdPmAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamIdPmAll`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamIdPmAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdPmAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **message** | **string** | The message to send to all your team members. | 

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


## TeamIdQuit

> AccountKidPost200Response TeamIdQuit(ctx, teamId).Execute()

Leave a team



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
	teamId := "coders" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamIdQuit(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamIdQuit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamIdQuit`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamIdQuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdQuitRequest struct via the builder pattern


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


## TeamIdUsers

> TeamIdUsers200Response TeamIdUsers(ctx, teamId).Full(full).Execute()

Get members of a team



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
	teamId := "coders" // string | 
	full := true // bool | Full user documents with performance ratings. This limits the response to 1,000 users.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamIdUsers(context.Background(), teamId).Full(full).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamIdUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamIdUsers`: TeamIdUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **full** | **bool** | Full user documents with performance ratings. This limits the response to 1,000 users.  | [default to false]

### Return type

[**TeamIdUsers200Response**](TeamIdUsers200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamOfUsername

> []TeamShow200Response TeamOfUsername(ctx, username).Execute()

Teams of a player



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
	resp, r, err := apiClient.TeamsAPI.TeamOfUsername(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamOfUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamOfUsername`: []TeamShow200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamOfUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamOfUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TeamShow200Response**](TeamShow200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamRequestAccept

> AccountKidPost200Response TeamRequestAccept(ctx, teamId, userId).Execute()

Accept join request



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
	teamId := "coders" // string | 
	userId := "neio" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamRequestAccept(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamRequestAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamRequestAccept`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamRequestAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestAcceptRequest struct via the builder pattern


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


## TeamRequestDecline

> AccountKidPost200Response TeamRequestDecline(ctx, teamId, userId).Execute()

Decline join request



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
	teamId := "coders" // string | 
	userId := "neio" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamRequestDecline(context.Background(), teamId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamRequestDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamRequestDecline`: AccountKidPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamRequestDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestDeclineRequest struct via the builder pattern


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


## TeamRequests

> []TeamRequests200ResponseInner TeamRequests(ctx, teamId).Declined(declined).Execute()

Get join requests



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
	teamId := "teamId_example" // string | 
	declined := true // bool | Get the declined join requests (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamRequests(context.Background(), teamId).Declined(declined).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamRequests`: []TeamRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **declined** | **bool** | Get the declined join requests | [default to false]

### Return type

[**[]TeamRequests200ResponseInner**](TeamRequests200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamSearch

> TeamAll200Response TeamSearch(ctx).Text(text).Page(page).Execute()

Search teams



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
	text := "coders" // string |  (optional)
	page := int32(1) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamSearch(context.Background()).Text(text).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamSearch`: TeamAll200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **page** | **int32** |  | [default to 1]

### Return type

[**TeamAll200Response**](TeamAll200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamShow

> TeamShow200Response TeamShow(ctx, teamId).Execute()

Get a single team



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
	teamId := "teamId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.TeamShow(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.TeamShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TeamShow`: TeamShow200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.TeamShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamShow200Response**](TeamShow200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

