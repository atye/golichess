# Verdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** |  | 
**Verdict** | **string** |  | 

## Methods

### NewVerdict

`func NewVerdict(condition string, verdict string, ) *Verdict`

NewVerdict instantiates a new Verdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerdictWithDefaults

`func NewVerdictWithDefaults() *Verdict`

NewVerdictWithDefaults instantiates a new Verdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *Verdict) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Verdict) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Verdict) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetVerdict

`func (o *Verdict) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *Verdict) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *Verdict) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


