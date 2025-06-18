package main

import (
	"fmt"
)

type OTPSupport interface {
	generateOtp() string
}

// PaymentMethod interface
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard type
type CreditCard struct {
	CardNumber string
}

// PayPal type
type PayPal struct {
	Email string
}

// UPI type
type UPI struct {
	UPIid string
}

// Pay method for CreditCard
func (c CreditCard) Pay(amount float64) string {
	last4digits := c.CardNumber[len(c.CardNumber)-4:]
	return fmt.Sprintf("[CreditCard] Paid %.2f using card ending with %s\n", amount, last4digits)
}

// Pay method for PayPal
func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("[PayPal] Paid %.2f using PayPal account: %s\n", amount, p.Email)
}

// Pay method for UPI
func (u UPI) Pay(amount float64) string {
	return fmt.Sprintf("[UPI] Paid %.2f using UPI: %s\n", amount, u.UPIid)
}

// generate otp functionality
func (c CreditCard) generateOtp() string {
	return "[CreditCard] OTP sent to registered number "
}
func (u UPI) generateOtp() string {
	return "[UPI] OTP sent to registered number "
}

func main() {
	fmt.Println("Interface Type 3 : Interface with Type Assertion & Custom Behavior")

	fmt.Println("Enter default value for Card Number")
	var cardNumber string
	fmt.Scan(&cardNumber)

	fmt.Println("Enter default value for PayPal Email")
	var emailid string
	fmt.Scan(&emailid)

	fmt.Println("Enter default value for UPI ID")
	var upiid string
	fmt.Scan(&upiid)

	methods := []PaymentMethod{
		CreditCard{cardNumber},
		PayPal{emailid},
		UPI{upiid},
	}

	var amount float64
	fmt.Println("Enter amount to pay : ")
	fmt.Scan(&amount)

	for _, method := range methods {
		if otpMethod, ok := method.(OTPSupport); ok == true { //first assigning value to ok , then validating it's if condition
			fmt.Println(otpMethod.generateOtp())
		}

		fmt.Println(method.Pay(amount))
	}

}
