# Timeline200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]Timeline200ResponseEntriesInner**](Timeline200ResponseEntriesInner.md) |  | 
**Users** | [**map[string]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | 

## Methods

### NewTimeline200Response

`func NewTimeline200Response(entries []Timeline200ResponseEntriesInner, users map[string]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, ) *Timeline200Response`

NewTimeline200Response instantiates a new Timeline200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeline200ResponseWithDefaults

`func NewTimeline200ResponseWithDefaults() *Timeline200Response`

NewTimeline200ResponseWithDefaults instantiates a new Timeline200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *Timeline200Response) GetEntries() []Timeline200ResponseEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Timeline200Response) GetEntriesOk() (*[]Timeline200ResponseEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Timeline200Response) SetEntries(v []Timeline200ResponseEntriesInner)`

SetEntries sets Entries field to given value.


### GetUsers

`func (o *Timeline200Response) GetUsers() map[string]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Timeline200Response) GetUsersOk() (*map[string]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Timeline200Response) SetUsers(v map[string]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


