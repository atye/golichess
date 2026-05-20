# \PuzzlesAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPuzzleActivity**](PuzzlesAPI.md#ApiPuzzleActivity) | **Get** /api/puzzle/activity | Get your puzzle activity
[**ApiPuzzleBatchSelect**](PuzzlesAPI.md#ApiPuzzleBatchSelect) | **Get** /api/puzzle/batch/{angle} | Get multiple puzzles at once
[**ApiPuzzleBatchSolve**](PuzzlesAPI.md#ApiPuzzleBatchSolve) | **Post** /api/puzzle/batch/{angle} | Solve multiple puzzles at once
[**ApiPuzzleDaily**](PuzzlesAPI.md#ApiPuzzleDaily) | **Get** /api/puzzle/daily | Get the daily puzzle
[**ApiPuzzleDashboard**](PuzzlesAPI.md#ApiPuzzleDashboard) | **Get** /api/puzzle/dashboard/{days} | Get your puzzle dashboard
[**ApiPuzzleId**](PuzzlesAPI.md#ApiPuzzleId) | **Get** /api/puzzle/{id} | Get a puzzle by its ID
[**ApiPuzzleNext**](PuzzlesAPI.md#ApiPuzzleNext) | **Get** /api/puzzle/next | Get a new puzzle
[**ApiPuzzleReplay**](PuzzlesAPI.md#ApiPuzzleReplay) | **Get** /api/puzzle/replay/{days}/{theme} | Get puzzles to replay
[**ApiStormDashboard**](PuzzlesAPI.md#ApiStormDashboard) | **Get** /api/storm/dashboard/{username} | Get the storm dashboard of a player
[**RacerGet**](PuzzlesAPI.md#RacerGet) | **Get** /api/racer/{id} | Get puzzle race results
[**RacerPost**](PuzzlesAPI.md#RacerPost) | **Post** /api/racer | Create and join a puzzle race



## ApiPuzzleActivity

> ApiPuzzleActivity200Response ApiPuzzleActivity(ctx).Max(max).Before(before).Since(since).Execute()

Get your puzzle activity



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
	max := int32(56) // int32 | How many entries to download. Leave empty to download all activity. (optional)
	before := int32(56) // int32 | Download entries before this timestamp. Defaults to now. Use `before` and `max` for pagination. (optional)
	since := int32(56) // int32 | Download entries since this timestamp. Defaults to account creation date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleActivity(context.Background()).Max(max).Before(before).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleActivity`: ApiPuzzleActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int32** | How many entries to download. Leave empty to download all activity. | 
 **before** | **int32** | Download entries before this timestamp. Defaults to now. Use &#x60;before&#x60; and &#x60;max&#x60; for pagination. | 
 **since** | **int32** | Download entries since this timestamp. Defaults to account creation date. | 

### Return type

[**ApiPuzzleActivity200Response**](ApiPuzzleActivity200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleBatchSelect

> ApiPuzzleBatchSelect200Response ApiPuzzleBatchSelect(ctx, angle).Difficulty(difficulty).Nb(nb).Color(color).Execute()

Get multiple puzzles at once



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
	angle := "mix" // string | The theme or opening to filter puzzles with. Recommended: `mix`.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
	difficulty := "difficulty_example" // string | The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous. (optional)
	nb := int32(10) // int32 | How many puzzles to fetch. Just set it to `1` if you only need one puzzle.  (optional) (default to 15)
	color := "color_example" // string | The color to play. Better left empty to automatically get 50% white. Currently only works when `nb=1`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleBatchSelect(context.Background(), angle).Difficulty(difficulty).Nb(nb).Color(color).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleBatchSelect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleBatchSelect`: ApiPuzzleBatchSelect200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleBatchSelect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**angle** | **string** | The theme or opening to filter puzzles with. Recommended: &#x60;mix&#x60;.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleBatchSelectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **difficulty** | **string** | The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous. | 
 **nb** | **int32** | How many puzzles to fetch. Just set it to &#x60;1&#x60; if you only need one puzzle.  | [default to 15]
 **color** | **string** | The color to play. Better left empty to automatically get 50% white. Currently only works when &#x60;nb&#x3D;1&#x60;.  | 

### Return type

[**ApiPuzzleBatchSelect200Response**](ApiPuzzleBatchSelect200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleBatchSolve

> ApiPuzzleBatchSolve200Response ApiPuzzleBatchSolve(ctx, angle).ApiPuzzleBatchSolveRequest(apiPuzzleBatchSolveRequest).Nb(nb).Execute()

Solve multiple puzzles at once



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
	angle := "mix" // string | The theme or opening of the solved puzzles.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
	apiPuzzleBatchSolveRequest := *openapiclient.NewApiPuzzleBatchSolveRequest() // ApiPuzzleBatchSolveRequest | List of solved puzzles
	nb := int32(1) // int32 | When > 0, the response includes a new puzzle batch with that many puzzles.  This is equivalent to calling [/api/puzzle/batch/{angle}](#tag/puzzles/GET/api/puzzle/batch/{angle}), and can sometimes save a request.  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleBatchSolve(context.Background(), angle).ApiPuzzleBatchSolveRequest(apiPuzzleBatchSolveRequest).Nb(nb).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleBatchSolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleBatchSolve`: ApiPuzzleBatchSolve200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleBatchSolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**angle** | **string** | The theme or opening of the solved puzzles.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleBatchSolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiPuzzleBatchSolveRequest** | [**ApiPuzzleBatchSolveRequest**](ApiPuzzleBatchSolveRequest.md) | List of solved puzzles | 
 **nb** | **int32** | When &gt; 0, the response includes a new puzzle batch with that many puzzles.  This is equivalent to calling [/api/puzzle/batch/{angle}](#tag/puzzles/GET/api/puzzle/batch/{angle}), and can sometimes save a request.  | [default to 0]

### Return type

[**ApiPuzzleBatchSolve200Response**](ApiPuzzleBatchSolve200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleDaily

> ApiPuzzleDaily200Response ApiPuzzleDaily(ctx).Execute()

Get the daily puzzle



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
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleDaily(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleDaily``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleDaily`: ApiPuzzleDaily200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleDaily`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleDailyRequest struct via the builder pattern


### Return type

[**ApiPuzzleDaily200Response**](ApiPuzzleDaily200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleDashboard

> ApiPuzzleDashboard200Response ApiPuzzleDashboard(ctx, days).Execute()

Get your puzzle dashboard



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
	days := int32(56) // int32 | How many days to look back when aggregating puzzle results. 30 is sensible.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleDashboard(context.Background(), days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleDashboard`: ApiPuzzleDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**days** | **int32** | How many days to look back when aggregating puzzle results. 30 is sensible. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPuzzleDashboard200Response**](ApiPuzzleDashboard200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleId

> ApiPuzzleId200Response ApiPuzzleId(ctx, id).Execute()

Get a puzzle by its ID



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
	id := "id_example" // string | The puzzle ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleId`: ApiPuzzleId200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The puzzle ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPuzzleId200Response**](ApiPuzzleId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleNext

> ApiPuzzleId200Response ApiPuzzleNext(ctx).Angle(angle).Difficulty(difficulty).Color(color).Execute()

Get a new puzzle



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
	angle := "angle_example" // string | The theme or opening to filter puzzles with.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes).  (optional)
	difficulty := "difficulty_example" // string | The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous. (optional)
	color := "color_example" // string | The color to play. Better left empty to automatically get 50% white. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleNext(context.Background()).Angle(angle).Difficulty(difficulty).Color(color).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleNext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleNext`: ApiPuzzleId200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleNext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **angle** | **string** | The theme or opening to filter puzzles with.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes).  | 
 **difficulty** | **string** | The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous. | 
 **color** | **string** | The color to play. Better left empty to automatically get 50% white. | 

### Return type

[**ApiPuzzleId200Response**](ApiPuzzleId200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPuzzleReplay

> ApiPuzzleReplay200Response ApiPuzzleReplay(ctx, days, theme).Execute()

Get puzzles to replay



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
	days := int32(56) // int32 | How many days to look back when aggregating puzzle results. 30 is sensible.
	theme := "theme_example" // string | The theme or opening to filter puzzles with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiPuzzleReplay(context.Background(), days, theme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiPuzzleReplay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPuzzleReplay`: ApiPuzzleReplay200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiPuzzleReplay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**days** | **int32** | How many days to look back when aggregating puzzle results. 30 is sensible. | 
**theme** | **string** | The theme or opening to filter puzzles with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPuzzleReplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiPuzzleReplay200Response**](ApiPuzzleReplay200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStormDashboard

> ApiStormDashboard200Response ApiStormDashboard(ctx, username).Days(days).Execute()

Get the storm dashboard of a player



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
	username := "username_example" // string | Username of the player
	days := int32(56) // int32 | How many days of history to return (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.ApiStormDashboard(context.Background(), username).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.ApiStormDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStormDashboard`: ApiStormDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.ApiStormDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the player | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStormDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | How many days of history to return | [default to 30]

### Return type

[**ApiStormDashboard200Response**](ApiStormDashboard200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacerGet

> RacerGet200Response RacerGet(ctx, id).Execute()

Get puzzle race results



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
	id := "id_example" // string | The puzzle race ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.RacerGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.RacerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacerGet`: RacerGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.RacerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The puzzle race ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RacerGet200Response**](RacerGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacerPost

> RacerPost200Response RacerPost(ctx).Execute()

Create and join a puzzle race



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
	resp, r, err := apiClient.PuzzlesAPI.RacerPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.RacerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacerPost`: RacerPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.RacerPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRacerPostRequest struct via the builder pattern


### Return type

[**RacerPost200Response**](RacerPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

