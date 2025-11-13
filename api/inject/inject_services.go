package inject

import (
	walletStorage "wallet/internal/repository/wallet"

	"wallet/internal/service/wallet"

	"github.com/google/wire"
)

// wire Set for loading the services.
var serviceSet = wire.NewSet( // nolint
	wallet.NewService,
	walletStorage.NewStore,
)
