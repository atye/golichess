# TablebaseStandard200ResponseMovesInner

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

### NewTablebaseStandard200ResponseMovesInner

`func NewTablebaseStandard200ResponseMovesInner(uci string, san string, category string, ) *TablebaseStandard200ResponseMovesInner`

NewTablebaseStandard200ResponseMovesInner instantiates a new TablebaseStandard200ResponseMovesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablebaseStandard200ResponseMovesInnerWithDefaults

`func NewTablebaseStandard200ResponseMovesInnerWithDefaults() *TablebaseStandard200ResponseMovesInner`

NewTablebaseStandard200ResponseMovesInnerWithDefaults instantiates a new TablebaseStandard200ResponseMovesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *TablebaseStandard200ResponseMovesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *TablebaseStandard200ResponseMovesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *TablebaseStandard200ResponseMovesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *TablebaseStandard200ResponseMovesInner) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *TablebaseStandard200ResponseMovesInner) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *TablebaseStandard200ResponseMovesInner) SetSan(v string)`

SetSan sets San field to given value.


### GetCategory

`func (o *TablebaseStandard200ResponseMovesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TablebaseStandard200ResponseMovesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TablebaseStandard200ResponseMovesInner) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDtz

`func (o *TablebaseStandard200ResponseMovesInner) GetDtz() int32`

GetDtz returns the Dtz field if non-nil, zero value otherwise.

### GetDtzOk

`func (o *TablebaseStandard200ResponseMovesInner) GetDtzOk() (*int32, bool)`

GetDtzOk returns a tuple with the Dtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtz

`func (o *TablebaseStandard200ResponseMovesInner) SetDtz(v int32)`

SetDtz sets Dtz field to given value.

### HasDtz

`func (o *TablebaseStandard200ResponseMovesInner) HasDtz() bool`

HasDtz returns a boolean if a field has been set.

### SetDtzNil

`func (o *TablebaseStandard200ResponseMovesInner) SetDtzNil(b bool)`

 SetDtzNil sets the value for Dtz to be an explicit nil

### UnsetDtz
`func (o *TablebaseStandard200ResponseMovesInner) UnsetDtz()`

UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
### GetPreciseDtz

`func (o *TablebaseStandard200ResponseMovesInner) GetPreciseDtz() int32`

GetPreciseDtz returns the PreciseDtz field if non-nil, zero value otherwise.

### GetPreciseDtzOk

`func (o *TablebaseStandard200ResponseMovesInner) GetPreciseDtzOk() (*int32, bool)`

GetPreciseDtzOk returns a tuple with the PreciseDtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseDtz

`func (o *TablebaseStandard200ResponseMovesInner) SetPreciseDtz(v int32)`

SetPreciseDtz sets PreciseDtz field to given value.

### HasPreciseDtz

`func (o *TablebaseStandard200ResponseMovesInner) HasPreciseDtz() bool`

HasPreciseDtz returns a boolean if a field has been set.

### SetPreciseDtzNil

`func (o *TablebaseStandard200ResponseMovesInner) SetPreciseDtzNil(b bool)`

 SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil

### UnsetPreciseDtz
`func (o *TablebaseStandard200ResponseMovesInner) UnsetPreciseDtz()`

UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
### GetDtc

`func (o *TablebaseStandard200ResponseMovesInner) GetDtc() int32`

GetDtc returns the Dtc field if non-nil, zero value otherwise.

### GetDtcOk

`func (o *TablebaseStandard200ResponseMovesInner) GetDtcOk() (*int32, bool)`

GetDtcOk returns a tuple with the Dtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtc

`func (o *TablebaseStandard200ResponseMovesInner) SetDtc(v int32)`

SetDtc sets Dtc field to given value.

### HasDtc

`func (o *TablebaseStandard200ResponseMovesInner) HasDtc() bool`

HasDtc returns a boolean if a field has been set.

### SetDtcNil

`func (o *TablebaseStandard200ResponseMovesInner) SetDtcNil(b bool)`

 SetDtcNil sets the value for Dtc to be an explicit nil

### UnsetDtc
`func (o *TablebaseStandard200ResponseMovesInner) UnsetDtc()`

UnsetDtc ensures that no value is present for Dtc, not even an explicit nil
### GetDtm

`func (o *TablebaseStandard200ResponseMovesInner) GetDtm() int32`

GetDtm returns the Dtm field if non-nil, zero value otherwise.

### GetDtmOk

`func (o *TablebaseStandard200ResponseMovesInner) GetDtmOk() (*int32, bool)`

GetDtmOk returns a tuple with the Dtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtm

`func (o *TablebaseStandard200ResponseMovesInner) SetDtm(v int32)`

SetDtm sets Dtm field to given value.

### HasDtm

`func (o *TablebaseStandard200ResponseMovesInner) HasDtm() bool`

HasDtm returns a boolean if a field has been set.

### SetDtmNil

`func (o *TablebaseStandard200ResponseMovesInner) SetDtmNil(b bool)`

 SetDtmNil sets the value for Dtm to be an explicit nil

### UnsetDtm
`func (o *TablebaseStandard200ResponseMovesInner) UnsetDtm()`

UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
### GetDtw

`func (o *TablebaseStandard200ResponseMovesInner) GetDtw() int32`

GetDtw returns the Dtw field if non-nil, zero value otherwise.

### GetDtwOk

`func (o *TablebaseStandard200ResponseMovesInner) GetDtwOk() (*int32, bool)`

GetDtwOk returns a tuple with the Dtw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtw

`func (o *TablebaseStandard200ResponseMovesInner) SetDtw(v int32)`

SetDtw sets Dtw field to given value.

### HasDtw

`func (o *TablebaseStandard200ResponseMovesInner) HasDtw() bool`

HasDtw returns a boolean if a field has been set.

### SetDtwNil

`func (o *TablebaseStandard200ResponseMovesInner) SetDtwNil(b bool)`

 SetDtwNil sets the value for Dtw to be an explicit nil

### UnsetDtw
`func (o *TablebaseStandard200ResponseMovesInner) UnsetDtw()`

UnsetDtw ensures that no value is present for Dtw, not even an explicit nil
### GetZeroing

`func (o *TablebaseStandard200ResponseMovesInner) GetZeroing() bool`

GetZeroing returns the Zeroing field if non-nil, zero value otherwise.

### GetZeroingOk

`func (o *TablebaseStandard200ResponseMovesInner) GetZeroingOk() (*bool, bool)`

GetZeroingOk returns a tuple with the Zeroing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroing

`func (o *TablebaseStandard200ResponseMovesInner) SetZeroing(v bool)`

SetZeroing sets Zeroing field to given value.

### HasZeroing

`func (o *TablebaseStandard200ResponseMovesInner) HasZeroing() bool`

HasZeroing returns a boolean if a field has been set.

### GetConversion

`func (o *TablebaseStandard200ResponseMovesInner) GetConversion() bool`

GetConversion returns the Conversion field if non-nil, zero value otherwise.

### GetConversionOk

`func (o *TablebaseStandard200ResponseMovesInner) GetConversionOk() (*bool, bool)`

GetConversionOk returns a tuple with the Conversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversion

`func (o *TablebaseStandard200ResponseMovesInner) SetConversion(v bool)`

SetConversion sets Conversion field to given value.

### HasConversion

`func (o *TablebaseStandard200ResponseMovesInner) HasConversion() bool`

HasConversion returns a boolean if a field has been set.

### GetCheckmate

`func (o *TablebaseStandard200ResponseMovesInner) GetCheckmate() bool`

GetCheckmate returns the Checkmate field if non-nil, zero value otherwise.

### GetCheckmateOk

`func (o *TablebaseStandard200ResponseMovesInner) GetCheckmateOk() (*bool, bool)`

GetCheckmateOk returns a tuple with the Checkmate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmate

`func (o *TablebaseStandard200ResponseMovesInner) SetCheckmate(v bool)`

SetCheckmate sets Checkmate field to given value.

### HasCheckmate

`func (o *TablebaseStandard200ResponseMovesInner) HasCheckmate() bool`

HasCheckmate returns a boolean if a field has been set.

### GetStalemate

`func (o *TablebaseStandard200ResponseMovesInner) GetStalemate() bool`

GetStalemate returns the Stalemate field if non-nil, zero value otherwise.

### GetStalemateOk

`func (o *TablebaseStandard200ResponseMovesInner) GetStalemateOk() (*bool, bool)`

GetStalemateOk returns a tuple with the Stalemate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalemate

`func (o *TablebaseStandard200ResponseMovesInner) SetStalemate(v bool)`

SetStalemate sets Stalemate field to given value.

### HasStalemate

`func (o *TablebaseStandard200ResponseMovesInner) HasStalemate() bool`

HasStalemate returns a boolean if a field has been set.

### GetVariantWin

`func (o *TablebaseStandard200ResponseMovesInner) GetVariantWin() bool`

GetVariantWin returns the VariantWin field if non-nil, zero value otherwise.

### GetVariantWinOk

`func (o *TablebaseStandard200ResponseMovesInner) GetVariantWinOk() (*bool, bool)`

GetVariantWinOk returns a tuple with the VariantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWin

`func (o *TablebaseStandard200ResponseMovesInner) SetVariantWin(v bool)`

SetVariantWin sets VariantWin field to given value.

### HasVariantWin

`func (o *TablebaseStandard200ResponseMovesInner) HasVariantWin() bool`

HasVariantWin returns a boolean if a field has been set.

### GetVariantLoss

`func (o *TablebaseStandard200ResponseMovesInner) GetVariantLoss() bool`

GetVariantLoss returns the VariantLoss field if non-nil, zero value otherwise.

### GetVariantLossOk

`func (o *TablebaseStandard200ResponseMovesInner) GetVariantLossOk() (*bool, bool)`

GetVariantLossOk returns a tuple with the VariantLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLoss

`func (o *TablebaseStandard200ResponseMovesInner) SetVariantLoss(v bool)`

SetVariantLoss sets VariantLoss field to given value.

### HasVariantLoss

`func (o *TablebaseStandard200ResponseMovesInner) HasVariantLoss() bool`

HasVariantLoss returns a boolean if a field has been set.

### GetInsufficientMaterial

`func (o *TablebaseStandard200ResponseMovesInner) GetInsufficientMaterial() bool`

GetInsufficientMaterial returns the InsufficientMaterial field if non-nil, zero value otherwise.

### GetInsufficientMaterialOk

`func (o *TablebaseStandard200ResponseMovesInner) GetInsufficientMaterialOk() (*bool, bool)`

GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientMaterial

`func (o *TablebaseStandard200ResponseMovesInner) SetInsufficientMaterial(v bool)`

SetInsufficientMaterial sets InsufficientMaterial field to given value.

### HasInsufficientMaterial

`func (o *TablebaseStandard200ResponseMovesInner) HasInsufficientMaterial() bool`

HasInsufficientMaterial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


