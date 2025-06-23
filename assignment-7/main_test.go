package main

import (
	"testing"
)

// Test CreditCard Pay
func TestCreditCardPay(t *testing.T) {
	card := CreditCard{CardNumber: "1234567890123456"}
	expected := "[CreditCard] Paid 100.00 using card ending with 3456\n"
	result := card.Pay(100.0)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test PayPal Pay
func TestPayPalPay(t *testing.T) {
	paypal := PayPal{Email: "user@example.com"}
	expected := "[PayPal] Paid 100.00 using PayPal account: user@example.com\n"
	result := paypal.Pay(100.0)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test UPI Pay
func TestUPIPay(t *testing.T) {
	upi := UPI{UPIid: "user@upi"}
	expected := "[UPI] Paid 100.00 using UPI: user@upi\n"
	result := upi.Pay(100.0)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test OTP generation for CreditCard
func TestCreditCardOTP(t *testing.T) {
	card := CreditCard{}
	expected := "[CreditCard] OTP sent to registered number "
	result := card.generateOtp()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test OTP generation for UPI
func TestUPIOTP(t *testing.T) {
	upi := UPI{}
	expected := "[UPI] OTP sent to registered number "
	result := upi.generateOtp()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
