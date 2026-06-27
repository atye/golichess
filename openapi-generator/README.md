# Go API client for openapigenerator

# Introduction
Welcome to the reference for the Lichess API! Lichess is free/libre,
open-source chess server powered by volunteers and donations.
- Get help in the [Lichess Discord channel](https://discord.gg/lichess)
- API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/)
- API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui)
- [Contribute to this documentation on Github](https://github.com/lichess-org/api)
- Check out [Lichess widgets to embed in your website](https://lichess.org/developers)
- [Download all Lichess rated games](https://database.lichess.org/)
- [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles)
- [Download all evaluated positions](https://database.lichess.org/#evals)

## Endpoint
All requests go to `https://lichess.org` (unless otherwise specified).

## Clients
- [Python general API](https://github.com/lichess-org/berserk)
- [MicroPython general API](https://github.com/mkomon/uberserk)
- [Python general API - async](https://pypi.org/project/async-lichess-sdk)
- [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot)
- [Python Board API for Certabo](https://github.com/haklein/certabo-lichess)
- [Java general API](https://github.com/tors42/chariot)
- [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine)
- [Rust general API](https://github.com/obazin/litchee)
- [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET)
- [.NET general API](https://github.com/Dblike/LichessSharp)

## Rate limiting
All requests are rate limited using various strategies,
to ensure the API remains responsive for everyone.
Only make one request at a time.
If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429),
you have exceded one of the rate limits.
In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer.
Reduce your request frequency before retrying.

## Streaming with ND-JSON
Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.

Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.

## Authentication
### Which authentication method is right for me?
[Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)

### Personal Access Token
Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow.
- [Generate a personal access token](https://lichess.org/account/oauth/token)
- `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"`
- [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)

### Token Security
- Keep your tokens secret. Do not share them in public repositories or public forums.
- Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account.
- Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users.
- If you suspect a token has been compromised, revoke it immediately.

To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).

### Authorization Code Flow with PKCE
The authorization code flow with PKCE allows your users to **login with Lichess**.
Lichess supports unregistered and public clients (no client authentication, choose any unique client id).
The only accepted code challenge method is `S256`.
Access tokens are long-lived (expect one year), unless they are revoked.
Refresh tokens are not supported.

See the [documentation for the OAuth endpoints](#tag/OAuth) or
the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.

- [Demo app](https://lichess-org.github.io/api-demo/)
- [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app)
- [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask)
- [Java example](https://github.com/tors42/lichess-oauth-pkce-app)
- [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)

#### Real life examples
- [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants))
- [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour))
- [English Chess Federation](https://ecf.octoknight.com/)
- [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)

### Token format
Access tokens and authorization codes match `^[A-Za-z0-9_]+$`.
The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters.
By convention tokens have a recognizable prefix, but do not rely on this.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.149
- Package version: 1.0.0
- Generator version: 7.23.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://lichess.org/api](https://lichess.org/api)

## Installation

Import the package in a go file in your project and run `go mod tidy`:

```go
import openapigenerator "github.com/atye/golichess/openapigenerator"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapigenerator.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapigenerator.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapigenerator.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapigenerator.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapigenerator.ContextOperationServerIndices` and `openapigenerator.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapigenerator.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapigenerator.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://lichess.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**Account**](docs/AccountAPI.md#account) | **Get** /api/account/preferences | Get my preferences
*AccountAPI* | [**AccountEmail**](docs/AccountAPI.md#accountemail) | **Get** /api/account/email | Get my email address
*AccountAPI* | [**AccountKid**](docs/AccountAPI.md#accountkid) | **Get** /api/account/kid | Get my kid mode status
*AccountAPI* | [**AccountKidPost**](docs/AccountAPI.md#accountkidpost) | **Post** /api/account/kid | Set my kid mode status
*AccountAPI* | [**AccountMe**](docs/AccountAPI.md#accountme) | **Get** /api/account | Get my profile
*AccountAPI* | [**Timeline**](docs/AccountAPI.md#timeline) | **Get** /api/timeline | Get my timeline
*AnalysisAPI* | [**ApiCloudEval**](docs/AnalysisAPI.md#apicloudeval) | **Get** /api/cloud-eval | Get cloud evaluation of a position.
*BoardAPI* | [**ApiBoardSeek**](docs/BoardAPI.md#apiboardseek) | **Post** /api/board/seek | Create a seek
*BoardAPI* | [**ApiStreamEvent**](docs/BoardAPI.md#apistreamevent) | **Get** /api/stream/event | Stream incoming events
*BoardAPI* | [**BoardGameAbort**](docs/BoardAPI.md#boardgameabort) | **Post** /api/board/game/{gameId}/abort | Abort a game
*BoardAPI* | [**BoardGameBerserk**](docs/BoardAPI.md#boardgameberserk) | **Post** /api/board/game/{gameId}/berserk | Berserk a tournament game
*BoardAPI* | [**BoardGameChatGet**](docs/BoardAPI.md#boardgamechatget) | **Get** /api/board/game/{gameId}/chat | Fetch the player chat
*BoardAPI* | [**BoardGameChatPost**](docs/BoardAPI.md#boardgamechatpost) | **Post** /api/board/game/{gameId}/chat | Write in the chat
*BoardAPI* | [**BoardGameClaimDraw**](docs/BoardAPI.md#boardgameclaimdraw) | **Post** /api/board/game/{gameId}/claim-draw | Claim draw of a game
*BoardAPI* | [**BoardGameClaimVictory**](docs/BoardAPI.md#boardgameclaimvictory) | **Post** /api/board/game/{gameId}/claim-victory | Claim victory of a game
*BoardAPI* | [**BoardGameDraw**](docs/BoardAPI.md#boardgamedraw) | **Post** /api/board/game/{gameId}/draw/{accept} | Handle draw offers
*BoardAPI* | [**BoardGameMove**](docs/BoardAPI.md#boardgamemove) | **Post** /api/board/game/{gameId}/move/{move} | Make a Board move
*BoardAPI* | [**BoardGameResign**](docs/BoardAPI.md#boardgameresign) | **Post** /api/board/game/{gameId}/resign | Resign a game
*BoardAPI* | [**BoardGameStream**](docs/BoardAPI.md#boardgamestream) | **Get** /api/board/game/stream/{gameId} | Stream Board game state
*BoardAPI* | [**BoardGameTakeback**](docs/BoardAPI.md#boardgametakeback) | **Post** /api/board/game/{gameId}/takeback/{accept} | Handle takeback offers
*BotAPI* | [**ApiBotOnline**](docs/BotAPI.md#apibotonline) | **Get** /api/bot/online | Get online bots
*BotAPI* | [**ApiStreamEvent**](docs/BotAPI.md#apistreamevent) | **Get** /api/stream/event | Stream incoming events
*BotAPI* | [**BotAccountUpgrade**](docs/BotAPI.md#botaccountupgrade) | **Post** /api/bot/account/upgrade | Upgrade to Bot account
*BotAPI* | [**BotGameAbort**](docs/BotAPI.md#botgameabort) | **Post** /api/bot/game/{gameId}/abort | Abort a game
*BotAPI* | [**BotGameChat**](docs/BotAPI.md#botgamechat) | **Post** /api/bot/game/{gameId}/chat | Write in the chat
*BotAPI* | [**BotGameChatGet**](docs/BotAPI.md#botgamechatget) | **Get** /api/bot/game/{gameId}/chat | Fetch the game chat
*BotAPI* | [**BotGameClaimDraw**](docs/BotAPI.md#botgameclaimdraw) | **Post** /api/bot/game/{gameId}/claim-draw | Claim draw of a game
*BotAPI* | [**BotGameClaimVictory**](docs/BotAPI.md#botgameclaimvictory) | **Post** /api/bot/game/{gameId}/claim-victory | Claim victory of a game
*BotAPI* | [**BotGameDraw**](docs/BotAPI.md#botgamedraw) | **Post** /api/bot/game/{gameId}/draw/{accept} | Handle draw offers
*BotAPI* | [**BotGameMove**](docs/BotAPI.md#botgamemove) | **Post** /api/bot/game/{gameId}/move/{move} | Make a Bot move
*BotAPI* | [**BotGameResign**](docs/BotAPI.md#botgameresign) | **Post** /api/bot/game/{gameId}/resign | Resign a game
*BotAPI* | [**BotGameStream**](docs/BotAPI.md#botgamestream) | **Get** /api/bot/game/stream/{gameId} | Stream Bot game state
*BotAPI* | [**BotGameTakeback**](docs/BotAPI.md#botgametakeback) | **Post** /api/bot/game/{gameId}/takeback/{accept} | Handle takeback offers
*BroadcastsAPI* | [**BroadcastAllRoundsPgn**](docs/BroadcastsAPI.md#broadcastallroundspgn) | **Get** /api/broadcast/{broadcastTournamentId}.pgn | Export all rounds as PGN
*BroadcastsAPI* | [**BroadcastMyRoundsGet**](docs/BroadcastsAPI.md#broadcastmyroundsget) | **Get** /api/broadcast/my-rounds | Get your broadcast rounds
*BroadcastsAPI* | [**BroadcastPlayerGet**](docs/BroadcastsAPI.md#broadcastplayerget) | **Get** /broadcast/{broadcastTournamentId}/players/{playerId} | Get a player of a broadcast
*BroadcastsAPI* | [**BroadcastPlayersGet**](docs/BroadcastsAPI.md#broadcastplayersget) | **Get** /broadcast/{broadcastTournamentId}/players | Get players of a broadcast
*BroadcastsAPI* | [**BroadcastPush**](docs/BroadcastsAPI.md#broadcastpush) | **Post** /api/broadcast/round/{broadcastRoundId}/push | Push PGN to a broadcast round
*BroadcastsAPI* | [**BroadcastRoundCreate**](docs/BroadcastsAPI.md#broadcastroundcreate) | **Post** /broadcast/{broadcastTournamentId}/new | Create a broadcast round
*BroadcastsAPI* | [**BroadcastRoundGet**](docs/BroadcastsAPI.md#broadcastroundget) | **Get** /api/broadcast/{broadcastTournamentSlug}/{broadcastRoundSlug}/{broadcastRoundId} | Get a broadcast round
*BroadcastsAPI* | [**BroadcastRoundPgn**](docs/BroadcastsAPI.md#broadcastroundpgn) | **Get** /api/broadcast/round/{broadcastRoundId}.pgn | Export one round as PGN
*BroadcastsAPI* | [**BroadcastRoundReset**](docs/BroadcastsAPI.md#broadcastroundreset) | **Post** /api/broadcast/round/{broadcastRoundId}/reset | Reset a broadcast round
*BroadcastsAPI* | [**BroadcastRoundUpdate**](docs/BroadcastsAPI.md#broadcastroundupdate) | **Post** /broadcast/round/{broadcastRoundId}/edit | Update a broadcast round
*BroadcastsAPI* | [**BroadcastStreamGroupPgn**](docs/BroadcastsAPI.md#broadcaststreamgrouppgn) | **Get** /api/stream/broadcast/group/{broadcastGroupId}.pgn | Stream ongoing broadcast rounds of a group as PGN
*BroadcastsAPI* | [**BroadcastStreamRoundPgn**](docs/BroadcastsAPI.md#broadcaststreamroundpgn) | **Get** /api/stream/broadcast/round/{broadcastRoundId}.pgn | Stream an ongoing broadcast round as PGN
*BroadcastsAPI* | [**BroadcastTeamLeaderboardGet**](docs/BroadcastsAPI.md#broadcastteamleaderboardget) | **Get** /broadcast/{broadcastTournamentId}/teams/standings | Get the team leaderboard of a broadcast
*BroadcastsAPI* | [**BroadcastTourCreate**](docs/BroadcastsAPI.md#broadcasttourcreate) | **Post** /broadcast/new | Create a broadcast tournament
*BroadcastsAPI* | [**BroadcastTourGet**](docs/BroadcastsAPI.md#broadcasttourget) | **Get** /api/broadcast/{broadcastTournamentId} | Get a broadcast tournament
*BroadcastsAPI* | [**BroadcastTourUpdate**](docs/BroadcastsAPI.md#broadcasttourupdate) | **Post** /broadcast/{broadcastTournamentId}/edit | Update your broadcast tournament
*BroadcastsAPI* | [**BroadcastsByUser**](docs/BroadcastsAPI.md#broadcastsbyuser) | **Get** /api/broadcast/by/{username} | Get broadcasts created by a user
*BroadcastsAPI* | [**BroadcastsOfficial**](docs/BroadcastsAPI.md#broadcastsofficial) | **Get** /api/broadcast | Get official broadcasts
*BroadcastsAPI* | [**BroadcastsSearch**](docs/BroadcastsAPI.md#broadcastssearch) | **Get** /api/broadcast/search | Search broadcasts
*BroadcastsAPI* | [**BroadcastsTop**](docs/BroadcastsAPI.md#broadcaststop) | **Get** /api/broadcast/top | Get paginated top broadcast previews
*BulkPairingsAPI* | [**BulkPairingCreate**](docs/BulkPairingsAPI.md#bulkpairingcreate) | **Post** /api/bulk-pairing | Create a bulk pairing
*BulkPairingsAPI* | [**BulkPairingDelete**](docs/BulkPairingsAPI.md#bulkpairingdelete) | **Delete** /api/bulk-pairing/{id} | Cancel a bulk pairing
*BulkPairingsAPI* | [**BulkPairingGet**](docs/BulkPairingsAPI.md#bulkpairingget) | **Get** /api/bulk-pairing/{id} | Show a bulk pairing
*BulkPairingsAPI* | [**BulkPairingIdGamesGet**](docs/BulkPairingsAPI.md#bulkpairingidgamesget) | **Get** /api/bulk-pairing/{id}/games | Export games of a bulk pairing
*BulkPairingsAPI* | [**BulkPairingList**](docs/BulkPairingsAPI.md#bulkpairinglist) | **Get** /api/bulk-pairing | View your bulk pairings
*BulkPairingsAPI* | [**BulkPairingStartClocks**](docs/BulkPairingsAPI.md#bulkpairingstartclocks) | **Post** /api/bulk-pairing/{id}/start-clocks | Manually start clocks
*ChallengesAPI* | [**AdminChallengeTokens**](docs/ChallengesAPI.md#adminchallengetokens) | **Post** /api/token/admin-challenge | Admin challenge tokens
*ChallengesAPI* | [**ChallengeAccept**](docs/ChallengesAPI.md#challengeaccept) | **Post** /api/challenge/{challengeId}/accept | Accept a challenge
*ChallengesAPI* | [**ChallengeAi**](docs/ChallengesAPI.md#challengeai) | **Post** /api/challenge/ai | Challenge the AI
*ChallengesAPI* | [**ChallengeCancel**](docs/ChallengesAPI.md#challengecancel) | **Post** /api/challenge/{challengeId}/cancel | Cancel a challenge
*ChallengesAPI* | [**ChallengeCreate**](docs/ChallengesAPI.md#challengecreate) | **Post** /api/challenge/{username} | Create a challenge
*ChallengesAPI* | [**ChallengeDecline**](docs/ChallengesAPI.md#challengedecline) | **Post** /api/challenge/{challengeId}/decline | Decline a challenge
*ChallengesAPI* | [**ChallengeList**](docs/ChallengesAPI.md#challengelist) | **Get** /api/challenge | List your challenges
*ChallengesAPI* | [**ChallengeOpen**](docs/ChallengesAPI.md#challengeopen) | **Post** /api/challenge/open | Open-ended challenge
*ChallengesAPI* | [**ChallengeShow**](docs/ChallengesAPI.md#challengeshow) | **Get** /api/challenge/{challengeId}/show | Show one challenge
*ChallengesAPI* | [**ChallengeStartClocks**](docs/ChallengesAPI.md#challengestartclocks) | **Post** /api/challenge/{gameId}/start-clocks | Start clocks of a game
*ChallengesAPI* | [**RoundAddTime**](docs/ChallengesAPI.md#roundaddtime) | **Post** /api/round/{gameId}/add-time/{seconds} | Add time to the opponent clock
*ExternalEngineAPI* | [**ApiExternalEngineAcquire**](docs/ExternalEngineAPI.md#apiexternalengineacquire) | **Post** /api/external-engine/work | Acquire analysis request
*ExternalEngineAPI* | [**ApiExternalEngineAnalyse**](docs/ExternalEngineAPI.md#apiexternalengineanalyse) | **Post** /api/external-engine/{id}/analyse | Analyse with external engine
*ExternalEngineAPI* | [**ApiExternalEngineCreate**](docs/ExternalEngineAPI.md#apiexternalenginecreate) | **Post** /api/external-engine | Create external engine
*ExternalEngineAPI* | [**ApiExternalEngineDelete**](docs/ExternalEngineAPI.md#apiexternalenginedelete) | **Delete** /api/external-engine/{id} | Delete external engine
*ExternalEngineAPI* | [**ApiExternalEngineGet**](docs/ExternalEngineAPI.md#apiexternalengineget) | **Get** /api/external-engine/{id} | Get external engine
*ExternalEngineAPI* | [**ApiExternalEngineList**](docs/ExternalEngineAPI.md#apiexternalenginelist) | **Get** /api/external-engine | List external engines
*ExternalEngineAPI* | [**ApiExternalEnginePut**](docs/ExternalEngineAPI.md#apiexternalengineput) | **Put** /api/external-engine/{id} | Update external engine
*ExternalEngineAPI* | [**ApiExternalEngineSubmit**](docs/ExternalEngineAPI.md#apiexternalenginesubmit) | **Post** /api/external-engine/work/{id} | Answer analysis request
*FIDEAPI* | [**FidePlayerGet**](docs/FIDEAPI.md#fideplayerget) | **Get** /api/fide/player/{playerId} | Get a FIDE player
*FIDEAPI* | [**FidePlayerRatings**](docs/FIDEAPI.md#fideplayerratings) | **Get** /api/fide/player/{playerId}/ratings | Get ratings history of a FIDE player
*FIDEAPI* | [**FidePlayerSearch**](docs/FIDEAPI.md#fideplayersearch) | **Get** /api/fide/player | Search FIDE players
*GamesAPI* | [**ApiAccountPlaying**](docs/GamesAPI.md#apiaccountplaying) | **Get** /api/account/playing | Get my ongoing games
*GamesAPI* | [**ApiExportBookmarks**](docs/GamesAPI.md#apiexportbookmarks) | **Get** /api/games/export/bookmarks | Export your bookmarked games
*GamesAPI* | [**ApiGamesUser**](docs/GamesAPI.md#apigamesuser) | **Get** /api/games/user/{username} | Export games of a user
*GamesAPI* | [**ApiImportedGamesUser**](docs/GamesAPI.md#apiimportedgamesuser) | **Get** /api/games/export/imports | Export your imported games
*GamesAPI* | [**ApiUserCurrentGame**](docs/GamesAPI.md#apiusercurrentgame) | **Get** /api/user/{username}/current-game | Export ongoing game of a user
*GamesAPI* | [**GameChatGet**](docs/GamesAPI.md#gamechatget) | **Get** /game/{gameId}/chat | Fetch the spectator game chat
*GamesAPI* | [**GameImport**](docs/GamesAPI.md#gameimport) | **Post** /api/import | Import one game
*GamesAPI* | [**GamePgn**](docs/GamesAPI.md#gamepgn) | **Get** /game/export/{gameId} | Export one game
*GamesAPI* | [**GamesByIds**](docs/GamesAPI.md#gamesbyids) | **Post** /api/stream/games/{streamId} | Stream games by IDs
*GamesAPI* | [**GamesByIdsAdd**](docs/GamesAPI.md#gamesbyidsadd) | **Post** /api/stream/games/{streamId}/add | Add game IDs to stream
*GamesAPI* | [**GamesByUsers**](docs/GamesAPI.md#gamesbyusers) | **Post** /api/stream/games-by-users | Stream games of users
*GamesAPI* | [**GamesExportIds**](docs/GamesAPI.md#gamesexportids) | **Post** /api/games/export/_ids | Export games by IDs
*GamesAPI* | [**StreamGame**](docs/GamesAPI.md#streamgame) | **Get** /api/stream/game/{id} | Stream moves of a game
*MessagingAPI* | [**InboxUsername**](docs/MessagingAPI.md#inboxusername) | **Post** /inbox/{username} | Send a private message
*OAuthAPI* | [**ApiToken**](docs/OAuthAPI.md#apitoken) | **Post** /api/token | Obtain access token
*OAuthAPI* | [**ApiTokenDelete**](docs/OAuthAPI.md#apitokendelete) | **Delete** /api/token | Revoke access token
*OAuthAPI* | [**Oauth**](docs/OAuthAPI.md#oauth) | **Get** /oauth | Request authorization code
*OAuthAPI* | [**TokenTest**](docs/OAuthAPI.md#tokentest) | **Post** /api/token/test | Test multiple OAuth tokens
*OpeningExplorerAPI* | [**OpeningExplorerLichess**](docs/OpeningExplorerAPI.md#openingexplorerlichess) | **Get** /lichess | Lichess games
*OpeningExplorerAPI* | [**OpeningExplorerMaster**](docs/OpeningExplorerAPI.md#openingexplorermaster) | **Get** /masters | Masters database
*OpeningExplorerAPI* | [**OpeningExplorerMasterGame**](docs/OpeningExplorerAPI.md#openingexplorermastergame) | **Get** /masters/pgn/{gameId} | OTB master game
*OpeningExplorerAPI* | [**OpeningExplorerPlayer**](docs/OpeningExplorerAPI.md#openingexplorerplayer) | **Get** /player | Player games
*PuzzlesAPI* | [**ApiPuzzleActivity**](docs/PuzzlesAPI.md#apipuzzleactivity) | **Get** /api/puzzle/activity | Get your puzzle activity
*PuzzlesAPI* | [**ApiPuzzleBatchSelect**](docs/PuzzlesAPI.md#apipuzzlebatchselect) | **Get** /api/puzzle/batch/{angle} | Get multiple puzzles at once
*PuzzlesAPI* | [**ApiPuzzleBatchSolve**](docs/PuzzlesAPI.md#apipuzzlebatchsolve) | **Post** /api/puzzle/batch/{angle} | Solve multiple puzzles at once
*PuzzlesAPI* | [**ApiPuzzleDaily**](docs/PuzzlesAPI.md#apipuzzledaily) | **Get** /api/puzzle/daily | Get the daily puzzle
*PuzzlesAPI* | [**ApiPuzzleDashboard**](docs/PuzzlesAPI.md#apipuzzledashboard) | **Get** /api/puzzle/dashboard/{days} | Get your puzzle dashboard
*PuzzlesAPI* | [**ApiPuzzleId**](docs/PuzzlesAPI.md#apipuzzleid) | **Get** /api/puzzle/{id} | Get a puzzle by its ID
*PuzzlesAPI* | [**ApiPuzzleNext**](docs/PuzzlesAPI.md#apipuzzlenext) | **Get** /api/puzzle/next | Get a new puzzle
*PuzzlesAPI* | [**ApiPuzzleReplay**](docs/PuzzlesAPI.md#apipuzzlereplay) | **Get** /api/puzzle/replay/{days}/{theme} | Get puzzles to replay
*PuzzlesAPI* | [**ApiStormDashboard**](docs/PuzzlesAPI.md#apistormdashboard) | **Get** /api/storm/dashboard/{username} | Get the storm dashboard of a player
*PuzzlesAPI* | [**RacerGet**](docs/PuzzlesAPI.md#racerget) | **Get** /api/racer/{id} | Get puzzle race results
*PuzzlesAPI* | [**RacerPost**](docs/PuzzlesAPI.md#racerpost) | **Post** /api/racer | Create and join a puzzle race
*RelationsAPI* | [**ApiUserFollowing**](docs/RelationsAPI.md#apiuserfollowing) | **Get** /api/rel/following | Get users followed by the logged in user
*RelationsAPI* | [**BlockUser**](docs/RelationsAPI.md#blockuser) | **Post** /api/rel/block/{username} | Block a player
*RelationsAPI* | [**FollowUser**](docs/RelationsAPI.md#followuser) | **Post** /api/rel/follow/{username} | Follow a player
*RelationsAPI* | [**UnblockUser**](docs/RelationsAPI.md#unblockuser) | **Post** /api/rel/unblock/{username} | Unblock a player
*RelationsAPI* | [**UnfollowUser**](docs/RelationsAPI.md#unfollowuser) | **Post** /api/rel/unfollow/{username} | Unfollow a player
*SimulsAPI* | [**ApiSimul**](docs/SimulsAPI.md#apisimul) | **Get** /api/simul | Get current simuls
*StudiesAPI* | [**ApiStudyChapterMoves**](docs/StudiesAPI.md#apistudychaptermoves) | **Post** /api/study/{studyId}/{chapterId}/moves | Update the moves of a study chapter
*StudiesAPI* | [**ApiStudyChapterTags**](docs/StudiesAPI.md#apistudychaptertags) | **Post** /api/study/{studyId}/{chapterId}/tags | Update PGN tags of a study chapter
*StudiesAPI* | [**ApiStudyImportPGN**](docs/StudiesAPI.md#apistudyimportpgn) | **Post** /api/study/{studyId}/import-pgn | Import PGN into a study
*StudiesAPI* | [**ApiStudyPost**](docs/StudiesAPI.md#apistudypost) | **Post** /api/study | Create a new Study
*StudiesAPI* | [**ApiStudyStudyIdChapterIdDelete**](docs/StudiesAPI.md#apistudystudyidchapteriddelete) | **Delete** /api/study/{studyId}/{chapterId} | Delete a study chapter
*StudiesAPI* | [**StudyAllChaptersHead**](docs/StudiesAPI.md#studyallchaptershead) | **Head** /api/study/{studyId}.pgn | Study metadata
*StudiesAPI* | [**StudyAllChaptersPgn**](docs/StudiesAPI.md#studyallchapterspgn) | **Get** /api/study/{studyId}.pgn | Export all chapters
*StudiesAPI* | [**StudyChapterPgn**](docs/StudiesAPI.md#studychapterpgn) | **Get** /api/study/{studyId}/{chapterId}.pgn | Export one study chapter
*StudiesAPI* | [**StudyExportAllPgn**](docs/StudiesAPI.md#studyexportallpgn) | **Get** /api/study/by/{username}/export.pgn | Export all studies of a user
*StudiesAPI* | [**StudyListMetadata**](docs/StudiesAPI.md#studylistmetadata) | **Get** /api/study/by/{username} | List studies of a user
*TVAPI* | [**TvChannelFeed**](docs/TVAPI.md#tvchannelfeed) | **Get** /api/tv/{channel}/feed | Stream current TV game of a TV channel
*TVAPI* | [**TvChannelGames**](docs/TVAPI.md#tvchannelgames) | **Get** /api/tv/{channel} | Get best ongoing games of a TV channel
*TVAPI* | [**TvChannels**](docs/TVAPI.md#tvchannels) | **Get** /api/tv/channels | Get current TV games
*TVAPI* | [**TvFeed**](docs/TVAPI.md#tvfeed) | **Get** /api/tv/feed | Stream current TV game
*TablebaseAPI* | [**AntichessAtomic**](docs/TablebaseAPI.md#antichessatomic) | **Get** /antichess | Tablebase lookup for Antichess
*TablebaseAPI* | [**TablebaseAtomic**](docs/TablebaseAPI.md#tablebaseatomic) | **Get** /atomic | Tablebase lookup for Atomic chess
*TablebaseAPI* | [**TablebaseStandard**](docs/TablebaseAPI.md#tablebasestandard) | **Get** /standard | Tablebase lookup
*TeamsAPI* | [**ApiTeamArena**](docs/TeamsAPI.md#apiteamarena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
*TeamsAPI* | [**ApiTeamSwiss**](docs/TeamsAPI.md#apiteamswiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
*TeamsAPI* | [**TeamAll**](docs/TeamsAPI.md#teamall) | **Get** /api/team/all | Get popular teams
*TeamsAPI* | [**TeamIdJoin**](docs/TeamsAPI.md#teamidjoin) | **Post** /team/{teamId}/join | Join a team
*TeamsAPI* | [**TeamIdKickUserId**](docs/TeamsAPI.md#teamidkickuserid) | **Post** /api/team/{teamId}/kick/{userId} | Kick a user from your team
*TeamsAPI* | [**TeamIdPmAll**](docs/TeamsAPI.md#teamidpmall) | **Post** /team/{teamId}/pm-all | Message all members
*TeamsAPI* | [**TeamIdQuit**](docs/TeamsAPI.md#teamidquit) | **Post** /team/{teamId}/quit | Leave a team
*TeamsAPI* | [**TeamIdUsers**](docs/TeamsAPI.md#teamidusers) | **Get** /api/team/{teamId}/users | Get members of a team
*TeamsAPI* | [**TeamOfUsername**](docs/TeamsAPI.md#teamofusername) | **Get** /api/team/of/{username} | Teams of a player
*TeamsAPI* | [**TeamRequestAccept**](docs/TeamsAPI.md#teamrequestaccept) | **Post** /api/team/{teamId}/request/{userId}/accept | Accept join request
*TeamsAPI* | [**TeamRequestDecline**](docs/TeamsAPI.md#teamrequestdecline) | **Post** /api/team/{teamId}/request/{userId}/decline | Decline join request
*TeamsAPI* | [**TeamRequests**](docs/TeamsAPI.md#teamrequests) | **Get** /api/team/{teamId}/requests | Get join requests
*TeamsAPI* | [**TeamSearch**](docs/TeamsAPI.md#teamsearch) | **Get** /api/team/search | Search teams
*TeamsAPI* | [**TeamShow**](docs/TeamsAPI.md#teamshow) | **Get** /api/team/{teamId} | Get a single team
*TournamentsArenaAPI* | [**ApiTeamArena**](docs/TournamentsArenaAPI.md#apiteamarena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
*TournamentsArenaAPI* | [**ApiTournament**](docs/TournamentsArenaAPI.md#apitournament) | **Get** /api/tournament | Get current tournaments
*TournamentsArenaAPI* | [**ApiTournamentJoin**](docs/TournamentsArenaAPI.md#apitournamentjoin) | **Post** /api/tournament/{id}/join | Join an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentPost**](docs/TournamentsArenaAPI.md#apitournamentpost) | **Post** /api/tournament | Create a new Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentTeamBattlePost**](docs/TournamentsArenaAPI.md#apitournamentteambattlepost) | **Post** /api/tournament/team-battle/{id} | Update a team battle
*TournamentsArenaAPI* | [**ApiTournamentTerminate**](docs/TournamentsArenaAPI.md#apitournamentterminate) | **Post** /api/tournament/{id}/terminate | Terminate an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentUpdate**](docs/TournamentsArenaAPI.md#apitournamentupdate) | **Post** /api/tournament/{id} | Update an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentWithdraw**](docs/TournamentsArenaAPI.md#apitournamentwithdraw) | **Post** /api/tournament/{id}/withdraw | Pause or leave an Arena tournament
*TournamentsArenaAPI* | [**ApiUserNameTournamentCreated**](docs/TournamentsArenaAPI.md#apiusernametournamentcreated) | **Get** /api/user/{username}/tournament/created | Get tournaments created by a user
*TournamentsArenaAPI* | [**ApiUserNameTournamentPlayed**](docs/TournamentsArenaAPI.md#apiusernametournamentplayed) | **Get** /api/user/{username}/tournament/played | Get tournaments played by a user
*TournamentsArenaAPI* | [**GamesByTournament**](docs/TournamentsArenaAPI.md#gamesbytournament) | **Get** /api/tournament/{id}/games | Export games of an Arena tournament
*TournamentsArenaAPI* | [**ResultsByTournament**](docs/TournamentsArenaAPI.md#resultsbytournament) | **Get** /api/tournament/{id}/results | Get results of an Arena tournament
*TournamentsArenaAPI* | [**TeamsByTournament**](docs/TournamentsArenaAPI.md#teamsbytournament) | **Get** /api/tournament/{id}/teams | Get team standing of a team battle
*TournamentsArenaAPI* | [**Tournament**](docs/TournamentsArenaAPI.md#tournament) | **Get** /api/tournament/{id} | Get info about an Arena tournament
*TournamentsSwissAPI* | [**ApiSwissJoin**](docs/TournamentsSwissAPI.md#apiswissjoin) | **Post** /api/swiss/{id}/join | Join a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissNew**](docs/TournamentsSwissAPI.md#apiswissnew) | **Post** /api/swiss/new/{teamId} | Create a new Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissScheduleNextRound**](docs/TournamentsSwissAPI.md#apiswissschedulenextround) | **Post** /api/swiss/{id}/schedule-next-round | Manually schedule the next round
*TournamentsSwissAPI* | [**ApiSwissTerminate**](docs/TournamentsSwissAPI.md#apiswissterminate) | **Post** /api/swiss/{id}/terminate | Terminate a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissUpdate**](docs/TournamentsSwissAPI.md#apiswissupdate) | **Post** /api/swiss/{id}/edit | Update a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissWithdraw**](docs/TournamentsSwissAPI.md#apiswisswithdraw) | **Post** /api/swiss/{id}/withdraw | Pause or leave a swiss tournament
*TournamentsSwissAPI* | [**ApiTeamSwiss**](docs/TournamentsSwissAPI.md#apiteamswiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
*TournamentsSwissAPI* | [**GamesBySwiss**](docs/TournamentsSwissAPI.md#gamesbyswiss) | **Get** /api/swiss/{id}/games | Export games of a Swiss tournament
*TournamentsSwissAPI* | [**ResultsBySwiss**](docs/TournamentsSwissAPI.md#resultsbyswiss) | **Get** /api/swiss/{id}/results | Get results of a swiss tournament
*TournamentsSwissAPI* | [**Swiss**](docs/TournamentsSwissAPI.md#swiss) | **Get** /api/swiss/{id} | Get info about a Swiss tournament
*TournamentsSwissAPI* | [**SwissTrf**](docs/TournamentsSwissAPI.md#swisstrf) | **Get** /swiss/{id}.trf | Export TRF of a Swiss tournament
*UsersAPI* | [**ApiCrosstable**](docs/UsersAPI.md#apicrosstable) | **Get** /api/crosstable/{user1}/{user2} | Get crosstable
*UsersAPI* | [**ApiPlayerAutocomplete**](docs/UsersAPI.md#apiplayerautocomplete) | **Get** /api/player/autocomplete | Autocomplete usernames
*UsersAPI* | [**ApiUser**](docs/UsersAPI.md#apiuser) | **Get** /api/user/{username} | Get user public data
*UsersAPI* | [**ApiUserActivity**](docs/UsersAPI.md#apiuseractivity) | **Get** /api/user/{username}/activity | Get user activity
*UsersAPI* | [**ApiUserPerf**](docs/UsersAPI.md#apiuserperf) | **Get** /api/user/{username}/perf/{perf} | Get performance statistics of a user
*UsersAPI* | [**ApiUserRatingHistory**](docs/UsersAPI.md#apiuserratinghistory) | **Get** /api/user/{username}/rating-history | Get rating history of a user
*UsersAPI* | [**ApiUsers**](docs/UsersAPI.md#apiusers) | **Post** /api/users | Get users by ID
*UsersAPI* | [**ApiUsersStatus**](docs/UsersAPI.md#apiusersstatus) | **Get** /api/users/status | Get real-time users status
*UsersAPI* | [**Player**](docs/UsersAPI.md#player) | **Get** /api/player | Get all top 10
*UsersAPI* | [**PlayerTopNbPerfType**](docs/UsersAPI.md#playertopnbperftype) | **Get** /api/player/top/{nb}/{perfType} | Get one leaderboard
*UsersAPI* | [**ReadNote**](docs/UsersAPI.md#readnote) | **Get** /api/user/{username}/note | Get notes for a user
*UsersAPI* | [**StreamerLive**](docs/UsersAPI.md#streamerlive) | **Get** /api/streamer/live | Get live streamers
*UsersAPI* | [**WriteNote**](docs/UsersAPI.md#writenote) | **Post** /api/user/{username}/note | Add a note for a user


## Documentation For Models

 - [AIOpponent](docs/AIOpponent.md)
 - [Account200Response](docs/Account200Response.md)
 - [AccountEmail200Response](docs/AccountEmail200Response.md)
 - [AccountKid200Response](docs/AccountKid200Response.md)
 - [ApiAccountPlaying200Response](docs/ApiAccountPlaying200Response.md)
 - [ApiAccountPlaying200ResponseNowPlayingInner](docs/ApiAccountPlaying200ResponseNowPlayingInner.md)
 - [ApiAccountPlaying200ResponseNowPlayingInnerOpponent](docs/ApiAccountPlaying200ResponseNowPlayingInnerOpponent.md)
 - [ApiBoardSeek200Response](docs/ApiBoardSeek200Response.md)
 - [ApiExternalEngineAcquire200Response](docs/ApiExternalEngineAcquire200Response.md)
 - [ApiExternalEngineAcquireRequest](docs/ApiExternalEngineAcquireRequest.md)
 - [ApiExternalEngineAnalyse200Response](docs/ApiExternalEngineAnalyse200Response.md)
 - [ApiExternalEngineAnalyse200ResponsePvsInner](docs/ApiExternalEngineAnalyse200ResponsePvsInner.md)
 - [ApiExternalEngineAnalyseRequest](docs/ApiExternalEngineAnalyseRequest.md)
 - [ApiPlayerAutocomplete200Response](docs/ApiPlayerAutocomplete200Response.md)
 - [ApiPlayerAutocomplete200ResponseOneOf](docs/ApiPlayerAutocomplete200ResponseOneOf.md)
 - [ApiPuzzleReplay404Response](docs/ApiPuzzleReplay404Response.md)
 - [ApiSimul200Response](docs/ApiSimul200Response.md)
 - [ApiStreamEvent200Response](docs/ApiStreamEvent200Response.md)
 - [ApiStudyPost200Response](docs/ApiStudyPost200Response.md)
 - [ApiToken200Response](docs/ApiToken200Response.md)
 - [ApiUsersStatus200ResponseInner](docs/ApiUsersStatus200ResponseInner.md)
 - [ArenaPerf](docs/ArenaPerf.md)
 - [ArenaPosition](docs/ArenaPosition.md)
 - [ArenaRatingObj](docs/ArenaRatingObj.md)
 - [ArenaSheet](docs/ArenaSheet.md)
 - [ArenaStatus](docs/ArenaStatus.md)
 - [ArenaStatusName](docs/ArenaStatusName.md)
 - [ArenaTournament](docs/ArenaTournament.md)
 - [ArenaTournamentFull](docs/ArenaTournamentFull.md)
 - [ArenaTournamentFullDuelsInner](docs/ArenaTournamentFullDuelsInner.md)
 - [ArenaTournamentFullDuelsInnerPInner](docs/ArenaTournamentFullDuelsInnerPInner.md)
 - [ArenaTournamentFullFeatured](docs/ArenaTournamentFullFeatured.md)
 - [ArenaTournamentFullFeaturedC](docs/ArenaTournamentFullFeaturedC.md)
 - [ArenaTournamentFullFeaturedWhite](docs/ArenaTournamentFullFeaturedWhite.md)
 - [ArenaTournamentFullGreatPlayer](docs/ArenaTournamentFullGreatPlayer.md)
 - [ArenaTournamentFullPerf](docs/ArenaTournamentFullPerf.md)
 - [ArenaTournamentFullPodiumInner](docs/ArenaTournamentFullPodiumInner.md)
 - [ArenaTournamentFullPodiumInnerNb](docs/ArenaTournamentFullPodiumInnerNb.md)
 - [ArenaTournamentFullQuote](docs/ArenaTournamentFullQuote.md)
 - [ArenaTournamentFullSchedule](docs/ArenaTournamentFullSchedule.md)
 - [ArenaTournamentFullSpotlight](docs/ArenaTournamentFullSpotlight.md)
 - [ArenaTournamentFullStanding](docs/ArenaTournamentFullStanding.md)
 - [ArenaTournamentFullStandingPlayersInner](docs/ArenaTournamentFullStandingPlayersInner.md)
 - [ArenaTournamentFullStats](docs/ArenaTournamentFullStats.md)
 - [ArenaTournamentMinRatedGames](docs/ArenaTournamentMinRatedGames.md)
 - [ArenaTournamentPlayed](docs/ArenaTournamentPlayed.md)
 - [ArenaTournamentPlayer](docs/ArenaTournamentPlayer.md)
 - [ArenaTournamentSchedule](docs/ArenaTournamentSchedule.md)
 - [ArenaTournamentTeamBattle](docs/ArenaTournamentTeamBattle.md)
 - [ArenaTournaments](docs/ArenaTournaments.md)
 - [BoardGameDrawAcceptParameter](docs/BoardGameDrawAcceptParameter.md)
 - [BoardGameStream200Response](docs/BoardGameStream200Response.md)
 - [BroadcastByUser](docs/BroadcastByUser.md)
 - [BroadcastCustomPointsPerColor](docs/BroadcastCustomPointsPerColor.md)
 - [BroadcastCustomScoring](docs/BroadcastCustomScoring.md)
 - [BroadcastFormGrouping](docs/BroadcastFormGrouping.md)
 - [BroadcastFormGroupingInfo](docs/BroadcastFormGroupingInfo.md)
 - [BroadcastGameEntry](docs/BroadcastGameEntry.md)
 - [BroadcastGroup](docs/BroadcastGroup.md)
 - [BroadcastGroupTour](docs/BroadcastGroupTour.md)
 - [BroadcastMyRound](docs/BroadcastMyRound.md)
 - [BroadcastPgnPush](docs/BroadcastPgnPush.md)
 - [BroadcastPgnPushGamesInner](docs/BroadcastPgnPushGamesInner.md)
 - [BroadcastPhotosValue](docs/BroadcastPhotosValue.md)
 - [BroadcastPlayerEntry](docs/BroadcastPlayerEntry.md)
 - [BroadcastPlayerEntryWithFideAndGames](docs/BroadcastPlayerEntryWithFideAndGames.md)
 - [BroadcastPlayerEntryWithFideAndGamesAllOfFide](docs/BroadcastPlayerEntryWithFideAndGamesAllOfFide.md)
 - [BroadcastPlayerTiebreak](docs/BroadcastPlayerTiebreak.md)
 - [BroadcastPlayerWithFed](docs/BroadcastPlayerWithFed.md)
 - [BroadcastPointStr](docs/BroadcastPointStr.md)
 - [BroadcastRound](docs/BroadcastRound.md)
 - [BroadcastRoundGame](docs/BroadcastRoundGame.md)
 - [BroadcastRoundGamePlayersInner](docs/BroadcastRoundGamePlayersInner.md)
 - [BroadcastRoundInfo](docs/BroadcastRoundInfo.md)
 - [BroadcastRoundNew](docs/BroadcastRoundNew.md)
 - [BroadcastRoundStudyInfo](docs/BroadcastRoundStudyInfo.md)
 - [BroadcastRoundStudyInfoFeatures](docs/BroadcastRoundStudyInfoFeatures.md)
 - [BroadcastTeamLeaderboardEntry](docs/BroadcastTeamLeaderboardEntry.md)
 - [BroadcastTeamPOVMatchEntry](docs/BroadcastTeamPOVMatchEntry.md)
 - [BroadcastTiebreakExtendedCode](docs/BroadcastTiebreakExtendedCode.md)
 - [BroadcastTop](docs/BroadcastTop.md)
 - [BroadcastTopPast](docs/BroadcastTopPast.md)
 - [BroadcastTour](docs/BroadcastTour.md)
 - [BroadcastTourInfo](docs/BroadcastTourInfo.md)
 - [BroadcastWithLastRound](docs/BroadcastWithLastRound.md)
 - [BroadcastWithRounds](docs/BroadcastWithRounds.md)
 - [BroadcastWithRoundsAndFullGroup](docs/BroadcastWithRoundsAndFullGroup.md)
 - [BroadcastsByUser200Response](docs/BroadcastsByUser200Response.md)
 - [BroadcastsSearch200Response](docs/BroadcastsSearch200Response.md)
 - [BulkPairing](docs/BulkPairing.md)
 - [BulkPairingGamesInner](docs/BulkPairingGamesInner.md)
 - [ChallengeAi201Response](docs/ChallengeAi201Response.md)
 - [ChallengeCanceledEvent](docs/ChallengeCanceledEvent.md)
 - [ChallengeColor](docs/ChallengeColor.md)
 - [ChallengeDeclinedEvent](docs/ChallengeDeclinedEvent.md)
 - [ChallengeDeclinedJson](docs/ChallengeDeclinedJson.md)
 - [ChallengeEvent](docs/ChallengeEvent.md)
 - [ChallengeJson](docs/ChallengeJson.md)
 - [ChallengeJsonPerf](docs/ChallengeJsonPerf.md)
 - [ChallengeList200Response](docs/ChallengeList200Response.md)
 - [ChallengeOpenJson](docs/ChallengeOpenJson.md)
 - [ChallengeOpenJsonOpen](docs/ChallengeOpenJsonOpen.md)
 - [ChallengeOpenJsonPerf](docs/ChallengeOpenJsonPerf.md)
 - [ChallengeStatus](docs/ChallengeStatus.md)
 - [ChallengeUser](docs/ChallengeUser.md)
 - [ChatLineEvent](docs/ChatLineEvent.md)
 - [Clock](docs/Clock.md)
 - [CloudEval](docs/CloudEval.md)
 - [CloudEvalPvsInner](docs/CloudEvalPvsInner.md)
 - [Correspondence](docs/Correspondence.md)
 - [Count](docs/Count.md)
 - [Crosstable](docs/Crosstable.md)
 - [CustomPosition](docs/CustomPosition.md)
 - [Error](docs/Error.md)
 - [ExternalEngine](docs/ExternalEngine.md)
 - [ExternalEngineRegistration](docs/ExternalEngineRegistration.md)
 - [ExternalEngineWork](docs/ExternalEngineWork.md)
 - [ExternalEngineWorkCommon](docs/ExternalEngineWorkCommon.md)
 - [ExternalEngineWorkOneOf](docs/ExternalEngineWorkOneOf.md)
 - [ExternalEngineWorkOneOf1](docs/ExternalEngineWorkOneOf1.md)
 - [ExternalEngineWorkOneOf2](docs/ExternalEngineWorkOneOf2.md)
 - [FIDEPlayer](docs/FIDEPlayer.md)
 - [FIDEPlayerPhoto](docs/FIDEPlayerPhoto.md)
 - [FIDEPlayerRatings](docs/FIDEPlayerRatings.md)
 - [Featured](docs/Featured.md)
 - [FeaturedPlayersInner](docs/FeaturedPlayersInner.md)
 - [Fen](docs/Fen.md)
 - [FideTimeControl](docs/FideTimeControl.md)
 - [GameColor](docs/GameColor.md)
 - [GameCompat](docs/GameCompat.md)
 - [GameEventInfo](docs/GameEventInfo.md)
 - [GameEventOpponent](docs/GameEventOpponent.md)
 - [GameEventPlayer](docs/GameEventPlayer.md)
 - [GameFinishEvent](docs/GameFinishEvent.md)
 - [GameFullEvent](docs/GameFullEvent.md)
 - [GameFullEventClock](docs/GameFullEventClock.md)
 - [GameFullEventPerf](docs/GameFullEventPerf.md)
 - [GameImport200Response](docs/GameImport200Response.md)
 - [GameJson](docs/GameJson.md)
 - [GameJsonClock](docs/GameJsonClock.md)
 - [GameJsonDivision](docs/GameJsonDivision.md)
 - [GameMoveAnalysis](docs/GameMoveAnalysis.md)
 - [GameMoveAnalysisJudgment](docs/GameMoveAnalysisJudgment.md)
 - [GameOpening](docs/GameOpening.md)
 - [GamePgn200Response](docs/GamePgn200Response.md)
 - [GamePlayerUser](docs/GamePlayerUser.md)
 - [GamePlayerUserAnalysis](docs/GamePlayerUserAnalysis.md)
 - [GamePlayers](docs/GamePlayers.md)
 - [GameSource](docs/GameSource.md)
 - [GameStartEvent](docs/GameStartEvent.md)
 - [GameStateEvent](docs/GameStateEvent.md)
 - [GameStateEventExpiration](docs/GameStateEventExpiration.md)
 - [GameStatus](docs/GameStatus.md)
 - [GameStatusId](docs/GameStatusId.md)
 - [GameStatusName](docs/GameStatusName.md)
 - [GameStreamGame](docs/GameStreamGame.md)
 - [GameStreamGameClock](docs/GameStreamGameClock.md)
 - [GameStreamGamePlayers](docs/GameStreamGamePlayers.md)
 - [GameStreamGamePlayersWhite](docs/GameStreamGamePlayersWhite.md)
 - [Leaderboard](docs/Leaderboard.md)
 - [LightUser](docs/LightUser.md)
 - [LightUserOnline](docs/LightUserOnline.md)
 - [MateVariation](docs/MateVariation.md)
 - [MoveStreamEntry](docs/MoveStreamEntry.md)
 - [MoveStreamEntryOneOf](docs/MoveStreamEntryOneOf.md)
 - [MoveStreamEntryOneOf1](docs/MoveStreamEntryOneOf1.md)
 - [NonMateVariation](docs/NonMateVariation.md)
 - [NotFound](docs/NotFound.md)
 - [OAuthError](docs/OAuthError.md)
 - [Ok](docs/Ok.md)
 - [OpeningExplorerGamePlayer](docs/OpeningExplorerGamePlayer.md)
 - [OpeningExplorerLichess](docs/OpeningExplorerLichess.md)
 - [OpeningExplorerLichessGame](docs/OpeningExplorerLichessGame.md)
 - [OpeningExplorerLichessHistoryInner](docs/OpeningExplorerLichessHistoryInner.md)
 - [OpeningExplorerLichessMovesInner](docs/OpeningExplorerLichessMovesInner.md)
 - [OpeningExplorerLichessTopGamesInner](docs/OpeningExplorerLichessTopGamesInner.md)
 - [OpeningExplorerMasters](docs/OpeningExplorerMasters.md)
 - [OpeningExplorerMastersGame](docs/OpeningExplorerMastersGame.md)
 - [OpeningExplorerMastersMovesInner](docs/OpeningExplorerMastersMovesInner.md)
 - [OpeningExplorerMastersTopGamesInner](docs/OpeningExplorerMastersTopGamesInner.md)
 - [OpeningExplorerOpening](docs/OpeningExplorerOpening.md)
 - [OpeningExplorerPlayer](docs/OpeningExplorerPlayer.md)
 - [OpeningExplorerPlayerGame](docs/OpeningExplorerPlayerGame.md)
 - [OpeningExplorerPlayerMovesInner](docs/OpeningExplorerPlayerMovesInner.md)
 - [OpeningExplorerPlayerRecentGamesInner](docs/OpeningExplorerPlayerRecentGamesInner.md)
 - [OpponentGoneEvent](docs/OpponentGoneEvent.md)
 - [Perf](docs/Perf.md)
 - [PerfStat](docs/PerfStat.md)
 - [PerfStatPerf](docs/PerfStatPerf.md)
 - [PerfStatPerfGlicko](docs/PerfStatPerfGlicko.md)
 - [PerfStatStat](docs/PerfStatStat.md)
 - [PerfStatStatBestWins](docs/PerfStatStatBestWins.md)
 - [PerfStatStatBestWinsResultsInner](docs/PerfStatStatBestWinsResultsInner.md)
 - [PerfStatStatCount](docs/PerfStatStatCount.md)
 - [PerfStatStatHighest](docs/PerfStatStatHighest.md)
 - [PerfStatStatPlayStreak](docs/PerfStatStatPlayStreak.md)
 - [PerfStatStatPlayStreakNb](docs/PerfStatStatPlayStreakNb.md)
 - [PerfStatStatPlayStreakNbCur](docs/PerfStatStatPlayStreakNbCur.md)
 - [PerfStatStatResultStreak](docs/PerfStatStatResultStreak.md)
 - [PerfStatStatResultStreakLoss](docs/PerfStatStatResultStreakLoss.md)
 - [PerfStatStatResultStreakLossMax](docs/PerfStatStatResultStreakLossMax.md)
 - [PerfStatStatResultStreakLossMaxFrom](docs/PerfStatStatResultStreakLossMaxFrom.md)
 - [PerfStatStatResultStreakWin](docs/PerfStatStatResultStreakWin.md)
 - [PerfStatStatResultStreakWinCur](docs/PerfStatStatResultStreakWinCur.md)
 - [PerfStatStatResultStreakWinCurFrom](docs/PerfStatStatResultStreakWinCurFrom.md)
 - [PerfStatUser](docs/PerfStatUser.md)
 - [PerfType](docs/PerfType.md)
 - [Perfs](docs/Perfs.md)
 - [PlayTime](docs/PlayTime.md)
 - [Player](docs/Player.md)
 - [Profile](docs/Profile.md)
 - [PuzzleActivity](docs/PuzzleActivity.md)
 - [PuzzleActivityPuzzle](docs/PuzzleActivityPuzzle.md)
 - [PuzzleAndGame](docs/PuzzleAndGame.md)
 - [PuzzleAndGameGame](docs/PuzzleAndGameGame.md)
 - [PuzzleAndGameGamePerf](docs/PuzzleAndGameGamePerf.md)
 - [PuzzleAndGameGamePlayersInner](docs/PuzzleAndGameGamePlayersInner.md)
 - [PuzzleAndGamePuzzle](docs/PuzzleAndGamePuzzle.md)
 - [PuzzleBatchSelect](docs/PuzzleBatchSelect.md)
 - [PuzzleBatchSolveRequest](docs/PuzzleBatchSolveRequest.md)
 - [PuzzleBatchSolveRequestSolutionsInner](docs/PuzzleBatchSolveRequestSolutionsInner.md)
 - [PuzzleBatchSolveResponse](docs/PuzzleBatchSolveResponse.md)
 - [PuzzleBatchSolveResponseRoundsInner](docs/PuzzleBatchSolveResponseRoundsInner.md)
 - [PuzzleDashboard](docs/PuzzleDashboard.md)
 - [PuzzleDashboardThemesValue](docs/PuzzleDashboardThemesValue.md)
 - [PuzzleGlicko](docs/PuzzleGlicko.md)
 - [PuzzleModePerf](docs/PuzzleModePerf.md)
 - [PuzzlePerformance](docs/PuzzlePerformance.md)
 - [PuzzleRaceResults](docs/PuzzleRaceResults.md)
 - [PuzzleRaceResultsPlayersInner](docs/PuzzleRaceResultsPlayersInner.md)
 - [PuzzleRaceResultsPuzzlesInner](docs/PuzzleRaceResultsPuzzlesInner.md)
 - [PuzzleRacer](docs/PuzzleRacer.md)
 - [PuzzleReplay](docs/PuzzleReplay.md)
 - [PuzzleReplayAngle](docs/PuzzleReplayAngle.md)
 - [PuzzleReplayReplay](docs/PuzzleReplayReplay.md)
 - [PuzzleStormDashboard](docs/PuzzleStormDashboard.md)
 - [PuzzleStormDashboardDaysInner](docs/PuzzleStormDashboardDaysInner.md)
 - [PuzzleStormDashboardHigh](docs/PuzzleStormDashboardHigh.md)
 - [RatingHistoryEntry](docs/RatingHistoryEntry.md)
 - [RealTime](docs/RealTime.md)
 - [ResultsBySwiss200Response](docs/ResultsBySwiss200Response.md)
 - [ResultsByTournament200Response](docs/ResultsByTournament200Response.md)
 - [Simul](docs/Simul.md)
 - [SimulHost](docs/SimulHost.md)
 - [SimulVariantsInner](docs/SimulVariantsInner.md)
 - [SpectatorGameChatInner](docs/SpectatorGameChatInner.md)
 - [Speed](docs/Speed.md)
 - [StatByFideTC](docs/StatByFideTC.md)
 - [StreamerLive200ResponseInner](docs/StreamerLive200ResponseInner.md)
 - [StreamerLive200ResponseInnerAllOfStream](docs/StreamerLive200ResponseInnerAllOfStream.md)
 - [StreamerLive200ResponseInnerAllOfStreamer](docs/StreamerLive200ResponseInnerAllOfStreamer.md)
 - [StudyImportPgnChapters](docs/StudyImportPgnChapters.md)
 - [StudyImportPgnChaptersChaptersInner](docs/StudyImportPgnChaptersChaptersInner.md)
 - [StudyImportPgnChaptersChaptersInnerPlayersInner](docs/StudyImportPgnChaptersChaptersInnerPlayersInner.md)
 - [StudyMetadata](docs/StudyMetadata.md)
 - [StudyUserSelection](docs/StudyUserSelection.md)
 - [SwissStatus](docs/SwissStatus.md)
 - [SwissTournament](docs/SwissTournament.md)
 - [SwissTournamentClock](docs/SwissTournamentClock.md)
 - [SwissTournamentNextRound](docs/SwissTournamentNextRound.md)
 - [SwissTournamentStats](docs/SwissTournamentStats.md)
 - [SwissUnauthorisedEdit](docs/SwissUnauthorisedEdit.md)
 - [TablebaseJson](docs/TablebaseJson.md)
 - [TablebaseMove](docs/TablebaseMove.md)
 - [Team](docs/Team.md)
 - [TeamIdUsers200Response](docs/TeamIdUsers200Response.md)
 - [TeamPaginatorJson](docs/TeamPaginatorJson.md)
 - [TeamRequest](docs/TeamRequest.md)
 - [TeamRequestWithUser](docs/TeamRequestWithUser.md)
 - [TeamsByTournament200Response](docs/TeamsByTournament200Response.md)
 - [TeamsByTournament200ResponseTeamsInner](docs/TeamsByTournament200ResponseTeamsInner.md)
 - [TeamsByTournament200ResponseTeamsInnerPlayersInner](docs/TeamsByTournament200ResponseTeamsInnerPlayersInner.md)
 - [Thematic](docs/Thematic.md)
 - [TimeControl](docs/TimeControl.md)
 - [Timeline](docs/Timeline.md)
 - [TimelineEntriesInner](docs/TimelineEntriesInner.md)
 - [TimelineEntryBlogPost](docs/TimelineEntryBlogPost.md)
 - [TimelineEntryBlogPostData](docs/TimelineEntryBlogPostData.md)
 - [TimelineEntryFollow](docs/TimelineEntryFollow.md)
 - [TimelineEntryFollowData](docs/TimelineEntryFollowData.md)
 - [TimelineEntryForumPost](docs/TimelineEntryForumPost.md)
 - [TimelineEntryForumPostData](docs/TimelineEntryForumPostData.md)
 - [TimelineEntryGameEnd](docs/TimelineEntryGameEnd.md)
 - [TimelineEntryGameEndData](docs/TimelineEntryGameEndData.md)
 - [TimelineEntryPlanRenew](docs/TimelineEntryPlanRenew.md)
 - [TimelineEntryPlanRenewData](docs/TimelineEntryPlanRenewData.md)
 - [TimelineEntryPlanStart](docs/TimelineEntryPlanStart.md)
 - [TimelineEntryPlanStartData](docs/TimelineEntryPlanStartData.md)
 - [TimelineEntrySimul](docs/TimelineEntrySimul.md)
 - [TimelineEntrySimulData](docs/TimelineEntrySimulData.md)
 - [TimelineEntryStreamStart](docs/TimelineEntryStreamStart.md)
 - [TimelineEntryStreamStartData](docs/TimelineEntryStreamStartData.md)
 - [TimelineEntryStudyLike](docs/TimelineEntryStudyLike.md)
 - [TimelineEntryStudyLikeData](docs/TimelineEntryStudyLikeData.md)
 - [TimelineEntryTeamCreate](docs/TimelineEntryTeamCreate.md)
 - [TimelineEntryTeamJoin](docs/TimelineEntryTeamJoin.md)
 - [TimelineEntryTeamJoinData](docs/TimelineEntryTeamJoinData.md)
 - [TimelineEntryTourJoin](docs/TimelineEntryTourJoin.md)
 - [TimelineEntryTourJoinData](docs/TimelineEntryTourJoinData.md)
 - [TimelineEntryUblogPost](docs/TimelineEntryUblogPost.md)
 - [TimelineEntryUblogPostData](docs/TimelineEntryUblogPostData.md)
 - [TimelineEntryUblogPostLike](docs/TimelineEntryUblogPostLike.md)
 - [TimelineEntryUblogPostLikeData](docs/TimelineEntryUblogPostLikeData.md)
 - [TimelineUsersValue](docs/TimelineUsersValue.md)
 - [Title](docs/Title.md)
 - [TokenTest200ResponseValue](docs/TokenTest200ResponseValue.md)
 - [TokenTest200ResponseValueOneOf](docs/TokenTest200ResponseValueOneOf.md)
 - [Top10s](docs/Top10s.md)
 - [TopUser](docs/TopUser.md)
 - [TopUserPerfsValue](docs/TopUserPerfsValue.md)
 - [TvChannels200Response](docs/TvChannels200Response.md)
 - [TvFeed](docs/TvFeed.md)
 - [TvFeedFeatured](docs/TvFeedFeatured.md)
 - [TvFeedFen](docs/TvFeedFen.md)
 - [TvGame](docs/TvGame.md)
 - [UciVariant](docs/UciVariant.md)
 - [Unlimited](docs/Unlimited.md)
 - [User](docs/User.md)
 - [UserActivity](docs/UserActivity.md)
 - [UserActivityCorrespondenceEnds](docs/UserActivityCorrespondenceEnds.md)
 - [UserActivityCorrespondenceEndsCorrespondence](docs/UserActivityCorrespondenceEndsCorrespondence.md)
 - [UserActivityCorrespondenceGame](docs/UserActivityCorrespondenceGame.md)
 - [UserActivityCorrespondenceGameOpponent](docs/UserActivityCorrespondenceGameOpponent.md)
 - [UserActivityCorrespondenceGameOpponentOneOf](docs/UserActivityCorrespondenceGameOpponentOneOf.md)
 - [UserActivityCorrespondenceGameOpponentOneOf1](docs/UserActivityCorrespondenceGameOpponentOneOf1.md)
 - [UserActivityCorrespondenceMoves](docs/UserActivityCorrespondenceMoves.md)
 - [UserActivityFollowList](docs/UserActivityFollowList.md)
 - [UserActivityFollows](docs/UserActivityFollows.md)
 - [UserActivityGames](docs/UserActivityGames.md)
 - [UserActivityInterval](docs/UserActivityInterval.md)
 - [UserActivityPatron](docs/UserActivityPatron.md)
 - [UserActivityPostsInner](docs/UserActivityPostsInner.md)
 - [UserActivityPostsInnerPostsInner](docs/UserActivityPostsInnerPostsInner.md)
 - [UserActivityPracticeInner](docs/UserActivityPracticeInner.md)
 - [UserActivityPuzzles](docs/UserActivityPuzzles.md)
 - [UserActivityScore](docs/UserActivityScore.md)
 - [UserActivityScoreRp](docs/UserActivityScoreRp.md)
 - [UserActivityTeamsInner](docs/UserActivityTeamsInner.md)
 - [UserActivityTournaments](docs/UserActivityTournaments.md)
 - [UserActivityTournamentsBestInner](docs/UserActivityTournamentsBestInner.md)
 - [UserActivityTournamentsBestInnerTournament](docs/UserActivityTournamentsBestInnerTournament.md)
 - [UserExtended](docs/UserExtended.md)
 - [UserNote](docs/UserNote.md)
 - [UserPreferences](docs/UserPreferences.md)
 - [UserStreamer](docs/UserStreamer.md)
 - [UserStreamerTwitch](docs/UserStreamerTwitch.md)
 - [Variant](docs/Variant.md)
 - [VariantKey](docs/VariantKey.md)
 - [Verdict](docs/Verdict.md)
 - [Verdicts](docs/Verdicts.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OAuth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://lichess.org/oauth
- **Scopes**: 
 - **preference:read**: Read your preferences
 - **preference:write**: Write your preferences
 - **email:read**: Read your email address
 - **engine:read**: Read your external engines
 - **engine:write**: Create, update, delete your external engines
 - **challenge:read**: Read incoming challenges
 - **challenge:write**: Create, accept, decline challenges
 - **challenge:bulk**: Create, delete, query bulk pairings
 - **study:read**: Read private studies and broadcasts
 - **study:write**: Create, update, delete studies and broadcasts
 - **tournament:write**: Create tournaments
 - **racer:write**: Create and join puzzle races
 - **puzzle:read**: Read puzzle activity
 - **puzzle:write**: Write puzzle activity
 - **team:read**: Read private team information
 - **team:write**: Join, leave teams
 - **team:lead**: Manage teams (kick members, send PMs)
 - **follow:read**: Read followed players
 - **follow:write**: Follow and unfollow other players
 - **msg:write**: Send private messages to other players
 - **board:play**: Play with the Board API
 - **bot:play**: Play with the Bot API. Only for [Bot accounts](#tag/bot/POST/api/bot/account/upgrade)
 - **web:mod**: Use moderator tools (within the bounds of your permissions)

Example

```go
auth := context.WithValue(context.Background(), openapigenerator.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapigenerator.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

contact@lichess.org

