# Variant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**VariantKey**](VariantKey.md) |  | [default to VARIANTKEY_STANDARD]
**Name** | **string** |  | 
**Short** | Pointer to **string** |  | [optional] 

## Methods

### NewVariant

`func NewVariant(key VariantKey, name string, ) *Variant`

NewVariant instantiates a new Variant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantWithDefaults

`func NewVariantWithDefaults() *Variant`

NewVariantWithDefaults instantiates a new Variant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Variant) GetKey() VariantKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Variant) GetKeyOk() (*VariantKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Variant) SetKey(v VariantKey)`

SetKey sets Key field to given value.


### GetName

`func (o *Variant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variant) SetName(v string)`

SetName sets Name field to given value.


### GetShort

`func (o *Variant) GetShort() string`

GetShort returns the Short field if non-nil, zero value otherwise.

### GetShortOk

`func (o *Variant) GetShortOk() (*string, bool)`

GetShortOk returns a tuple with the Short field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShort

`func (o *Variant) SetShort(v string)`

SetShort sets Short field to given value.

### HasShort

`func (o *Variant) HasShort() bool`

HasShort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


