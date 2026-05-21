# BroadcastTour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**CreatedAt** | **int32** |  | 
**Dates** | Pointer to **[]int64** | Start and end dates of the tournament, as Unix timestamps in milliseconds | [optional] 
**Info** | Pointer to [**BroadcastTourInfo**](BroadcastTourInfo.md) |  | [optional] 
**Tier** | Pointer to **int32** | Used to designate featured tournaments on Lichess | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Full tournament description in markdown format, or in HTML if the html&#x3D;1 query parameter is set. | [optional] 
**TeamTable** | Pointer to **bool** |  | [optional] 
**ShowTeamScores** | Pointer to **bool** |  | [optional] 
**Url** | **string** |  | 
**CommunityOwner** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 

## Methods

### NewBroadcastTour

`func NewBroadcastTour(id string, name string, slug string, createdAt int32, url string, ) *BroadcastTour`

NewBroadcastTour instantiates a new BroadcastTour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTourWithDefaults

`func NewBroadcastTourWithDefaults() *BroadcastTour`

NewBroadcastTourWithDefaults instantiates a new BroadcastTour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastTour) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastTour) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastTour) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastTour) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastTour) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastTour) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BroadcastTour) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BroadcastTour) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BroadcastTour) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetCreatedAt

`func (o *BroadcastTour) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BroadcastTour) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BroadcastTour) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDates

`func (o *BroadcastTour) GetDates() []int64`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *BroadcastTour) GetDatesOk() (*[]int64, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *BroadcastTour) SetDates(v []int64)`

SetDates sets Dates field to given value.

### HasDates

`func (o *BroadcastTour) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetInfo

`func (o *BroadcastTour) GetInfo() BroadcastTourInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BroadcastTour) GetInfoOk() (*BroadcastTourInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BroadcastTour) SetInfo(v BroadcastTourInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BroadcastTour) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTier

`func (o *BroadcastTour) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *BroadcastTour) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *BroadcastTour) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *BroadcastTour) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetImage

`func (o *BroadcastTour) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BroadcastTour) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BroadcastTour) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BroadcastTour) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDescription

`func (o *BroadcastTour) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BroadcastTour) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BroadcastTour) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BroadcastTour) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTeamTable

`func (o *BroadcastTour) GetTeamTable() bool`

GetTeamTable returns the TeamTable field if non-nil, zero value otherwise.

### GetTeamTableOk

`func (o *BroadcastTour) GetTeamTableOk() (*bool, bool)`

GetTeamTableOk returns a tuple with the TeamTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamTable

`func (o *BroadcastTour) SetTeamTable(v bool)`

SetTeamTable sets TeamTable field to given value.

### HasTeamTable

`func (o *BroadcastTour) HasTeamTable() bool`

HasTeamTable returns a boolean if a field has been set.

### GetShowTeamScores

`func (o *BroadcastTour) GetShowTeamScores() bool`

GetShowTeamScores returns the ShowTeamScores field if non-nil, zero value otherwise.

### GetShowTeamScoresOk

`func (o *BroadcastTour) GetShowTeamScoresOk() (*bool, bool)`

GetShowTeamScoresOk returns a tuple with the ShowTeamScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTeamScores

`func (o *BroadcastTour) SetShowTeamScores(v bool)`

SetShowTeamScores sets ShowTeamScores field to given value.

### HasShowTeamScores

`func (o *BroadcastTour) HasShowTeamScores() bool`

HasShowTeamScores returns a boolean if a field has been set.

### GetUrl

`func (o *BroadcastTour) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BroadcastTour) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BroadcastTour) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCommunityOwner

`func (o *BroadcastTour) GetCommunityOwner() LightUser`

GetCommunityOwner returns the CommunityOwner field if non-nil, zero value otherwise.

### GetCommunityOwnerOk

`func (o *BroadcastTour) GetCommunityOwnerOk() (*LightUser, bool)`

GetCommunityOwnerOk returns a tuple with the CommunityOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityOwner

`func (o *BroadcastTour) SetCommunityOwner(v LightUser)`

SetCommunityOwner sets CommunityOwner field to given value.

### HasCommunityOwner

`func (o *BroadcastTour) HasCommunityOwner() bool`

HasCommunityOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


