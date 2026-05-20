# BroadcastsOfficial200ResponseTour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**CreatedAt** | **int32** |  | 
**Dates** | Pointer to **[]int64** | Start and end dates of the tournament, as Unix timestamps in milliseconds | [optional] 
**Info** | Pointer to [**BroadcastsOfficial200ResponseTourInfo**](BroadcastsOfficial200ResponseTourInfo.md) |  | [optional] 
**Tier** | Pointer to **int32** | Used to designate featured tournaments on Lichess | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Full tournament description in markdown format, or in HTML if the html&#x3D;1 query parameter is set. | [optional] 
**TeamTable** | Pointer to **bool** |  | [optional] 
**ShowTeamScores** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 
**CommunityOwner** | Pointer to [**ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | [optional] 

## Methods

### NewBroadcastsOfficial200ResponseTour

`func NewBroadcastsOfficial200ResponseTour(id string, name string, slug string, createdAt int32, url string, ) *BroadcastsOfficial200ResponseTour`

NewBroadcastsOfficial200ResponseTour instantiates a new BroadcastsOfficial200ResponseTour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastsOfficial200ResponseTourWithDefaults

`func NewBroadcastsOfficial200ResponseTourWithDefaults() *BroadcastsOfficial200ResponseTour`

NewBroadcastsOfficial200ResponseTourWithDefaults instantiates a new BroadcastsOfficial200ResponseTour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastsOfficial200ResponseTour) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastsOfficial200ResponseTour) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastsOfficial200ResponseTour) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastsOfficial200ResponseTour) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastsOfficial200ResponseTour) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastsOfficial200ResponseTour) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BroadcastsOfficial200ResponseTour) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BroadcastsOfficial200ResponseTour) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BroadcastsOfficial200ResponseTour) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetCreatedAt

`func (o *BroadcastsOfficial200ResponseTour) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BroadcastsOfficial200ResponseTour) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BroadcastsOfficial200ResponseTour) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDates

`func (o *BroadcastsOfficial200ResponseTour) GetDates() []int64`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *BroadcastsOfficial200ResponseTour) GetDatesOk() (*[]int64, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *BroadcastsOfficial200ResponseTour) SetDates(v []int64)`

SetDates sets Dates field to given value.

### HasDates

`func (o *BroadcastsOfficial200ResponseTour) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetInfo

`func (o *BroadcastsOfficial200ResponseTour) GetInfo() BroadcastsOfficial200ResponseTourInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BroadcastsOfficial200ResponseTour) GetInfoOk() (*BroadcastsOfficial200ResponseTourInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BroadcastsOfficial200ResponseTour) SetInfo(v BroadcastsOfficial200ResponseTourInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BroadcastsOfficial200ResponseTour) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTier

`func (o *BroadcastsOfficial200ResponseTour) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *BroadcastsOfficial200ResponseTour) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *BroadcastsOfficial200ResponseTour) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *BroadcastsOfficial200ResponseTour) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetImage

`func (o *BroadcastsOfficial200ResponseTour) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BroadcastsOfficial200ResponseTour) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BroadcastsOfficial200ResponseTour) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BroadcastsOfficial200ResponseTour) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDescription

`func (o *BroadcastsOfficial200ResponseTour) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BroadcastsOfficial200ResponseTour) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BroadcastsOfficial200ResponseTour) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BroadcastsOfficial200ResponseTour) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTeamTable

`func (o *BroadcastsOfficial200ResponseTour) GetTeamTable() bool`

GetTeamTable returns the TeamTable field if non-nil, zero value otherwise.

### GetTeamTableOk

`func (o *BroadcastsOfficial200ResponseTour) GetTeamTableOk() (*bool, bool)`

GetTeamTableOk returns a tuple with the TeamTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamTable

`func (o *BroadcastsOfficial200ResponseTour) SetTeamTable(v bool)`

SetTeamTable sets TeamTable field to given value.

### HasTeamTable

`func (o *BroadcastsOfficial200ResponseTour) HasTeamTable() bool`

HasTeamTable returns a boolean if a field has been set.

### GetShowTeamScores

`func (o *BroadcastsOfficial200ResponseTour) GetShowTeamScores() bool`

GetShowTeamScores returns the ShowTeamScores field if non-nil, zero value otherwise.

### GetShowTeamScoresOk

`func (o *BroadcastsOfficial200ResponseTour) GetShowTeamScoresOk() (*bool, bool)`

GetShowTeamScoresOk returns a tuple with the ShowTeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTeamScores

`func (o *BroadcastsOfficial200ResponseTour) SetShowTeamScores(v bool)`

SetShowTeamScores sets ShowTeamScores field to given value.

### HasShowTeamScores

`func (o *BroadcastsOfficial200ResponseTour) HasShowTeamScores() bool`

HasShowTeamScores returns a boolean if a field has been set.

### GetUrl

`func (o *BroadcastsOfficial200ResponseTour) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BroadcastsOfficial200ResponseTour) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BroadcastsOfficial200ResponseTour) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCommunityOwner

`func (o *BroadcastsOfficial200ResponseTour) GetCommunityOwner() ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetCommunityOwner returns the CommunityOwner field if non-nil, zero value otherwise.

### GetCommunityOwnerOk

`func (o *BroadcastsOfficial200ResponseTour) GetCommunityOwnerOk() (*ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetCommunityOwnerOk returns a tuple with the CommunityOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityOwner

`func (o *BroadcastsOfficial200ResponseTour) SetCommunityOwner(v ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetCommunityOwner sets CommunityOwner field to given value.

### HasCommunityOwner

`func (o *BroadcastsOfficial200ResponseTour) HasCommunityOwner() bool`

HasCommunityOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


