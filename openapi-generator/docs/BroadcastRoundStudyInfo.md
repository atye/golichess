# BroadcastRoundStudyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Writeable** | Pointer to **bool** | Whether the currently authenticated user has permission to update the study | [optional] 
**Features** | Pointer to [**BroadcastRoundStudyInfoFeatures**](BroadcastRoundStudyInfoFeatures.md) |  | [optional] 

## Methods

### NewBroadcastRoundStudyInfo

`func NewBroadcastRoundStudyInfo() *BroadcastRoundStudyInfo`

NewBroadcastRoundStudyInfo instantiates a new BroadcastRoundStudyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundStudyInfoWithDefaults

`func NewBroadcastRoundStudyInfoWithDefaults() *BroadcastRoundStudyInfo`

NewBroadcastRoundStudyInfoWithDefaults instantiates a new BroadcastRoundStudyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWriteable

`func (o *BroadcastRoundStudyInfo) GetWriteable() bool`

GetWriteable returns the Writeable field if non-nil, zero value otherwise.

### GetWriteableOk

`func (o *BroadcastRoundStudyInfo) GetWriteableOk() (*bool, bool)`

GetWriteableOk returns a tuple with the Writeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteable

`func (o *BroadcastRoundStudyInfo) SetWriteable(v bool)`

SetWriteable sets Writeable field to given value.

### HasWriteable

`func (o *BroadcastRoundStudyInfo) HasWriteable() bool`

HasWriteable returns a boolean if a field has been set.

### GetFeatures

`func (o *BroadcastRoundStudyInfo) GetFeatures() BroadcastRoundStudyInfoFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BroadcastRoundStudyInfo) GetFeaturesOk() (*BroadcastRoundStudyInfoFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BroadcastRoundStudyInfo) SetFeatures(v BroadcastRoundStudyInfoFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BroadcastRoundStudyInfo) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


