# BroadcastRound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Round** | [**BroadcastRoundInfo**](BroadcastRoundInfo.md) |  | 
**Tour** | [**BroadcastTour**](BroadcastTour.md) |  | 
**Study** | [**BroadcastRoundStudyInfo**](BroadcastRoundStudyInfo.md) |  | 
**Games** | [**[]BroadcastRoundGame**](BroadcastRoundGame.md) |  | 
**Group** | Pointer to [**BroadcastGroup**](BroadcastGroup.md) |  | [optional] 
**IsSubscribed** | Pointer to **bool** | Indicates if the user making the request is subscribed to the broadcast | [optional] 
**Photos** | [**map[string]BroadcastPhotosValue**](BroadcastPhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | 

## Methods

### NewBroadcastRound

`func NewBroadcastRound(round BroadcastRoundInfo, tour BroadcastTour, study BroadcastRoundStudyInfo, games []BroadcastRoundGame, photos map[string]BroadcastPhotosValue, ) *BroadcastRound`

NewBroadcastRound instantiates a new BroadcastRound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundWithDefaults

`func NewBroadcastRoundWithDefaults() *BroadcastRound`

NewBroadcastRoundWithDefaults instantiates a new BroadcastRound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRound

`func (o *BroadcastRound) GetRound() BroadcastRoundInfo`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BroadcastRound) GetRoundOk() (*BroadcastRoundInfo, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BroadcastRound) SetRound(v BroadcastRoundInfo)`

SetRound sets Round field to given value.


### GetTour

`func (o *BroadcastRound) GetTour() BroadcastTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastRound) GetTourOk() (*BroadcastTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastRound) SetTour(v BroadcastTour)`

SetTour sets Tour field to given value.


### GetStudy

`func (o *BroadcastRound) GetStudy() BroadcastRoundStudyInfo`

GetStudy returns the Study field if non-nil, zero value otherwise.

### GetStudyOk

`func (o *BroadcastRound) GetStudyOk() (*BroadcastRoundStudyInfo, bool)`

GetStudyOk returns a tuple with the Study field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudy

`func (o *BroadcastRound) SetStudy(v BroadcastRoundStudyInfo)`

SetStudy sets Study field to given value.


### GetGames

`func (o *BroadcastRound) GetGames() []BroadcastRoundGame`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BroadcastRound) GetGamesOk() (*[]BroadcastRoundGame, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BroadcastRound) SetGames(v []BroadcastRoundGame)`

SetGames sets Games field to given value.


### GetGroup

`func (o *BroadcastRound) GetGroup() BroadcastGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastRound) GetGroupOk() (*BroadcastGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastRound) SetGroup(v BroadcastGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastRound) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *BroadcastRound) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *BroadcastRound) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *BroadcastRound) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *BroadcastRound) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastRound) GetPhotos() map[string]BroadcastPhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastRound) GetPhotosOk() (*map[string]BroadcastPhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastRound) SetPhotos(v map[string]BroadcastPhotosValue)`

SetPhotos sets Photos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


