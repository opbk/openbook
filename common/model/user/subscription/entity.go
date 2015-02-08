package subscription

import (
	"time"

	"github.com/opbk/openbook/common/model/subscription"
)

type UserSubscription struct {
	subscription.Subscription
	UserId     int64
	Expiration time.Time
}
