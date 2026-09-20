# TeamIdUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JoinedTeamAt** | Pointer to **int64** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamIdUsers200Response

`func NewTeamIdUsers200Response(id string, name string, ) *TeamIdUsers200Response`

NewTeamIdUsers200Response instantiates a new TeamIdUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamIdUsers200ResponseWithDefaults

`func NewTeamIdUsers200ResponseWithDefaults() *TeamIdUsers200Response`

NewTeamIdUsers200ResponseWithDefaults instantiates a new TeamIdUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoinedTeamAt

`func (o *TeamIdUsers200Response) GetJoinedTeamAt() int64`

GetJoinedTeamAt returns the JoinedTeamAt field if non-nil, zero value otherwise.

### GetJoinedTeamAtOk

`func (o *TeamIdUsers200Response) GetJoinedTeamAtOk() (*int64, bool)`

GetJoinedTeamAtOk returns a tuple with the JoinedTeamAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedTeamAt

`func (o *TeamIdUsers200Response) SetJoinedTeamAt(v int64)`

SetJoinedTeamAt sets JoinedTeamAt field to given value.

### HasJoinedTeamAt

`func (o *TeamIdUsers200Response) HasJoinedTeamAt() bool`

HasJoinedTeamAt returns a boolean if a field has been set.

### GetId

`func (o *TeamIdUsers200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamIdUsers200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamIdUsers200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TeamIdUsers200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamIdUsers200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamIdUsers200Response) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *TeamIdUsers200Response) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *TeamIdUsers200Response) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *TeamIdUsers200Response) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *TeamIdUsers200Response) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *TeamIdUsers200Response) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TeamIdUsers200Response) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TeamIdUsers200Response) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TeamIdUsers200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *TeamIdUsers200Response) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *TeamIdUsers200Response) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *TeamIdUsers200Response) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *TeamIdUsers200Response) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *TeamIdUsers200Response) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *TeamIdUsers200Response) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *TeamIdUsers200Response) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *TeamIdUsers200Response) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetUrl

`func (o *TeamIdUsers200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamIdUsers200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamIdUsers200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TeamIdUsers200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


