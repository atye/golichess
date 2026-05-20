/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-generator

import (
	"encoding/json"
)

// checks if the Account200ResponsePrefs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account200ResponsePrefs{}

// Account200ResponsePrefs struct for Account200ResponsePrefs
type Account200ResponsePrefs struct {
	Dark *bool `json:"dark,omitempty"`
	Transp *bool `json:"transp,omitempty"`
	BgImg *string `json:"bgImg,omitempty"`
	Is3d *bool `json:"is3d,omitempty"`
	Theme *string `json:"theme,omitempty"`
	PieceSet *string `json:"pieceSet,omitempty"`
	Theme3d *string `json:"theme3d,omitempty"`
	PieceSet3d *string `json:"pieceSet3d,omitempty"`
	SoundSet *string `json:"soundSet,omitempty"`
	Blindfold *int32 `json:"blindfold,omitempty"`
	// 1 = Never, 2 = When premoving, 3 = Always
	AutoQueen *int32 `json:"autoQueen,omitempty"`
	// 0 = Never, 2 = When time remaining < 30 seconds,  3 = Always
	AutoThreefold *int32 `json:"autoThreefold,omitempty"`
	// 1 = Never, 2 = In casual games only, 3 = Always
	Takeback *int32 `json:"takeback,omitempty"`
	// 1 = Never, 2 = In casual games only, 3 = Always
	Moretime *int32 `json:"moretime,omitempty"`
	// 0 = Never, 1 = When remaining time less than 10 seconds, 2 = Always
	ClockTenths *int32 `json:"clockTenths,omitempty"`
	ClockBar *bool `json:"clockBar,omitempty"`
	ClockSound *bool `json:"clockSound,omitempty"`
	Premove *bool `json:"premove,omitempty"`
	// 0 = None, 1 = Fast, 2 = Normal, 3 = Slow
	Animation *int32 `json:"animation,omitempty"`
	// 0 = Chess piece symbol, 1 = KQRBN Letter
	PieceNotation *int32 `json:"pieceNotation,omitempty"`
	Captured *bool `json:"captured,omitempty"`
	Follow *bool `json:"follow,omitempty"`
	Highlight *bool `json:"highlight,omitempty"`
	Destination *bool `json:"destination,omitempty"`
	// 0 = No, 1 = Inside the board, 2 = Outside the board, 3 = All squares
	Coords *int32 `json:"coords,omitempty"`
	Replay *int32 `json:"replay,omitempty"`
	Challenge *int32 `json:"challenge,omitempty"`
	Message *int32 `json:"message,omitempty"`
	SubmitMove *int32 `json:"submitMove,omitempty"`
	// 1 = Confirm resignation and draw offers, 0 = Do not confirm
	ConfirmResign *int32 `json:"confirmResign,omitempty"`
	InsightShare *int32 `json:"insightShare,omitempty"`
	// 1 = input moves with the keyboard
	KeyboardMove *int32 `json:"keyboardMove,omitempty"`
	VoiceMove *bool `json:"voiceMove,omitempty"`
	// 0 = No, 1 = yes, 2 = in-game only
	Zen *int32 `json:"zen,omitempty"`
	// 0 = Hide ratings, 1 = Show ratings, 2 = Show ratings except in-game
	Ratings *int32 `json:"ratings,omitempty"`
	MoveEvent *int32 `json:"moveEvent,omitempty"`
	// 0 = Move king two squares, 1 = Move king onto rook
	RookCastle *int32 `json:"rookCastle,omitempty"`
	// Show player flairs
	Flairs *bool `json:"flairs,omitempty"`
	// 0 = No, 1 = When losing, 2 = When losing or drawing
	SayGG *int32 `json:"sayGG,omitempty"`
}

// NewAccount200ResponsePrefs instantiates a new Account200ResponsePrefs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount200ResponsePrefs() *Account200ResponsePrefs {
	this := Account200ResponsePrefs{}
	return &this
}

// NewAccount200ResponsePrefsWithDefaults instantiates a new Account200ResponsePrefs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccount200ResponsePrefsWithDefaults() *Account200ResponsePrefs {
	this := Account200ResponsePrefs{}
	return &this
}

// GetDark returns the Dark field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetDark() bool {
	if o == nil || IsNil(o.Dark) {
		var ret bool
		return ret
	}
	return *o.Dark
}

// GetDarkOk returns a tuple with the Dark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetDarkOk() (*bool, bool) {
	if o == nil || IsNil(o.Dark) {
		return nil, false
	}
	return o.Dark, true
}

// HasDark returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasDark() bool {
	if o != nil && !IsNil(o.Dark) {
		return true
	}

	return false
}

// SetDark gets a reference to the given bool and assigns it to the Dark field.
func (o *Account200ResponsePrefs) SetDark(v bool) {
	o.Dark = &v
}

// GetTransp returns the Transp field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetTransp() bool {
	if o == nil || IsNil(o.Transp) {
		var ret bool
		return ret
	}
	return *o.Transp
}

// GetTranspOk returns a tuple with the Transp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetTranspOk() (*bool, bool) {
	if o == nil || IsNil(o.Transp) {
		return nil, false
	}
	return o.Transp, true
}

// HasTransp returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasTransp() bool {
	if o != nil && !IsNil(o.Transp) {
		return true
	}

	return false
}

// SetTransp gets a reference to the given bool and assigns it to the Transp field.
func (o *Account200ResponsePrefs) SetTransp(v bool) {
	o.Transp = &v
}

// GetBgImg returns the BgImg field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetBgImg() string {
	if o == nil || IsNil(o.BgImg) {
		var ret string
		return ret
	}
	return *o.BgImg
}

// GetBgImgOk returns a tuple with the BgImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetBgImgOk() (*string, bool) {
	if o == nil || IsNil(o.BgImg) {
		return nil, false
	}
	return o.BgImg, true
}

// HasBgImg returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasBgImg() bool {
	if o != nil && !IsNil(o.BgImg) {
		return true
	}

	return false
}

// SetBgImg gets a reference to the given string and assigns it to the BgImg field.
func (o *Account200ResponsePrefs) SetBgImg(v string) {
	o.BgImg = &v
}

// GetIs3d returns the Is3d field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetIs3d() bool {
	if o == nil || IsNil(o.Is3d) {
		var ret bool
		return ret
	}
	return *o.Is3d
}

// GetIs3dOk returns a tuple with the Is3d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetIs3dOk() (*bool, bool) {
	if o == nil || IsNil(o.Is3d) {
		return nil, false
	}
	return o.Is3d, true
}

// HasIs3d returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasIs3d() bool {
	if o != nil && !IsNil(o.Is3d) {
		return true
	}

	return false
}

// SetIs3d gets a reference to the given bool and assigns it to the Is3d field.
func (o *Account200ResponsePrefs) SetIs3d(v bool) {
	o.Is3d = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *Account200ResponsePrefs) SetTheme(v string) {
	o.Theme = &v
}

// GetPieceSet returns the PieceSet field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetPieceSet() string {
	if o == nil || IsNil(o.PieceSet) {
		var ret string
		return ret
	}
	return *o.PieceSet
}

// GetPieceSetOk returns a tuple with the PieceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetPieceSetOk() (*string, bool) {
	if o == nil || IsNil(o.PieceSet) {
		return nil, false
	}
	return o.PieceSet, true
}

// HasPieceSet returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasPieceSet() bool {
	if o != nil && !IsNil(o.PieceSet) {
		return true
	}

	return false
}

// SetPieceSet gets a reference to the given string and assigns it to the PieceSet field.
func (o *Account200ResponsePrefs) SetPieceSet(v string) {
	o.PieceSet = &v
}

// GetTheme3d returns the Theme3d field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetTheme3d() string {
	if o == nil || IsNil(o.Theme3d) {
		var ret string
		return ret
	}
	return *o.Theme3d
}

// GetTheme3dOk returns a tuple with the Theme3d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetTheme3dOk() (*string, bool) {
	if o == nil || IsNil(o.Theme3d) {
		return nil, false
	}
	return o.Theme3d, true
}

// HasTheme3d returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasTheme3d() bool {
	if o != nil && !IsNil(o.Theme3d) {
		return true
	}

	return false
}

// SetTheme3d gets a reference to the given string and assigns it to the Theme3d field.
func (o *Account200ResponsePrefs) SetTheme3d(v string) {
	o.Theme3d = &v
}

// GetPieceSet3d returns the PieceSet3d field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetPieceSet3d() string {
	if o == nil || IsNil(o.PieceSet3d) {
		var ret string
		return ret
	}
	return *o.PieceSet3d
}

// GetPieceSet3dOk returns a tuple with the PieceSet3d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetPieceSet3dOk() (*string, bool) {
	if o == nil || IsNil(o.PieceSet3d) {
		return nil, false
	}
	return o.PieceSet3d, true
}

// HasPieceSet3d returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasPieceSet3d() bool {
	if o != nil && !IsNil(o.PieceSet3d) {
		return true
	}

	return false
}

// SetPieceSet3d gets a reference to the given string and assigns it to the PieceSet3d field.
func (o *Account200ResponsePrefs) SetPieceSet3d(v string) {
	o.PieceSet3d = &v
}

// GetSoundSet returns the SoundSet field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetSoundSet() string {
	if o == nil || IsNil(o.SoundSet) {
		var ret string
		return ret
	}
	return *o.SoundSet
}

// GetSoundSetOk returns a tuple with the SoundSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetSoundSetOk() (*string, bool) {
	if o == nil || IsNil(o.SoundSet) {
		return nil, false
	}
	return o.SoundSet, true
}

// HasSoundSet returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasSoundSet() bool {
	if o != nil && !IsNil(o.SoundSet) {
		return true
	}

	return false
}

// SetSoundSet gets a reference to the given string and assigns it to the SoundSet field.
func (o *Account200ResponsePrefs) SetSoundSet(v string) {
	o.SoundSet = &v
}

// GetBlindfold returns the Blindfold field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetBlindfold() int32 {
	if o == nil || IsNil(o.Blindfold) {
		var ret int32
		return ret
	}
	return *o.Blindfold
}

// GetBlindfoldOk returns a tuple with the Blindfold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetBlindfoldOk() (*int32, bool) {
	if o == nil || IsNil(o.Blindfold) {
		return nil, false
	}
	return o.Blindfold, true
}

// HasBlindfold returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasBlindfold() bool {
	if o != nil && !IsNil(o.Blindfold) {
		return true
	}

	return false
}

// SetBlindfold gets a reference to the given int32 and assigns it to the Blindfold field.
func (o *Account200ResponsePrefs) SetBlindfold(v int32) {
	o.Blindfold = &v
}

// GetAutoQueen returns the AutoQueen field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetAutoQueen() int32 {
	if o == nil || IsNil(o.AutoQueen) {
		var ret int32
		return ret
	}
	return *o.AutoQueen
}

// GetAutoQueenOk returns a tuple with the AutoQueen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetAutoQueenOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoQueen) {
		return nil, false
	}
	return o.AutoQueen, true
}

// HasAutoQueen returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasAutoQueen() bool {
	if o != nil && !IsNil(o.AutoQueen) {
		return true
	}

	return false
}

// SetAutoQueen gets a reference to the given int32 and assigns it to the AutoQueen field.
func (o *Account200ResponsePrefs) SetAutoQueen(v int32) {
	o.AutoQueen = &v
}

// GetAutoThreefold returns the AutoThreefold field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetAutoThreefold() int32 {
	if o == nil || IsNil(o.AutoThreefold) {
		var ret int32
		return ret
	}
	return *o.AutoThreefold
}

// GetAutoThreefoldOk returns a tuple with the AutoThreefold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetAutoThreefoldOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoThreefold) {
		return nil, false
	}
	return o.AutoThreefold, true
}

// HasAutoThreefold returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasAutoThreefold() bool {
	if o != nil && !IsNil(o.AutoThreefold) {
		return true
	}

	return false
}

// SetAutoThreefold gets a reference to the given int32 and assigns it to the AutoThreefold field.
func (o *Account200ResponsePrefs) SetAutoThreefold(v int32) {
	o.AutoThreefold = &v
}

// GetTakeback returns the Takeback field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetTakeback() int32 {
	if o == nil || IsNil(o.Takeback) {
		var ret int32
		return ret
	}
	return *o.Takeback
}

// GetTakebackOk returns a tuple with the Takeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetTakebackOk() (*int32, bool) {
	if o == nil || IsNil(o.Takeback) {
		return nil, false
	}
	return o.Takeback, true
}

// HasTakeback returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasTakeback() bool {
	if o != nil && !IsNil(o.Takeback) {
		return true
	}

	return false
}

// SetTakeback gets a reference to the given int32 and assigns it to the Takeback field.
func (o *Account200ResponsePrefs) SetTakeback(v int32) {
	o.Takeback = &v
}

// GetMoretime returns the Moretime field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetMoretime() int32 {
	if o == nil || IsNil(o.Moretime) {
		var ret int32
		return ret
	}
	return *o.Moretime
}

// GetMoretimeOk returns a tuple with the Moretime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetMoretimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Moretime) {
		return nil, false
	}
	return o.Moretime, true
}

// HasMoretime returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasMoretime() bool {
	if o != nil && !IsNil(o.Moretime) {
		return true
	}

	return false
}

// SetMoretime gets a reference to the given int32 and assigns it to the Moretime field.
func (o *Account200ResponsePrefs) SetMoretime(v int32) {
	o.Moretime = &v
}

// GetClockTenths returns the ClockTenths field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetClockTenths() int32 {
	if o == nil || IsNil(o.ClockTenths) {
		var ret int32
		return ret
	}
	return *o.ClockTenths
}

// GetClockTenthsOk returns a tuple with the ClockTenths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetClockTenthsOk() (*int32, bool) {
	if o == nil || IsNil(o.ClockTenths) {
		return nil, false
	}
	return o.ClockTenths, true
}

// HasClockTenths returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasClockTenths() bool {
	if o != nil && !IsNil(o.ClockTenths) {
		return true
	}

	return false
}

// SetClockTenths gets a reference to the given int32 and assigns it to the ClockTenths field.
func (o *Account200ResponsePrefs) SetClockTenths(v int32) {
	o.ClockTenths = &v
}

// GetClockBar returns the ClockBar field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetClockBar() bool {
	if o == nil || IsNil(o.ClockBar) {
		var ret bool
		return ret
	}
	return *o.ClockBar
}

// GetClockBarOk returns a tuple with the ClockBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetClockBarOk() (*bool, bool) {
	if o == nil || IsNil(o.ClockBar) {
		return nil, false
	}
	return o.ClockBar, true
}

// HasClockBar returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasClockBar() bool {
	if o != nil && !IsNil(o.ClockBar) {
		return true
	}

	return false
}

// SetClockBar gets a reference to the given bool and assigns it to the ClockBar field.
func (o *Account200ResponsePrefs) SetClockBar(v bool) {
	o.ClockBar = &v
}

// GetClockSound returns the ClockSound field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetClockSound() bool {
	if o == nil || IsNil(o.ClockSound) {
		var ret bool
		return ret
	}
	return *o.ClockSound
}

// GetClockSoundOk returns a tuple with the ClockSound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetClockSoundOk() (*bool, bool) {
	if o == nil || IsNil(o.ClockSound) {
		return nil, false
	}
	return o.ClockSound, true
}

// HasClockSound returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasClockSound() bool {
	if o != nil && !IsNil(o.ClockSound) {
		return true
	}

	return false
}

// SetClockSound gets a reference to the given bool and assigns it to the ClockSound field.
func (o *Account200ResponsePrefs) SetClockSound(v bool) {
	o.ClockSound = &v
}

// GetPremove returns the Premove field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetPremove() bool {
	if o == nil || IsNil(o.Premove) {
		var ret bool
		return ret
	}
	return *o.Premove
}

// GetPremoveOk returns a tuple with the Premove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetPremoveOk() (*bool, bool) {
	if o == nil || IsNil(o.Premove) {
		return nil, false
	}
	return o.Premove, true
}

// HasPremove returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasPremove() bool {
	if o != nil && !IsNil(o.Premove) {
		return true
	}

	return false
}

// SetPremove gets a reference to the given bool and assigns it to the Premove field.
func (o *Account200ResponsePrefs) SetPremove(v bool) {
	o.Premove = &v
}

// GetAnimation returns the Animation field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetAnimation() int32 {
	if o == nil || IsNil(o.Animation) {
		var ret int32
		return ret
	}
	return *o.Animation
}

// GetAnimationOk returns a tuple with the Animation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetAnimationOk() (*int32, bool) {
	if o == nil || IsNil(o.Animation) {
		return nil, false
	}
	return o.Animation, true
}

// HasAnimation returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasAnimation() bool {
	if o != nil && !IsNil(o.Animation) {
		return true
	}

	return false
}

// SetAnimation gets a reference to the given int32 and assigns it to the Animation field.
func (o *Account200ResponsePrefs) SetAnimation(v int32) {
	o.Animation = &v
}

// GetPieceNotation returns the PieceNotation field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetPieceNotation() int32 {
	if o == nil || IsNil(o.PieceNotation) {
		var ret int32
		return ret
	}
	return *o.PieceNotation
}

// GetPieceNotationOk returns a tuple with the PieceNotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetPieceNotationOk() (*int32, bool) {
	if o == nil || IsNil(o.PieceNotation) {
		return nil, false
	}
	return o.PieceNotation, true
}

// HasPieceNotation returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasPieceNotation() bool {
	if o != nil && !IsNil(o.PieceNotation) {
		return true
	}

	return false
}

// SetPieceNotation gets a reference to the given int32 and assigns it to the PieceNotation field.
func (o *Account200ResponsePrefs) SetPieceNotation(v int32) {
	o.PieceNotation = &v
}

// GetCaptured returns the Captured field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetCaptured() bool {
	if o == nil || IsNil(o.Captured) {
		var ret bool
		return ret
	}
	return *o.Captured
}

// GetCapturedOk returns a tuple with the Captured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetCapturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Captured) {
		return nil, false
	}
	return o.Captured, true
}

// HasCaptured returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasCaptured() bool {
	if o != nil && !IsNil(o.Captured) {
		return true
	}

	return false
}

// SetCaptured gets a reference to the given bool and assigns it to the Captured field.
func (o *Account200ResponsePrefs) SetCaptured(v bool) {
	o.Captured = &v
}

// GetFollow returns the Follow field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetFollow() bool {
	if o == nil || IsNil(o.Follow) {
		var ret bool
		return ret
	}
	return *o.Follow
}

// GetFollowOk returns a tuple with the Follow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetFollowOk() (*bool, bool) {
	if o == nil || IsNil(o.Follow) {
		return nil, false
	}
	return o.Follow, true
}

// HasFollow returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasFollow() bool {
	if o != nil && !IsNil(o.Follow) {
		return true
	}

	return false
}

// SetFollow gets a reference to the given bool and assigns it to the Follow field.
func (o *Account200ResponsePrefs) SetFollow(v bool) {
	o.Follow = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetHighlight() bool {
	if o == nil || IsNil(o.Highlight) {
		var ret bool
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetHighlightOk() (*bool, bool) {
	if o == nil || IsNil(o.Highlight) {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given bool and assigns it to the Highlight field.
func (o *Account200ResponsePrefs) SetHighlight(v bool) {
	o.Highlight = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetDestination() bool {
	if o == nil || IsNil(o.Destination) {
		var ret bool
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetDestinationOk() (*bool, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given bool and assigns it to the Destination field.
func (o *Account200ResponsePrefs) SetDestination(v bool) {
	o.Destination = &v
}

// GetCoords returns the Coords field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetCoords() int32 {
	if o == nil || IsNil(o.Coords) {
		var ret int32
		return ret
	}
	return *o.Coords
}

// GetCoordsOk returns a tuple with the Coords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetCoordsOk() (*int32, bool) {
	if o == nil || IsNil(o.Coords) {
		return nil, false
	}
	return o.Coords, true
}

// HasCoords returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasCoords() bool {
	if o != nil && !IsNil(o.Coords) {
		return true
	}

	return false
}

// SetCoords gets a reference to the given int32 and assigns it to the Coords field.
func (o *Account200ResponsePrefs) SetCoords(v int32) {
	o.Coords = &v
}

// GetReplay returns the Replay field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetReplay() int32 {
	if o == nil || IsNil(o.Replay) {
		var ret int32
		return ret
	}
	return *o.Replay
}

// GetReplayOk returns a tuple with the Replay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetReplayOk() (*int32, bool) {
	if o == nil || IsNil(o.Replay) {
		return nil, false
	}
	return o.Replay, true
}

// HasReplay returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasReplay() bool {
	if o != nil && !IsNil(o.Replay) {
		return true
	}

	return false
}

// SetReplay gets a reference to the given int32 and assigns it to the Replay field.
func (o *Account200ResponsePrefs) SetReplay(v int32) {
	o.Replay = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetChallenge() int32 {
	if o == nil || IsNil(o.Challenge) {
		var ret int32
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetChallengeOk() (*int32, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given int32 and assigns it to the Challenge field.
func (o *Account200ResponsePrefs) SetChallenge(v int32) {
	o.Challenge = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetMessage() int32 {
	if o == nil || IsNil(o.Message) {
		var ret int32
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetMessageOk() (*int32, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given int32 and assigns it to the Message field.
func (o *Account200ResponsePrefs) SetMessage(v int32) {
	o.Message = &v
}

// GetSubmitMove returns the SubmitMove field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetSubmitMove() int32 {
	if o == nil || IsNil(o.SubmitMove) {
		var ret int32
		return ret
	}
	return *o.SubmitMove
}

// GetSubmitMoveOk returns a tuple with the SubmitMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetSubmitMoveOk() (*int32, bool) {
	if o == nil || IsNil(o.SubmitMove) {
		return nil, false
	}
	return o.SubmitMove, true
}

// HasSubmitMove returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasSubmitMove() bool {
	if o != nil && !IsNil(o.SubmitMove) {
		return true
	}

	return false
}

// SetSubmitMove gets a reference to the given int32 and assigns it to the SubmitMove field.
func (o *Account200ResponsePrefs) SetSubmitMove(v int32) {
	o.SubmitMove = &v
}

// GetConfirmResign returns the ConfirmResign field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetConfirmResign() int32 {
	if o == nil || IsNil(o.ConfirmResign) {
		var ret int32
		return ret
	}
	return *o.ConfirmResign
}

// GetConfirmResignOk returns a tuple with the ConfirmResign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetConfirmResignOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfirmResign) {
		return nil, false
	}
	return o.ConfirmResign, true
}

// HasConfirmResign returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasConfirmResign() bool {
	if o != nil && !IsNil(o.ConfirmResign) {
		return true
	}

	return false
}

// SetConfirmResign gets a reference to the given int32 and assigns it to the ConfirmResign field.
func (o *Account200ResponsePrefs) SetConfirmResign(v int32) {
	o.ConfirmResign = &v
}

// GetInsightShare returns the InsightShare field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetInsightShare() int32 {
	if o == nil || IsNil(o.InsightShare) {
		var ret int32
		return ret
	}
	return *o.InsightShare
}

// GetInsightShareOk returns a tuple with the InsightShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetInsightShareOk() (*int32, bool) {
	if o == nil || IsNil(o.InsightShare) {
		return nil, false
	}
	return o.InsightShare, true
}

// HasInsightShare returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasInsightShare() bool {
	if o != nil && !IsNil(o.InsightShare) {
		return true
	}

	return false
}

// SetInsightShare gets a reference to the given int32 and assigns it to the InsightShare field.
func (o *Account200ResponsePrefs) SetInsightShare(v int32) {
	o.InsightShare = &v
}

// GetKeyboardMove returns the KeyboardMove field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetKeyboardMove() int32 {
	if o == nil || IsNil(o.KeyboardMove) {
		var ret int32
		return ret
	}
	return *o.KeyboardMove
}

// GetKeyboardMoveOk returns a tuple with the KeyboardMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetKeyboardMoveOk() (*int32, bool) {
	if o == nil || IsNil(o.KeyboardMove) {
		return nil, false
	}
	return o.KeyboardMove, true
}

// HasKeyboardMove returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasKeyboardMove() bool {
	if o != nil && !IsNil(o.KeyboardMove) {
		return true
	}

	return false
}

// SetKeyboardMove gets a reference to the given int32 and assigns it to the KeyboardMove field.
func (o *Account200ResponsePrefs) SetKeyboardMove(v int32) {
	o.KeyboardMove = &v
}

// GetVoiceMove returns the VoiceMove field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetVoiceMove() bool {
	if o == nil || IsNil(o.VoiceMove) {
		var ret bool
		return ret
	}
	return *o.VoiceMove
}

// GetVoiceMoveOk returns a tuple with the VoiceMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetVoiceMoveOk() (*bool, bool) {
	if o == nil || IsNil(o.VoiceMove) {
		return nil, false
	}
	return o.VoiceMove, true
}

// HasVoiceMove returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasVoiceMove() bool {
	if o != nil && !IsNil(o.VoiceMove) {
		return true
	}

	return false
}

// SetVoiceMove gets a reference to the given bool and assigns it to the VoiceMove field.
func (o *Account200ResponsePrefs) SetVoiceMove(v bool) {
	o.VoiceMove = &v
}

// GetZen returns the Zen field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetZen() int32 {
	if o == nil || IsNil(o.Zen) {
		var ret int32
		return ret
	}
	return *o.Zen
}

// GetZenOk returns a tuple with the Zen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetZenOk() (*int32, bool) {
	if o == nil || IsNil(o.Zen) {
		return nil, false
	}
	return o.Zen, true
}

// HasZen returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasZen() bool {
	if o != nil && !IsNil(o.Zen) {
		return true
	}

	return false
}

// SetZen gets a reference to the given int32 and assigns it to the Zen field.
func (o *Account200ResponsePrefs) SetZen(v int32) {
	o.Zen = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetRatings() int32 {
	if o == nil || IsNil(o.Ratings) {
		var ret int32
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetRatingsOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given int32 and assigns it to the Ratings field.
func (o *Account200ResponsePrefs) SetRatings(v int32) {
	o.Ratings = &v
}

// GetMoveEvent returns the MoveEvent field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetMoveEvent() int32 {
	if o == nil || IsNil(o.MoveEvent) {
		var ret int32
		return ret
	}
	return *o.MoveEvent
}

// GetMoveEventOk returns a tuple with the MoveEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetMoveEventOk() (*int32, bool) {
	if o == nil || IsNil(o.MoveEvent) {
		return nil, false
	}
	return o.MoveEvent, true
}

// HasMoveEvent returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasMoveEvent() bool {
	if o != nil && !IsNil(o.MoveEvent) {
		return true
	}

	return false
}

// SetMoveEvent gets a reference to the given int32 and assigns it to the MoveEvent field.
func (o *Account200ResponsePrefs) SetMoveEvent(v int32) {
	o.MoveEvent = &v
}

// GetRookCastle returns the RookCastle field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetRookCastle() int32 {
	if o == nil || IsNil(o.RookCastle) {
		var ret int32
		return ret
	}
	return *o.RookCastle
}

// GetRookCastleOk returns a tuple with the RookCastle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetRookCastleOk() (*int32, bool) {
	if o == nil || IsNil(o.RookCastle) {
		return nil, false
	}
	return o.RookCastle, true
}

// HasRookCastle returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasRookCastle() bool {
	if o != nil && !IsNil(o.RookCastle) {
		return true
	}

	return false
}

// SetRookCastle gets a reference to the given int32 and assigns it to the RookCastle field.
func (o *Account200ResponsePrefs) SetRookCastle(v int32) {
	o.RookCastle = &v
}

// GetFlairs returns the Flairs field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetFlairs() bool {
	if o == nil || IsNil(o.Flairs) {
		var ret bool
		return ret
	}
	return *o.Flairs
}

// GetFlairsOk returns a tuple with the Flairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetFlairsOk() (*bool, bool) {
	if o == nil || IsNil(o.Flairs) {
		return nil, false
	}
	return o.Flairs, true
}

// HasFlairs returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasFlairs() bool {
	if o != nil && !IsNil(o.Flairs) {
		return true
	}

	return false
}

// SetFlairs gets a reference to the given bool and assigns it to the Flairs field.
func (o *Account200ResponsePrefs) SetFlairs(v bool) {
	o.Flairs = &v
}

// GetSayGG returns the SayGG field value if set, zero value otherwise.
func (o *Account200ResponsePrefs) GetSayGG() int32 {
	if o == nil || IsNil(o.SayGG) {
		var ret int32
		return ret
	}
	return *o.SayGG
}

// GetSayGGOk returns a tuple with the SayGG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account200ResponsePrefs) GetSayGGOk() (*int32, bool) {
	if o == nil || IsNil(o.SayGG) {
		return nil, false
	}
	return o.SayGG, true
}

// HasSayGG returns a boolean if a field has been set.
func (o *Account200ResponsePrefs) HasSayGG() bool {
	if o != nil && !IsNil(o.SayGG) {
		return true
	}

	return false
}

// SetSayGG gets a reference to the given int32 and assigns it to the SayGG field.
func (o *Account200ResponsePrefs) SetSayGG(v int32) {
	o.SayGG = &v
}

func (o Account200ResponsePrefs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account200ResponsePrefs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dark) {
		toSerialize["dark"] = o.Dark
	}
	if !IsNil(o.Transp) {
		toSerialize["transp"] = o.Transp
	}
	if !IsNil(o.BgImg) {
		toSerialize["bgImg"] = o.BgImg
	}
	if !IsNil(o.Is3d) {
		toSerialize["is3d"] = o.Is3d
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.PieceSet) {
		toSerialize["pieceSet"] = o.PieceSet
	}
	if !IsNil(o.Theme3d) {
		toSerialize["theme3d"] = o.Theme3d
	}
	if !IsNil(o.PieceSet3d) {
		toSerialize["pieceSet3d"] = o.PieceSet3d
	}
	if !IsNil(o.SoundSet) {
		toSerialize["soundSet"] = o.SoundSet
	}
	if !IsNil(o.Blindfold) {
		toSerialize["blindfold"] = o.Blindfold
	}
	if !IsNil(o.AutoQueen) {
		toSerialize["autoQueen"] = o.AutoQueen
	}
	if !IsNil(o.AutoThreefold) {
		toSerialize["autoThreefold"] = o.AutoThreefold
	}
	if !IsNil(o.Takeback) {
		toSerialize["takeback"] = o.Takeback
	}
	if !IsNil(o.Moretime) {
		toSerialize["moretime"] = o.Moretime
	}
	if !IsNil(o.ClockTenths) {
		toSerialize["clockTenths"] = o.ClockTenths
	}
	if !IsNil(o.ClockBar) {
		toSerialize["clockBar"] = o.ClockBar
	}
	if !IsNil(o.ClockSound) {
		toSerialize["clockSound"] = o.ClockSound
	}
	if !IsNil(o.Premove) {
		toSerialize["premove"] = o.Premove
	}
	if !IsNil(o.Animation) {
		toSerialize["animation"] = o.Animation
	}
	if !IsNil(o.PieceNotation) {
		toSerialize["pieceNotation"] = o.PieceNotation
	}
	if !IsNil(o.Captured) {
		toSerialize["captured"] = o.Captured
	}
	if !IsNil(o.Follow) {
		toSerialize["follow"] = o.Follow
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Coords) {
		toSerialize["coords"] = o.Coords
	}
	if !IsNil(o.Replay) {
		toSerialize["replay"] = o.Replay
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.SubmitMove) {
		toSerialize["submitMove"] = o.SubmitMove
	}
	if !IsNil(o.ConfirmResign) {
		toSerialize["confirmResign"] = o.ConfirmResign
	}
	if !IsNil(o.InsightShare) {
		toSerialize["insightShare"] = o.InsightShare
	}
	if !IsNil(o.KeyboardMove) {
		toSerialize["keyboardMove"] = o.KeyboardMove
	}
	if !IsNil(o.VoiceMove) {
		toSerialize["voiceMove"] = o.VoiceMove
	}
	if !IsNil(o.Zen) {
		toSerialize["zen"] = o.Zen
	}
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if !IsNil(o.MoveEvent) {
		toSerialize["moveEvent"] = o.MoveEvent
	}
	if !IsNil(o.RookCastle) {
		toSerialize["rookCastle"] = o.RookCastle
	}
	if !IsNil(o.Flairs) {
		toSerialize["flairs"] = o.Flairs
	}
	if !IsNil(o.SayGG) {
		toSerialize["sayGG"] = o.SayGG
	}
	return toSerialize, nil
}

type NullableAccount200ResponsePrefs struct {
	value *Account200ResponsePrefs
	isSet bool
}

func (v NullableAccount200ResponsePrefs) Get() *Account200ResponsePrefs {
	return v.value
}

func (v *NullableAccount200ResponsePrefs) Set(val *Account200ResponsePrefs) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount200ResponsePrefs) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount200ResponsePrefs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount200ResponsePrefs(val *Account200ResponsePrefs) *NullableAccount200ResponsePrefs {
	return &NullableAccount200ResponsePrefs{value: val, isSet: true}
}

func (v NullableAccount200ResponsePrefs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount200ResponsePrefs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


