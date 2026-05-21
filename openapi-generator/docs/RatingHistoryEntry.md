# RatingHistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Points** | Pointer to **[][]int32** |  | [optional] 

## Methods

### NewRatingHistoryEntry

`func NewRatingHistoryEntry() *RatingHistoryEntry`

NewRatingHistoryEntry instantiates a new RatingHistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingHistoryEntryWithDefaults

`func NewRatingHistoryEntryWithDefaults() *RatingHistoryEntry`

NewRatingHistoryEntryWithDefaults instantiates a new RatingHistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RatingHistoryEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RatingHistoryEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RatingHistoryEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RatingHistoryEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoints

`func (o *RatingHistoryEntry) GetPoints() [][]int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RatingHistoryEntry) GetPointsOk() (*[][]int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RatingHistoryEntry) SetPoints(v [][]int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RatingHistoryEntry) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


