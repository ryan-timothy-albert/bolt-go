// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type PaymentMethodExtendedType string

const (
	PaymentMethodExtendedTypeID            PaymentMethodExtendedType = "id"
	PaymentMethodExtendedTypeCreditCard    PaymentMethodExtendedType = "credit_card"
	PaymentMethodExtendedTypePaypal        PaymentMethodExtendedType = "paypal"
	PaymentMethodExtendedTypeAffirm        PaymentMethodExtendedType = "affirm"
	PaymentMethodExtendedTypeAfterpay      PaymentMethodExtendedType = "afterpay"
	PaymentMethodExtendedTypeKlarna        PaymentMethodExtendedType = "klarna"
	PaymentMethodExtendedTypeKlarnaAccount PaymentMethodExtendedType = "klarna_account"
	PaymentMethodExtendedTypeKlarnaPaynow  PaymentMethodExtendedType = "klarna_paynow"
)

type PaymentMethodExtended struct {
	PaymentMethodReference       *PaymentMethodReference
	PaymentMethodCreditCardInput *PaymentMethodCreditCardInput
	PaymentMethodPaypal          *PaymentMethodPaypal
	PaymentMethodAffirm          *PaymentMethodAffirm
	PaymentMethodAfterpay        *PaymentMethodAfterpay
	PaymentMethodKlarna          *PaymentMethodKlarna
	PaymentMethodKlarnaAccount   *PaymentMethodKlarnaAccount
	PaymentMethodKlarnaPaynow    *PaymentMethodKlarnaPaynow

	Type PaymentMethodExtendedType
}

func CreatePaymentMethodExtendedID(id PaymentMethodReference) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeID

	typStr := PaymentMethodReferenceTag(typ)
	id.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodReference: &id,
		Type:                   typ,
	}
}

func CreatePaymentMethodExtendedCreditCard(creditCard PaymentMethodCreditCardInput) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeCreditCard

	typStr := DotTag(typ)
	creditCard.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodCreditCardInput: &creditCard,
		Type:                         typ,
	}
}

func CreatePaymentMethodExtendedPaypal(paypal PaymentMethodPaypal) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypePaypal

	typStr := PaymentMethodPaypalTag(typ)
	paypal.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodPaypal: &paypal,
		Type:                typ,
	}
}

func CreatePaymentMethodExtendedAffirm(affirm PaymentMethodAffirm) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeAffirm

	typStr := PaymentMethodAffirmTag(typ)
	affirm.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodAffirm: &affirm,
		Type:                typ,
	}
}

func CreatePaymentMethodExtendedAfterpay(afterpay PaymentMethodAfterpay) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeAfterpay

	typStr := PaymentMethodAfterpayTag(typ)
	afterpay.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodAfterpay: &afterpay,
		Type:                  typ,
	}
}

func CreatePaymentMethodExtendedKlarna(klarna PaymentMethodKlarna) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeKlarna

	typStr := PaymentMethodKlarnaTag(typ)
	klarna.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodKlarna: &klarna,
		Type:                typ,
	}
}

func CreatePaymentMethodExtendedKlarnaAccount(klarnaAccount PaymentMethodKlarnaAccount) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeKlarnaAccount

	typStr := PaymentMethodKlarnaAccountTag(typ)
	klarnaAccount.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodKlarnaAccount: &klarnaAccount,
		Type:                       typ,
	}
}

func CreatePaymentMethodExtendedKlarnaPaynow(klarnaPaynow PaymentMethodKlarnaPaynow) PaymentMethodExtended {
	typ := PaymentMethodExtendedTypeKlarnaPaynow

	typStr := PaymentMethodKlarnaPaynowTag(typ)
	klarnaPaynow.DotTag = typStr

	return PaymentMethodExtended{
		PaymentMethodKlarnaPaynow: &klarnaPaynow,
		Type:                      typ,
	}
}

func (u *PaymentMethodExtended) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string `json:".tag"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "id":
		paymentMethodReference := new(PaymentMethodReference)
		if err := utils.UnmarshalJSON(data, &paymentMethodReference, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodReference = paymentMethodReference
		u.Type = PaymentMethodExtendedTypeID
		return nil
	case "credit_card":
		paymentMethodCreditCardInput := new(PaymentMethodCreditCardInput)
		if err := utils.UnmarshalJSON(data, &paymentMethodCreditCardInput, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodCreditCardInput = paymentMethodCreditCardInput
		u.Type = PaymentMethodExtendedTypeCreditCard
		return nil
	case "paypal":
		paymentMethodPaypal := new(PaymentMethodPaypal)
		if err := utils.UnmarshalJSON(data, &paymentMethodPaypal, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodPaypal = paymentMethodPaypal
		u.Type = PaymentMethodExtendedTypePaypal
		return nil
	case "affirm":
		paymentMethodAffirm := new(PaymentMethodAffirm)
		if err := utils.UnmarshalJSON(data, &paymentMethodAffirm, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodAffirm = paymentMethodAffirm
		u.Type = PaymentMethodExtendedTypeAffirm
		return nil
	case "afterpay":
		paymentMethodAfterpay := new(PaymentMethodAfterpay)
		if err := utils.UnmarshalJSON(data, &paymentMethodAfterpay, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodAfterpay = paymentMethodAfterpay
		u.Type = PaymentMethodExtendedTypeAfterpay
		return nil
	case "klarna":
		paymentMethodKlarna := new(PaymentMethodKlarna)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarna, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodKlarna = paymentMethodKlarna
		u.Type = PaymentMethodExtendedTypeKlarna
		return nil
	case "klarna_account":
		paymentMethodKlarnaAccount := new(PaymentMethodKlarnaAccount)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaAccount, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodKlarnaAccount = paymentMethodKlarnaAccount
		u.Type = PaymentMethodExtendedTypeKlarnaAccount
		return nil
	case "klarna_paynow":
		paymentMethodKlarnaPaynow := new(PaymentMethodKlarnaPaynow)
		if err := utils.UnmarshalJSON(data, &paymentMethodKlarnaPaynow, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PaymentMethodKlarnaPaynow = paymentMethodKlarnaPaynow
		u.Type = PaymentMethodExtendedTypeKlarnaPaynow
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PaymentMethodExtended) MarshalJSON() ([]byte, error) {
	if u.PaymentMethodReference != nil {
		return utils.MarshalJSON(u.PaymentMethodReference, "", true)
	}

	if u.PaymentMethodCreditCardInput != nil {
		return utils.MarshalJSON(u.PaymentMethodCreditCardInput, "", true)
	}

	if u.PaymentMethodPaypal != nil {
		return utils.MarshalJSON(u.PaymentMethodPaypal, "", true)
	}

	if u.PaymentMethodAffirm != nil {
		return utils.MarshalJSON(u.PaymentMethodAffirm, "", true)
	}

	if u.PaymentMethodAfterpay != nil {
		return utils.MarshalJSON(u.PaymentMethodAfterpay, "", true)
	}

	if u.PaymentMethodKlarna != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarna, "", true)
	}

	if u.PaymentMethodKlarnaAccount != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaAccount, "", true)
	}

	if u.PaymentMethodKlarnaPaynow != nil {
		return utils.MarshalJSON(u.PaymentMethodKlarnaPaynow, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
