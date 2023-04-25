// Code generated by "stringer -type=HandlerID -output=handlers_string.go"; DO NOT EDIT.

package handlers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoHandler-0]
	_ = x[HelpHandler-1]
	_ = x[InfaHandler-2]
	_ = x[WhoHandler-3]
	_ = x[ListHandler-4]
	_ = x[TopHandler-5]
	_ = x[GameHandler-6]
	_ = x[WeatherHandler-7]
	_ = x[AvatarHandler-8]
	_ = x[TurnOnHandler-9]
	_ = x[TurnOffHandler-10]
	_ = x[BanHandler-11]
	_ = x[UnbanHandler-12]
	_ = x[CalculatorHandler-13]
	_ = x[DailyEblanHandler-14]
	_ = x[DailyAdminHandler-15]
	_ = x[DailyPairHandler-16]
	_ = x[NameHandler-17]
	_ = x[InventoryHandler-18]
	_ = x[SortHandler-19]
	_ = x[CatchHandler-20]
	_ = x[DropHandler-21]
	_ = x[PickHandler-22]
	_ = x[FloorHandler-23]
	_ = x[MarketHandler-24]
	_ = x[PriceListHandler-25]
	_ = x[NameMarketHandler-26]
	_ = x[GetJobHandler-27]
	_ = x[QuitJobHandler-28]
	_ = x[BuyHandler-29]
	_ = x[EatHandler-30]
	_ = x[EatQuickHandler-31]
	_ = x[FishHandler-32]
	_ = x[CastNetHandler-33]
	_ = x[DrawNetHandler-34]
	_ = x[NetHandler-35]
	_ = x[FishingRecordsHandler-36]
	_ = x[CraftHandler-37]
	_ = x[StatusHandler-38]
	_ = x[SellHandler-39]
	_ = x[SellQuickHandler-40]
	_ = x[StackHandler-41]
	_ = x[SplitHandler-42]
	_ = x[CashoutHandler-43]
	_ = x[FightHandler-44]
	_ = x[PvPHandler-45]
	_ = x[ProfileHandler-46]
	_ = x[TopStrongHandler-47]
	_ = x[TopRatingHandler-48]
	_ = x[TopRichHandler-49]
	_ = x[CapitalHandler-50]
	_ = x[BalanceHandler-51]
	_ = x[FundsHandler-52]
	_ = x[EnergyHandler-53]
	_ = x[NamePetHandler-54]
	_ = x[ReceiveSMSHandler-55]
	_ = x[SendSMSHandler-56]
	_ = x[ContactsHandler-57]
	_ = x[SpamHandler-58]
	_ = x[FarmHandler-59]
	_ = x[PlantHandler-60]
	_ = x[HarvestHandler-61]
	_ = x[HarvestInlineHandler-62]
	_ = x[UpgradeFarmHandler-63]
	_ = x[NameFarmHandler-64]
	_ = x[AuctionHandler-65]
	_ = x[AuctionBuyHandler-66]
	_ = x[AuctionSellHandler-67]
	_ = x[FriendsHandler-68]
	_ = x[TransferHandler-69]
	_ = x[UseHandler-70]
}

const _HandlerID_name = "NoHandlerHelpHandlerInfaHandlerWhoHandlerListHandlerTopHandlerGameHandlerWeatherHandlerAvatarHandlerTurnOnHandlerTurnOffHandlerBanHandlerUnbanHandlerCalculatorHandlerDailyEblanHandlerDailyAdminHandlerDailyPairHandlerNameHandlerInventoryHandlerSortHandlerCatchHandlerDropHandlerPickHandlerFloorHandlerMarketHandlerPriceListHandlerNameMarketHandlerGetJobHandlerQuitJobHandlerBuyHandlerEatHandlerEatQuickHandlerFishHandlerCastNetHandlerDrawNetHandlerNetHandlerFishingRecordsHandlerCraftHandlerStatusHandlerSellHandlerSellQuickHandlerStackHandlerSplitHandlerCashoutHandlerFightHandlerPvPHandlerProfileHandlerTopStrongHandlerTopRatingHandlerTopRichHandlerCapitalHandlerBalanceHandlerFundsHandlerEnergyHandlerNamePetHandlerReceiveSMSHandlerSendSMSHandlerContactsHandlerSpamHandlerFarmHandlerPlantHandlerHarvestHandlerHarvestInlineHandlerUpgradeFarmHandlerNameFarmHandlerAuctionHandlerAuctionBuyHandlerAuctionSellHandlerFriendsHandlerTransferHandlerUseHandler"

var _HandlerID_index = [...]uint16{0, 9, 20, 31, 41, 52, 62, 73, 87, 100, 113, 127, 137, 149, 166, 183, 200, 216, 227, 243, 254, 266, 277, 288, 300, 313, 329, 346, 359, 373, 383, 393, 408, 419, 433, 447, 457, 478, 490, 503, 514, 530, 542, 554, 568, 580, 590, 604, 620, 636, 650, 664, 678, 690, 703, 717, 734, 748, 763, 774, 785, 797, 811, 831, 849, 864, 878, 895, 913, 927, 942, 952}

func (i HandlerID) String() string {
	if i < 0 || i >= HandlerID(len(_HandlerID_index)-1) {
		return "HandlerID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _HandlerID_name[_HandlerID_index[i]:_HandlerID_index[i+1]]
}
