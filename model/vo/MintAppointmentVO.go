package vo

// user make an appmointment to join airdrop
type MintAppointmentVO struct {
	WalletAddress  string `json:walletAddress`
	AirdropTokenId string `json:airdropTokenId`
}
