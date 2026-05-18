# ApiExternalEngineAnalyseRequestWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Movetime** | **int32** | Amount of time to analyse the position, in milliseconds. | 
**SessionId** | **string** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **int32** | Number of threads to use for analysis. | 
**Hash** | **int32** | Hash table size to use for analysis, in MiB. | 
**MultiPv** | **int32** | Requested number of principal variations. | 
**Variant** | **string** |  | [default to "chess"]
**InitialFen** | **string** | Initial position of the game. | 
**Moves** | **[]string** | List of moves played from the initial position, in UCI notation. | 
**Depth** | **int32** | Analysis target depth | 
**Nodes** | **int32** | Number of nodes to analyse in the position | 

## Methods

### NewApiExternalEngineAnalyseRequestWork

`func NewApiExternalEngineAnalyseRequestWork(movetime int32, sessionId string, threads int32, hash int32, multiPv int32, variant string, initialFen string, moves []string, depth int32, nodes int32, ) *ApiExternalEngineAnalyseRequestWork`

NewApiExternalEngineAnalyseRequestWork instantiates a new ApiExternalEngineAnalyseRequestWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWorkWithDefaults

`func NewApiExternalEngineAnalyseRequestWorkWithDefaults() *ApiExternalEngineAnalyseRequestWork`

NewApiExternalEngineAnalyseRequestWorkWithDefaults instantiates a new ApiExternalEngineAnalyseRequestWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovetime

`func (o *ApiExternalEngineAnalyseRequestWork) GetMovetime() int32`

GetMovetime returns the Movetime field if non-nil, zero value otherwise.

### GetMovetimeOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetMovetimeOk() (*int32, bool)`

GetMovetimeOk returns a tuple with the Movetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovetime

`func (o *ApiExternalEngineAnalyseRequestWork) SetMovetime(v int32)`

SetMovetime sets Movetime field to given value.


### GetSessionId

`func (o *ApiExternalEngineAnalyseRequestWork) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApiExternalEngineAnalyseRequestWork) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ApiExternalEngineAnalyseRequestWork) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ApiExternalEngineAnalyseRequestWork) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ApiExternalEngineAnalyseRequestWork) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiExternalEngineAnalyseRequestWork) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWork) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWork) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ApiExternalEngineAnalyseRequestWork) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiExternalEngineAnalyseRequestWork) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWork) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWork) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ApiExternalEngineAnalyseRequestWork) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyseRequestWork) SetMoves(v []string)`

SetMoves sets Moves field to given value.


### GetDepth

`func (o *ApiExternalEngineAnalyseRequestWork) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyseRequestWork) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetNodes

`func (o *ApiExternalEngineAnalyseRequestWork) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ApiExternalEngineAnalyseRequestWork) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ApiExternalEngineAnalyseRequestWork) SetNodes(v int32)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


