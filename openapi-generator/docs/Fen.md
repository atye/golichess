# Fen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fen** | **string** | The X-FEN of the current position | 
**Lm** | **string** | The last move in UCI format (King to rook for Chess960-compatible castling notation)  | 
**Wc** | **int32** | White&#39;s clock in seconds | 
**Bc** | **int32** | Black&#39;s clock in seconds | 

## Methods

### NewFen

`func NewFen(fen string, lm string, wc int32, bc int32, ) *Fen`

NewFen instantiates a new Fen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFenWithDefaults

`func NewFenWithDefaults() *Fen`

NewFenWithDefaults instantiates a new Fen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFen

`func (o *Fen) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *Fen) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *Fen) SetFen(v string)`

SetFen sets Fen field to given value.


### GetLm

`func (o *Fen) GetLm() string`

GetLm returns the Lm field if non-nil, zero value otherwise.

### GetLmOk

`func (o *Fen) GetLmOk() (*string, bool)`

GetLmOk returns a tuple with the Lm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLm

`func (o *Fen) SetLm(v string)`

SetLm sets Lm field to given value.


### GetWc

`func (o *Fen) GetWc() int32`

GetWc returns the Wc field if non-nil, zero value otherwise.

### GetWcOk

`func (o *Fen) GetWcOk() (*int32, bool)`

GetWcOk returns a tuple with the Wc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWc

`func (o *Fen) SetWc(v int32)`

SetWc sets Wc field to given value.


### GetBc

`func (o *Fen) GetBc() int32`

GetBc returns the Bc field if non-nil, zero value otherwise.

### GetBcOk

`func (o *Fen) GetBcOk() (*int32, bool)`

GetBcOk returns a tuple with the Bc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBc

`func (o *Fen) SetBc(v int32)`

SetBc sets Bc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


