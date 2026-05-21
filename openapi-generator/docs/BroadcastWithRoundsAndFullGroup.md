# BroadcastWithRoundsAndFullGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tour** | [**BroadcastTour**](BroadcastTour.md) |  | 
**Group** | Pointer to [**BroadcastGroup**](BroadcastGroup.md) |  | [optional] 
**Rounds** | [**[]BroadcastRoundInfo**](BroadcastRoundInfo.md) |  | 
**DefaultRoundId** | Pointer to **string** |  | [optional] 
**Photos** | Pointer to [**map[string]BroadcastPhotosValue**](BroadcastPhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | [optional] 

## Methods

### NewBroadcastWithRoundsAndFullGroup

`func NewBroadcastWithRoundsAndFullGroup(tour BroadcastTour, rounds []BroadcastRoundInfo, ) *BroadcastWithRoundsAndFullGroup`

NewBroadcastWithRoundsAndFullGroup instantiates a new BroadcastWithRoundsAndFullGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastWithRoundsAndFullGroupWithDefaults

`func NewBroadcastWithRoundsAndFullGroupWithDefaults() *BroadcastWithRoundsAndFullGroup`

NewBroadcastWithRoundsAndFullGroupWithDefaults instantiates a new BroadcastWithRoundsAndFullGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTour

`func (o *BroadcastWithRoundsAndFullGroup) GetTour() BroadcastTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastWithRoundsAndFullGroup) GetTourOk() (*BroadcastTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastWithRoundsAndFullGroup) SetTour(v BroadcastTour)`

SetTour sets Tour field to given value.


### GetGroup

`func (o *BroadcastWithRoundsAndFullGroup) GetGroup() BroadcastGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastWithRoundsAndFullGroup) GetGroupOk() (*BroadcastGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastWithRoundsAndFullGroup) SetGroup(v BroadcastGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastWithRoundsAndFullGroup) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRounds

`func (o *BroadcastWithRoundsAndFullGroup) GetRounds() []BroadcastRoundInfo`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *BroadcastWithRoundsAndFullGroup) GetRoundsOk() (*[]BroadcastRoundInfo, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *BroadcastWithRoundsAndFullGroup) SetRounds(v []BroadcastRoundInfo)`

SetRounds sets Rounds field to given value.


### GetDefaultRoundId

`func (o *BroadcastWithRoundsAndFullGroup) GetDefaultRoundId() string`

GetDefaultRoundId returns the DefaultRoundId field if non-nil, zero value otherwise.

### GetDefaultRoundIdOk

`func (o *BroadcastWithRoundsAndFullGroup) GetDefaultRoundIdOk() (*string, bool)`

GetDefaultRoundIdOk returns a tuple with the DefaultRoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoundId

`func (o *BroadcastWithRoundsAndFullGroup) SetDefaultRoundId(v string)`

SetDefaultRoundId sets DefaultRoundId field to given value.

### HasDefaultRoundId

`func (o *BroadcastWithRoundsAndFullGroup) HasDefaultRoundId() bool`

HasDefaultRoundId returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastWithRoundsAndFullGroup) GetPhotos() map[string]BroadcastPhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastWithRoundsAndFullGroup) GetPhotosOk() (*map[string]BroadcastPhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastWithRoundsAndFullGroup) SetPhotos(v map[string]BroadcastPhotosValue)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *BroadcastWithRoundsAndFullGroup) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


