package wallet

import (
	"context"

	"github.com/google/uuid"

	"wallet/internal/core"
)

type accountService struct {
	accStg core.AccountStore
}

func NewService(
	account core.AccountStore,

) core.WalletService {
	return &accountService{
		accStg: account,
	}
}

// core.AccountService interface
func (s *accountService) Create(ctx context.Context, data core.RegistrationData) (core.CreationOk, error) {
	// random UUID v4
	accID := uuid.New()

	// // FixMe: It should be a transaction with usrStg
	// if err = s.accStg.Create(ctx, core.AccountCreationData{AccountID: accID, Password: pwd}); err != nil {
	// 	return core.CreationOk{}, err
	// }

	// if err = s.infoStg.LinkToAccount(ctx, accID, data.Info); err != nil {
	// 	return core.CreationOk{}, err
	// }

	return core.CreationOk{
		UserID: accID,
	}, nil
}

// core.AccountService interface
func (s *accountService) GetInfo(ctx context.Context, userID uuid.UUID) (core.PersonalInfoEntity, error) {
	// return s.infoStg.ReadInfo(ctx, userID)

	return core.PersonalInfoEntity{}, nil
}

// core.AccountService interface
func (s *accountService) SearchAccounts(ctx context.Context, parameters core.SearchAccountsInfoParams) ([]core.PersonalInfoEntity, error) {
	// people, err := s.infoStg.GetInfoList(ctx, parameters)
	// if err != nil {
	// 	return []core.PersonalInfoEntity{}, err
	// }

	// sort.Slice(people, func(i, j int) bool {
	// 	for index := range len(people[i].UserID) {
	// 		if people[i].UserID[index] < people[j].UserID[index] {
	// 			return true
	// 		} else if people[i].UserID[index] < people[j].UserID[index] {
	// 			return false
	// 		}
	// 	}
	// 	return false
	// })

	// return people, nil

	return []core.PersonalInfoEntity{}, nil
}
