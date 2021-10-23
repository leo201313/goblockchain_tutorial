//constcoe.go
package constcoe

const (
	Difficulty          = 12
	InitCoin            = 1000
	TransactionPoolFile = "./tmp/transaction_pool.data"
	BCPath              = "./tmp/blocks"
	BCFile              = "./tmp/blocks/MANIFEST"
	ChecksumLength      = 4                 // new
	NetworkVersion      = byte(0x00)        // new
	Wallets             = "./tmp/wallets/"  // new
	WalletsRefList      = "./tmp/ref_list/" // new
)
