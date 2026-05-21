# UserActivityTeamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 

## Methods

### NewUserActivityTeamsInner

`func NewUserActivityTeamsInner(url string, name string, ) *UserActivityTeamsInner`

NewUserActivityTeamsInner instantiates a new UserActivityTeamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityTeamsInnerWithDefaults

`func NewUserActivityTeamsInnerWithDefaults() *UserActivityTeamsInner`

NewUserActivityTeamsInnerWithDefaults instantiates a new UserActivityTeamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserActivityTeamsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserActivityTeamsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserActivityTeamsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *UserActivityTeamsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserActivityTeamsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserActivityTeamsInner) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *UserActivityTeamsInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *UserActivityTeamsInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *UserActivityTeamsInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *UserActivityTeamsInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


