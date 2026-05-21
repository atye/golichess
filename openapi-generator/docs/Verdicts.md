# Verdicts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **bool** |  | 
**List** | [**[]Verdict**](Verdict.md) |  | 

## Methods

### NewVerdicts

`func NewVerdicts(accepted bool, list []Verdict, ) *Verdicts`

NewVerdicts instantiates a new Verdicts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerdictsWithDefaults

`func NewVerdictsWithDefaults() *Verdicts`

NewVerdictsWithDefaults instantiates a new Verdicts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *Verdicts) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *Verdicts) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *Verdicts) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.


### GetList

`func (o *Verdicts) GetList() []Verdict`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *Verdicts) GetListOk() (*[]Verdict, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *Verdicts) SetList(v []Verdict)`

SetList sets List field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


