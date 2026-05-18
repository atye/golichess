# Tournament200ResponseStanding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**Players** | Pointer to [**[]Tournament200ResponseStandingPlayersInner**](Tournament200ResponseStandingPlayersInner.md) |  | [optional] 

## Methods

### NewTournament200ResponseStanding

`func NewTournament200ResponseStanding() *Tournament200ResponseStanding`

NewTournament200ResponseStanding instantiates a new Tournament200ResponseStanding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTournament200ResponseStandingWithDefaults

`func NewTournament200ResponseStandingWithDefaults() *Tournament200ResponseStanding`

NewTournament200ResponseStandingWithDefaults instantiates a new Tournament200ResponseStanding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *Tournament200ResponseStanding) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Tournament200ResponseStanding) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Tournament200ResponseStanding) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *Tournament200ResponseStanding) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlayers

`func (o *Tournament200ResponseStanding) GetPlayers() []Tournament200ResponseStandingPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *Tournament200ResponseStanding) GetPlayersOk() (*[]Tournament200ResponseStandingPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *Tournament200ResponseStanding) SetPlayers(v []Tournament200ResponseStandingPlayersInner)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *Tournament200ResponseStanding) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


