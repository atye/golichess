# BroadcastRoundGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Round** | [**BroadcastsOfficial200ResponseRoundsInner**](BroadcastsOfficial200ResponseRoundsInner.md) |  | 
**Tour** | [**BroadcastsOfficial200ResponseTour**](BroadcastsOfficial200ResponseTour.md) |  | 
**Study** | [**BroadcastRoundCreate200ResponseStudy**](BroadcastRoundCreate200ResponseStudy.md) |  | 
**Games** | [**[]BroadcastRoundGet200ResponseGamesInner**](BroadcastRoundGet200ResponseGamesInner.md) |  | 
**Group** | Pointer to [**BroadcastsOfficial200ResponseGroup**](BroadcastsOfficial200ResponseGroup.md) |  | [optional] 
**IsSubscribed** | Pointer to **bool** | Indicates if the user making the request is subscribed to the broadcast | [optional] 
**Photos** | [**map[string]BroadcastsOfficial200ResponsePhotosValue**](BroadcastsOfficial200ResponsePhotosValue.md) | Photos of players, when available. The object keys are FIDE IDs | 

## Methods

### NewBroadcastRoundGet200Response

`func NewBroadcastRoundGet200Response(round BroadcastsOfficial200ResponseRoundsInner, tour BroadcastsOfficial200ResponseTour, study BroadcastRoundCreate200ResponseStudy, games []BroadcastRoundGet200ResponseGamesInner, photos map[string]BroadcastsOfficial200ResponsePhotosValue, ) *BroadcastRoundGet200Response`

NewBroadcastRoundGet200Response instantiates a new BroadcastRoundGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundGet200ResponseWithDefaults

`func NewBroadcastRoundGet200ResponseWithDefaults() *BroadcastRoundGet200Response`

NewBroadcastRoundGet200ResponseWithDefaults instantiates a new BroadcastRoundGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRound

`func (o *BroadcastRoundGet200Response) GetRound() BroadcastsOfficial200ResponseRoundsInner`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BroadcastRoundGet200Response) GetRoundOk() (*BroadcastsOfficial200ResponseRoundsInner, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BroadcastRoundGet200Response) SetRound(v BroadcastsOfficial200ResponseRoundsInner)`

SetRound sets Round field to given value.


### GetTour

`func (o *BroadcastRoundGet200Response) GetTour() BroadcastsOfficial200ResponseTour`

GetTour returns the Tour field if non-nil, zero value otherwise.

### GetTourOk

`func (o *BroadcastRoundGet200Response) GetTourOk() (*BroadcastsOfficial200ResponseTour, bool)`

GetTourOk returns a tuple with the Tour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTour

`func (o *BroadcastRoundGet200Response) SetTour(v BroadcastsOfficial200ResponseTour)`

SetTour sets Tour field to given value.


### GetStudy

`func (o *BroadcastRoundGet200Response) GetStudy() BroadcastRoundCreate200ResponseStudy`

GetStudy returns the Study field if non-nil, zero value otherwise.

### GetStudyOk

`func (o *BroadcastRoundGet200Response) GetStudyOk() (*BroadcastRoundCreate200ResponseStudy, bool)`

GetStudyOk returns a tuple with the Study field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudy

`func (o *BroadcastRoundGet200Response) SetStudy(v BroadcastRoundCreate200ResponseStudy)`

SetStudy sets Study field to given value.


### GetGames

`func (o *BroadcastRoundGet200Response) GetGames() []BroadcastRoundGet200ResponseGamesInner`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BroadcastRoundGet200Response) GetGamesOk() (*[]BroadcastRoundGet200ResponseGamesInner, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BroadcastRoundGet200Response) SetGames(v []BroadcastRoundGet200ResponseGamesInner)`

SetGames sets Games field to given value.


### GetGroup

`func (o *BroadcastRoundGet200Response) GetGroup() BroadcastsOfficial200ResponseGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BroadcastRoundGet200Response) GetGroupOk() (*BroadcastsOfficial200ResponseGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BroadcastRoundGet200Response) SetGroup(v BroadcastsOfficial200ResponseGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BroadcastRoundGet200Response) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *BroadcastRoundGet200Response) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *BroadcastRoundGet200Response) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *BroadcastRoundGet200Response) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *BroadcastRoundGet200Response) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetPhotos

`func (o *BroadcastRoundGet200Response) GetPhotos() map[string]BroadcastsOfficial200ResponsePhotosValue`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *BroadcastRoundGet200Response) GetPhotosOk() (*map[string]BroadcastsOfficial200ResponsePhotosValue, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *BroadcastRoundGet200Response) SetPhotos(v map[string]BroadcastsOfficial200ResponsePhotosValue)`

SetPhotos sets Photos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


