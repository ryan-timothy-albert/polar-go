// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Timeframe string

const (
	TimeframeYear  Timeframe = "year"
	TimeframeMonth Timeframe = "month"
	TimeframeDay   Timeframe = "day"
)

func (e Timeframe) ToPointer() *Timeframe {
	return &e
}
func (e *Timeframe) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "year":
		fallthrough
	case "month":
		fallthrough
	case "day":
		*e = Timeframe(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Timeframe: %v", v)
	}
}

type BenefitLicenseKeyExpirationProperties struct {
	TTL       int64     `json:"ttl"`
	Timeframe Timeframe `json:"timeframe"`
}

func (o *BenefitLicenseKeyExpirationProperties) GetTTL() int64 {
	if o == nil {
		return 0
	}
	return o.TTL
}

func (o *BenefitLicenseKeyExpirationProperties) GetTimeframe() Timeframe {
	if o == nil {
		return Timeframe("")
	}
	return o.Timeframe
}
