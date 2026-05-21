# UserActivityTournaments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nb** | Pointer to **int32** |  | [optional] 
**Best** | Pointer to [**[]UserActivityTournamentsBestInner**](UserActivityTournamentsBestInner.md) |  | [optional] 

## Methods

### NewUserActivityTournaments

`func NewUserActivityTournaments() *UserActivityTournaments`

NewUserActivityTournaments instantiates a new UserActivityTournaments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityTournamentsWithDefaults

`func NewUserActivityTournamentsWithDefaults() *UserActivityTournaments`

NewUserActivityTournamentsWithDefaults instantiates a new UserActivityTournaments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNb

`func (o *UserActivityTournaments) GetNb() int32`

GetNb returns the Nb field if non-nil, zero value otherwise.

### GetNbOk

`func (o *UserActivityTournaments) GetNbOk() (*int32, bool)`

GetNbOk returns a tuple with the Nb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNb

`func (o *UserActivityTournaments) SetNb(v int32)`

SetNb sets Nb field to given value.

### HasNb

`func (o *UserActivityTournaments) HasNb() bool`

HasNb returns a boolean if a field has been set.

### GetBest

`func (o *UserActivityTournaments) GetBest() []UserActivityTournamentsBestInner`

GetBest returns the Best field if non-nil, zero value otherwise.

### GetBestOk

`func (o *UserActivityTournaments) GetBestOk() (*[]UserActivityTournamentsBestInner, bool)`

GetBestOk returns a tuple with the Best field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBest

`func (o *UserActivityTournaments) SetBest(v []UserActivityTournamentsBestInner)`

SetBest sets Best field to given value.

### HasBest

`func (o *UserActivityTournaments) HasBest() bool`

HasBest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


