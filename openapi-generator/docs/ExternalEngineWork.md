# ExternalEngineWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Movetime** | **int32** | Amount of time to analyse the position, in milliseconds. | 
**SessionId** | **string** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **int32** | Number of threads to use for analysis. | 
**Hash** | **int32** | Hash table size to use for analysis, in MiB. | 
**MultiPv** | **int32** | Requested number of principal variations. | 
**Variant** | [**UciVariant**](UciVariant.md) |  | [default to UCIVARIANT_CHESS]
**InitialFen** | **string** | Initial position of the game. | 
**Moves** | **[]string** | List of moves played from the initial position, in UCI notation. | 
**Depth** | **int32** | Analysis target depth | 
**Nodes** | **int32** | Number of nodes to analyse in the position | 

## Methods

### NewExternalEngineWork

`func NewExternalEngineWork(movetime int32, sessionId string, threads int32, hash int32, multiPv int32, variant UciVariant, initialFen string, moves []string, depth int32, nodes int32, ) *ExternalEngineWork`

NewExternalEngineWork instantiates a new ExternalEngineWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWorkWithDefaults

`func NewExternalEngineWorkWithDefaults() *ExternalEngineWork`

NewExternalEngineWorkWithDefaults instantiates a new ExternalEngineWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovetime

`func (o *ExternalEngineWork) GetMovetime() int32`

GetMovetime returns the Movetime field if non-nil, zero value otherwise.

### GetMovetimeOk

`func (o *ExternalEngineWork) GetMovetimeOk() (*int32, bool)`

GetMovetimeOk returns a tuple with the Movetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovetime

`func (o *ExternalEngineWork) SetMovetime(v int32)`

SetMovetime sets Movetime field to given value.


### GetSessionId

`func (o *ExternalEngineWork) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExternalEngineWork) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExternalEngineWork) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ExternalEngineWork) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ExternalEngineWork) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ExternalEngineWork) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ExternalEngineWork) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ExternalEngineWork) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ExternalEngineWork) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ExternalEngineWork) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ExternalEngineWork) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ExternalEngineWork) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ExternalEngineWork) GetVariant() UciVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ExternalEngineWork) GetVariantOk() (*UciVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ExternalEngineWork) SetVariant(v UciVariant)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ExternalEngineWork) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ExternalEngineWork) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ExternalEngineWork) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ExternalEngineWork) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ExternalEngineWork) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ExternalEngineWork) SetMoves(v []string)`

SetMoves sets Moves field to given value.


### GetDepth

`func (o *ExternalEngineWork) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ExternalEngineWork) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ExternalEngineWork) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetNodes

`func (o *ExternalEngineWork) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ExternalEngineWork) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ExternalEngineWork) SetNodes(v int32)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


