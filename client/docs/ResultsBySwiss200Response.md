# ResultsBySwiss200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absent** | Pointer to **bool** |  | [optional] 
**Rank** | **int32** |  | 
**Points** | **float32** |  | 
**TieBreak** | **int32** |  | 
**Rating** | **int32** |  | 
**Username** | **string** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Performance** | **int32** |  | 

## Methods

### NewResultsBySwiss200Response

`func NewResultsBySwiss200Response(rank int32, points float32, tieBreak int32, rating int32, username string, performance int32, ) *ResultsBySwiss200Response`

NewResultsBySwiss200Response instantiates a new ResultsBySwiss200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsBySwiss200ResponseWithDefaults

`func NewResultsBySwiss200ResponseWithDefaults() *ResultsBySwiss200Response`

NewResultsBySwiss200ResponseWithDefaults instantiates a new ResultsBySwiss200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsent

`func (o *ResultsBySwiss200Response) GetAbsent() bool`

GetAbsent returns the Absent field if non-nil, zero value otherwise.

### GetAbsentOk

`func (o *ResultsBySwiss200Response) GetAbsentOk() (*bool, bool)`

GetAbsentOk returns a tuple with the Absent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsent

`func (o *ResultsBySwiss200Response) SetAbsent(v bool)`

SetAbsent sets Absent field to given value.

### HasAbsent

`func (o *ResultsBySwiss200Response) HasAbsent() bool`

HasAbsent returns a boolean if a field has been set.

### GetRank

`func (o *ResultsBySwiss200Response) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ResultsBySwiss200Response) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ResultsBySwiss200Response) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetPoints

`func (o *ResultsBySwiss200Response) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ResultsBySwiss200Response) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ResultsBySwiss200Response) SetPoints(v float32)`

SetPoints sets Points field to given value.


### GetTieBreak

`func (o *ResultsBySwiss200Response) GetTieBreak() int32`

GetTieBreak returns the TieBreak field if non-nil, zero value otherwise.

### GetTieBreakOk

`func (o *ResultsBySwiss200Response) GetTieBreakOk() (*int32, bool)`

GetTieBreakOk returns a tuple with the TieBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieBreak

`func (o *ResultsBySwiss200Response) SetTieBreak(v int32)`

SetTieBreak sets TieBreak field to given value.


### GetRating

`func (o *ResultsBySwiss200Response) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ResultsBySwiss200Response) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ResultsBySwiss200Response) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetUsername

`func (o *ResultsBySwiss200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResultsBySwiss200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResultsBySwiss200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetTitle

`func (o *ResultsBySwiss200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResultsBySwiss200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResultsBySwiss200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResultsBySwiss200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResultsBySwiss200Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResultsBySwiss200Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPerformance

`func (o *ResultsBySwiss200Response) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *ResultsBySwiss200Response) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *ResultsBySwiss200Response) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


