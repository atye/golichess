# BroadcastTourInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Tournament format. Example: &#x60;\&quot;8-player round-robin\&quot; or \&quot;5-round Swiss\&quot;&#x60;  | [optional] 
**Tc** | Pointer to **string** | Time control. Example: &#x60;\&quot;Classical\&quot; or \&quot;Rapid\&quot; or \&quot;Rapid &amp; Blitz\&quot;&#x60;  | [optional] 
**FideTC** | Pointer to [**FideTimeControl**](FideTimeControl.md) |  | [optional] 
**TimeZone** | Pointer to **string** | Timezone of the tournament. Example: &#x60;America/New_York&#x60;. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more.  | [optional] 
**Location** | Pointer to **string** | Tournament location | [optional] 
**Players** | Pointer to **string** | Mentioning up to 4 of the best players participating.  | [optional] 
**Website** | Pointer to **string** | Official website. External website URL | [optional] 
**Standings** | Pointer to **string** | Official standings website. External website URL, e.g. chess-results.com, info64.org  | [optional] 
**Regulations** | Pointer to **string** | External URL to the official tournament regulations.  | [optional] 

## Methods

### NewBroadcastTourInfo

`func NewBroadcastTourInfo() *BroadcastTourInfo`

NewBroadcastTourInfo instantiates a new BroadcastTourInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTourInfoWithDefaults

`func NewBroadcastTourInfoWithDefaults() *BroadcastTourInfo`

NewBroadcastTourInfoWithDefaults instantiates a new BroadcastTourInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *BroadcastTourInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BroadcastTourInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BroadcastTourInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BroadcastTourInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTc

`func (o *BroadcastTourInfo) GetTc() string`

GetTc returns the Tc field if non-nil, zero value otherwise.

### GetTcOk

`func (o *BroadcastTourInfo) GetTcOk() (*string, bool)`

GetTcOk returns a tuple with the Tc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTc

`func (o *BroadcastTourInfo) SetTc(v string)`

SetTc sets Tc field to given value.

### HasTc

`func (o *BroadcastTourInfo) HasTc() bool`

HasTc returns a boolean if a field has been set.

### GetFideTC

`func (o *BroadcastTourInfo) GetFideTC() FideTimeControl`

GetFideTC returns the FideTC field if non-nil, zero value otherwise.

### GetFideTCOk

`func (o *BroadcastTourInfo) GetFideTCOk() (*FideTimeControl, bool)`

GetFideTCOk returns a tuple with the FideTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideTC

`func (o *BroadcastTourInfo) SetFideTC(v FideTimeControl)`

SetFideTC sets FideTC field to given value.

### HasFideTC

`func (o *BroadcastTourInfo) HasFideTC() bool`

HasFideTC returns a boolean if a field has been set.

### GetTimeZone

`func (o *BroadcastTourInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BroadcastTourInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BroadcastTourInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BroadcastTourInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocation

`func (o *BroadcastTourInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BroadcastTourInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BroadcastTourInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BroadcastTourInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPlayers

`func (o *BroadcastTourInfo) GetPlayers() string`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastTourInfo) GetPlayersOk() (*string, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastTourInfo) SetPlayers(v string)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *BroadcastTourInfo) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetWebsite

`func (o *BroadcastTourInfo) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *BroadcastTourInfo) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *BroadcastTourInfo) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *BroadcastTourInfo) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetStandings

`func (o *BroadcastTourInfo) GetStandings() string`

GetStandings returns the Standings field if non-nil, zero value otherwise.

### GetStandingsOk

`func (o *BroadcastTourInfo) GetStandingsOk() (*string, bool)`

GetStandingsOk returns a tuple with the Standings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandings

`func (o *BroadcastTourInfo) SetStandings(v string)`

SetStandings sets Standings field to given value.

### HasStandings

`func (o *BroadcastTourInfo) HasStandings() bool`

HasStandings returns a boolean if a field has been set.

### GetRegulations

`func (o *BroadcastTourInfo) GetRegulations() string`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *BroadcastTourInfo) GetRegulationsOk() (*string, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *BroadcastTourInfo) SetRegulations(v string)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *BroadcastTourInfo) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


