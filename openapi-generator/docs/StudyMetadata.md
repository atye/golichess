# StudyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The study ID | 
**Name** | **string** | The study name | 
**CreatedAt** | **int64** | The study creation date | 
**UpdatedAt** | **int64** | The study last update date | 

## Methods

### NewStudyMetadata

`func NewStudyMetadata(id string, name string, createdAt int64, updatedAt int64, ) *StudyMetadata`

NewStudyMetadata instantiates a new StudyMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyMetadataWithDefaults

`func NewStudyMetadataWithDefaults() *StudyMetadata`

NewStudyMetadataWithDefaults instantiates a new StudyMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StudyMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudyMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudyMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StudyMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudyMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudyMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *StudyMetadata) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StudyMetadata) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StudyMetadata) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StudyMetadata) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StudyMetadata) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StudyMetadata) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


