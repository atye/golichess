# TeamsByTournament200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Teams** | [**[]TeamsByTournament200ResponseTeamsInner**](TeamsByTournament200ResponseTeamsInner.md) |  | 

## Methods

### NewTeamsByTournament200Response

`func NewTeamsByTournament200Response(id string, teams []TeamsByTournament200ResponseTeamsInner, ) *TeamsByTournament200Response`

NewTeamsByTournament200Response instantiates a new TeamsByTournament200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsByTournament200ResponseWithDefaults

`func NewTeamsByTournament200ResponseWithDefaults() *TeamsByTournament200Response`

NewTeamsByTournament200ResponseWithDefaults instantiates a new TeamsByTournament200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamsByTournament200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamsByTournament200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamsByTournament200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTeams

`func (o *TeamsByTournament200Response) GetTeams() []TeamsByTournament200ResponseTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TeamsByTournament200Response) GetTeamsOk() (*[]TeamsByTournament200ResponseTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TeamsByTournament200Response) SetTeams(v []TeamsByTournament200ResponseTeamsInner)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


