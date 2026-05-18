# ApiUserNameTournamentPlayed200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tournament** | [**ApiTournament200ResponseCreatedInner**](ApiTournament200ResponseCreatedInner.md) |  | 
**Player** | [**ApiUserNameTournamentPlayed200ResponsePlayer**](ApiUserNameTournamentPlayed200ResponsePlayer.md) |  | 

## Methods

### NewApiUserNameTournamentPlayed200Response

`func NewApiUserNameTournamentPlayed200Response(tournament ApiTournament200ResponseCreatedInner, player ApiUserNameTournamentPlayed200ResponsePlayer, ) *ApiUserNameTournamentPlayed200Response`

NewApiUserNameTournamentPlayed200Response instantiates a new ApiUserNameTournamentPlayed200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserNameTournamentPlayed200ResponseWithDefaults

`func NewApiUserNameTournamentPlayed200ResponseWithDefaults() *ApiUserNameTournamentPlayed200Response`

NewApiUserNameTournamentPlayed200ResponseWithDefaults instantiates a new ApiUserNameTournamentPlayed200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTournament

`func (o *ApiUserNameTournamentPlayed200Response) GetTournament() ApiTournament200ResponseCreatedInner`

GetTournament returns the Tournament field if non-nil, zero value otherwise.

### GetTournamentOk

`func (o *ApiUserNameTournamentPlayed200Response) GetTournamentOk() (*ApiTournament200ResponseCreatedInner, bool)`

GetTournamentOk returns a tuple with the Tournament field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournament

`func (o *ApiUserNameTournamentPlayed200Response) SetTournament(v ApiTournament200ResponseCreatedInner)`

SetTournament sets Tournament field to given value.


### GetPlayer

`func (o *ApiUserNameTournamentPlayed200Response) GetPlayer() ApiUserNameTournamentPlayed200ResponsePlayer`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *ApiUserNameTournamentPlayed200Response) GetPlayerOk() (*ApiUserNameTournamentPlayed200ResponsePlayer, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *ApiUserNameTournamentPlayed200Response) SetPlayer(v ApiUserNameTournamentPlayed200ResponsePlayer)`

SetPlayer sets Player field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


