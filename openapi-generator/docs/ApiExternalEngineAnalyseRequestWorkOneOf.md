# ApiExternalEngineAnalyseRequestWorkOneOf

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

## Methods

### NewApiExternalEngineAnalyseRequestWorkOneOf

`func NewApiExternalEngineAnalyseRequestWorkOneOf(movetime int32, sessionId string, threads int32, hash int32, multiPv int32, variant string, initialFen string, moves []string, ) *ApiExternalEngineAnalyseRequestWorkOneOf`

NewApiExternalEngineAnalyseRequestWorkOneOf instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWorkOneOfWithDefaults

`func NewApiExternalEngineAnalyseRequestWorkOneOfWithDefaults() *ApiExternalEngineAnalyseRequestWorkOneOf`

NewApiExternalEngineAnalyseRequestWorkOneOfWithDefaults instantiates a new ApiExternalEngineAnalyseRequestWorkOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovetime

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMovetime() int32`

GetMovetime returns the Movetime field if non-nil, zero value otherwise.

### GetMovetimeOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMovetimeOk() (*int32, bool)`

GetMovetimeOk returns a tuple with the Movetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovetime

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetMovetime(v int32)`

SetMovetime sets Movetime field to given value.


### GetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetThreads(v int32)`

SetThreads sets Threads field to given value.


### GetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetHash() int32`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetHashOk() (*int32, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetHash(v int32)`

SetHash sets Hash field to given value.


### GetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMultiPv() int32`

GetMultiPv returns the MultiPv field if non-nil, zero value otherwise.

### GetMultiPvOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMultiPvOk() (*int32, bool)`

GetMultiPvOk returns a tuple with the MultiPv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPv

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetMultiPv(v int32)`

SetMultiPv sets MultiPv field to given value.


### GetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyseRequestWorkOneOf) SetMoves(v []string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


