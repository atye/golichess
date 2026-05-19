# ApiExternalEngineAnalyseRequestWorkOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | **int32** | Number of nodes to analyse in the position | 
**SessionId** | **string** | Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions.  | 
**Threads** | **int32** | Number of threads to use for analysis. | 
**Hash** | **int32** | Hash table size to use for analysis, in MiB. | 
**MultiPv** | **int32** | Requested number of principal variations. | 
**Variant** | **string** |  | [default to "chess"]
**InitialFen** | **string** | Initial position of the game. | 
**Moves** | **[]string** | List of moves played from the initial position, in UCI notation. | 

## Methods

### NewApiExternalEngineAnalyseRequestWorkOneOf2

`func NewApiExternalEngineAnalyseRequestWorkOneOf2(nodes int32, sessionId string, threads int32, hash int32, multiPv int32, variant string, initialFen string, moves []string, ) *ApiExternalEngineAnalyseRequestWorkOneOf2`

NewApiExternalEngineAnalyseRequestWorkOneOf2 instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWorkOneOf2WithDefaults

`func NewApiExternalEngineAnalyseRequestWorkOneOf2WithDefaults() *ApiExternalEngineAnalyseRequestWorkOneOf2`

NewApiExternalEngineAnalyseRequestWorkOneOf2WithDefaults instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetNodes(v int32)`

SetNodes sets Nodes field to given value.


### GetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf2) SetMoves(v []string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


