# ArenaPerf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**PerfType**](PerfType.md) |  | 
**Name** | **string** |  | 
**Position** | **int32** |  | 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewArenaPerf

`func NewArenaPerf(key PerfType, name string, position int32, ) *ArenaPerf`

NewArenaPerf instantiates a new ArenaPerf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaPerfWithDefaults

`func NewArenaPerfWithDefaults() *ArenaPerf`

NewArenaPerfWithDefaults instantiates a new ArenaPerf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ArenaPerf) GetKey() PerfType`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ArenaPerf) GetKeyOk() (*PerfType, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ArenaPerf) SetKey(v PerfType)`

SetKey sets Key field to given value.


### GetName

`func (o *ArenaPerf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArenaPerf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArenaPerf) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *ArenaPerf) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ArenaPerf) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ArenaPerf) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetIcon

`func (o *ArenaPerf) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ArenaPerf) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ArenaPerf) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ArenaPerf) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


