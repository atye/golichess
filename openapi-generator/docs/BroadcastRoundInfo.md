# BroadcastRoundInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Ongoing** | Pointer to **bool** |  | [optional] 
**StartsAt** | Pointer to **int64** |  | [optional] 
**StartsAfterPrevious** | Pointer to **bool** | The start date/time is unknown and the round will start automatically when the previous round completes | [optional] 
**FinishedAt** | Pointer to **int64** |  | [optional] 
**Finished** | Pointer to **bool** | Use finishedAt instead | [optional] 
**Url** | **string** |  | 
**Rated** | Pointer to **bool** | Whether the round is used for rating calculations | [optional] 
**CustomScoring** | Pointer to [**BroadcastCustomScoring**](BroadcastCustomScoring.md) |  | [optional] 

## Methods

### NewBroadcastRoundInfo

`func NewBroadcastRoundInfo(id string, name string, slug string, url string, ) *BroadcastRoundInfo`

NewBroadcastRoundInfo instantiates a new BroadcastRoundInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundInfoWithDefaults

`func NewBroadcastRoundInfoWithDefaults() *BroadcastRoundInfo`

NewBroadcastRoundInfoWithDefaults instantiates a new BroadcastRoundInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastRoundInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastRoundInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastRoundInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastRoundInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastRoundInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastRoundInfo) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BroadcastRoundInfo) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BroadcastRoundInfo) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BroadcastRoundInfo) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetOngoing

`func (o *BroadcastRoundInfo) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *BroadcastRoundInfo) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *BroadcastRoundInfo) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *BroadcastRoundInfo) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetStartsAt

`func (o *BroadcastRoundInfo) GetStartsAt() int64`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *BroadcastRoundInfo) GetStartsAtOk() (*int64, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *BroadcastRoundInfo) SetStartsAt(v int64)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *BroadcastRoundInfo) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStartsAfterPrevious

`func (o *BroadcastRoundInfo) GetStartsAfterPrevious() bool`

GetStartsAfterPrevious returns the StartsAfterPrevious field if non-nil, zero value otherwise.

### GetStartsAfterPreviousOk

`func (o *BroadcastRoundInfo) GetStartsAfterPreviousOk() (*bool, bool)`

GetStartsAfterPreviousOk returns a tuple with the StartsAfterPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAfterPrevious

`func (o *BroadcastRoundInfo) SetStartsAfterPrevious(v bool)`

SetStartsAfterPrevious sets StartsAfterPrevious field to given value.

### HasStartsAfterPrevious

`func (o *BroadcastRoundInfo) HasStartsAfterPrevious() bool`

HasStartsAfterPrevious returns a boolean if a field has been set.

### GetFinishedAt

`func (o *BroadcastRoundInfo) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *BroadcastRoundInfo) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *BroadcastRoundInfo) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *BroadcastRoundInfo) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetFinished

`func (o *BroadcastRoundInfo) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *BroadcastRoundInfo) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *BroadcastRoundInfo) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *BroadcastRoundInfo) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetUrl

`func (o *BroadcastRoundInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BroadcastRoundInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BroadcastRoundInfo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRated

`func (o *BroadcastRoundInfo) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *BroadcastRoundInfo) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *BroadcastRoundInfo) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *BroadcastRoundInfo) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetCustomScoring

`func (o *BroadcastRoundInfo) GetCustomScoring() BroadcastCustomScoring`

GetCustomScoring returns the CustomScoring field if non-nil, zero value otherwise.

### GetCustomScoringOk

`func (o *BroadcastRoundInfo) GetCustomScoringOk() (*BroadcastCustomScoring, bool)`

GetCustomScoringOk returns a tuple with the CustomScoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomScoring

`func (o *BroadcastRoundInfo) SetCustomScoring(v BroadcastCustomScoring)`

SetCustomScoring sets CustomScoring field to given value.

### HasCustomScoring

`func (o *BroadcastRoundInfo) HasCustomScoring() bool`

HasCustomScoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


