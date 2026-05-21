# TablebaseMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**San** | **string** |  | 
**Category** | **string** |  | 
**Dtz** | Pointer to **NullableInt32** |  | [optional] 
**PreciseDtz** | Pointer to **NullableInt32** |  | [optional] 
**Dtc** | Pointer to **NullableInt32** |  | [optional] 
**Dtm** | Pointer to **NullableInt32** |  | [optional] 
**Dtw** | Pointer to **NullableInt32** |  | [optional] 
**Zeroing** | Pointer to **bool** |  | [optional] 
**Conversion** | Pointer to **bool** |  | [optional] 
**Checkmate** | Pointer to **bool** |  | [optional] 
**Stalemate** | Pointer to **bool** |  | [optional] 
**VariantWin** | Pointer to **bool** |  | [optional] 
**VariantLoss** | Pointer to **bool** |  | [optional] 
**InsufficientMaterial** | Pointer to **bool** |  | [optional] 

## Methods

### NewTablebaseMove

`func NewTablebaseMove(uci string, san string, category string, ) *TablebaseMove`

NewTablebaseMove instantiates a new TablebaseMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablebaseMoveWithDefaults

`func NewTablebaseMoveWithDefaults() *TablebaseMove`

NewTablebaseMoveWithDefaults instantiates a new TablebaseMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *TablebaseMove) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *TablebaseMove) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *TablebaseMove) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *TablebaseMove) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *TablebaseMove) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *TablebaseMove) SetSan(v string)`

SetSan sets San field to given value.


### GetCategory

`func (o *TablebaseMove) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TablebaseMove) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TablebaseMove) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDtz

`func (o *TablebaseMove) GetDtz() int32`

GetDtz returns the Dtz field if non-nil, zero value otherwise.

### GetDtzOk

`func (o *TablebaseMove) GetDtzOk() (*int32, bool)`

GetDtzOk returns a tuple with the Dtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtz

`func (o *TablebaseMove) SetDtz(v int32)`

SetDtz sets Dtz field to given value.

### HasDtz

`func (o *TablebaseMove) HasDtz() bool`

HasDtz returns a boolean if a field has been set.

### SetDtzNil

`func (o *TablebaseMove) SetDtzNil(b bool)`

 SetDtzNil sets the value for Dtz to be an explicit nil

### UnsetDtz
`func (o *TablebaseMove) UnsetDtz()`

UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
### GetPreciseDtz

`func (o *TablebaseMove) GetPreciseDtz() int32`

GetPreciseDtz returns the PreciseDtz field if non-nil, zero value otherwise.

### GetPreciseDtzOk

`func (o *TablebaseMove) GetPreciseDtzOk() (*int32, bool)`

GetPreciseDtzOk returns a tuple with the PreciseDtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseDtz

`func (o *TablebaseMove) SetPreciseDtz(v int32)`

SetPreciseDtz sets PreciseDtz field to given value.

### HasPreciseDtz

`func (o *TablebaseMove) HasPreciseDtz() bool`

HasPreciseDtz returns a boolean if a field has been set.

### SetPreciseDtzNil

`func (o *TablebaseMove) SetPreciseDtzNil(b bool)`

 SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil

### UnsetPreciseDtz
`func (o *TablebaseMove) UnsetPreciseDtz()`

UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
### GetDtc

`func (o *TablebaseMove) GetDtc() int32`

GetDtc returns the Dtc field if non-nil, zero value otherwise.

### GetDtcOk

`func (o *TablebaseMove) GetDtcOk() (*int32, bool)`

GetDtcOk returns a tuple with the Dtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtc

`func (o *TablebaseMove) SetDtc(v int32)`

SetDtc sets Dtc field to given value.

### HasDtc

`func (o *TablebaseMove) HasDtc() bool`

HasDtc returns a boolean if a field has been set.

### SetDtcNil

`func (o *TablebaseMove) SetDtcNil(b bool)`

 SetDtcNil sets the value for Dtc to be an explicit nil

### UnsetDtc
`func (o *TablebaseMove) UnsetDtc()`

UnsetDtc ensures that no value is present for Dtc, not even an explicit nil
### GetDtm

`func (o *TablebaseMove) GetDtm() int32`

GetDtm returns the Dtm field if non-nil, zero value otherwise.

### GetDtmOk

`func (o *TablebaseMove) GetDtmOk() (*int32, bool)`

GetDtmOk returns a tuple with the Dtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtm

`func (o *TablebaseMove) SetDtm(v int32)`

SetDtm sets Dtm field to given value.

### HasDtm

`func (o *TablebaseMove) HasDtm() bool`

HasDtm returns a boolean if a field has been set.

### SetDtmNil

`func (o *TablebaseMove) SetDtmNil(b bool)`

 SetDtmNil sets the value for Dtm to be an explicit nil

### UnsetDtm
`func (o *TablebaseMove) UnsetDtm()`

UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
### GetDtw

`func (o *TablebaseMove) GetDtw() int32`

GetDtw returns the Dtw field if non-nil, zero value otherwise.

### GetDtwOk

`func (o *TablebaseMove) GetDtwOk() (*int32, bool)`

GetDtwOk returns a tuple with the Dtw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtw

`func (o *TablebaseMove) SetDtw(v int32)`

SetDtw sets Dtw field to given value.

### HasDtw

`func (o *TablebaseMove) HasDtw() bool`

HasDtw returns a boolean if a field has been set.

### SetDtwNil

`func (o *TablebaseMove) SetDtwNil(b bool)`

 SetDtwNil sets the value for Dtw to be an explicit nil

### UnsetDtw
`func (o *TablebaseMove) UnsetDtw()`

UnsetDtw ensures that no value is present for Dtw, not even an explicit nil
### GetZeroing

`func (o *TablebaseMove) GetZeroing() bool`

GetZeroing returns the Zeroing field if non-nil, zero value otherwise.

### GetZeroingOk

`func (o *TablebaseMove) GetZeroingOk() (*bool, bool)`

GetZeroingOk returns a tuple with the Zeroing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroing

`func (o *TablebaseMove) SetZeroing(v bool)`

SetZeroing sets Zeroing field to given value.

### HasZeroing

`func (o *TablebaseMove) HasZeroing() bool`

HasZeroing returns a boolean if a field has been set.

### GetConversion

`func (o *TablebaseMove) GetConversion() bool`

GetConversion returns the Conversion field if non-nil, zero value otherwise.

### GetConversionOk

`func (o *TablebaseMove) GetConversionOk() (*bool, bool)`

GetConversionOk returns a tuple with the Conversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversion

`func (o *TablebaseMove) SetConversion(v bool)`

SetConversion sets Conversion field to given value.

### HasConversion

`func (o *TablebaseMove) HasConversion() bool`

HasConversion returns a boolean if a field has been set.

### GetCheckmate

`func (o *TablebaseMove) GetCheckmate() bool`

GetCheckmate returns the Checkmate field if non-nil, zero value otherwise.

### GetCheckmateOk

`func (o *TablebaseMove) GetCheckmateOk() (*bool, bool)`

GetCheckmateOk returns a tuple with the Checkmate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmate

`func (o *TablebaseMove) SetCheckmate(v bool)`

SetCheckmate sets Checkmate field to given value.

### HasCheckmate

`func (o *TablebaseMove) HasCheckmate() bool`

HasCheckmate returns a boolean if a field has been set.

### GetStalemate

`func (o *TablebaseMove) GetStalemate() bool`

GetStalemate returns the Stalemate field if non-nil, zero value otherwise.

### GetStalemateOk

`func (o *TablebaseMove) GetStalemateOk() (*bool, bool)`

GetStalemateOk returns a tuple with the Stalemate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalemate

`func (o *TablebaseMove) SetStalemate(v bool)`

SetStalemate sets Stalemate field to given value.

### HasStalemate

`func (o *TablebaseMove) HasStalemate() bool`

HasStalemate returns a boolean if a field has been set.

### GetVariantWin

`func (o *TablebaseMove) GetVariantWin() bool`

GetVariantWin returns the VariantWin field if non-nil, zero value otherwise.

### GetVariantWinOk

`func (o *TablebaseMove) GetVariantWinOk() (*bool, bool)`

GetVariantWinOk returns a tuple with the VariantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWin

`func (o *TablebaseMove) SetVariantWin(v bool)`

SetVariantWin sets VariantWin field to given value.

### HasVariantWin

`func (o *TablebaseMove) HasVariantWin() bool`

HasVariantWin returns a boolean if a field has been set.

### GetVariantLoss

`func (o *TablebaseMove) GetVariantLoss() bool`

GetVariantLoss returns the VariantLoss field if non-nil, zero value otherwise.

### GetVariantLossOk

`func (o *TablebaseMove) GetVariantLossOk() (*bool, bool)`

GetVariantLossOk returns a tuple with the VariantLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLoss

`func (o *TablebaseMove) SetVariantLoss(v bool)`

SetVariantLoss sets VariantLoss field to given value.

### HasVariantLoss

`func (o *TablebaseMove) HasVariantLoss() bool`

HasVariantLoss returns a boolean if a field has been set.

### GetInsufficientMaterial

`func (o *TablebaseMove) GetInsufficientMaterial() bool`

GetInsufficientMaterial returns the InsufficientMaterial field if non-nil, zero value otherwise.

### GetInsufficientMaterialOk

`func (o *TablebaseMove) GetInsufficientMaterialOk() (*bool, bool)`

GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientMaterial

`func (o *TablebaseMove) SetInsufficientMaterial(v bool)`

SetInsufficientMaterial sets InsufficientMaterial field to given value.

### HasInsufficientMaterial

`func (o *TablebaseMove) HasInsufficientMaterial() bool`

HasInsufficientMaterial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


