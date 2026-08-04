/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.161
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type StudiesAPI interface {

	/*
	ApiStudyChapterMoves Update the moves of a study chapter

	Replaces the moves tree of a study chapter.
No tags will be modified.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@param chapterId The chapter ID
	@return StudiesAPIApiStudyChapterMovesRequest
	*/
	ApiStudyChapterMoves(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyChapterMovesRequest

	// ApiStudyChapterMovesExecute executes the request
	ApiStudyChapterMovesExecute(r StudiesAPIApiStudyChapterMovesRequest) (*http.Response, error)

	/*
	ApiStudyChapterTags Update PGN tags of a study chapter

	Add, update and delete the PGN tags of a study.
By providing a list of PGN tags in the usual PGN format, you can:
- Add new tags if the chapter doesn't have them yet
- Update existing chapter tags
- Delete existing chapter tags, by providing a tag with an empty value.

The chapter keeps the tags that you don't provide.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@param chapterId The chapter ID
	@return StudiesAPIApiStudyChapterTagsRequest
	*/
	ApiStudyChapterTags(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyChapterTagsRequest

	// ApiStudyChapterTagsExecute executes the request
	ApiStudyChapterTagsExecute(r StudiesAPIApiStudyChapterTagsRequest) (*http.Response, error)

	/*
	ApiStudyImportPGN Import PGN into a study

	Imports arbitrary PGN into an existing [study](https://lichess.org/study). Creates a new chapter in the study.
If the PGN contains multiple games (separated by 2 or more newlines)
then multiple chapters will be created within the study.
Note that a study can contain at most 64 chapters.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId ID of the study
	@return StudiesAPIApiStudyImportPGNRequest
	*/
	ApiStudyImportPGN(ctx context.Context, studyId string) StudiesAPIApiStudyImportPGNRequest

	// ApiStudyImportPGNExecute executes the request
	//  @return StudyImportPgnChapters
	ApiStudyImportPGNExecute(r StudiesAPIApiStudyImportPGNRequest) (*StudyImportPgnChapters, *http.Response, error)

	/*
	ApiStudyPost Create a new Study

	Create a [study](https://lichess.org/study), and a new empty chapter within it.
You can make up to 30 new studies per day.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StudiesAPIApiStudyPostRequest
	*/
	ApiStudyPost(ctx context.Context) StudiesAPIApiStudyPostRequest

	// ApiStudyPostExecute executes the request
	//  @return ApiStudyPost200Response
	ApiStudyPostExecute(r StudiesAPIApiStudyPostRequest) (*ApiStudyPost200Response, *http.Response, error)

	/*
	ApiStudyStudyIdChapterIdDelete Delete a study chapter

	Delete a chapter of a study you own. This is definitive.
A study must have at least one chapter; so if you delete the last chapter,
an empty one will be automatically created to replace it.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@param chapterId The chapter ID
	@return StudiesAPIApiStudyStudyIdChapterIdDeleteRequest
	*/
	ApiStudyStudyIdChapterIdDelete(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyStudyIdChapterIdDeleteRequest

	// ApiStudyStudyIdChapterIdDeleteExecute executes the request
	ApiStudyStudyIdChapterIdDeleteExecute(r StudiesAPIApiStudyStudyIdChapterIdDeleteRequest) (*http.Response, error)

	/*
	StudyAllChaptersHead Study metadata

	Only get the study headers, including `Last-Modified`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@return StudiesAPIStudyAllChaptersHeadRequest
	*/
	StudyAllChaptersHead(ctx context.Context, studyId string) StudiesAPIStudyAllChaptersHeadRequest

	// StudyAllChaptersHeadExecute executes the request
	StudyAllChaptersHeadExecute(r StudiesAPIStudyAllChaptersHeadRequest) (*http.Response, error)

	/*
	StudyAllChaptersPgn Export all chapters

	Download all chapters of a study in PGN format.
If authenticated, then all public, unlisted, and private study chapters are read.
If not, only public (non-unlisted) study chapters are read.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@return StudiesAPIStudyAllChaptersPgnRequest
	*/
	StudyAllChaptersPgn(ctx context.Context, studyId string) StudiesAPIStudyAllChaptersPgnRequest

	// StudyAllChaptersPgnExecute executes the request
	//  @return string
	StudyAllChaptersPgnExecute(r StudiesAPIStudyAllChaptersPgnRequest) (string, *http.Response, error)

	/*
	StudyChapterPgn Export one study chapter

	Download one study chapter in PGN format.
If authenticated, then all public, unlisted, and private study chapters are read.
If not, only public (non-unlisted) study chapters are read.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param studyId The study ID
	@param chapterId The chapter ID
	@return StudiesAPIStudyChapterPgnRequest
	*/
	StudyChapterPgn(ctx context.Context, studyId string, chapterId string) StudiesAPIStudyChapterPgnRequest

	// StudyChapterPgnExecute executes the request
	//  @return string
	StudyChapterPgnExecute(r StudiesAPIStudyChapterPgnRequest) (string, *http.Response, error)

	/*
	StudyExportAllPgn Export all studies of a user

	Download all chapters of all studies of a user in PGN format.
If authenticated, then all public, unlisted, and private studies are included.
If not, only public (non-unlisted) studies are included.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The user whose studies we export
	@return StudiesAPIStudyExportAllPgnRequest
	*/
	StudyExportAllPgn(ctx context.Context, username string) StudiesAPIStudyExportAllPgnRequest

	// StudyExportAllPgnExecute executes the request
	//  @return string
	StudyExportAllPgnExecute(r StudiesAPIStudyExportAllPgnRequest) (string, *http.Response, error)

	/*
	StudyListMetadata List studies of a user

	Get metadata (name and dates) of all studies of a user.
If authenticated, then all public, unlisted, and private studies are included.
If not, only public (non-unlisted) studies are included.
Studies are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The user whose studies we list
	@return StudiesAPIStudyListMetadataRequest
	*/
	StudyListMetadata(ctx context.Context, username string) StudiesAPIStudyListMetadataRequest

	// StudyListMetadataExecute executes the request
	//  @return StudyMetadata
	StudyListMetadataExecute(r StudiesAPIStudyListMetadataRequest) (*StudyMetadata, *http.Response, error)
}

// StudiesAPIService StudiesAPI service
type StudiesAPIService service

type StudiesAPIApiStudyChapterMovesRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	chapterId string
	pgn *string
}

// PGN text containing the moves that will replace the chapter&#39;s existing moves. Any provided tags are ignored. 
func (r StudiesAPIApiStudyChapterMovesRequest) Pgn(pgn string) StudiesAPIApiStudyChapterMovesRequest {
	r.pgn = &pgn
	return r
}

func (r StudiesAPIApiStudyChapterMovesRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiStudyChapterMovesExecute(r)
}

/*
ApiStudyChapterMoves Update the moves of a study chapter

Replaces the moves tree of a study chapter.
No tags will be modified.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @param chapterId The chapter ID
 @return StudiesAPIApiStudyChapterMovesRequest
*/
func (a *StudiesAPIService) ApiStudyChapterMoves(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyChapterMovesRequest {
	return StudiesAPIApiStudyChapterMovesRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
		chapterId: chapterId,
	}
}

// Execute executes the request
func (a *StudiesAPIService) ApiStudyChapterMovesExecute(r StudiesAPIApiStudyChapterMovesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.ApiStudyChapterMoves")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}/{chapterId}/moves"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chapterId"+"}", url.PathEscape(parameterValueToString(r.chapterId, "chapterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return nil, reportError("studyId must have less than 8 elements")
	}
	if strlen(r.chapterId) < 8 {
		return nil, reportError("chapterId must have at least 8 elements")
	}
	if strlen(r.chapterId) > 8 {
		return nil, reportError("chapterId must have less than 8 elements")
	}
	if r.pgn == nil {
		return nil, reportError("pgn is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pgn", r.pgn, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type StudiesAPIApiStudyChapterTagsRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	chapterId string
	pgn *string
}

// PGN text containing the tags. Only the tags are used. Moves are just ignored. 
func (r StudiesAPIApiStudyChapterTagsRequest) Pgn(pgn string) StudiesAPIApiStudyChapterTagsRequest {
	r.pgn = &pgn
	return r
}

func (r StudiesAPIApiStudyChapterTagsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiStudyChapterTagsExecute(r)
}

/*
ApiStudyChapterTags Update PGN tags of a study chapter

Add, update and delete the PGN tags of a study.
By providing a list of PGN tags in the usual PGN format, you can:
- Add new tags if the chapter doesn't have them yet
- Update existing chapter tags
- Delete existing chapter tags, by providing a tag with an empty value.

The chapter keeps the tags that you don't provide.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @param chapterId The chapter ID
 @return StudiesAPIApiStudyChapterTagsRequest
*/
func (a *StudiesAPIService) ApiStudyChapterTags(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyChapterTagsRequest {
	return StudiesAPIApiStudyChapterTagsRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
		chapterId: chapterId,
	}
}

// Execute executes the request
func (a *StudiesAPIService) ApiStudyChapterTagsExecute(r StudiesAPIApiStudyChapterTagsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.ApiStudyChapterTags")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}/{chapterId}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chapterId"+"}", url.PathEscape(parameterValueToString(r.chapterId, "chapterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return nil, reportError("studyId must have less than 8 elements")
	}
	if strlen(r.chapterId) < 8 {
		return nil, reportError("chapterId must have at least 8 elements")
	}
	if strlen(r.chapterId) > 8 {
		return nil, reportError("chapterId must have less than 8 elements")
	}
	if r.pgn == nil {
		return nil, reportError("pgn is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pgn", r.pgn, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type StudiesAPIApiStudyImportPGNRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	pgn *string
	name *string
	orientation *string
	variant *VariantKey
	mode *string
}

// PGN to import. Can contain multiple games separated by 2 or more newlines. 
func (r StudiesAPIApiStudyImportPGNRequest) Pgn(pgn string) StudiesAPIApiStudyImportPGNRequest {
	r.pgn = &pgn
	return r
}

// Name of the new chapter. If not specified, or if multiple chapters are created, the names will be inferred from the PGN tags. 
func (r StudiesAPIApiStudyImportPGNRequest) Name(name string) StudiesAPIApiStudyImportPGNRequest {
	r.name = &name
	return r
}

// Board orientation. If not specified, the orientation is automatically determined. 
func (r StudiesAPIApiStudyImportPGNRequest) Orientation(orientation string) StudiesAPIApiStudyImportPGNRequest {
	r.orientation = &orientation
	return r
}

func (r StudiesAPIApiStudyImportPGNRequest) Variant(variant VariantKey) StudiesAPIApiStudyImportPGNRequest {
	r.variant = &variant
	return r
}

// Analysis mode. If not specified, Normal analysis. * practice - Practise with Computer * conceal - Hide next moves * gamebook - Interactive lesson 
func (r StudiesAPIApiStudyImportPGNRequest) Mode(mode string) StudiesAPIApiStudyImportPGNRequest {
	r.mode = &mode
	return r
}

func (r StudiesAPIApiStudyImportPGNRequest) Execute() (*StudyImportPgnChapters, *http.Response, error) {
	return r.ApiService.ApiStudyImportPGNExecute(r)
}

/*
ApiStudyImportPGN Import PGN into a study

Imports arbitrary PGN into an existing [study](https://lichess.org/study). Creates a new chapter in the study.
If the PGN contains multiple games (separated by 2 or more newlines)
then multiple chapters will be created within the study.
Note that a study can contain at most 64 chapters.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId ID of the study
 @return StudiesAPIApiStudyImportPGNRequest
*/
func (a *StudiesAPIService) ApiStudyImportPGN(ctx context.Context, studyId string) StudiesAPIApiStudyImportPGNRequest {
	return StudiesAPIApiStudyImportPGNRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
	}
}

// Execute executes the request
//  @return StudyImportPgnChapters
func (a *StudiesAPIService) ApiStudyImportPGNExecute(r StudiesAPIApiStudyImportPGNRequest) (*StudyImportPgnChapters, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StudyImportPgnChapters
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.ApiStudyImportPGN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}/import-pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pgn == nil {
		return localVarReturnValue, nil, reportError("pgn is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pgn", r.pgn, "", "")
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.orientation != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "orientation", r.orientation, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mode", r.mode, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type StudiesAPIApiStudyPostRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	name *string
	visibility *string
	computer *StudyUserSelection
	explorer *StudyUserSelection
	cloneable *StudyUserSelection
	shareable *StudyUserSelection
	chat *StudyUserSelection
	flair *string
	sticky *bool
	description *bool
}

// The study name.
func (r StudiesAPIApiStudyPostRequest) Name(name string) StudiesAPIApiStudyPostRequest {
	r.name = &name
	return r
}

// Who can view the study. * &#x60;public&#x60;: Default. Anyone can view the study, it appears on public listings * &#x60;unlisted&#x60;: Only people with the link can view the study, it doesn&#39;t appear on public listings * &#x60;private&#x60;: Only the study members can view the study 
func (r StudiesAPIApiStudyPostRequest) Visibility(visibility string) StudiesAPIApiStudyPostRequest {
	r.visibility = &visibility
	return r
}

func (r StudiesAPIApiStudyPostRequest) Computer(computer StudyUserSelection) StudiesAPIApiStudyPostRequest {
	r.computer = &computer
	return r
}

func (r StudiesAPIApiStudyPostRequest) Explorer(explorer StudyUserSelection) StudiesAPIApiStudyPostRequest {
	r.explorer = &explorer
	return r
}

func (r StudiesAPIApiStudyPostRequest) Cloneable(cloneable StudyUserSelection) StudiesAPIApiStudyPostRequest {
	r.cloneable = &cloneable
	return r
}

func (r StudiesAPIApiStudyPostRequest) Shareable(shareable StudyUserSelection) StudiesAPIApiStudyPostRequest {
	r.shareable = &shareable
	return r
}

func (r StudiesAPIApiStudyPostRequest) Chat(chat StudyUserSelection) StudiesAPIApiStudyPostRequest {
	r.chat = &chat
	return r
}

// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
func (r StudiesAPIApiStudyPostRequest) Flair(flair string) StudiesAPIApiStudyPostRequest {
	r.flair = &flair
	return r
}

// Keep everyone on the same chapter and position.
func (r StudiesAPIApiStudyPostRequest) Sticky(sticky bool) StudiesAPIApiStudyPostRequest {
	r.sticky = &sticky
	return r
}

// Add pinned study comment right under the board.
func (r StudiesAPIApiStudyPostRequest) Description(description bool) StudiesAPIApiStudyPostRequest {
	r.description = &description
	return r
}

func (r StudiesAPIApiStudyPostRequest) Execute() (*ApiStudyPost200Response, *http.Response, error) {
	return r.ApiService.ApiStudyPostExecute(r)
}

/*
ApiStudyPost Create a new Study

Create a [study](https://lichess.org/study), and a new empty chapter within it.
You can make up to 30 new studies per day.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StudiesAPIApiStudyPostRequest
*/
func (a *StudiesAPIService) ApiStudyPost(ctx context.Context) StudiesAPIApiStudyPostRequest {
	return StudiesAPIApiStudyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStudyPost200Response
func (a *StudiesAPIService) ApiStudyPostExecute(r StudiesAPIApiStudyPostRequest) (*ApiStudyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStudyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.ApiStudyPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 2 {
		return localVarReturnValue, nil, reportError("name must have at least 2 elements")
	}
	if strlen(*r.name) > 100 {
		return localVarReturnValue, nil, reportError("name must have less than 100 elements")
	}
	if r.visibility == nil {
		return localVarReturnValue, nil, reportError("visibility is required and must be specified")
	}
	if r.computer == nil {
		return localVarReturnValue, nil, reportError("computer is required and must be specified")
	}
	if r.explorer == nil {
		return localVarReturnValue, nil, reportError("explorer is required and must be specified")
	}
	if r.cloneable == nil {
		return localVarReturnValue, nil, reportError("cloneable is required and must be specified")
	}
	if r.shareable == nil {
		return localVarReturnValue, nil, reportError("shareable is required and must be specified")
	}
	if r.chat == nil {
		return localVarReturnValue, nil, reportError("chat is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "visibility", r.visibility, "", "")
	if r.flair != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "flair", r.flair, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "computer", r.computer, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "explorer", r.explorer, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "cloneable", r.cloneable, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "shareable", r.shareable, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "chat", r.chat, "", "")
	if r.sticky != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sticky", r.sticky, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type StudiesAPIApiStudyStudyIdChapterIdDeleteRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	chapterId string
}

func (r StudiesAPIApiStudyStudyIdChapterIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiStudyStudyIdChapterIdDeleteExecute(r)
}

/*
ApiStudyStudyIdChapterIdDelete Delete a study chapter

Delete a chapter of a study you own. This is definitive.
A study must have at least one chapter; so if you delete the last chapter,
an empty one will be automatically created to replace it.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @param chapterId The chapter ID
 @return StudiesAPIApiStudyStudyIdChapterIdDeleteRequest
*/
func (a *StudiesAPIService) ApiStudyStudyIdChapterIdDelete(ctx context.Context, studyId string, chapterId string) StudiesAPIApiStudyStudyIdChapterIdDeleteRequest {
	return StudiesAPIApiStudyStudyIdChapterIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
		chapterId: chapterId,
	}
}

// Execute executes the request
func (a *StudiesAPIService) ApiStudyStudyIdChapterIdDeleteExecute(r StudiesAPIApiStudyStudyIdChapterIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.ApiStudyStudyIdChapterIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}/{chapterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chapterId"+"}", url.PathEscape(parameterValueToString(r.chapterId, "chapterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return nil, reportError("studyId must have less than 8 elements")
	}
	if strlen(r.chapterId) < 8 {
		return nil, reportError("chapterId must have at least 8 elements")
	}
	if strlen(r.chapterId) > 8 {
		return nil, reportError("chapterId must have less than 8 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type StudiesAPIStudyAllChaptersHeadRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
}

func (r StudiesAPIStudyAllChaptersHeadRequest) Execute() (*http.Response, error) {
	return r.ApiService.StudyAllChaptersHeadExecute(r)
}

/*
StudyAllChaptersHead Study metadata

Only get the study headers, including `Last-Modified`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @return StudiesAPIStudyAllChaptersHeadRequest
*/
func (a *StudiesAPIService) StudyAllChaptersHead(ctx context.Context, studyId string) StudiesAPIStudyAllChaptersHeadRequest {
	return StudiesAPIStudyAllChaptersHeadRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
	}
}

// Execute executes the request
func (a *StudiesAPIService) StudyAllChaptersHeadExecute(r StudiesAPIStudyAllChaptersHeadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.StudyAllChaptersHead")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return nil, reportError("studyId must have less than 8 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type StudiesAPIStudyAllChaptersPgnRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	clocks *bool
	comments *bool
	variations *bool
	orientation *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r StudiesAPIStudyAllChaptersPgnRequest) Clocks(clocks bool) StudiesAPIStudyAllChaptersPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r StudiesAPIStudyAllChaptersPgnRequest) Comments(comments bool) StudiesAPIStudyAllChaptersPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r StudiesAPIStudyAllChaptersPgnRequest) Variations(variations bool) StudiesAPIStudyAllChaptersPgnRequest {
	r.variations = &variations
	return r
}

// Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60; 
func (r StudiesAPIStudyAllChaptersPgnRequest) Orientation(orientation bool) StudiesAPIStudyAllChaptersPgnRequest {
	r.orientation = &orientation
	return r
}

func (r StudiesAPIStudyAllChaptersPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.StudyAllChaptersPgnExecute(r)
}

/*
StudyAllChaptersPgn Export all chapters

Download all chapters of a study in PGN format.
If authenticated, then all public, unlisted, and private study chapters are read.
If not, only public (non-unlisted) study chapters are read.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @return StudiesAPIStudyAllChaptersPgnRequest
*/
func (a *StudiesAPIService) StudyAllChaptersPgn(ctx context.Context, studyId string) StudiesAPIStudyAllChaptersPgnRequest {
	return StudiesAPIStudyAllChaptersPgnRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
	}
}

// Execute executes the request
//  @return string
func (a *StudiesAPIService) StudyAllChaptersPgnExecute(r StudiesAPIStudyAllChaptersPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.StudyAllChaptersPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return localVarReturnValue, nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return localVarReturnValue, nil, reportError("studyId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
	if r.variations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", r.variations, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", defaultValue, "form", "")
		r.variations = &defaultValue
	}
	if r.orientation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", r.orientation, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", defaultValue, "form", "")
		r.orientation = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type StudiesAPIStudyChapterPgnRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	studyId string
	chapterId string
	clocks *bool
	comments *bool
	variations *bool
	orientation *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r StudiesAPIStudyChapterPgnRequest) Clocks(clocks bool) StudiesAPIStudyChapterPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r StudiesAPIStudyChapterPgnRequest) Comments(comments bool) StudiesAPIStudyChapterPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r StudiesAPIStudyChapterPgnRequest) Variations(variations bool) StudiesAPIStudyChapterPgnRequest {
	r.variations = &variations
	return r
}

// Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60; 
func (r StudiesAPIStudyChapterPgnRequest) Orientation(orientation bool) StudiesAPIStudyChapterPgnRequest {
	r.orientation = &orientation
	return r
}

func (r StudiesAPIStudyChapterPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.StudyChapterPgnExecute(r)
}

/*
StudyChapterPgn Export one study chapter

Download one study chapter in PGN format.
If authenticated, then all public, unlisted, and private study chapters are read.
If not, only public (non-unlisted) study chapters are read.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param studyId The study ID
 @param chapterId The chapter ID
 @return StudiesAPIStudyChapterPgnRequest
*/
func (a *StudiesAPIService) StudyChapterPgn(ctx context.Context, studyId string, chapterId string) StudiesAPIStudyChapterPgnRequest {
	return StudiesAPIStudyChapterPgnRequest{
		ApiService: a,
		ctx: ctx,
		studyId: studyId,
		chapterId: chapterId,
	}
}

// Execute executes the request
//  @return string
func (a *StudiesAPIService) StudyChapterPgnExecute(r StudiesAPIStudyChapterPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.StudyChapterPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/{studyId}/{chapterId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"studyId"+"}", url.PathEscape(parameterValueToString(r.studyId, "studyId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chapterId"+"}", url.PathEscape(parameterValueToString(r.chapterId, "chapterId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.studyId) < 8 {
		return localVarReturnValue, nil, reportError("studyId must have at least 8 elements")
	}
	if strlen(r.studyId) > 8 {
		return localVarReturnValue, nil, reportError("studyId must have less than 8 elements")
	}
	if strlen(r.chapterId) < 8 {
		return localVarReturnValue, nil, reportError("chapterId must have at least 8 elements")
	}
	if strlen(r.chapterId) > 8 {
		return localVarReturnValue, nil, reportError("chapterId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
	if r.variations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", r.variations, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", defaultValue, "form", "")
		r.variations = &defaultValue
	}
	if r.orientation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", r.orientation, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", defaultValue, "form", "")
		r.orientation = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type StudiesAPIStudyExportAllPgnRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	username string
	clocks *bool
	comments *bool
	variations *bool
	orientation *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r StudiesAPIStudyExportAllPgnRequest) Clocks(clocks bool) StudiesAPIStudyExportAllPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis and annotator comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] } a3 { White is in a pickle. }&#x60; 
func (r StudiesAPIStudyExportAllPgnRequest) Comments(comments bool) StudiesAPIStudyExportAllPgnRequest {
	r.comments = &comments
	return r
}

// Include non-mainline moves, when available. Example: &#x60;4. d4 Bb4+ (4... Nc6 5. Nf3 Bb4+ 6. Bd2 (6. Nbd2 O-O 7. O-O) 6... Bd6) 5. Nd2&#x60; 
func (r StudiesAPIStudyExportAllPgnRequest) Variations(variations bool) StudiesAPIStudyExportAllPgnRequest {
	r.variations = &variations
	return r
}

// Add a &#x60;Orientation&#x60; PGN tag with the chapter predefined orientation. Example: &#x60;[Orientation \&quot;white\&quot;]&#x60; 
func (r StudiesAPIStudyExportAllPgnRequest) Orientation(orientation bool) StudiesAPIStudyExportAllPgnRequest {
	r.orientation = &orientation
	return r
}

func (r StudiesAPIStudyExportAllPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.StudyExportAllPgnExecute(r)
}

/*
StudyExportAllPgn Export all studies of a user

Download all chapters of all studies of a user in PGN format.
If authenticated, then all public, unlisted, and private studies are included.
If not, only public (non-unlisted) studies are included.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user whose studies we export
 @return StudiesAPIStudyExportAllPgnRequest
*/
func (a *StudiesAPIService) StudyExportAllPgn(ctx context.Context, username string) StudiesAPIStudyExportAllPgnRequest {
	return StudiesAPIStudyExportAllPgnRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return string
func (a *StudiesAPIService) StudyExportAllPgnExecute(r StudiesAPIStudyExportAllPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.StudyExportAllPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/by/{username}/export.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
	if r.variations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", r.variations, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "variations", defaultValue, "form", "")
		r.variations = &defaultValue
	}
	if r.orientation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", r.orientation, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "orientation", defaultValue, "form", "")
		r.orientation = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type StudiesAPIStudyListMetadataRequest struct {
	ctx context.Context
	ApiService StudiesAPI
	username string
}

func (r StudiesAPIStudyListMetadataRequest) Execute() (*StudyMetadata, *http.Response, error) {
	return r.ApiService.StudyListMetadataExecute(r)
}

/*
StudyListMetadata List studies of a user

Get metadata (name and dates) of all studies of a user.
If authenticated, then all public, unlisted, and private studies are included.
If not, only public (non-unlisted) studies are included.
Studies are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user whose studies we list
 @return StudiesAPIStudyListMetadataRequest
*/
func (a *StudiesAPIService) StudyListMetadata(ctx context.Context, username string) StudiesAPIStudyListMetadataRequest {
	return StudiesAPIStudyListMetadataRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return StudyMetadata
func (a *StudiesAPIService) StudyListMetadataExecute(r StudiesAPIStudyListMetadataRequest) (*StudyMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StudyMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiesAPIService.StudyListMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/study/by/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
