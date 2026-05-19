# ApiExternalEngineAnalyseRequestWorkOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **int32** | Analysis target depth | 
**SessionId** | **string** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **int32** | Number of threads to use for analysis. | 
**Hash** | **int32** | Hash table size to use for analysis, in MiB. | 
**MultiPv** | **int32** | Requested number of principal variations. | 
**Variant** | **string** |  | [default to "chess"]
**InitialFen** | **string** | Initial position of the game. | 
**Moves** | **[]string** | List of moves played from the initial position, in UCI notation. | 

## Methods

### NewApiExternalEngineAnalyseRequestWorkOneOf1

`func NewApiExternalEngineAnalyseRequestWorkOneOf1(depth int32, sessionId string, threads int32, hash int32, multiPv int32, variant string, initialFen string, moves []string, ) *ApiExternalEngineAnalyseRequestWorkOneOf1`

NewApiExternalEngineAnalyseRequestWorkOneOf1 instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWorkOneOf1WithDefaults

`func NewApiExternalEngineAnalyseRequestWorkOneOf1WithDefaults() *ApiExternalEngineAnalyseRequestWorkOneOf1`

NewApiExternalEngineAnalyseRequestWorkOneOf1WithDefaults instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf1) SetMoves(v []string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


