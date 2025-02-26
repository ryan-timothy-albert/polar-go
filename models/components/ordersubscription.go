// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/polarsource/polar-go/internal/utils"
	"time"
)

type OrderSubscriptionMetadataType string

const (
	OrderSubscriptionMetadataTypeStr     OrderSubscriptionMetadataType = "str"
	OrderSubscriptionMetadataTypeInteger OrderSubscriptionMetadataType = "integer"
	OrderSubscriptionMetadataTypeBoolean OrderSubscriptionMetadataType = "boolean"
)

type OrderSubscriptionMetadata struct {
	Str     *string `queryParam:"inline"`
	Integer *int64  `queryParam:"inline"`
	Boolean *bool   `queryParam:"inline"`

	Type OrderSubscriptionMetadataType
}

func CreateOrderSubscriptionMetadataStr(str string) OrderSubscriptionMetadata {
	typ := OrderSubscriptionMetadataTypeStr

	return OrderSubscriptionMetadata{
		Str:  &str,
		Type: typ,
	}
}

func CreateOrderSubscriptionMetadataInteger(integer int64) OrderSubscriptionMetadata {
	typ := OrderSubscriptionMetadataTypeInteger

	return OrderSubscriptionMetadata{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateOrderSubscriptionMetadataBoolean(boolean bool) OrderSubscriptionMetadata {
	typ := OrderSubscriptionMetadataTypeBoolean

	return OrderSubscriptionMetadata{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *OrderSubscriptionMetadata) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = OrderSubscriptionMetadataTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = OrderSubscriptionMetadataTypeInteger
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = OrderSubscriptionMetadataTypeBoolean
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for OrderSubscriptionMetadata", string(data))
}

func (u OrderSubscriptionMetadata) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type OrderSubscriptionMetadata: all fields are null")
}

type OrderSubscription struct {
	Metadata map[string]OrderSubscriptionMetadata `json:"metadata"`
	// Creation timestamp of the object.
	CreatedAt time.Time `json:"created_at"`
	// Last modification timestamp of the object.
	ModifiedAt *time.Time `json:"modified_at"`
	// The ID of the object.
	ID                 string                        `json:"id"`
	Amount             *int64                        `json:"amount"`
	Currency           *string                       `json:"currency"`
	RecurringInterval  SubscriptionRecurringInterval `json:"recurring_interval"`
	Status             SubscriptionStatus            `json:"status"`
	CurrentPeriodStart time.Time                     `json:"current_period_start"`
	CurrentPeriodEnd   *time.Time                    `json:"current_period_end"`
	CancelAtPeriodEnd  bool                          `json:"cancel_at_period_end"`
	StartedAt          *time.Time                    `json:"started_at"`
	EndedAt            *time.Time                    `json:"ended_at"`
	CustomerID         string                        `json:"customer_id"`
	ProductID          string                        `json:"product_id"`
	PriceID            string                        `json:"price_id"`
	DiscountID         *string                       `json:"discount_id"`
	CheckoutID         *string                       `json:"checkout_id"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	UserID string `json:"user_id"`
}

func (o OrderSubscription) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrderSubscription) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrderSubscription) GetMetadata() map[string]OrderSubscriptionMetadata {
	if o == nil {
		return map[string]OrderSubscriptionMetadata{}
	}
	return o.Metadata
}

func (o *OrderSubscription) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *OrderSubscription) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *OrderSubscription) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OrderSubscription) GetAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *OrderSubscription) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *OrderSubscription) GetRecurringInterval() SubscriptionRecurringInterval {
	if o == nil {
		return SubscriptionRecurringInterval("")
	}
	return o.RecurringInterval
}

func (o *OrderSubscription) GetStatus() SubscriptionStatus {
	if o == nil {
		return SubscriptionStatus("")
	}
	return o.Status
}

func (o *OrderSubscription) GetCurrentPeriodStart() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CurrentPeriodStart
}

func (o *OrderSubscription) GetCurrentPeriodEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.CurrentPeriodEnd
}

func (o *OrderSubscription) GetCancelAtPeriodEnd() bool {
	if o == nil {
		return false
	}
	return o.CancelAtPeriodEnd
}

func (o *OrderSubscription) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *OrderSubscription) GetEndedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndedAt
}

func (o *OrderSubscription) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *OrderSubscription) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *OrderSubscription) GetPriceID() string {
	if o == nil {
		return ""
	}
	return o.PriceID
}

func (o *OrderSubscription) GetDiscountID() *string {
	if o == nil {
		return nil
	}
	return o.DiscountID
}

func (o *OrderSubscription) GetCheckoutID() *string {
	if o == nil {
		return nil
	}
	return o.CheckoutID
}

func (o *OrderSubscription) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
