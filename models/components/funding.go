// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Funding struct {
	FundingGoal *CurrencyAmount `json:"funding_goal,omitempty"`
	// Sum of pledges to this isuse (including currently open pledges and pledges that have been paid out). Always in USD.
	PledgesSum *CurrencyAmount `json:"pledges_sum,omitempty"`
}

func (o *Funding) GetFundingGoal() *CurrencyAmount {
	if o == nil {
		return nil
	}
	return o.FundingGoal
}

func (o *Funding) GetPledgesSum() *CurrencyAmount {
	if o == nil {
		return nil
	}
	return o.PledgesSum
}
