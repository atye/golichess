# ApiSimul200ResponsePendingInnerHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiSimul200ResponsePendingInnerHost

`func NewApiSimul200ResponsePendingInnerHost(id string, name string, ) *ApiSimul200ResponsePendingInnerHost`

NewApiSimul200ResponsePendingInnerHost instantiates a new ApiSimul200ResponsePendingInnerHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSimul200ResponsePendingInnerHostWithDefaults

`func NewApiSimul200ResponsePendingInnerHostWithDefaults() *ApiSimul200ResponsePendingInnerHost`

NewApiSimul200ResponsePendingInnerHostWithDefaults instantiates a new ApiSimul200ResponsePendingInnerHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiSimul200ResponsePendingInnerHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSimul200ResponsePendingInnerHost) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiSimul200ResponsePendingInnerHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSimul200ResponsePendingInnerHost) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *ApiSimul200ResponsePendingInnerHost) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiSimul200ResponsePendingInnerHost) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiSimul200ResponsePendingInnerHost) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *ApiSimul200ResponsePendingInnerHost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiSimul200ResponsePendingInnerHost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiSimul200ResponsePendingInnerHost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *ApiSimul200ResponsePendingInnerHost) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiSimul200ResponsePendingInnerHost) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiSimul200ResponsePendingInnerHost) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiSimul200ResponsePendingInnerHost) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiSimul200ResponsePendingInnerHost) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiSimul200ResponsePendingInnerHost) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetRating

`func (o *ApiSimul200ResponsePendingInnerHost) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiSimul200ResponsePendingInnerHost) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiSimul200ResponsePendingInnerHost) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetProvisional

`func (o *ApiSimul200ResponsePendingInnerHost) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *ApiSimul200ResponsePendingInnerHost) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *ApiSimul200ResponsePendingInnerHost) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetGameId

`func (o *ApiSimul200ResponsePendingInnerHost) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *ApiSimul200ResponsePendingInnerHost) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *ApiSimul200ResponsePendingInnerHost) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetOnline

`func (o *ApiSimul200ResponsePendingInnerHost) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApiSimul200ResponsePendingInnerHost) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApiSimul200ResponsePendingInnerHost) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ApiSimul200ResponsePendingInnerHost) HasOnline() bool`

HasOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


