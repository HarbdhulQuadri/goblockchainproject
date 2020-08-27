package entities

type Bitcoin struct{
	//geting started with our variables
	address string `json:"address"`
	wallet string `json:"wallet"`
	hdWallet string `json:"hd_wallet"`
	totalReceived string `json:"total_received"`
	totalSent string `json:"total_sent"`
	balance string `json:"balance"`
	unconfirmedBalance string `json:"unconfirmed_balance"`
	finalBalance string `json:"final_balance"`
}