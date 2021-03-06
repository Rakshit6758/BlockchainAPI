package models

// block 
type Block struct{
	BlockNumber     		int64        	`json:"blocknumber"`
	Timestamp       		uint64       	`json:"timestamp"`
	Difficulty      		uint64       	`json:"difficulty"`
	Hash            		string       	`json:"hash"`
	TransactionsCount     	int       	 	`json:"transactionsCount"`
	Transactions            []Transaction   `json:"transactions"`
}

// transaction
type Transaction struct{
	Hash 		string		`json:"hash"`
	Value		string		`json:"value"`
	Gas			uint64		`json:"gas"`
	GasPrice	uint64		`json:"gasprice"`
	Nonce		uint64		`json:"nonce"`	
	To 			string		`json:"to"`
	Pending		bool		`json:"pending"`
	Hero		string		`json:"hero"`
}

// transafer ethereum request
type TransferEthRequest struct {
	PrivKey string `json:"privKey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// hash response
type HashResponse struct {
	Hash string `json:"hash"`
}

// balance response
type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"units"`
}

//error
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}