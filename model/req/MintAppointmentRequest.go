package req

// user make an appmointment to join airdrop
type MintAppointmentRequest struct {
	WalletAddress  string `json:walletAddress`
	AirdropTokenId string `json:airdropTokenId`
}
