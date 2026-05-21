# PuzzleDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**Global** | [**PuzzlePerformance**](PuzzlePerformance.md) |  | 
**Themes** | [**map[string]PuzzleDashboardThemesValue**](PuzzleDashboardThemesValue.md) |  | 

## Methods

### NewPuzzleDashboard

`func NewPuzzleDashboard(days int32, global PuzzlePerformance, themes map[string]PuzzleDashboardThemesValue, ) *PuzzleDashboard`

NewPuzzleDashboard instantiates a new PuzzleDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleDashboardWithDefaults

`func NewPuzzleDashboardWithDefaults() *PuzzleDashboard`

NewPuzzleDashboardWithDefaults instantiates a new PuzzleDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *PuzzleDashboard) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PuzzleDashboard) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PuzzleDashboard) SetDays(v int32)`

SetDays sets Days field to given value.


### GetGlobal

`func (o *PuzzleDashboard) GetGlobal() PuzzlePerformance`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *PuzzleDashboard) GetGlobalOk() (*PuzzlePerformance, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *PuzzleDashboard) SetGlobal(v PuzzlePerformance)`

SetGlobal sets Global field to given value.


### GetThemes

`func (o *PuzzleDashboard) GetThemes() map[string]PuzzleDashboardThemesValue`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *PuzzleDashboard) GetThemesOk() (*map[string]PuzzleDashboardThemesValue, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *PuzzleDashboard) SetThemes(v map[string]PuzzleDashboardThemesValue)`

SetThemes sets Themes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


