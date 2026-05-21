# TeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **string** |  | 
**UserId** | **string** |  | 
**Date** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamRequest

`func NewTeamRequest(teamId string, userId string, date int32, ) *TeamRequest`

NewTeamRequest instantiates a new TeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRequestWithDefaults

`func NewTeamRequestWithDefaults() *TeamRequest`

NewTeamRequestWithDefaults instantiates a new TeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *TeamRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetUserId

`func (o *TeamRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TeamRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TeamRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDate

`func (o *TeamRequest) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TeamRequest) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TeamRequest) SetDate(v int32)`

SetDate sets Date field to given value.


### GetMessage

`func (o *TeamRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TeamRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TeamRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TeamRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


