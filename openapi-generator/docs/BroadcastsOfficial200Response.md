# BroadcastsOfficial200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tour** | [**BroadcastsOfficial200ResponseTour**](BroadcastsOfficial200ResponseTour.md) |  | 
**Group** | Pointer to **string** |  | [optional] 
**Rounds** | [**[]BroadcastsOfficial200ResponseRoundsInner**](BroadcastsOfficial200ResponseRoundsInner.md) |  | 
**DefaultRoundId** | Pointer to **string** |  | [optional] 
**Photos** | Pointer to [**map[string]BroadcastsOfficial200ResponsePhotosValue**](BroadcastsOfficial200ResponsePhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | [optional] 

## Methods

### NewBroadcastsOfficial200Response

`func NewBroadcastsOfficial200Response(tour BroadcastsOfficial200ResponseTour, rounds []BroadcastsOfficial200ResponseRoundsInner, ) *BroadcastsOfficial200Response`

NewBroadcastsOfficial200Response instantiates a new BroadcastsOfficial200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastsOfficial200ResponseWithDefaults

`func NewBroadcastsOfficial200ResponseWithDefaults() *BroadcastsOfficial200Response`

NewBroadcastsOfficial200ResponseWithDefaults instantiates a new BroadcastsOfficial200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTour

`func (o *BroadcastsOfficial200Response) GetTour() BroadcastsOfficial200ResponseTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastsOfficial200Response) GetTourOk() (*BroadcastsOfficial200ResponseTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastsOfficial200Response) SetTour(v BroadcastsOfficial200ResponseTour)`

SetTour sets Tour field to given value.


### GetGroup

`func (o *BroadcastsOfficial200Response) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastsOfficial200Response) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastsOfficial200Response) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastsOfficial200Response) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRounds

`func (o *BroadcastsOfficial200Response) GetRounds() []BroadcastsOfficial200ResponseRoundsInner`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *BroadcastsOfficial200Response) GetRoundsOk() (*[]BroadcastsOfficial200ResponseRoundsInner, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *BroadcastsOfficial200Response) SetRounds(v []BroadcastsOfficial200ResponseRoundsInner)`

SetRounds sets Rounds field to given value.


### GetDefaultRoundId

`func (o *BroadcastsOfficial200Response) GetDefaultRoundId() string`

GetDefaultRoundId returns the DefaultRoundId field if non-nil, zero value otherwise.

### GetDefaultRoundIdOk

`func (o *BroadcastsOfficial200Response) GetDefaultRoundIdOk() (*string, bool)`

GetDefaultRoundIdOk returns a tuple with the DefaultRoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoundId

`func (o *BroadcastsOfficial200Response) SetDefaultRoundId(v string)`

SetDefaultRoundId sets DefaultRoundId field to given value.

### HasDefaultRoundId

`func (o *BroadcastsOfficial200Response) HasDefaultRoundId() bool`

HasDefaultRoundId returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastsOfficial200Response) GetPhotos() map[string]BroadcastsOfficial200ResponsePhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastsOfficial200Response) GetPhotosOk() (*map[string]BroadcastsOfficial200ResponsePhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastsOfficial200Response) SetPhotos(v map[string]BroadcastsOfficial200ResponsePhotosValue)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *BroadcastsOfficial200Response) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


