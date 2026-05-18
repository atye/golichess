# ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpId

`func NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpId(id string, name string, ) *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpId instantiates a new ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpIdWithDefaults

`func NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpIdWithDefaults() *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

NewApiUserPerf200ResponseStatWorstLossesResultsInnerOpIdWithDefaults instantiates a new ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


