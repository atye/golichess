# UserActivityScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Win** | **int32** |  | 
**Loss** | **int32** |  | 
**Draw** | **int32** |  | 
**Rp** | [**UserActivityScoreRp**](UserActivityScoreRp.md) |  | 

## Methods

### NewUserActivityScore

`func NewUserActivityScore(win int32, loss int32, draw int32, rp UserActivityScoreRp, ) *UserActivityScore`

NewUserActivityScore instantiates a new UserActivityScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityScoreWithDefaults

`func NewUserActivityScoreWithDefaults() *UserActivityScore`

NewUserActivityScoreWithDefaults instantiates a new UserActivityScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWin

`func (o *UserActivityScore) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *UserActivityScore) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *UserActivityScore) SetWin(v int32)`

SetWin sets Win field to given value.


### GetLoss

`func (o *UserActivityScore) GetLoss() int32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *UserActivityScore) GetLossOk() (*int32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *UserActivityScore) SetLoss(v int32)`

SetLoss sets Loss field to given value.


### GetDraw

`func (o *UserActivityScore) GetDraw() int32`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *UserActivityScore) GetDrawOk() (*int32, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *UserActivityScore) SetDraw(v int32)`

SetDraw sets Draw field to given value.


### GetRp

`func (o *UserActivityScore) GetRp() UserActivityScoreRp`

GetRp returns the Rp field if non-nil, zero value otherwise.

### GetRpOk

`func (o *UserActivityScore) GetRpOk() (*UserActivityScoreRp, bool)`

GetRpOk returns a tuple with the Rp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRp

`func (o *UserActivityScore) SetRp(v UserActivityScoreRp)`

SetRp sets Rp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


