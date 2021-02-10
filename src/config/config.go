package config

// store path
const DatasourceDir = "data/"

// database path
const DatabaseName = "naivechain.db"
const DatabaseURI = DatasourceDir + DatabaseName

// block database
const BlockBucketName = "block"
const BlockCursorName = "latest"

// ket and address database
const KeyBucketName = "key"

// UTxO database
const UTxOBucketName = "unspent"

// block and blockchain
const InitialTarget = int64(14)

// address coding
const KeyVersion = byte(0x00)
const AddressChecksumLen = 4

// transaction
const CoinbaseReward = 10
