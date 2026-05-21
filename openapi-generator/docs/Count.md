# Count

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **int32** |  | 
**Rated** | **int32** |  | 
**Ai** | Pointer to **int32** |  | [optional] 
**Draw** | **int32** |  | 
**DrawH** | Pointer to **int32** |  | [optional] 
**Loss** | **int32** |  | 
**LossH** | Pointer to **int32** |  | [optional] 
**Win** | **int32** |  | 
**WinH** | Pointer to **int32** |  | [optional] 
**Bookmark** | **int32** |  | 
**Playing** | **int32** |  | 
**Import** | **int32** |  | 
**Me** | **int32** |  | 

## Methods

### NewCount

`func NewCount(all int32, rated int32, draw int32, loss int32, win int32, bookmark int32, playing int32, import_ int32, me int32, ) *Count`

NewCount instantiates a new Count object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountWithDefaults

`func NewCountWithDefaults() *Count`

NewCountWithDefaults instantiates a new Count object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *Count) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *Count) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *Count) SetAll(v int32)`

SetAll sets All field to given value.


### GetRated

`func (o *Count) GetRated() int32`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *Count) GetRatedOk() (*int32, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *Count) SetRated(v int32)`

SetRated sets Rated field to given value.


### GetAi

`func (o *Count) GetAi() int32`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *Count) GetAiOk() (*int32, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *Count) SetAi(v int32)`

SetAi sets Ai field to given value.

### HasAi

`func (o *Count) HasAi() bool`

HasAi returns a boolean if a field has been set.

### GetDraw

`func (o *Count) GetDraw() int32`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *Count) GetDrawOk() (*int32, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *Count) SetDraw(v int32)`

SetDraw sets Draw field to given value.


### GetDrawH

`func (o *Count) GetDrawH() int32`

GetDrawH returns the DrawH field if non-nil, zero value otherwise.

### GetDrawHOk

`func (o *Count) GetDrawHOk() (*int32, bool)`

GetDrawHOk returns a tuple with the DrawH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawH

`func (o *Count) SetDrawH(v int32)`

SetDrawH sets DrawH field to given value.

### HasDrawH

`func (o *Count) HasDrawH() bool`

HasDrawH returns a boolean if a field has been set.

### GetLoss

`func (o *Count) GetLoss() int32`

GetLoss returns the Loss field if non-nil, zero value otherwise.

### GetLossOk

`func (o *Count) GetLossOk() (*int32, bool)`

GetLossOk returns a tuple with the Loss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoss

`func (o *Count) SetLoss(v int32)`

SetLoss sets Loss field to given value.


### GetLossH

`func (o *Count) GetLossH() int32`

GetLossH returns the LossH field if non-nil, zero value otherwise.

### GetLossHOk

`func (o *Count) GetLossHOk() (*int32, bool)`

GetLossHOk returns a tuple with the LossH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossH

`func (o *Count) SetLossH(v int32)`

SetLossH sets LossH field to given value.

### HasLossH

`func (o *Count) HasLossH() bool`

HasLossH returns a boolean if a field has been set.

### GetWin

`func (o *Count) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *Count) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *Count) SetWin(v int32)`

SetWin sets Win field to given value.


### GetWinH

`func (o *Count) GetWinH() int32`

GetWinH returns the WinH field if non-nil, zero value otherwise.

### GetWinHOk

`func (o *Count) GetWinHOk() (*int32, bool)`

GetWinHOk returns a tuple with the WinH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinH

`func (o *Count) SetWinH(v int32)`

SetWinH sets WinH field to given value.

### HasWinH

`func (o *Count) HasWinH() bool`

HasWinH returns a boolean if a field has been set.

### GetBookmark

`func (o *Count) GetBookmark() int32`

GetBookmark returns the Bookmark field if non-nil, zero value otherwise.

### GetBookmarkOk

`func (o *Count) GetBookmarkOk() (*int32, bool)`

GetBookmarkOk returns a tuple with the Bookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmark

`func (o *Count) SetBookmark(v int32)`

SetBookmark sets Bookmark field to given value.


### GetPlaying

`func (o *Count) GetPlaying() int32`

GetPlaying returns the Playing field if non-nil, zero value otherwise.

### GetPlayingOk

`func (o *Count) GetPlayingOk() (*int32, bool)`

GetPlayingOk returns a tuple with the Playing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaying

`func (o *Count) SetPlaying(v int32)`

SetPlaying sets Playing field to given value.


### GetImport

`func (o *Count) GetImport() int32`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *Count) GetImportOk() (*int32, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *Count) SetImport(v int32)`

SetImport sets Import field to given value.


### GetMe

`func (o *Count) GetMe() int32`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *Count) GetMeOk() (*int32, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMe

`func (o *Count) SetMe(v int32)`

SetMe sets Me field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


