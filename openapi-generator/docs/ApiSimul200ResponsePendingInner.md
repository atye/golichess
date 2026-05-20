# ApiSimul200ResponsePendingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Host** | [**ApiSimul200ResponsePendingInnerHost**](ApiSimul200ResponsePendingInnerHost.md) |  | 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Variants** | [**[]ApiSimul200ResponsePendingInnerVariantsInner**](ApiSimul200ResponsePendingInnerVariantsInner.md) |  | 
**IsCreated** | **bool** |  | 
**IsFinished** | **bool** |  | 
**IsRunning** | **bool** |  | 
**Text** | Pointer to **string** |  | [optional] 
**EstimatedStartAt** | Pointer to **int32** |  | [optional] 
**StartedAt** | Pointer to **int32** |  | [optional] 
**FinishedAt** | Pointer to **int32** |  | [optional] 
**NbApplicants** | **int32** |  | 
**NbPairings** | **int32** |  | 

## Methods

### NewApiSimul200ResponsePendingInner

`func NewApiSimul200ResponsePendingInner(id string, host ApiSimul200ResponsePendingInnerHost, name string, fullName string, variants []ApiSimul200ResponsePendingInnerVariantsInner, isCreated bool, isFinished bool, isRunning bool, nbApplicants int32, nbPairings int32, ) *ApiSimul200ResponsePendingInner`

NewApiSimul200ResponsePendingInner instantiates a new ApiSimul200ResponsePendingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSimul200ResponsePendingInnerWithDefaults

`func NewApiSimul200ResponsePendingInnerWithDefaults() *ApiSimul200ResponsePendingInner`

NewApiSimul200ResponsePendingInnerWithDefaults instantiates a new ApiSimul200ResponsePendingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiSimul200ResponsePendingInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSimul200ResponsePendingInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSimul200ResponsePendingInner) SetId(v string)`

SetId sets Id field to given value.


### GetHost

`func (o *ApiSimul200ResponsePendingInner) GetHost() ApiSimul200ResponsePendingInnerHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApiSimul200ResponsePendingInner) GetHostOk() (*ApiSimul200ResponsePendingInnerHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApiSimul200ResponsePendingInner) SetHost(v ApiSimul200ResponsePendingInnerHost)`

SetHost sets Host field to given value.


### GetName

`func (o *ApiSimul200ResponsePendingInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSimul200ResponsePendingInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSimul200ResponsePendingInner) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *ApiSimul200ResponsePendingInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApiSimul200ResponsePendingInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApiSimul200ResponsePendingInner) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetVariants

`func (o *ApiSimul200ResponsePendingInner) GetVariants() []ApiSimul200ResponsePendingInnerVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ApiSimul200ResponsePendingInner) GetVariantsOk() (*[]ApiSimul200ResponsePendingInnerVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ApiSimul200ResponsePendingInner) SetVariants(v []ApiSimul200ResponsePendingInnerVariantsInner)`

SetVariants sets Variants field to given value.


### GetIsCreated

`func (o *ApiSimul200ResponsePendingInner) GetIsCreated() bool`

GetIsCreated returns the IsCreated field if non-nil, zero value otherwise.

### GetIsCreatedOk

`func (o *ApiSimul200ResponsePendingInner) GetIsCreatedOk() (*bool, bool)`

GetIsCreatedOk returns a tuple with the IsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreated

`func (o *ApiSimul200ResponsePendingInner) SetIsCreated(v bool)`

SetIsCreated sets IsCreated field to given value.


### GetIsFinished

`func (o *ApiSimul200ResponsePendingInner) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *ApiSimul200ResponsePendingInner) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *ApiSimul200ResponsePendingInner) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.


### GetIsRunning

`func (o *ApiSimul200ResponsePendingInner) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *ApiSimul200ResponsePendingInner) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *ApiSimul200ResponsePendingInner) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetText

`func (o *ApiSimul200ResponsePendingInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ApiSimul200ResponsePendingInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ApiSimul200ResponsePendingInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ApiSimul200ResponsePendingInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEstimatedStartAt

`func (o *ApiSimul200ResponsePendingInner) GetEstimatedStartAt() int32`

GetEstimatedStartAt returns the EstimatedStartAt field if non-nil, zero value otherwise.

### GetEstimatedStartAtOk

`func (o *ApiSimul200ResponsePendingInner) GetEstimatedStartAtOk() (*int32, bool)`

GetEstimatedStartAtOk returns a tuple with the EstimatedStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartAt

`func (o *ApiSimul200ResponsePendingInner) SetEstimatedStartAt(v int32)`

SetEstimatedStartAt sets EstimatedStartAt field to given value.

### HasEstimatedStartAt

`func (o *ApiSimul200ResponsePendingInner) HasEstimatedStartAt() bool`

HasEstimatedStartAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ApiSimul200ResponsePendingInner) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApiSimul200ResponsePendingInner) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApiSimul200ResponsePendingInner) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ApiSimul200ResponsePendingInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiSimul200ResponsePendingInner) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiSimul200ResponsePendingInner) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiSimul200ResponsePendingInner) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiSimul200ResponsePendingInner) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetNbApplicants

`func (o *ApiSimul200ResponsePendingInner) GetNbApplicants() int32`

GetNbApplicants returns the NbApplicants field if non-nil, zero value otherwise.

### GetNbApplicantsOk

`func (o *ApiSimul200ResponsePendingInner) GetNbApplicantsOk() (*int32, bool)`

GetNbApplicantsOk returns a tuple with the NbApplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbApplicants

`func (o *ApiSimul200ResponsePendingInner) SetNbApplicants(v int32)`

SetNbApplicants sets NbApplicants field to given value.


### GetNbPairings

`func (o *ApiSimul200ResponsePendingInner) GetNbPairings() int32`

GetNbPairings returns the NbPairings field if non-nil, zero value otherwise.

### GetNbPairingsOk

`func (o *ApiSimul200ResponsePendingInner) GetNbPairingsOk() (*int32, bool)`

GetNbPairingsOk returns a tuple with the NbPairings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPairings

`func (o *ApiSimul200ResponsePendingInner) SetNbPairings(v int32)`

SetNbPairings sets NbPairings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


