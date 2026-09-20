# StudyImportPgnChapters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chapters** | Pointer to [**[]StudyImportPgnChaptersChaptersInner**](StudyImportPgnChaptersChaptersInner.md) |  | [optional] 
**Error** | Pointer to **NullableString** | An error message, if some of the games could not be imported. | [optional] 

## Methods

### NewStudyImportPgnChapters

`func NewStudyImportPgnChapters() *StudyImportPgnChapters`

NewStudyImportPgnChapters instantiates a new StudyImportPgnChapters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudyImportPgnChaptersWithDefaults

`func NewStudyImportPgnChaptersWithDefaults() *StudyImportPgnChapters`

NewStudyImportPgnChaptersWithDefaults instantiates a new StudyImportPgnChapters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChapters

`func (o *StudyImportPgnChapters) GetChapters() []StudyImportPgnChaptersChaptersInner`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *StudyImportPgnChapters) GetChaptersOk() (*[]StudyImportPgnChaptersChaptersInner, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *StudyImportPgnChapters) SetChapters(v []StudyImportPgnChaptersChaptersInner)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *StudyImportPgnChapters) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetError

`func (o *StudyImportPgnChapters) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StudyImportPgnChapters) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StudyImportPgnChapters) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *StudyImportPgnChapters) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *StudyImportPgnChapters) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *StudyImportPgnChapters) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


