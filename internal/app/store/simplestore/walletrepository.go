package simplestore

import (
	"errors"
	"github.com/adilku/vote_server/internal/app/model"
)

type WalletRepository struct {
	store  *Store
	wallet map[int]int
}

func (r *WalletRepository) Transfer(idSender int, idReceiver int, delta int) error {
	if r.wallet[idSender] >= delta {
		r.wallet[idSender] -= delta
		r.wallet[idReceiver] += delta
		return nil
	} else {
		return errors.New("Bad transaction")
	}
}

func (r *WalletRepository) ChangeBalance(id int, delta int) error {
	if _, ok := r.wallet[id]; !ok {
		r.wallet[id] = 0
	}
	old := r.wallet[id]
	if old + delta < 0 {
		return errors.New("insufficient funds")
	}
	r.wallet[id] = old + delta
	return nil
}

func (r *WalletRepository) Create(u *model.Wallet) error {
	//u.ID = len(r.wallet)
	r.wallet[u.ID] = u.Balance
	return nil
}

func (r *WalletRepository) FindById(id int) (*model.Wallet, error) {
	u, exist := r.wallet[id]
	if !exist {
		return nil, errors.New("Not found")
	}

	return &model.Wallet{ID : id, Balance: u}, nil
}


