# ApiPuzzleDashboard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**Global** | [**ApiPuzzleDashboard200ResponseGlobal**](ApiPuzzleDashboard200ResponseGlobal.md) |  | 
**Themes** | [**map[string]ApiPuzzleDashboard200ResponseThemesValue**](ApiPuzzleDashboard200ResponseThemesValue.md) |  | 

## Methods

### NewApiPuzzleDashboard200Response

`func NewApiPuzzleDashboard200Response(days int32, global ApiPuzzleDashboard200ResponseGlobal, themes map[string]ApiPuzzleDashboard200ResponseThemesValue, ) *ApiPuzzleDashboard200Response`

NewApiPuzzleDashboard200Response instantiates a new ApiPuzzleDashboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleDashboard200ResponseWithDefaults

`func NewApiPuzzleDashboard200ResponseWithDefaults() *ApiPuzzleDashboard200Response`

NewApiPuzzleDashboard200ResponseWithDefaults instantiates a new ApiPuzzleDashboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ApiPuzzleDashboard200Response) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ApiPuzzleDashboard200Response) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ApiPuzzleDashboard200Response) SetDays(v int32)`

SetDays sets Days field to given value.


### GetGlobal

`func (o *ApiPuzzleDashboard200Response) GetGlobal() ApiPuzzleDashboard200ResponseGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ApiPuzzleDashboard200Response) GetGlobalOk() (*ApiPuzzleDashboard200ResponseGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ApiPuzzleDashboard200Response) SetGlobal(v ApiPuzzleDashboard200ResponseGlobal)`

SetGlobal sets Global field to given value.


### GetThemes

`func (o *ApiPuzzleDashboard200Response) GetThemes() map[string]ApiPuzzleDashboard200ResponseThemesValue`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *ApiPuzzleDashboard200Response) GetThemesOk() (*map[string]ApiPuzzleDashboard200ResponseThemesValue, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *ApiPuzzleDashboard200Response) SetThemes(v map[string]ApiPuzzleDashboard200ResponseThemesValue)`

SetThemes sets Themes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


