# ExternalEngineWorkOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **int32** | Number of threads to use for analysis. | 
**Hash** | **int32** | Hash table size to use for analysis, in MiB. | 
**MultiPv** | **int32** | Requested number of principal variations. | 
**Variant** | [**UciVariant**](UciVariant.md) |  | [default to UCIVARIANT_CHESS]
**InitialFen** | **string** | Initial position of the game. | 
**Moves** | **[]string** | List of moves played from the initial position, in UCI notation. | 
**Depth** | **int32** | Analysis target depth | 

## Methods

### NewExternalEngineWorkOneOf1

`func NewExternalEngineWorkOneOf1(sessionId string, threads int32, hash int32, multiPv int32, variant UciVariant, initialFen string, moves []string, depth int32, ) *ExternalEngineWorkOneOf1`

NewExternalEngineWorkOneOf1 instantiates a new ExternalEngineWorkOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWorkOneOf1WithDefaults

`func NewExternalEngineWorkOneOf1WithDefaults() *ExternalEngineWorkOneOf1`

NewExternalEngineWorkOneOf1WithDefaults instantiates a new ExternalEngineWorkOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ExternalEngineWorkOneOf1) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExternalEngineWorkOneOf1) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExternalEngineWorkOneOf1) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ExternalEngineWorkOneOf1) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ExternalEngineWorkOneOf1) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ExternalEngineWorkOneOf1) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ExternalEngineWorkOneOf1) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ExternalEngineWorkOneOf1) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ExternalEngineWorkOneOf1) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ExternalEngineWorkOneOf1) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ExternalEngineWorkOneOf1) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ExternalEngineWorkOneOf1) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ExternalEngineWorkOneOf1) GetVariant() UciVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ExternalEngineWorkOneOf1) GetVariantOk() (*UciVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ExternalEngineWorkOneOf1) SetVariant(v UciVariant)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ExternalEngineWorkOneOf1) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ExternalEngineWorkOneOf1) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ExternalEngineWorkOneOf1) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ExternalEngineWorkOneOf1) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ExternalEngineWorkOneOf1) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ExternalEngineWorkOneOf1) SetMoves(v []string)`

SetMoves sets Moves field to given value.


### GetDepth

`func (o *ExternalEngineWorkOneOf1) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ExternalEngineWorkOneOf1) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ExternalEngineWorkOneOf1) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


