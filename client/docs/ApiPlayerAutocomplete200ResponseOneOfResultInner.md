# ApiPlayerAutocomplete200ResponseOneOfResultInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiPlayerAutocomplete200ResponseOneOfResultInner

`func NewApiPlayerAutocomplete200ResponseOneOfResultInner(id string, name string, ) *ApiPlayerAutocomplete200ResponseOneOfResultInner`

NewApiPlayerAutocomplete200ResponseOneOfResultInner instantiates a new ApiPlayerAutocomplete200ResponseOneOfResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlayerAutocomplete200ResponseOneOfResultInnerWithDefaults

`func NewApiPlayerAutocomplete200ResponseOneOfResultInnerWithDefaults() *ApiPlayerAutocomplete200ResponseOneOfResultInner`

NewApiPlayerAutocomplete200ResponseOneOfResultInnerWithDefaults instantiates a new ApiPlayerAutocomplete200ResponseOneOfResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetOnline

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ApiPlayerAutocomplete200ResponseOneOfResultInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


