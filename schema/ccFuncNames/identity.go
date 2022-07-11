package ccfuncnames

// Transacciones del contrato de prueba "chaincode-go" de test-network-optativo-nanobash\chaincodes
const (
	CC1InitLedger  = "InitLedger"
	CC1ReadAsset   = "ReadAsset"
	CC1CreateAsset = "CreateAsset"

	ContactFileCreate        = "CreateFile"
	ContactFileGetById       = "ReadFile"
	ContactFileUpdate        = "UpdateFile"
	ContactFileDelete        = "DeleteFile"
	ContactFileTransfer      = "TransferFile"
	ContactFileGetAll        = "GetAllFiles"
	ContactFileGetAllByOwner = "GetAllFilesByOwner"
	ContactFileHistory       = "FilesHistory"

	//======== UserContact ==========================================================

	ContactUserCreate = "CreateUser"
	ContactUserRead   = "ReadUser"
	ContactUserGetAll = "GetAllUsers"
	ContactUserDelete = "DeleteUser"

	MYCCInitLedger  = "InitLedger"
	MYCCReadAsset   = "ReadAsset"
	MYCCCreateAsset = "CreateAsset"
	MYCCUpdateAsset = "UpdateAsset"
)

const (
	ContractNameCC1 = "mycc"
)
