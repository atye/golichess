# Simul

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Host** | [**SimulHost**](SimulHost.md) |  | 
**Name** | **string** |  | 
**FullName** | **string** |  | 
**Variants** | [**[]SimulVariantsInner**](SimulVariantsInner.md) |  | 
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

### NewSimul

`func NewSimul(id string, host SimulHost, name string, fullName string, variants []SimulVariantsInner, isCreated bool, isFinished bool, isRunning bool, nbApplicants int32, nbPairings int32, ) *Simul`

NewSimul instantiates a new Simul object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulWithDefaults

`func NewSimulWithDefaults() *Simul`

NewSimulWithDefaults instantiates a new Simul object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Simul) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Simul) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Simul) SetId(v string)`

SetId sets Id field to given value.


### GetHost

`func (o *Simul) GetHost() SimulHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Simul) GetHostOk() (*SimulHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Simul) SetHost(v SimulHost)`

SetHost sets Host field to given value.


### GetName

`func (o *Simul) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Simul) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Simul) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *Simul) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Simul) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Simul) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetVariants

`func (o *Simul) GetVariants() []SimulVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Simul) GetVariantsOk() (*[]SimulVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Simul) SetVariants(v []SimulVariantsInner)`

SetVariants sets Variants field to given value.


### GetIsCreated

`func (o *Simul) GetIsCreated() bool`

GetIsCreated returns the IsCreated field if non-nil, zero value otherwise.

### GetIsCreatedOk

`func (o *Simul) GetIsCreatedOk() (*bool, bool)`

GetIsCreatedOk returns a tuple with the IsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreated

`func (o *Simul) SetIsCreated(v bool)`

SetIsCreated sets IsCreated field to given value.


### GetIsFinished

`func (o *Simul) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *Simul) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *Simul) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.


### GetIsRunning

`func (o *Simul) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *Simul) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *Simul) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetText

`func (o *Simul) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Simul) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Simul) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Simul) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEstimatedStartAt

`func (o *Simul) GetEstimatedStartAt() int32`

GetEstimatedStartAt returns the EstimatedStartAt field if non-nil, zero value otherwise.

### GetEstimatedStartAtOk

`func (o *Simul) GetEstimatedStartAtOk() (*int32, bool)`

GetEstimatedStartAtOk returns a tuple with the EstimatedStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartAt

`func (o *Simul) SetEstimatedStartAt(v int32)`

SetEstimatedStartAt sets EstimatedStartAt field to given value.

### HasEstimatedStartAt

`func (o *Simul) HasEstimatedStartAt() bool`

HasEstimatedStartAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Simul) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Simul) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Simul) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Simul) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *Simul) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Simul) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Simul) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Simul) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetNbApplicants

`func (o *Simul) GetNbApplicants() int32`

GetNbApplicants returns the NbApplicants field if non-nil, zero value otherwise.

### GetNbApplicantsOk

`func (o *Simul) GetNbApplicantsOk() (*int32, bool)`

GetNbApplicantsOk returns a tuple with the NbApplicants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbApplicants

`func (o *Simul) SetNbApplicants(v int32)`

SetNbApplicants sets NbApplicants field to given value.


### GetNbPairings

`func (o *Simul) GetNbPairings() int32`

GetNbPairings returns the NbPairings field if non-nil, zero value otherwise.

### GetNbPairingsOk

`func (o *Simul) GetNbPairingsOk() (*int32, bool)`

GetNbPairingsOk returns a tuple with the NbPairings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPairings

`func (o *Simul) SetNbPairings(v int32)`

SetNbPairings sets NbPairings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


