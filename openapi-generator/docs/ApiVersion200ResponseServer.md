# ApiVersion200ResponseServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | When the deployed server was built | 
**Commit** | **string** | Git commit hash of the deployed lila version | 

## Methods

### NewApiVersion200ResponseServer

`func NewApiVersion200ResponseServer(date time.Time, commit string, ) *ApiVersion200ResponseServer`

NewApiVersion200ResponseServer instantiates a new ApiVersion200ResponseServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVersion200ResponseServerWithDefaults

`func NewApiVersion200ResponseServerWithDefaults() *ApiVersion200ResponseServer`

NewApiVersion200ResponseServerWithDefaults instantiates a new ApiVersion200ResponseServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ApiVersion200ResponseServer) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ApiVersion200ResponseServer) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ApiVersion200ResponseServer) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetCommit

`func (o *ApiVersion200ResponseServer) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ApiVersion200ResponseServer) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ApiVersion200ResponseServer) SetCommit(v string)`

SetCommit sets Commit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


