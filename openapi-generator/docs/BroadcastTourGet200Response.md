# BroadcastTourGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tour** | [**BroadcastsOfficial200ResponseTour**](BroadcastsOfficial200ResponseTour.md) |  | 
**Group** | Pointer to [**BroadcastTourGet200ResponseGroup**](BroadcastTourGet200ResponseGroup.md) |  | [optional] 
**Rounds** | [**[]BroadcastsOfficial200ResponseRoundsInner**](BroadcastsOfficial200ResponseRoundsInner.md) |  | 
**DefaultRoundId** | Pointer to **string** |  | [optional] 
**Photos** | Pointer to [**map[string]BroadcastsOfficial200ResponsePhotosValue**](BroadcastsOfficial200ResponsePhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | [optional] 

## Methods

### NewBroadcastTourGet200Response

`func NewBroadcastTourGet200Response(tour BroadcastsOfficial200ResponseTour, rounds []BroadcastsOfficial200ResponseRoundsInner, ) *BroadcastTourGet200Response`

NewBroadcastTourGet200Response instantiates a new BroadcastTourGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTourGet200ResponseWithDefaults

`func NewBroadcastTourGet200ResponseWithDefaults() *BroadcastTourGet200Response`

NewBroadcastTourGet200ResponseWithDefaults instantiates a new BroadcastTourGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTour

`func (o *BroadcastTourGet200Response) GetTour() BroadcastsOfficial200ResponseTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastTourGet200Response) GetTourOk() (*BroadcastsOfficial200ResponseTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastTourGet200Response) SetTour(v BroadcastsOfficial200ResponseTour)`

SetTour sets Tour field to given value.


### GetGroup

`func (o *BroadcastTourGet200Response) GetGroup() BroadcastTourGet200ResponseGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastTourGet200Response) GetGroupOk() (*BroadcastTourGet200ResponseGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastTourGet200Response) SetGroup(v BroadcastTourGet200ResponseGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastTourGet200Response) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRounds

`func (o *BroadcastTourGet200Response) GetRounds() []BroadcastsOfficial200ResponseRoundsInner`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *BroadcastTourGet200Response) GetRoundsOk() (*[]BroadcastsOfficial200ResponseRoundsInner, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *BroadcastTourGet200Response) SetRounds(v []BroadcastsOfficial200ResponseRoundsInner)`

SetRounds sets Rounds field to given value.


### GetDefaultRoundId

`func (o *BroadcastTourGet200Response) GetDefaultRoundId() string`

GetDefaultRoundId returns the DefaultRoundId field if non-nil, zero value otherwise.

### GetDefaultRoundIdOk

`func (o *BroadcastTourGet200Response) GetDefaultRoundIdOk() (*string, bool)`

GetDefaultRoundIdOk returns a tuple with the DefaultRoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoundId

`func (o *BroadcastTourGet200Response) SetDefaultRoundId(v string)`

SetDefaultRoundId sets DefaultRoundId field to given value.

### HasDefaultRoundId

`func (o *BroadcastTourGet200Response) HasDefaultRoundId() bool`

HasDefaultRoundId returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastTourGet200Response) GetPhotos() map[string]BroadcastsOfficial200ResponsePhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastTourGet200Response) GetPhotosOk() (*map[string]BroadcastsOfficial200ResponsePhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastTourGet200Response) SetPhotos(v map[string]BroadcastsOfficial200ResponsePhotosValue)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *BroadcastTourGet200Response) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


