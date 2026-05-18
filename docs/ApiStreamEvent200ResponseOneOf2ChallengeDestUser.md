# ApiStreamEvent200ResponseOneOf2ChallengeDestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Rating** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Lag** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOf2ChallengeDestUser

`func NewApiStreamEvent200ResponseOneOf2ChallengeDestUser(id string, name string, ) *ApiStreamEvent200ResponseOneOf2ChallengeDestUser`

NewApiStreamEvent200ResponseOneOf2ChallengeDestUser instantiates a new ApiStreamEvent200ResponseOneOf2ChallengeDestUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOf2ChallengeDestUserWithDefaults

`func NewApiStreamEvent200ResponseOneOf2ChallengeDestUserWithDefaults() *ApiStreamEvent200ResponseOneOf2ChallengeDestUser`

NewApiStreamEvent200ResponseOneOf2ChallengeDestUserWithDefaults instantiates a new ApiStreamEvent200ResponseOneOf2ChallengeDestUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetName(v string)`

SetName sets Name field to given value.


### GetRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeDestUser) HasLag() bool`

HasLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


