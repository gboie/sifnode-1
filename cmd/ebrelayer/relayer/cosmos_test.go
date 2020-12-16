package relayer

import (
	"testing"

	"github.com/Sifchain/sifnode/pkg/siflogger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	tmProvider      = "Node"
	ethProvider     = "ws://127.0.0.1:7545/"
	contractAddress = "0x00"
	privateKeyStr   = "ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f"
)

func TestNewCosmosSub(t *testing.T) {

	privateKey, _ := crypto.HexToECDSA(privateKeyStr)
	logger := siflogger.New(siflogger.TDFmt)
	registryContractAddress := common.HexToAddress(contractAddress)
	sub := NewCosmosSub(tmProvider, ethProvider, registryContractAddress,
		privateKey, logger)
	require.NotEqual(t, sub, nil)
}
