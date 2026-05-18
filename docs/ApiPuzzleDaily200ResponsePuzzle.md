# ApiPuzzleDaily200ResponsePuzzle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InitialPly** | **int32** |  | 
**Plays** | **int32** |  | 
**Rating** | **int32** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**LastMove** | Pointer to **string** | In UCI format, e.g. \&quot;e2e4\&quot; | [optional] 
**Solution** | **[]string** |  | 
**Themes** | **[]string** |  | 

## Methods

### NewApiPuzzleDaily200ResponsePuzzle

`func NewApiPuzzleDaily200ResponsePuzzle(id string, initialPly int32, plays int32, rating int32, solution []string, themes []string, ) *ApiPuzzleDaily200ResponsePuzzle`

NewApiPuzzleDaily200ResponsePuzzle instantiates a new ApiPuzzleDaily200ResponsePuzzle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleDaily200ResponsePuzzleWithDefaults

`func NewApiPuzzleDaily200ResponsePuzzleWithDefaults() *ApiPuzzleDaily200ResponsePuzzle`

NewApiPuzzleDaily200ResponsePuzzleWithDefaults instantiates a new ApiPuzzleDaily200ResponsePuzzle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetId(v string)`

SetId sets Id field to given value.


### GetInitialPly

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetInitialPly() int32`

GetInitialPly returns the InitialPly field if non-nil, zero value otherwise.

### GetInitialPlyOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetInitialPlyOk() (*int32, bool)`

GetInitialPlyOk returns a tuple with the InitialPly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPly

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetInitialPly(v int32)`

SetInitialPly sets InitialPly field to given value.


### GetPlays

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetPlays() int32`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetPlaysOk() (*int32, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetPlays(v int32)`

SetPlays sets Plays field to given value.


### GetRating

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetFen

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *ApiPuzzleDaily200ResponsePuzzle) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetLastMove

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *ApiPuzzleDaily200ResponsePuzzle) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetSolution

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetSolution() []string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetSolutionOk() (*[]string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetSolution(v []string)`

SetSolution sets Solution field to given value.


### GetThemes

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetThemes() []string`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *ApiPuzzleDaily200ResponsePuzzle) GetThemesOk() (*[]string, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *ApiPuzzleDaily200ResponsePuzzle) SetThemes(v []string)`

SetThemes sets Themes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


