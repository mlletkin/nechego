// Code generated by "stringer -type=HandlerID -output=handlers_string.go"; DO NOT EDIT.

package handlers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoHandler-0]
	_ = x[InfaHandler-1]
	_ = x[WhoHandler-2]
	_ = x[ListHandler-3]
	_ = x[TopHandler-4]
	_ = x[MouseHandler-5]
	_ = x[TiktokHandler-6]
	_ = x[GameHandler-7]
	_ = x[WeatherHandler-8]
	_ = x[CatHandler-9]
	_ = x[AnimeHandler-10]
	_ = x[FurryHandler-11]
	_ = x[FlagHandler-12]
	_ = x[PersonHandler-13]
	_ = x[HorseHandler-14]
	_ = x[ArtHandler-15]
	_ = x[CarHandler-16]
	_ = x[SoyHandler-17]
	_ = x[DanbooruHandler-18]
	_ = x[FapHandler-19]
	_ = x[MasyunyaHandler-20]
	_ = x[PoppyHandler-21]
	_ = x[SimaHandler-22]
	_ = x[LageonaHandler-23]
	_ = x[HelloHandler-24]
	_ = x[BasiliHandler-25]
	_ = x[CasperHandler-26]
	_ = x[ZeusHandler-27]
	_ = x[PicHandler-28]
	_ = x[AvatarHandler-29]
	_ = x[TurnOnHandler-30]
	_ = x[TurnOffHandler-31]
	_ = x[BanHandler-32]
	_ = x[UnbanHandler-33]
	_ = x[CalculatorHandler-34]
	_ = x[DailyEblanHandler-35]
	_ = x[DailyAdminHandler-36]
	_ = x[DailyPairHandler-37]
	_ = x[NameHandler-38]
	_ = x[InventoryHandler-39]
	_ = x[SortHandler-40]
	_ = x[CatchHandler-41]
	_ = x[DropHandler-42]
	_ = x[PickHandler-43]
	_ = x[FloorHandler-44]
	_ = x[MarketHandler-45]
	_ = x[NameMarketHandler-46]
	_ = x[GetJobHandler-47]
	_ = x[QuitJobHandler-48]
	_ = x[BuyHandler-49]
	_ = x[EatHandler-50]
	_ = x[EatQuickHandler-51]
	_ = x[FishHandler-52]
	_ = x[CastNetHandler-53]
	_ = x[DrawNetHandler-54]
	_ = x[NetHandler-55]
	_ = x[FishingRecordsHandler-56]
	_ = x[CraftHandler-57]
	_ = x[StatusHandler-58]
	_ = x[SellHandler-59]
	_ = x[SellQuickHandler-60]
	_ = x[StackHandler-61]
	_ = x[CashoutHandler-62]
	_ = x[FightHandler-63]
	_ = x[PvPHandler-64]
	_ = x[ProfileHandler-65]
	_ = x[DiceHandler-66]
	_ = x[RollHandler-67]
	_ = x[TopStrongHandler-68]
	_ = x[TopRatingHandler-69]
	_ = x[TopRichHandler-70]
	_ = x[CapitalHandler-71]
	_ = x[BalanceHandler-72]
	_ = x[FundsHandler-73]
	_ = x[EnergyHandler-74]
	_ = x[NamePetHandler-75]
	_ = x[ReceiveSMSHandler-76]
	_ = x[SendSMSHandler-77]
	_ = x[ContactsHandler-78]
	_ = x[SpamHandler-79]
}

const _HandlerID_name = "NoHandlerInfaHandlerWhoHandlerListHandlerTopHandlerMouseHandlerTiktokHandlerGameHandlerWeatherHandlerCatHandlerAnimeHandlerFurryHandlerFlagHandlerPersonHandlerHorseHandlerArtHandlerCarHandlerSoyHandlerDanbooruHandlerFapHandlerMasyunyaHandlerPoppyHandlerSimaHandlerLageonaHandlerHelloHandlerBasiliHandlerCasperHandlerZeusHandlerPicHandlerAvatarHandlerTurnOnHandlerTurnOffHandlerBanHandlerUnbanHandlerCalculatorHandlerDailyEblanHandlerDailyAdminHandlerDailyPairHandlerNameHandlerInventoryHandlerSortHandlerCatchHandlerDropHandlerPickHandlerFloorHandlerMarketHandlerNameMarketHandlerGetJobHandlerQuitJobHandlerBuyHandlerEatHandlerEatQuickHandlerFishHandlerCastNetHandlerDrawNetHandlerNetHandlerFishingRecordsHandlerCraftHandlerStatusHandlerSellHandlerSellQuickHandlerStackHandlerCashoutHandlerFightHandlerPvPHandlerProfileHandlerDiceHandlerRollHandlerTopStrongHandlerTopRatingHandlerTopRichHandlerCapitalHandlerBalanceHandlerFundsHandlerEnergyHandlerNamePetHandlerReceiveSMSHandlerSendSMSHandlerContactsHandlerSpamHandler"

var _HandlerID_index = [...]uint16{0, 9, 20, 30, 41, 51, 63, 76, 87, 101, 111, 123, 135, 146, 159, 171, 181, 191, 201, 216, 226, 241, 253, 264, 278, 290, 303, 316, 327, 337, 350, 363, 377, 387, 399, 416, 433, 450, 466, 477, 493, 504, 516, 527, 538, 550, 563, 580, 593, 607, 617, 627, 642, 653, 667, 681, 691, 712, 724, 737, 748, 764, 776, 790, 802, 812, 826, 837, 848, 864, 880, 894, 908, 922, 934, 947, 961, 978, 992, 1007, 1018}

func (i HandlerID) String() string {
	if i < 0 || i >= HandlerID(len(_HandlerID_index)-1) {
		return "HandlerID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _HandlerID_name[_HandlerID_index[i]:_HandlerID_index[i+1]]
}
