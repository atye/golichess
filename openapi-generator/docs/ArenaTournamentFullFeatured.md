# ArenaTournamentFullFeatured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fen** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**White** | Pointer to [**ArenaTournamentFullFeaturedWhite**](ArenaTournamentFullFeaturedWhite.md) |  | [optional] 
**Black** | Pointer to [**ArenaTournamentFullFeaturedWhite**](ArenaTournamentFullFeaturedWhite.md) |  | [optional] 
**C** | Pointer to [**ArenaTournamentFullFeaturedC**](ArenaTournamentFullFeaturedC.md) |  | [optional] 

## Methods

### NewArenaTournamentFullFeatured

`func NewArenaTournamentFullFeatured() *ArenaTournamentFullFeatured`

NewArenaTournamentFullFeatured instantiates a new ArenaTournamentFullFeatured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaTournamentFullFeaturedWithDefaults

`func NewArenaTournamentFullFeaturedWithDefaults() *ArenaTournamentFullFeatured`

NewArenaTournamentFullFeaturedWithDefaults instantiates a new ArenaTournamentFullFeatured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArenaTournamentFullFeatured) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArenaTournamentFullFeatured) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArenaTournamentFullFeatured) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArenaTournamentFullFeatured) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFen

`func (o *ArenaTournamentFullFeatured) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ArenaTournamentFullFeatured) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ArenaTournamentFullFeatured) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *ArenaTournamentFullFeatured) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetOrientation

`func (o *ArenaTournamentFullFeatured) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ArenaTournamentFullFeatured) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ArenaTournamentFullFeatured) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ArenaTournamentFullFeatured) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetColor

`func (o *ArenaTournamentFullFeatured) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ArenaTournamentFullFeatured) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ArenaTournamentFullFeatured) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ArenaTournamentFullFeatured) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLastMove

`func (o *ArenaTournamentFullFeatured) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *ArenaTournamentFullFeatured) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *ArenaTournamentFullFeatured) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *ArenaTournamentFullFeatured) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetWhite

`func (o *ArenaTournamentFullFeatured) GetWhite() ArenaTournamentFullFeaturedWhite`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *ArenaTournamentFullFeatured) GetWhiteOk() (*ArenaTournamentFullFeaturedWhite, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *ArenaTournamentFullFeatured) SetWhite(v ArenaTournamentFullFeaturedWhite)`

SetWhite sets White field to given value.

### HasWhite

`func (o *ArenaTournamentFullFeatured) HasWhite() bool`

HasWhite returns a boolean if a field has been set.

### GetBlack

`func (o *ArenaTournamentFullFeatured) GetBlack() ArenaTournamentFullFeaturedWhite`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *ArenaTournamentFullFeatured) GetBlackOk() (*ArenaTournamentFullFeaturedWhite, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *ArenaTournamentFullFeatured) SetBlack(v ArenaTournamentFullFeaturedWhite)`

SetBlack sets Black field to given value.

### HasBlack

`func (o *ArenaTournamentFullFeatured) HasBlack() bool`

HasBlack returns a boolean if a field has been set.

### GetC

`func (o *ArenaTournamentFullFeatured) GetC() ArenaTournamentFullFeaturedC`

GetC returns the C field if non-nil, zero value otherwise.

### GetCOk

`func (o *ArenaTournamentFullFeatured) GetCOk() (*ArenaTournamentFullFeaturedC, bool)`

GetCOk returns a tuple with the C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC

`func (o *ArenaTournamentFullFeatured) SetC(v ArenaTournamentFullFeaturedC)`

SetC sets C field to given value.

### HasC

`func (o *ArenaTournamentFullFeatured) HasC() bool`

HasC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


