package csgsi

type Side string

const (
	T  Side = "T"
	CT Side = "CT"
)

type PlayerActivity string

const (
	PlayerActivityActive    PlayerActivity = "active"
	PlayerActivityMenu      PlayerActivity = "menu"
	PlayerActivityTextInput PlayerActivity = "textinput"
)

type MapPhase string

const (
	MapPhaseWarmup       MapPhase = "warmup"
	MapPhaseLive         MapPhase = "live"
	MapPhaseIntermission MapPhase = "intermission"
	MapPhaseGameOver     MapPhase = "gameover"
)

type RoundOutcome string

const (
	CTWinElimination RoundOutcome = "ct_win_elimination"
	TWinElimination  RoundOutcome = "t_win_elimination"
	CTWinTimeLimit   RoundOutcome = "ct_win_time"
	CTWinDefuse      RoundOutcome = "ct_win_defuse"
	TWinBomb         RoundOutcome = "t_win_bomb"
)

type BombRoundState string

const (
	BombRoundStatePlanted  BombRoundState = "planted"
	BombRoundStateExploded BombRoundState = "exploded"
	BombRoundStateDefused  BombRoundState = "defused"
)

type RoundPhase string

const (
	RoundPhaseFreezeTime RoundPhase = "freezetime"
	RoundPhaseLive       RoundPhase = "live"
	RoundPhaseOver       RoundPhase = "over"
)

type BombState string

const (
	BombStateCarried  BombState = "carried"
	BombStateDropped  BombState = "dropped"
	BombStatePlanted  BombState = "planted"
	BombStateExploded BombState = "exploded"
	BombStateDefused  BombState = "defused"
	BombStateDefusing BombState = "defusing"
	BombStatePlanting BombState = "planting"
)

type PhaseType string

const (
	PhaseTypeFreezetime PhaseType = "freezetime"
	PhaseTypeBomb       PhaseType = "bomb"
	PhaseTypeWarmup     PhaseType = "warmup"
	PhaseTypeLive       PhaseType = "live"
	PhaseTypeOver       PhaseType = "over"
	PhaseTypeDefuse     PhaseType = "defuse"
	PhaseTypePaused     PhaseType = "paused"
	PhaseTypeTimeoutCT  PhaseType = "timeout_ct"
	PhaseTypeTimeoutT   PhaseType = "timeout_t"
)

type GrenadeType string

const (
	GrenadeTypeFlash      GrenadeType = "flash"
	GrenadeTypeDecoy      GrenadeType = "decoy"
	GrenadeTypeFrag       GrenadeType = "frag"
	GrenadeTypeSmoke      GrenadeType = "smoke"
	GrenadeTypeMolotov    GrenadeType = "firebomb"
	GrenadeTypeIncendiary GrenadeType = "inferno"
)

type WeaponState string

const (
	WeaponStateActive    WeaponState = "active"
	WeaponStateHolstered WeaponState = "holstered"
	WeaponStateReloading WeaponState = "reloading"
)

type WeaponType string

const (
	WeaponTypeKnife         WeaponType = "Knife"
	WeaponTypePistol        WeaponType = "Pistol"
	WeaponTypeGrenade       WeaponType = "Grenade"
	WeaponTypeRifle         WeaponType = "Rifle"
	WeaponTypeSniperRifle   WeaponType = "SniperRifle"
	WeaponTypeC4            WeaponType = "C4"
	WeaponTypeSubmachineGun WeaponType = "Submachine Gun"
	WeaponTypeShotgun       WeaponType = "Shotgun"
	WeaponTypeMachineGun    WeaponType = "Machine Gun"
)

type State struct {
	Provider         *Provider
	Map              *CsMap
	Round            *Round
	Player           *Player
	AllPlayers       map[string]*Player // allplayers_*: steamid64 ...
	Bomb             *Bomb
	Grenades         map[string]*Grenade
	Previously       *State
	Added            *State
	Phase_countdowns *PhaseCountdown
	Auth             *Auth
}

// provider
type Provider struct {
	Name      string
	AppId     int
	Version   int
	SteamId   string
	Timestamp float32
}

// map
type CsMap struct {
	Mode                      string
	Name                      string
	Phase                     *MapPhase
	Round                     int
	Team_ct                   *Team
	Team_t                    *Team
	Num_matches_to_win_series int
	Current_spectators        int
	Souvenirs_total           int
	Round_wins                map[string]*RoundOutcome
}

// round
type Round struct {
	Phase    RoundPhase
	Win_team Side
	Bomb     BombRoundState
}

// player_id
type Player struct {
	SteamId       string
	Clan          string
	Name          string
	Observer_slot int
	Team          *Side
	Activity      PlayerActivity
	State         *PlayerState
	Weapons       map[string]*Weapon
	Match_stats   *PlayerMatchStats
	Position      [3]float32
	Forward       [3]float32
	Spectarget    string
}

// team
type Team struct {
	Logo                     string
	Score                    int
	Consecutive_round_losses int
	Timeouts_remaining       int
	Matches_won_this_series  int
	Name                     string
	Flag                     string
}

// player_state
type PlayerState struct {
	Health         int
	Armor          int
	Helmet         bool
	DefuseKit      bool
	Flashed        int
	Smoked         int
	Burning        int
	Money          int
	Round_kills    int
	Round_killhs   int
	Round_totaldmg int
	Equip_value    int
}

// player_weapons: weapon_0, weapon_1, weapon_2 ...
type Weapon struct {
	Name          string
	PaintKit      string
	Type          WeaponType
	State         WeaponState
	Ammo_clip     int
	Ammo_clip_max int
	Ammo_reserve  int
}

// player_match_stats
type PlayerMatchStats struct {
	Kills   int
	Assists int
	Deaths  int
	Mvps    int
	Score   int
}

type Auth struct {
	Token string
}

type Bomb struct {
	State     BombState
	Countdown string
	Player    string
	Position  [3]float32
}

type PhaseCountdown struct {
	Phase         PhaseType
	Phase_ends_in string
}

type Grenade struct {
	Owner      int
	Position   [3]float32
	Velocity   [3]float32
	Type       GrenadeType
	Lifetime   string
	EffectTime float32
}