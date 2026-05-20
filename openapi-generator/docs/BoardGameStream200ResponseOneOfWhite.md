# BoardGameStream200ResponseOneOfWhite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiLevel** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 

## Methods

### NewBoardGameStream200ResponseOneOfWhite

`func NewBoardGameStream200ResponseOneOfWhite(id string, name string, ) *BoardGameStream200ResponseOneOfWhite`

NewBoardGameStream200ResponseOneOfWhite instantiates a new BoardGameStream200ResponseOneOfWhite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardGameStream200ResponseOneOfWhiteWithDefaults

`func NewBoardGameStream200ResponseOneOfWhiteWithDefaults() *BoardGameStream200ResponseOneOfWhite`

NewBoardGameStream200ResponseOneOfWhiteWithDefaults instantiates a new BoardGameStream200ResponseOneOfWhite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiLevel

`func (o *BoardGameStream200ResponseOneOfWhite) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *BoardGameStream200ResponseOneOfWhite) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.

### HasAiLevel

`func (o *BoardGameStream200ResponseOneOfWhite) HasAiLevel() bool`

HasAiLevel returns a boolean if a field has been set.

### GetId

`func (o *BoardGameStream200ResponseOneOfWhite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardGameStream200ResponseOneOfWhite) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BoardGameStream200ResponseOneOfWhite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardGameStream200ResponseOneOfWhite) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BoardGameStream200ResponseOneOfWhite) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BoardGameStream200ResponseOneOfWhite) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BoardGameStream200ResponseOneOfWhite) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BoardGameStream200ResponseOneOfWhite) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BoardGameStream200ResponseOneOfWhite) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRating

`func (o *BoardGameStream200ResponseOneOfWhite) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BoardGameStream200ResponseOneOfWhite) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BoardGameStream200ResponseOneOfWhite) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetProvisional

`func (o *BoardGameStream200ResponseOneOfWhite) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *BoardGameStream200ResponseOneOfWhite) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *BoardGameStream200ResponseOneOfWhite) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *BoardGameStream200ResponseOneOfWhite) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


