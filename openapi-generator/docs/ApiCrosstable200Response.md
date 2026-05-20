# ApiCrosstable200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** |  |  | 
**NbGames** | **int32** |  | 

## Methods

### NewApiCrosstable200Response

`func NewApiCrosstable200Response(users map[string]float32, nbGames int32, ) *ApiCrosstable200Response`

NewApiCrosstable200Response instantiates a new ApiCrosstable200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCrosstable200ResponseWithDefaults

`func NewApiCrosstable200ResponseWithDefaults() *ApiCrosstable200Response`

NewApiCrosstable200ResponseWithDefaults instantiates a new ApiCrosstable200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ApiCrosstable200Response) GetUsers() map[string]float32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ApiCrosstable200Response) GetUsersOk() (*map[string]float32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ApiCrosstable200Response) SetUsers(v map[string]float32)`

SetUsers sets Users field to given value.


### SetUsersNil

`func (o *ApiCrosstable200Response) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ApiCrosstable200Response) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetNbGames

`func (o *ApiCrosstable200Response) GetNbGames() int32`

GetNbGames returns the NbGames field if non-nil, zero value otherwise.

### GetNbGamesOk

`func (o *ApiCrosstable200Response) GetNbGamesOk() (*int32, bool)`

GetNbGamesOk returns a tuple with the NbGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbGames

`func (o *ApiCrosstable200Response) SetNbGames(v int32)`

SetNbGames sets NbGames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


