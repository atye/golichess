# Account200ResponsePrefs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dark** | Pointer to **bool** |  | [optional] 
**Transp** | Pointer to **bool** |  | [optional] 
**BgImg** | Pointer to **string** |  | [optional] 
**Is3d** | Pointer to **bool** |  | [optional] 
**Theme** | Pointer to **string** |  | [optional] 
**PieceSet** | Pointer to **string** |  | [optional] 
**Theme3d** | Pointer to **string** |  | [optional] 
**PieceSet3d** | Pointer to **string** |  | [optional] 
**SoundSet** | Pointer to **string** |  | [optional] 
**Blindfold** | Pointer to **int32** |  | [optional] 
**AutoQueen** | Pointer to **int32** | 1 &#x3D; Never, 2 &#x3D; When premoving, 3 &#x3D; Always | [optional] 
**AutoThreefold** | Pointer to **int32** | 0 &#x3D; Never, 2 &#x3D; When time remaining &lt; 30 seconds,  3 &#x3D; Always | [optional] 
**Takeback** | Pointer to **int32** | 1 &#x3D; Never, 2 &#x3D; In casual games only, 3 &#x3D; Always | [optional] 
**Moretime** | Pointer to **int32** | 1 &#x3D; Never, 2 &#x3D; In casual games only, 3 &#x3D; Always | [optional] 
**ClockTenths** | Pointer to **int32** | 0 &#x3D; Never, 1 &#x3D; When remaining time less than 10 seconds, 2 &#x3D; Always | [optional] 
**ClockBar** | Pointer to **bool** |  | [optional] 
**ClockSound** | Pointer to **bool** |  | [optional] 
**Premove** | Pointer to **bool** |  | [optional] 
**Animation** | Pointer to **int32** | 0 &#x3D; None, 1 &#x3D; Fast, 2 &#x3D; Normal, 3 &#x3D; Slow | [optional] 
**PieceNotation** | Pointer to **int32** | 0 &#x3D; Chess piece symbol, 1 &#x3D; KQRBN Letter | [optional] 
**Captured** | Pointer to **bool** |  | [optional] 
**Follow** | Pointer to **bool** |  | [optional] 
**Highlight** | Pointer to **bool** |  | [optional] 
**Destination** | Pointer to **bool** |  | [optional] 
**Coords** | Pointer to **int32** | 0 &#x3D; No, 1 &#x3D; Inside the board, 2 &#x3D; Outside the board, 3 &#x3D; All squares | [optional] 
**Replay** | Pointer to **int32** |  | [optional] 
**Challenge** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **int32** |  | [optional] 
**SubmitMove** | Pointer to **int32** |  | [optional] 
**ConfirmResign** | Pointer to **int32** | 1 &#x3D; Confirm resignation and draw offers, 0 &#x3D; Do not confirm | [optional] 
**InsightShare** | Pointer to **int32** |  | [optional] 
**KeyboardMove** | Pointer to **int32** | 1 &#x3D; input moves with the keyboard | [optional] 
**VoiceMove** | Pointer to **bool** |  | [optional] 
**Zen** | Pointer to **int32** | 0 &#x3D; No, 1 &#x3D; yes, 2 &#x3D; in-game only | [optional] 
**Ratings** | Pointer to **int32** | 0 &#x3D; Hide ratings, 1 &#x3D; Show ratings, 2 &#x3D; Show ratings except in-game | [optional] 
**MoveEvent** | Pointer to **int32** |  | [optional] 
**RookCastle** | Pointer to **int32** | 0 &#x3D; Move king two squares, 1 &#x3D; Move king onto rook | [optional] 
**Flairs** | Pointer to **bool** | Show player flairs | [optional] 
**SayGG** | Pointer to **int32** | 0 &#x3D; No, 1 &#x3D; When losing, 2 &#x3D; When losing or drawing | [optional] 

## Methods

### NewAccount200ResponsePrefs

`func NewAccount200ResponsePrefs() *Account200ResponsePrefs`

NewAccount200ResponsePrefs instantiates a new Account200ResponsePrefs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccount200ResponsePrefsWithDefaults

`func NewAccount200ResponsePrefsWithDefaults() *Account200ResponsePrefs`

NewAccount200ResponsePrefsWithDefaults instantiates a new Account200ResponsePrefs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDark

`func (o *Account200ResponsePrefs) GetDark() bool`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *Account200ResponsePrefs) GetDarkOk() (*bool, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *Account200ResponsePrefs) SetDark(v bool)`

SetDark sets Dark field to given value.

### HasDark

`func (o *Account200ResponsePrefs) HasDark() bool`

HasDark returns a boolean if a field has been set.

### GetTransp

`func (o *Account200ResponsePrefs) GetTransp() bool`

GetTransp returns the Transp field if non-nil, zero value otherwise.

### GetTranspOk

`func (o *Account200ResponsePrefs) GetTranspOk() (*bool, bool)`

GetTranspOk returns a tuple with the Transp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransp

`func (o *Account200ResponsePrefs) SetTransp(v bool)`

SetTransp sets Transp field to given value.

### HasTransp

`func (o *Account200ResponsePrefs) HasTransp() bool`

HasTransp returns a boolean if a field has been set.

### GetBgImg

`func (o *Account200ResponsePrefs) GetBgImg() string`

GetBgImg returns the BgImg field if non-nil, zero value otherwise.

### GetBgImgOk

`func (o *Account200ResponsePrefs) GetBgImgOk() (*string, bool)`

GetBgImgOk returns a tuple with the BgImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImg

`func (o *Account200ResponsePrefs) SetBgImg(v string)`

SetBgImg sets BgImg field to given value.

### HasBgImg

`func (o *Account200ResponsePrefs) HasBgImg() bool`

HasBgImg returns a boolean if a field has been set.

### GetIs3d

`func (o *Account200ResponsePrefs) GetIs3d() bool`

GetIs3d returns the Is3d field if non-nil, zero value otherwise.

### GetIs3dOk

`func (o *Account200ResponsePrefs) GetIs3dOk() (*bool, bool)`

GetIs3dOk returns a tuple with the Is3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs3d

`func (o *Account200ResponsePrefs) SetIs3d(v bool)`

SetIs3d sets Is3d field to given value.

### HasIs3d

`func (o *Account200ResponsePrefs) HasIs3d() bool`

HasIs3d returns a boolean if a field has been set.

### GetTheme

`func (o *Account200ResponsePrefs) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Account200ResponsePrefs) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Account200ResponsePrefs) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Account200ResponsePrefs) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetPieceSet

`func (o *Account200ResponsePrefs) GetPieceSet() string`

GetPieceSet returns the PieceSet field if non-nil, zero value otherwise.

### GetPieceSetOk

`func (o *Account200ResponsePrefs) GetPieceSetOk() (*string, bool)`

GetPieceSetOk returns a tuple with the PieceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceSet

`func (o *Account200ResponsePrefs) SetPieceSet(v string)`

SetPieceSet sets PieceSet field to given value.

### HasPieceSet

`func (o *Account200ResponsePrefs) HasPieceSet() bool`

HasPieceSet returns a boolean if a field has been set.

### GetTheme3d

`func (o *Account200ResponsePrefs) GetTheme3d() string`

GetTheme3d returns the Theme3d field if non-nil, zero value otherwise.

### GetTheme3dOk

`func (o *Account200ResponsePrefs) GetTheme3dOk() (*string, bool)`

GetTheme3dOk returns a tuple with the Theme3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme3d

`func (o *Account200ResponsePrefs) SetTheme3d(v string)`

SetTheme3d sets Theme3d field to given value.

### HasTheme3d

`func (o *Account200ResponsePrefs) HasTheme3d() bool`

HasTheme3d returns a boolean if a field has been set.

### GetPieceSet3d

`func (o *Account200ResponsePrefs) GetPieceSet3d() string`

GetPieceSet3d returns the PieceSet3d field if non-nil, zero value otherwise.

### GetPieceSet3dOk

`func (o *Account200ResponsePrefs) GetPieceSet3dOk() (*string, bool)`

GetPieceSet3dOk returns a tuple with the PieceSet3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceSet3d

`func (o *Account200ResponsePrefs) SetPieceSet3d(v string)`

SetPieceSet3d sets PieceSet3d field to given value.

### HasPieceSet3d

`func (o *Account200ResponsePrefs) HasPieceSet3d() bool`

HasPieceSet3d returns a boolean if a field has been set.

### GetSoundSet

`func (o *Account200ResponsePrefs) GetSoundSet() string`

GetSoundSet returns the SoundSet field if non-nil, zero value otherwise.

### GetSoundSetOk

`func (o *Account200ResponsePrefs) GetSoundSetOk() (*string, bool)`

GetSoundSetOk returns a tuple with the SoundSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundSet

`func (o *Account200ResponsePrefs) SetSoundSet(v string)`

SetSoundSet sets SoundSet field to given value.

### HasSoundSet

`func (o *Account200ResponsePrefs) HasSoundSet() bool`

HasSoundSet returns a boolean if a field has been set.

### GetBlindfold

`func (o *Account200ResponsePrefs) GetBlindfold() int32`

GetBlindfold returns the Blindfold field if non-nil, zero value otherwise.

### GetBlindfoldOk

`func (o *Account200ResponsePrefs) GetBlindfoldOk() (*int32, bool)`

GetBlindfoldOk returns a tuple with the Blindfold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlindfold

`func (o *Account200ResponsePrefs) SetBlindfold(v int32)`

SetBlindfold sets Blindfold field to given value.

### HasBlindfold

`func (o *Account200ResponsePrefs) HasBlindfold() bool`

HasBlindfold returns a boolean if a field has been set.

### GetAutoQueen

`func (o *Account200ResponsePrefs) GetAutoQueen() int32`

GetAutoQueen returns the AutoQueen field if non-nil, zero value otherwise.

### GetAutoQueenOk

`func (o *Account200ResponsePrefs) GetAutoQueenOk() (*int32, bool)`

GetAutoQueenOk returns a tuple with the AutoQueen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoQueen

`func (o *Account200ResponsePrefs) SetAutoQueen(v int32)`

SetAutoQueen sets AutoQueen field to given value.

### HasAutoQueen

`func (o *Account200ResponsePrefs) HasAutoQueen() bool`

HasAutoQueen returns a boolean if a field has been set.

### GetAutoThreefold

`func (o *Account200ResponsePrefs) GetAutoThreefold() int32`

GetAutoThreefold returns the AutoThreefold field if non-nil, zero value otherwise.

### GetAutoThreefoldOk

`func (o *Account200ResponsePrefs) GetAutoThreefoldOk() (*int32, bool)`

GetAutoThreefoldOk returns a tuple with the AutoThreefold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoThreefold

`func (o *Account200ResponsePrefs) SetAutoThreefold(v int32)`

SetAutoThreefold sets AutoThreefold field to given value.

### HasAutoThreefold

`func (o *Account200ResponsePrefs) HasAutoThreefold() bool`

HasAutoThreefold returns a boolean if a field has been set.

### GetTakeback

`func (o *Account200ResponsePrefs) GetTakeback() int32`

GetTakeback returns the Takeback field if non-nil, zero value otherwise.

### GetTakebackOk

`func (o *Account200ResponsePrefs) GetTakebackOk() (*int32, bool)`

GetTakebackOk returns a tuple with the Takeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeback

`func (o *Account200ResponsePrefs) SetTakeback(v int32)`

SetTakeback sets Takeback field to given value.

### HasTakeback

`func (o *Account200ResponsePrefs) HasTakeback() bool`

HasTakeback returns a boolean if a field has been set.

### GetMoretime

`func (o *Account200ResponsePrefs) GetMoretime() int32`

GetMoretime returns the Moretime field if non-nil, zero value otherwise.

### GetMoretimeOk

`func (o *Account200ResponsePrefs) GetMoretimeOk() (*int32, bool)`

GetMoretimeOk returns a tuple with the Moretime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoretime

`func (o *Account200ResponsePrefs) SetMoretime(v int32)`

SetMoretime sets Moretime field to given value.

### HasMoretime

`func (o *Account200ResponsePrefs) HasMoretime() bool`

HasMoretime returns a boolean if a field has been set.

### GetClockTenths

`func (o *Account200ResponsePrefs) GetClockTenths() int32`

GetClockTenths returns the ClockTenths field if non-nil, zero value otherwise.

### GetClockTenthsOk

`func (o *Account200ResponsePrefs) GetClockTenthsOk() (*int32, bool)`

GetClockTenthsOk returns a tuple with the ClockTenths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockTenths

`func (o *Account200ResponsePrefs) SetClockTenths(v int32)`

SetClockTenths sets ClockTenths field to given value.

### HasClockTenths

`func (o *Account200ResponsePrefs) HasClockTenths() bool`

HasClockTenths returns a boolean if a field has been set.

### GetClockBar

`func (o *Account200ResponsePrefs) GetClockBar() bool`

GetClockBar returns the ClockBar field if non-nil, zero value otherwise.

### GetClockBarOk

`func (o *Account200ResponsePrefs) GetClockBarOk() (*bool, bool)`

GetClockBarOk returns a tuple with the ClockBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockBar

`func (o *Account200ResponsePrefs) SetClockBar(v bool)`

SetClockBar sets ClockBar field to given value.

### HasClockBar

`func (o *Account200ResponsePrefs) HasClockBar() bool`

HasClockBar returns a boolean if a field has been set.

### GetClockSound

`func (o *Account200ResponsePrefs) GetClockSound() bool`

GetClockSound returns the ClockSound field if non-nil, zero value otherwise.

### GetClockSoundOk

`func (o *Account200ResponsePrefs) GetClockSoundOk() (*bool, bool)`

GetClockSoundOk returns a tuple with the ClockSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSound

`func (o *Account200ResponsePrefs) SetClockSound(v bool)`

SetClockSound sets ClockSound field to given value.

### HasClockSound

`func (o *Account200ResponsePrefs) HasClockSound() bool`

HasClockSound returns a boolean if a field has been set.

### GetPremove

`func (o *Account200ResponsePrefs) GetPremove() bool`

GetPremove returns the Premove field if non-nil, zero value otherwise.

### GetPremoveOk

`func (o *Account200ResponsePrefs) GetPremoveOk() (*bool, bool)`

GetPremoveOk returns a tuple with the Premove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremove

`func (o *Account200ResponsePrefs) SetPremove(v bool)`

SetPremove sets Premove field to given value.

### HasPremove

`func (o *Account200ResponsePrefs) HasPremove() bool`

HasPremove returns a boolean if a field has been set.

### GetAnimation

`func (o *Account200ResponsePrefs) GetAnimation() int32`

GetAnimation returns the Animation field if non-nil, zero value otherwise.

### GetAnimationOk

`func (o *Account200ResponsePrefs) GetAnimationOk() (*int32, bool)`

GetAnimationOk returns a tuple with the Animation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimation

`func (o *Account200ResponsePrefs) SetAnimation(v int32)`

SetAnimation sets Animation field to given value.

### HasAnimation

`func (o *Account200ResponsePrefs) HasAnimation() bool`

HasAnimation returns a boolean if a field has been set.

### GetPieceNotation

`func (o *Account200ResponsePrefs) GetPieceNotation() int32`

GetPieceNotation returns the PieceNotation field if non-nil, zero value otherwise.

### GetPieceNotationOk

`func (o *Account200ResponsePrefs) GetPieceNotationOk() (*int32, bool)`

GetPieceNotationOk returns a tuple with the PieceNotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceNotation

`func (o *Account200ResponsePrefs) SetPieceNotation(v int32)`

SetPieceNotation sets PieceNotation field to given value.

### HasPieceNotation

`func (o *Account200ResponsePrefs) HasPieceNotation() bool`

HasPieceNotation returns a boolean if a field has been set.

### GetCaptured

`func (o *Account200ResponsePrefs) GetCaptured() bool`

GetCaptured returns the Captured field if non-nil, zero value otherwise.

### GetCapturedOk

`func (o *Account200ResponsePrefs) GetCapturedOk() (*bool, bool)`

GetCapturedOk returns a tuple with the Captured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptured

`func (o *Account200ResponsePrefs) SetCaptured(v bool)`

SetCaptured sets Captured field to given value.

### HasCaptured

`func (o *Account200ResponsePrefs) HasCaptured() bool`

HasCaptured returns a boolean if a field has been set.

### GetFollow

`func (o *Account200ResponsePrefs) GetFollow() bool`

GetFollow returns the Follow field if non-nil, zero value otherwise.

### GetFollowOk

`func (o *Account200ResponsePrefs) GetFollowOk() (*bool, bool)`

GetFollowOk returns a tuple with the Follow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollow

`func (o *Account200ResponsePrefs) SetFollow(v bool)`

SetFollow sets Follow field to given value.

### HasFollow

`func (o *Account200ResponsePrefs) HasFollow() bool`

HasFollow returns a boolean if a field has been set.

### GetHighlight

`func (o *Account200ResponsePrefs) GetHighlight() bool`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *Account200ResponsePrefs) GetHighlightOk() (*bool, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *Account200ResponsePrefs) SetHighlight(v bool)`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *Account200ResponsePrefs) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetDestination

`func (o *Account200ResponsePrefs) GetDestination() bool`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Account200ResponsePrefs) GetDestinationOk() (*bool, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Account200ResponsePrefs) SetDestination(v bool)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Account200ResponsePrefs) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetCoords

`func (o *Account200ResponsePrefs) GetCoords() int32`

GetCoords returns the Coords field if non-nil, zero value otherwise.

### GetCoordsOk

`func (o *Account200ResponsePrefs) GetCoordsOk() (*int32, bool)`

GetCoordsOk returns a tuple with the Coords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoords

`func (o *Account200ResponsePrefs) SetCoords(v int32)`

SetCoords sets Coords field to given value.

### HasCoords

`func (o *Account200ResponsePrefs) HasCoords() bool`

HasCoords returns a boolean if a field has been set.

### GetReplay

`func (o *Account200ResponsePrefs) GetReplay() int32`

GetReplay returns the Replay field if non-nil, zero value otherwise.

### GetReplayOk

`func (o *Account200ResponsePrefs) GetReplayOk() (*int32, bool)`

GetReplayOk returns a tuple with the Replay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplay

`func (o *Account200ResponsePrefs) SetReplay(v int32)`

SetReplay sets Replay field to given value.

### HasReplay

`func (o *Account200ResponsePrefs) HasReplay() bool`

HasReplay returns a boolean if a field has been set.

### GetChallenge

`func (o *Account200ResponsePrefs) GetChallenge() int32`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *Account200ResponsePrefs) GetChallengeOk() (*int32, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *Account200ResponsePrefs) SetChallenge(v int32)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *Account200ResponsePrefs) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetMessage

`func (o *Account200ResponsePrefs) GetMessage() int32`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Account200ResponsePrefs) GetMessageOk() (*int32, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Account200ResponsePrefs) SetMessage(v int32)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Account200ResponsePrefs) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSubmitMove

`func (o *Account200ResponsePrefs) GetSubmitMove() int32`

GetSubmitMove returns the SubmitMove field if non-nil, zero value otherwise.

### GetSubmitMoveOk

`func (o *Account200ResponsePrefs) GetSubmitMoveOk() (*int32, bool)`

GetSubmitMoveOk returns a tuple with the SubmitMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitMove

`func (o *Account200ResponsePrefs) SetSubmitMove(v int32)`

SetSubmitMove sets SubmitMove field to given value.

### HasSubmitMove

`func (o *Account200ResponsePrefs) HasSubmitMove() bool`

HasSubmitMove returns a boolean if a field has been set.

### GetConfirmResign

`func (o *Account200ResponsePrefs) GetConfirmResign() int32`

GetConfirmResign returns the ConfirmResign field if non-nil, zero value otherwise.

### GetConfirmResignOk

`func (o *Account200ResponsePrefs) GetConfirmResignOk() (*int32, bool)`

GetConfirmResignOk returns a tuple with the ConfirmResign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmResign

`func (o *Account200ResponsePrefs) SetConfirmResign(v int32)`

SetConfirmResign sets ConfirmResign field to given value.

### HasConfirmResign

`func (o *Account200ResponsePrefs) HasConfirmResign() bool`

HasConfirmResign returns a boolean if a field has been set.

### GetInsightShare

`func (o *Account200ResponsePrefs) GetInsightShare() int32`

GetInsightShare returns the InsightShare field if non-nil, zero value otherwise.

### GetInsightShareOk

`func (o *Account200ResponsePrefs) GetInsightShareOk() (*int32, bool)`

GetInsightShareOk returns a tuple with the InsightShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightShare

`func (o *Account200ResponsePrefs) SetInsightShare(v int32)`

SetInsightShare sets InsightShare field to given value.

### HasInsightShare

`func (o *Account200ResponsePrefs) HasInsightShare() bool`

HasInsightShare returns a boolean if a field has been set.

### GetKeyboardMove

`func (o *Account200ResponsePrefs) GetKeyboardMove() int32`

GetKeyboardMove returns the KeyboardMove field if non-nil, zero value otherwise.

### GetKeyboardMoveOk

`func (o *Account200ResponsePrefs) GetKeyboardMoveOk() (*int32, bool)`

GetKeyboardMoveOk returns a tuple with the KeyboardMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboardMove

`func (o *Account200ResponsePrefs) SetKeyboardMove(v int32)`

SetKeyboardMove sets KeyboardMove field to given value.

### HasKeyboardMove

`func (o *Account200ResponsePrefs) HasKeyboardMove() bool`

HasKeyboardMove returns a boolean if a field has been set.

### GetVoiceMove

`func (o *Account200ResponsePrefs) GetVoiceMove() bool`

GetVoiceMove returns the VoiceMove field if non-nil, zero value otherwise.

### GetVoiceMoveOk

`func (o *Account200ResponsePrefs) GetVoiceMoveOk() (*bool, bool)`

GetVoiceMoveOk returns a tuple with the VoiceMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMove

`func (o *Account200ResponsePrefs) SetVoiceMove(v bool)`

SetVoiceMove sets VoiceMove field to given value.

### HasVoiceMove

`func (o *Account200ResponsePrefs) HasVoiceMove() bool`

HasVoiceMove returns a boolean if a field has been set.

### GetZen

`func (o *Account200ResponsePrefs) GetZen() int32`

GetZen returns the Zen field if non-nil, zero value otherwise.

### GetZenOk

`func (o *Account200ResponsePrefs) GetZenOk() (*int32, bool)`

GetZenOk returns a tuple with the Zen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZen

`func (o *Account200ResponsePrefs) SetZen(v int32)`

SetZen sets Zen field to given value.

### HasZen

`func (o *Account200ResponsePrefs) HasZen() bool`

HasZen returns a boolean if a field has been set.

### GetRatings

`func (o *Account200ResponsePrefs) GetRatings() int32`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Account200ResponsePrefs) GetRatingsOk() (*int32, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Account200ResponsePrefs) SetRatings(v int32)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Account200ResponsePrefs) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMoveEvent

`func (o *Account200ResponsePrefs) GetMoveEvent() int32`

GetMoveEvent returns the MoveEvent field if non-nil, zero value otherwise.

### GetMoveEventOk

`func (o *Account200ResponsePrefs) GetMoveEventOk() (*int32, bool)`

GetMoveEventOk returns a tuple with the MoveEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveEvent

`func (o *Account200ResponsePrefs) SetMoveEvent(v int32)`

SetMoveEvent sets MoveEvent field to given value.

### HasMoveEvent

`func (o *Account200ResponsePrefs) HasMoveEvent() bool`

HasMoveEvent returns a boolean if a field has been set.

### GetRookCastle

`func (o *Account200ResponsePrefs) GetRookCastle() int32`

GetRookCastle returns the RookCastle field if non-nil, zero value otherwise.

### GetRookCastleOk

`func (o *Account200ResponsePrefs) GetRookCastleOk() (*int32, bool)`

GetRookCastleOk returns a tuple with the RookCastle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRookCastle

`func (o *Account200ResponsePrefs) SetRookCastle(v int32)`

SetRookCastle sets RookCastle field to given value.

### HasRookCastle

`func (o *Account200ResponsePrefs) HasRookCastle() bool`

HasRookCastle returns a boolean if a field has been set.

### GetFlairs

`func (o *Account200ResponsePrefs) GetFlairs() bool`

GetFlairs returns the Flairs field if non-nil, zero value otherwise.

### GetFlairsOk

`func (o *Account200ResponsePrefs) GetFlairsOk() (*bool, bool)`

GetFlairsOk returns a tuple with the Flairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairs

`func (o *Account200ResponsePrefs) SetFlairs(v bool)`

SetFlairs sets Flairs field to given value.

### HasFlairs

`func (o *Account200ResponsePrefs) HasFlairs() bool`

HasFlairs returns a boolean if a field has been set.

### GetSayGG

`func (o *Account200ResponsePrefs) GetSayGG() int32`

GetSayGG returns the SayGG field if non-nil, zero value otherwise.

### GetSayGGOk

`func (o *Account200ResponsePrefs) GetSayGGOk() (*int32, bool)`

GetSayGGOk returns a tuple with the SayGG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSayGG

`func (o *Account200ResponsePrefs) SetSayGG(v int32)`

SetSayGG sets SayGG field to given value.

### HasSayGG

`func (o *Account200ResponsePrefs) HasSayGG() bool`

HasSayGG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


