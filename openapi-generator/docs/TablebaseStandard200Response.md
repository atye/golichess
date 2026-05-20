# TablebaseStandard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | &#x60;cursed-win&#x60; and &#x60;blessed-loss&#x60; means the 50-move rule prevents the decisive result.  &#x60;syzygy-win&#x60; and &#x60;syzygy-loss&#x60; means exact result is unknown due to [DTZ rounding](https://syzygy-tables.info/metrics#dtz), i.e., the win or loss could also be prevented by the 50-move rule if the user has deviated from the tablebase recommendation since the last pawn move or capture.  &#x60;maybe-win&#x60; and &#x60;maybe-loss&#x60; means the result with regard to the 50-move rule is unknown, because DTZ is unknown and the DTC tablebase does not guarantee to reach a zeroing move as soon as possible.  | 
**Dtz** | Pointer to **NullableInt32** | [DTZ50&#39;&#39; with rounding](https://syzygy-tables.info/metrics#dtz) in plies (for Standard chess positions with not more than 7 pieces and variant positions not more than 6 pieces)  | [optional] 
**PreciseDtz** | Pointer to **NullableInt32** | DTZ50&#39;&#39; in plies, only if guaranteed to not be rounded, or absent if unknown  | [optional] 
**Dtc** | Pointer to **NullableInt32** | Depth to Conversion: Moves to next capture, promotion, or checkmate. Available for: * Standard chess positions with 8 pieces, more than one pawn of material   value for each side, and at least one pair of opposing pawns,   short *op1*, if query parameter &#x60;dtc&#x60; is &#x60;auxiliary&#x60; or &#x60;always&#x60;. * Some standard chess positions with up to 7 pieces, if query parameter   &#x60;dtc&#x60; is &#x60;always&#x60;. Work in progress.  | [optional] 
**Dtm** | Pointer to **NullableInt32** | Depth To Mate: Plies to mate (available only for Standard positions with not more than 6 pieces)  | [optional] 
**Dtw** | Pointer to **NullableInt32** | Depth To Win: Plies to win (available only for Antichess positions with not more than 4 pieces)  | [optional] 
**Checkmate** | Pointer to **bool** |  | [optional] 
**Stalemate** | Pointer to **bool** |  | [optional] 
**VariantWin** | Pointer to **bool** | Only in chess variants | [optional] 
**VariantLoss** | Pointer to **bool** | Only in chess variants | [optional] 
**InsufficientMaterial** | Pointer to **bool** |  | [optional] 
**Moves** | [**[]TablebaseStandard200ResponseMovesInner**](TablebaseStandard200ResponseMovesInner.md) | Information about legal moves, best first | 

## Methods

### NewTablebaseStandard200Response

`func NewTablebaseStandard200Response(category string, moves []TablebaseStandard200ResponseMovesInner, ) *TablebaseStandard200Response`

NewTablebaseStandard200Response instantiates a new TablebaseStandard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablebaseStandard200ResponseWithDefaults

`func NewTablebaseStandard200ResponseWithDefaults() *TablebaseStandard200Response`

NewTablebaseStandard200ResponseWithDefaults instantiates a new TablebaseStandard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TablebaseStandard200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TablebaseStandard200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TablebaseStandard200Response) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDtz

`func (o *TablebaseStandard200Response) GetDtz() int32`

GetDtz returns the Dtz field if non-nil, zero value otherwise.

### GetDtzOk

`func (o *TablebaseStandard200Response) GetDtzOk() (*int32, bool)`

GetDtzOk returns a tuple with the Dtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtz

`func (o *TablebaseStandard200Response) SetDtz(v int32)`

SetDtz sets Dtz field to given value.

### HasDtz

`func (o *TablebaseStandard200Response) HasDtz() bool`

HasDtz returns a boolean if a field has been set.

### SetDtzNil

`func (o *TablebaseStandard200Response) SetDtzNil(b bool)`

 SetDtzNil sets the value for Dtz to be an explicit nil

### UnsetDtz
`func (o *TablebaseStandard200Response) UnsetDtz()`

UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
### GetPreciseDtz

`func (o *TablebaseStandard200Response) GetPreciseDtz() int32`

GetPreciseDtz returns the PreciseDtz field if non-nil, zero value otherwise.

### GetPreciseDtzOk

`func (o *TablebaseStandard200Response) GetPreciseDtzOk() (*int32, bool)`

GetPreciseDtzOk returns a tuple with the PreciseDtz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseDtz

`func (o *TablebaseStandard200Response) SetPreciseDtz(v int32)`

SetPreciseDtz sets PreciseDtz field to given value.

### HasPreciseDtz

`func (o *TablebaseStandard200Response) HasPreciseDtz() bool`

HasPreciseDtz returns a boolean if a field has been set.

### SetPreciseDtzNil

`func (o *TablebaseStandard200Response) SetPreciseDtzNil(b bool)`

 SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil

### UnsetPreciseDtz
`func (o *TablebaseStandard200Response) UnsetPreciseDtz()`

UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
### GetDtc

`func (o *TablebaseStandard200Response) GetDtc() int32`

GetDtc returns the Dtc field if non-nil, zero value otherwise.

### GetDtcOk

`func (o *TablebaseStandard200Response) GetDtcOk() (*int32, bool)`

GetDtcOk returns a tuple with the Dtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtc

`func (o *TablebaseStandard200Response) SetDtc(v int32)`

SetDtc sets Dtc field to given value.

### HasDtc

`func (o *TablebaseStandard200Response) HasDtc() bool`

HasDtc returns a boolean if a field has been set.

### SetDtcNil

`func (o *TablebaseStandard200Response) SetDtcNil(b bool)`

 SetDtcNil sets the value for Dtc to be an explicit nil

### UnsetDtc
`func (o *TablebaseStandard200Response) UnsetDtc()`

UnsetDtc ensures that no value is present for Dtc, not even an explicit nil
### GetDtm

`func (o *TablebaseStandard200Response) GetDtm() int32`

GetDtm returns the Dtm field if non-nil, zero value otherwise.

### GetDtmOk

`func (o *TablebaseStandard200Response) GetDtmOk() (*int32, bool)`

GetDtmOk returns a tuple with the Dtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtm

`func (o *TablebaseStandard200Response) SetDtm(v int32)`

SetDtm sets Dtm field to given value.

### HasDtm

`func (o *TablebaseStandard200Response) HasDtm() bool`

HasDtm returns a boolean if a field has been set.

### SetDtmNil

`func (o *TablebaseStandard200Response) SetDtmNil(b bool)`

 SetDtmNil sets the value for Dtm to be an explicit nil

### UnsetDtm
`func (o *TablebaseStandard200Response) UnsetDtm()`

UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
### GetDtw

`func (o *TablebaseStandard200Response) GetDtw() int32`

GetDtw returns the Dtw field if non-nil, zero value otherwise.

### GetDtwOk

`func (o *TablebaseStandard200Response) GetDtwOk() (*int32, bool)`

GetDtwOk returns a tuple with the Dtw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtw

`func (o *TablebaseStandard200Response) SetDtw(v int32)`

SetDtw sets Dtw field to given value.

### HasDtw

`func (o *TablebaseStandard200Response) HasDtw() bool`

HasDtw returns a boolean if a field has been set.

### SetDtwNil

`func (o *TablebaseStandard200Response) SetDtwNil(b bool)`

 SetDtwNil sets the value for Dtw to be an explicit nil

### UnsetDtw
`func (o *TablebaseStandard200Response) UnsetDtw()`

UnsetDtw ensures that no value is present for Dtw, not even an explicit nil
### GetCheckmate

`func (o *TablebaseStandard200Response) GetCheckmate() bool`

GetCheckmate returns the Checkmate field if non-nil, zero value otherwise.

### GetCheckmateOk

`func (o *TablebaseStandard200Response) GetCheckmateOk() (*bool, bool)`

GetCheckmateOk returns a tuple with the Checkmate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckmate

`func (o *TablebaseStandard200Response) SetCheckmate(v bool)`

SetCheckmate sets Checkmate field to given value.

### HasCheckmate

`func (o *TablebaseStandard200Response) HasCheckmate() bool`

HasCheckmate returns a boolean if a field has been set.

### GetStalemate

`func (o *TablebaseStandard200Response) GetStalemate() bool`

GetStalemate returns the Stalemate field if non-nil, zero value otherwise.

### GetStalemateOk

`func (o *TablebaseStandard200Response) GetStalemateOk() (*bool, bool)`

GetStalemateOk returns a tuple with the Stalemate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalemate

`func (o *TablebaseStandard200Response) SetStalemate(v bool)`

SetStalemate sets Stalemate field to given value.

### HasStalemate

`func (o *TablebaseStandard200Response) HasStalemate() bool`

HasStalemate returns a boolean if a field has been set.

### GetVariantWin

`func (o *TablebaseStandard200Response) GetVariantWin() bool`

GetVariantWin returns the VariantWin field if non-nil, zero value otherwise.

### GetVariantWinOk

`func (o *TablebaseStandard200Response) GetVariantWinOk() (*bool, bool)`

GetVariantWinOk returns a tuple with the VariantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantWin

`func (o *TablebaseStandard200Response) SetVariantWin(v bool)`

SetVariantWin sets VariantWin field to given value.

### HasVariantWin

`func (o *TablebaseStandard200Response) HasVariantWin() bool`

HasVariantWin returns a boolean if a field has been set.

### GetVariantLoss

`func (o *TablebaseStandard200Response) GetVariantLoss() bool`

GetVariantLoss returns the VariantLoss field if non-nil, zero value otherwise.

### GetVariantLossOk

`func (o *TablebaseStandard200Response) GetVariantLossOk() (*bool, bool)`

GetVariantLossOk returns a tuple with the VariantLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantLoss

`func (o *TablebaseStandard200Response) SetVariantLoss(v bool)`

SetVariantLoss sets VariantLoss field to given value.

### HasVariantLoss

`func (o *TablebaseStandard200Response) HasVariantLoss() bool`

HasVariantLoss returns a boolean if a field has been set.

### GetInsufficientMaterial

`func (o *TablebaseStandard200Response) GetInsufficientMaterial() bool`

GetInsufficientMaterial returns the InsufficientMaterial field if non-nil, zero value otherwise.

### GetInsufficientMaterialOk

`func (o *TablebaseStandard200Response) GetInsufficientMaterialOk() (*bool, bool)`

GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientMaterial

`func (o *TablebaseStandard200Response) SetInsufficientMaterial(v bool)`

SetInsufficientMaterial sets InsufficientMaterial field to given value.

### HasInsufficientMaterial

`func (o *TablebaseStandard200Response) HasInsufficientMaterial() bool`

HasInsufficientMaterial returns a boolean if a field has been set.

### GetMoves

`func (o *TablebaseStandard200Response) GetMoves() []TablebaseStandard200ResponseMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *TablebaseStandard200Response) GetMovesOk() (*[]TablebaseStandard200ResponseMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *TablebaseStandard200Response) SetMoves(v []TablebaseStandard200ResponseMovesInner)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


