// Code generated by "handlerplate basic.go calc.go daily.go game.go phone.go farm.go auction.go actions.go"; DO NOT EDIT.
package handlers

type HandlerID int

const (
	NoHandler HandlerID = iota
	HelpHandler
	InfaHandler
	WhoHandler
	ListHandler
	TopHandler
	MouseHandler
	TiktokHandler
	GameHandler
	WeatherHandler
	CatHandler
	AnimeHandler
	FurryHandler
	FlagHandler
	PersonHandler
	HorseHandler
	ArtHandler
	CarHandler
	SoyHandler
	DanbooruHandler
	FapHandler
	MasyunyaHandler
	PoppyHandler
	SimaHandler
	LageonaHandler
	HelloHandler
	BasiliHandler
	CasperHandler
	ZeusHandler
	PicHandler
	AvatarHandler
	TurnOnHandler
	TurnOffHandler
	BanHandler
	UnbanHandler
	CalculatorHandler
	DailyEblanHandler
	DailyAdminHandler
	DailyPairHandler
	NameHandler
	InventoryHandler
	SortHandler
	CatchHandler
	DropHandler
	PickHandler
	FloorHandler
	MarketHandler
	PriceListHandler
	NameMarketHandler
	GetJobHandler
	QuitJobHandler
	BuyHandler
	EatHandler
	EatQuickHandler
	FishHandler
	CastNetHandler
	DrawNetHandler
	NetHandler
	FishingRecordsHandler
	CraftHandler
	StatusHandler
	SellHandler
	SellQuickHandler
	StackHandler
	SplitHandler
	CashoutHandler
	FightHandler
	PvPHandler
	ProfileHandler
	DiceHandler
	RollHandler
	TopStrongHandler
	TopRatingHandler
	TopRichHandler
	CapitalHandler
	BalanceHandler
	FundsHandler
	EnergyHandler
	NamePetHandler
	ReceiveSMSHandler
	SendSMSHandler
	ContactsHandler
	SpamHandler
	FarmHandler
	PlantHandler
	HarvestHandler
	HarvestInlineHandler
	UpgradeFarmHandler
	AuctionHandler
	AuctionBuyHandler
	AuctionSellHandler
	FriendsHandler
	TransferHandler
)

func (_ *Help) Self() HandlerID           { return HelpHandler }
func (_ *Infa) Self() HandlerID           { return InfaHandler }
func (_ *Who) Self() HandlerID            { return WhoHandler }
func (_ *List) Self() HandlerID           { return ListHandler }
func (_ *Top) Self() HandlerID            { return TopHandler }
func (_ *Mouse) Self() HandlerID          { return MouseHandler }
func (_ *Tiktok) Self() HandlerID         { return TiktokHandler }
func (_ *Game) Self() HandlerID           { return GameHandler }
func (_ *Weather) Self() HandlerID        { return WeatherHandler }
func (_ *Cat) Self() HandlerID            { return CatHandler }
func (_ *Anime) Self() HandlerID          { return AnimeHandler }
func (_ *Furry) Self() HandlerID          { return FurryHandler }
func (_ *Flag) Self() HandlerID           { return FlagHandler }
func (_ *Person) Self() HandlerID         { return PersonHandler }
func (_ *Horse) Self() HandlerID          { return HorseHandler }
func (_ *Art) Self() HandlerID            { return ArtHandler }
func (_ *Car) Self() HandlerID            { return CarHandler }
func (_ *Soy) Self() HandlerID            { return SoyHandler }
func (_ *Danbooru) Self() HandlerID       { return DanbooruHandler }
func (_ *Fap) Self() HandlerID            { return FapHandler }
func (_ *Masyunya) Self() HandlerID       { return MasyunyaHandler }
func (_ *Poppy) Self() HandlerID          { return PoppyHandler }
func (_ *Sima) Self() HandlerID           { return SimaHandler }
func (_ *Lageona) Self() HandlerID        { return LageonaHandler }
func (_ *Hello) Self() HandlerID          { return HelloHandler }
func (_ *Basili) Self() HandlerID         { return BasiliHandler }
func (_ *Casper) Self() HandlerID         { return CasperHandler }
func (_ *Zeus) Self() HandlerID           { return ZeusHandler }
func (_ *Pic) Self() HandlerID            { return PicHandler }
func (_ *Avatar) Self() HandlerID         { return AvatarHandler }
func (_ *TurnOn) Self() HandlerID         { return TurnOnHandler }
func (_ *TurnOff) Self() HandlerID        { return TurnOffHandler }
func (_ *Ban) Self() HandlerID            { return BanHandler }
func (_ *Unban) Self() HandlerID          { return UnbanHandler }
func (_ *Calculator) Self() HandlerID     { return CalculatorHandler }
func (_ *DailyEblan) Self() HandlerID     { return DailyEblanHandler }
func (_ *DailyAdmin) Self() HandlerID     { return DailyAdminHandler }
func (_ *DailyPair) Self() HandlerID      { return DailyPairHandler }
func (_ *Name) Self() HandlerID           { return NameHandler }
func (_ *Inventory) Self() HandlerID      { return InventoryHandler }
func (_ *Sort) Self() HandlerID           { return SortHandler }
func (_ *Catch) Self() HandlerID          { return CatchHandler }
func (_ *Drop) Self() HandlerID           { return DropHandler }
func (_ *Pick) Self() HandlerID           { return PickHandler }
func (_ *Floor) Self() HandlerID          { return FloorHandler }
func (_ *Market) Self() HandlerID         { return MarketHandler }
func (_ *PriceList) Self() HandlerID      { return PriceListHandler }
func (_ *NameMarket) Self() HandlerID     { return NameMarketHandler }
func (_ *GetJob) Self() HandlerID         { return GetJobHandler }
func (_ *QuitJob) Self() HandlerID        { return QuitJobHandler }
func (_ *Buy) Self() HandlerID            { return BuyHandler }
func (_ *Eat) Self() HandlerID            { return EatHandler }
func (_ *EatQuick) Self() HandlerID       { return EatQuickHandler }
func (_ *Fish) Self() HandlerID           { return FishHandler }
func (_ *CastNet) Self() HandlerID        { return CastNetHandler }
func (_ *DrawNet) Self() HandlerID        { return DrawNetHandler }
func (_ *Net) Self() HandlerID            { return NetHandler }
func (_ *FishingRecords) Self() HandlerID { return FishingRecordsHandler }
func (_ *Craft) Self() HandlerID          { return CraftHandler }
func (_ *Status) Self() HandlerID         { return StatusHandler }
func (_ *Sell) Self() HandlerID           { return SellHandler }
func (_ *SellQuick) Self() HandlerID      { return SellQuickHandler }
func (_ *Stack) Self() HandlerID          { return StackHandler }
func (_ *Split) Self() HandlerID          { return SplitHandler }
func (_ *Cashout) Self() HandlerID        { return CashoutHandler }
func (_ *Fight) Self() HandlerID          { return FightHandler }
func (_ *PvP) Self() HandlerID            { return PvPHandler }
func (_ *Profile) Self() HandlerID        { return ProfileHandler }
func (_ *Dice) Self() HandlerID           { return DiceHandler }
func (_ *Roll) Self() HandlerID           { return RollHandler }
func (_ *TopStrong) Self() HandlerID      { return TopStrongHandler }
func (_ *TopRating) Self() HandlerID      { return TopRatingHandler }
func (_ *TopRich) Self() HandlerID        { return TopRichHandler }
func (_ *Capital) Self() HandlerID        { return CapitalHandler }
func (_ *Balance) Self() HandlerID        { return BalanceHandler }
func (_ *Funds) Self() HandlerID          { return FundsHandler }
func (_ *Energy) Self() HandlerID         { return EnergyHandler }
func (_ *NamePet) Self() HandlerID        { return NamePetHandler }
func (_ *ReceiveSMS) Self() HandlerID     { return ReceiveSMSHandler }
func (_ *SendSMS) Self() HandlerID        { return SendSMSHandler }
func (_ *Contacts) Self() HandlerID       { return ContactsHandler }
func (_ *Spam) Self() HandlerID           { return SpamHandler }
func (_ *Farm) Self() HandlerID           { return FarmHandler }
func (_ *Plant) Self() HandlerID          { return PlantHandler }
func (_ *Harvest) Self() HandlerID        { return HarvestHandler }
func (_ *HarvestInline) Self() HandlerID  { return HarvestInlineHandler }
func (_ *UpgradeFarm) Self() HandlerID    { return UpgradeFarmHandler }
func (_ *Auction) Self() HandlerID        { return AuctionHandler }
func (_ *AuctionBuy) Self() HandlerID     { return AuctionBuyHandler }
func (_ *AuctionSell) Self() HandlerID    { return AuctionSellHandler }
func (_ *Friends) Self() HandlerID        { return FriendsHandler }
func (_ *Transfer) Self() HandlerID       { return TransferHandler }
