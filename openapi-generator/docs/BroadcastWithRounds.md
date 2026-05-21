# BroadcastWithRounds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tour** | [**BroadcastTour**](BroadcastTour.md) |  | 
**Group** | Pointer to **string** |  | [optional] 
**Rounds** | [**[]BroadcastRoundInfo**](BroadcastRoundInfo.md) |  | 
**DefaultRoundId** | Pointer to **string** |  | [optional] 
**Photos** | Pointer to [**map[string]BroadcastPhotosValue**](BroadcastPhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | [optional] 

## Methods

### NewBroadcastWithRounds

`func NewBroadcastWithRounds(tour BroadcastTour, rounds []BroadcastRoundInfo, ) *BroadcastWithRounds`

NewBroadcastWithRounds instantiates a new BroadcastWithRounds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastWithRoundsWithDefaults

`func NewBroadcastWithRoundsWithDefaults() *BroadcastWithRounds`

NewBroadcastWithRoundsWithDefaults instantiates a new BroadcastWithRounds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTour

`func (o *BroadcastWithRounds) GetTour() BroadcastTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastWithRounds) GetTourOk() (*BroadcastTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastWithRounds) SetTour(v BroadcastTour)`

SetTour sets Tour field to given value.


### GetGroup

`func (o *BroadcastWithRounds) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastWithRounds) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastWithRounds) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastWithRounds) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRounds

`func (o *BroadcastWithRounds) GetRounds() []BroadcastRoundInfo`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *BroadcastWithRounds) GetRoundsOk() (*[]BroadcastRoundInfo, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *BroadcastWithRounds) SetRounds(v []BroadcastRoundInfo)`

SetRounds sets Rounds field to given value.


### GetDefaultRoundId

`func (o *BroadcastWithRounds) GetDefaultRoundId() string`

GetDefaultRoundId returns the DefaultRoundId field if non-nil, zero value otherwise.

### GetDefaultRoundIdOk

`func (o *BroadcastWithRounds) GetDefaultRoundIdOk() (*string, bool)`

GetDefaultRoundIdOk returns a tuple with the DefaultRoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoundId

`func (o *BroadcastWithRounds) SetDefaultRoundId(v string)`

SetDefaultRoundId sets DefaultRoundId field to given value.

### HasDefaultRoundId

`func (o *BroadcastWithRounds) HasDefaultRoundId() bool`

HasDefaultRoundId returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastWithRounds) GetPhotos() map[string]BroadcastPhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastWithRounds) GetPhotosOk() (*map[string]BroadcastPhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastWithRounds) SetPhotos(v map[string]BroadcastPhotosValue)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *BroadcastWithRounds) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


