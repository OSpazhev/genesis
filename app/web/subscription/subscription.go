package subscription

import (
	"errors"
	"github.com/ospazhev/genesis/app/db"
)

func Subscribe(email string) error {
	if db.Db.IsSubscriber(email) == true {
		return errors.New("subscription is already subscribed")
	}
	db.Db.AddSubscriber(email)

	return nil
}
