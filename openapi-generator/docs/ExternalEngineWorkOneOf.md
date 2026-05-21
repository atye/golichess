# ExternalEngineWorkOneOf

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
**Movetime** | **int32** | Amount of time to analyse the position, in milliseconds. | 

## Methods

### NewExternalEngineWorkOneOf

`func NewExternalEngineWorkOneOf(sessionId string, threads int32, hash int32, multiPv int32, variant UciVariant, initialFen string, moves []string, movetime int32, ) *ExternalEngineWorkOneOf`

NewExternalEngineWorkOneOf instantiates a new ExternalEngineWorkOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWorkOneOfWithDefaults

`func NewExternalEngineWorkOneOfWithDefaults() *ExternalEngineWorkOneOf`

NewExternalEngineWorkOneOfWithDefaults instantiates a new ExternalEngineWorkOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ExternalEngineWorkOneOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExternalEngineWorkOneOf) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExternalEngineWorkOneOf) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ExternalEngineWorkOneOf) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ExternalEngineWorkOneOf) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ExternalEngineWorkOneOf) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ExternalEngineWorkOneOf) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ExternalEngineWorkOneOf) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ExternalEngineWorkOneOf) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ExternalEngineWorkOneOf) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ExternalEngineWorkOneOf) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ExternalEngineWorkOneOf) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ExternalEngineWorkOneOf) GetVariant() UciVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ExternalEngineWorkOneOf) GetVariantOk() (*UciVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ExternalEngineWorkOneOf) SetVariant(v UciVariant)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ExternalEngineWorkOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ExternalEngineWorkOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ExternalEngineWorkOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ExternalEngineWorkOneOf) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ExternalEngineWorkOneOf) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ExternalEngineWorkOneOf) SetMoves(v []string)`

SetMoves sets Moves field to given value.


### GetMovetime

`func (o *ExternalEngineWorkOneOf) GetMovetime() int32`

GetMovetime returns the Movetime field if non-nil, zero value otherwise.

### GetMovetimeOk

`func (o *ExternalEngineWorkOneOf) GetMovetimeOk() (*int32, bool)`

GetMovetimeOk returns a tuple with the Movetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovetime

`func (o *ExternalEngineWorkOneOf) SetMovetime(v int32)`

SetMovetime sets Movetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


