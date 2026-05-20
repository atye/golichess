# \StudiesAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStudyChapterMoves**](StudiesAPI.md#ApiStudyChapterMoves) | **Post** /api/study/{studyId}/{chapterId}/moves | Update the moves of a study chapter
[**ApiStudyChapterTags**](StudiesAPI.md#ApiStudyChapterTags) | **Post** /api/study/{studyId}/{chapterId}/tags | Update PGN tags of a study chapter
[**ApiStudyImportPGN**](StudiesAPI.md#ApiStudyImportPGN) | **Post** /api/study/{studyId}/import-pgn | Import PGN into a study
[**ApiStudyPost**](StudiesAPI.md#ApiStudyPost) | **Post** /api/study | Create a new Study
[**ApiStudyStudyIdChapterIdDelete**](StudiesAPI.md#ApiStudyStudyIdChapterIdDelete) | **Delete** /api/study/{studyId}/{chapterId} | Delete a study chapter
[**StudyAllChaptersHead**](StudiesAPI.md#StudyAllChaptersHead) | **Head** /api/study/{studyId}.pgn | Study metadata
[**StudyAllChaptersPgn**](StudiesAPI.md#StudyAllChaptersPgn) | **Get** /api/study/{studyId}.pgn | Export all chapters
[**StudyChapterPgn**](StudiesAPI.md#StudyChapterPgn) | **Get** /api/study/{studyId}/{chapterId}.pgn | Export one study chapter
[**StudyExportAllPgn**](StudiesAPI.md#StudyExportAllPgn) | **Get** /api/study/by/{username}/export.pgn | Export all studies of a user
[**StudyListMetadata**](StudiesAPI.md#StudyListMetadata) | **Get** /api/study/by/{username} | List studies of a user



## ApiStudyChapterMoves

> ApiStudyChapterMoves(ctx, studyId, chapterId).Pgn(pgn).Execute()

Update the moves of a study chapter



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
	studyId := "studyId_example" // string | The study ID
	chapterId := "chapterId_example" // string | The chapter ID
	pgn := "pgn_example" // string | PGN text containing the moves that will replace the chapter's existing moves. Any provided tags are ignored. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.ApiStudyChapterMoves(context.Background(), studyId, chapterId).Pgn(pgn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ApiStudyChapterMoves``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 
**chapterId** | **string** | The chapter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStudyChapterMovesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pgn** | **string** | PGN text containing the moves that will replace the chapter&#39;s existing moves. Any provided tags are ignored.  | 

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


## ApiStudyChapterTags

> ApiStudyChapterTags(ctx, studyId, chapterId).Pgn(pgn).Execute()

Update PGN tags of a study chapter



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
	studyId := "studyId_example" // string | The study ID
	chapterId := "chapterId_example" // string | The chapter ID
	pgn := "pgn_example" // string | PGN text containing the tags. Only the tags are used. Moves are just ignored. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.ApiStudyChapterTags(context.Background(), studyId, chapterId).Pgn(pgn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ApiStudyChapterTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 
**chapterId** | **string** | The chapter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStudyChapterTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pgn** | **string** | PGN text containing the tags. Only the tags are used. Moves are just ignored.  | 

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


## ApiStudyImportPGN

> ApiStudyImportPGN200Response ApiStudyImportPGN(ctx, studyId).Pgn(pgn).Name(name).Orientation(orientation).Variant(variant).Mode(mode).Execute()

Import PGN into a study



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
	studyId := "studyId_example" // string | ID of the study
	pgn := "pgn_example" // string | PGN to import. Can contain multiple games separated by 2 or more newlines. 
	name := "name_example" // string | Name of the new chapter. If not specified, or if multiple chapters are created, the names will be inferred from the PGN tags.  (optional)
	orientation := "orientation_example" // string | Default board orientation. (optional) (default to "white")
	variant := "variant_example" // string |  (optional) (default to "standard")
	mode := "mode_example" // string | Analysis mode. If not specified, Normal analysis. * practice - Practise with Computer * conceal - Hide next moves * gamebook - Interactive lesson  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.ApiStudyImportPGN(context.Background(), studyId).Pgn(pgn).Name(name).Orientation(orientation).Variant(variant).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ApiStudyImportPGN``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStudyImportPGN`: ApiStudyImportPGN200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.ApiStudyImportPGN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | ID of the study | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStudyImportPGNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pgn** | **string** | PGN to import. Can contain multiple games separated by 2 or more newlines.  | 
 **name** | **string** | Name of the new chapter. If not specified, or if multiple chapters are created, the names will be inferred from the PGN tags.  | 
 **orientation** | **string** | Default board orientation. | [default to &quot;white&quot;]
 **variant** | **string** |  | [default to &quot;standard&quot;]
 **mode** | **string** | Analysis mode. If not specified, Normal analysis. * practice - Practise with Computer * conceal - Hide next moves * gamebook - Interactive lesson  | 

### Return type

[**ApiStudyImportPGN200Response**](ApiStudyImportPGN200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStudyPost

> ApiStudyPost200Response ApiStudyPost(ctx).Name(name).Visibility(visibility).Computer(computer).Explorer(explorer).Cloneable(cloneable).Shareable(shareable).Chat(chat).Sticky(sticky).Execute()

Create a new Study



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
	name := "name_example" // string | The study name.
	visibility := "visibility_example" // string | Who can view the study. * `public`: Default. Anyone can view the study, it appears on public listings * `unlisted`: Only people with the link can view the study, it doesn't appear on public listings * `private`: Only the study members can view the study  (default to "unlisted")
	computer := "computer_example" // string | 
	explorer := "explorer_example" // string | 
	cloneable := "cloneable_example" // string | 
	shareable := "shareable_example" // string | 
	chat := "chat_example" // string | 
	sticky := "sticky_example" // string | Keep everyone on the same chapter and position (optional) (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.ApiStudyPost(context.Background()).Name(name).Visibility(visibility).Computer(computer).Explorer(explorer).Cloneable(cloneable).Shareable(shareable).Chat(chat).Sticky(sticky).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ApiStudyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiStudyPost`: ApiStudyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.ApiStudyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStudyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The study name. | 
 **visibility** | **string** | Who can view the study. * &#x60;public&#x60;: Default. Anyone can view the study, it appears on public listings * &#x60;unlisted&#x60;: Only people with the link can view the study, it doesn&#39;t appear on public listings * &#x60;private&#x60;: Only the study members can view the study  | [default to &quot;unlisted&quot;]
 **computer** | **string** |  | 
 **explorer** | **string** |  | 
 **cloneable** | **string** |  | 
 **shareable** | **string** |  | 
 **chat** | **string** |  | 
 **sticky** | **string** | Keep everyone on the same chapter and position | [default to &quot;true&quot;]

### Return type

[**ApiStudyPost200Response**](ApiStudyPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStudyStudyIdChapterIdDelete

> ApiStudyStudyIdChapterIdDelete(ctx, studyId, chapterId).Execute()

Delete a study chapter



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
	studyId := "studyId_example" // string | The study ID
	chapterId := "chapterId_example" // string | The chapter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.ApiStudyStudyIdChapterIdDelete(context.Background(), studyId, chapterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.ApiStudyStudyIdChapterIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 
**chapterId** | **string** | The chapter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStudyStudyIdChapterIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StudyAllChaptersHead

> StudyAllChaptersHead(ctx, studyId).Execute()

Study metadata



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
	studyId := "studyId_example" // string | The study ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudiesAPI.StudyAllChaptersHead(context.Background(), studyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.StudyAllChaptersHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyAllChaptersHeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StudyAllChaptersPgn

> string StudyAllChaptersPgn(ctx, studyId).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()

Export all chapters



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
	studyId := "studyId_example" // string | The study ID
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis and annotator comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
	variations := true // bool | Include non-mainline moves, when available. Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)
	orientation := true // bool | Add a `Orientation` PGN tag with the chapter predefined orientation. Example: `[Orientation \"white\"]`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.StudyAllChaptersPgn(context.Background(), studyId).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.StudyAllChaptersPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StudyAllChaptersPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.StudyAllChaptersPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyAllChaptersPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | **bool** | Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]
 **orientation** | **bool** | Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60;  | [default to false]

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


## StudyChapterPgn

> string StudyChapterPgn(ctx, studyId, chapterId).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()

Export one study chapter



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
	studyId := "studyId_example" // string | The study ID
	chapterId := "chapterId_example" // string | The chapter ID
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis and annotator comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
	variations := true // bool | Include non-mainline moves, when available. Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)
	orientation := true // bool | Add a `Orientation` PGN tag with the chapter predefined orientation. Example: `[Orientation \"white\"]`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.StudyChapterPgn(context.Background(), studyId, chapterId).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.StudyChapterPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StudyChapterPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.StudyChapterPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studyId** | **string** | The study ID | 
**chapterId** | **string** | The chapter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyChapterPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | **bool** | Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]
 **orientation** | **bool** | Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60;  | [default to false]

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


## StudyExportAllPgn

> string StudyExportAllPgn(ctx, username).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()

Export all studies of a user



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
	username := "username_example" // string | The user whose studies we export
	clocks := true // bool | Include clock comments in the PGN moves, when available. Example: `2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }`  (optional) (default to true)
	comments := true // bool | Include analysis and annotator comments in the PGN moves, when available. Example: `12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }`  (optional) (default to true)
	variations := true // bool | Include non-mainline moves, when available. Example: `4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2`  (optional) (default to true)
	orientation := true // bool | Add a `Orientation` PGN tag with the chapter predefined orientation. Example: `[Orientation \"white\"]`  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.StudyExportAllPgn(context.Background(), username).Clocks(clocks).Comments(comments).Variations(variations).Orientation(orientation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.StudyExportAllPgn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StudyExportAllPgn`: string
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.StudyExportAllPgn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user whose studies we export | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyExportAllPgnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clocks** | **bool** | Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60;  | [default to true]
 **comments** | **bool** | Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60;  | [default to true]
 **variations** | **bool** | Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60;  | [default to true]
 **orientation** | **bool** | Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60;  | [default to false]

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


## StudyListMetadata

> StudyListMetadata200Response StudyListMetadata(ctx, username).Execute()

List studies of a user



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
	username := "username_example" // string | The user whose studies we list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StudiesAPI.StudyListMetadata(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudiesAPI.StudyListMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StudyListMetadata`: StudyListMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `StudiesAPI.StudyListMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user whose studies we list | 

### Other Parameters

Other parameters are passed through a pointer to a apiStudyListMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StudyListMetadata200Response**](StudyListMetadata200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

