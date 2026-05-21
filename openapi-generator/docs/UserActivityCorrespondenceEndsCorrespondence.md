# UserActivityCorrespondenceEndsCorrespondence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | [**UserActivityScore**](UserActivityScore.md) |  | 
**Games** | [**[]UserActivityCorrespondenceGame**](UserActivityCorrespondenceGame.md) |  | 

## Methods

### NewUserActivityCorrespondenceEndsCorrespondence

`func NewUserActivityCorrespondenceEndsCorrespondence(score UserActivityScore, games []UserActivityCorrespondenceGame, ) *UserActivityCorrespondenceEndsCorrespondence`

NewUserActivityCorrespondenceEndsCorrespondence instantiates a new UserActivityCorrespondenceEndsCorrespondence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityCorrespondenceEndsCorrespondenceWithDefaults

`func NewUserActivityCorrespondenceEndsCorrespondenceWithDefaults() *UserActivityCorrespondenceEndsCorrespondence`

NewUserActivityCorrespondenceEndsCorrespondenceWithDefaults instantiates a new UserActivityCorrespondenceEndsCorrespondence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *UserActivityCorrespondenceEndsCorrespondence) GetScore() UserActivityScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UserActivityCorrespondenceEndsCorrespondence) GetScoreOk() (*UserActivityScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UserActivityCorrespondenceEndsCorrespondence) SetScore(v UserActivityScore)`

SetScore sets Score field to given value.


### GetGames

`func (o *UserActivityCorrespondenceEndsCorrespondence) GetGames() []UserActivityCorrespondenceGame`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *UserActivityCorrespondenceEndsCorrespondence) GetGamesOk() (*[]UserActivityCorrespondenceGame, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *UserActivityCorrespondenceEndsCorrespondence) SetGames(v []UserActivityCorrespondenceGame)`

SetGames sets Games field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


