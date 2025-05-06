package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
two types of payment
1. card
2. wallet

2 regions
1. india
2, SEA
*/

// Product Interface
type cardPayment interface {
	pay(rs retryStrategy)
}
type walletPayment interface {
	pay(rs retryStrategy)
}

// Factory Interface
type paymentProcessorFactory interface {
	createCardProcessor() cardPayment
	createWalletPayment() walletPayment
}

func GetPaymentFactory(country string) paymentProcessorFactory {
	switch country {
	case "india":
		return &indianProcessor{}

	case "sea":
		return &seaProcessor{}
	default:
		return nil
	}
}

// Concrete factory

type indianProcessor struct {
}

func (i *indianProcessor) createCardProcessor() cardPayment {
	return &razorpay{}
}

func (i *indianProcessor) createWalletPayment() walletPayment {
	return &paytm{}
}

type seaProcessor struct {
}

func (i *seaProcessor) createCardProcessor() cardPayment {
	return &stripe{}
}

func (i *seaProcessor) createWalletPayment() walletPayment {
	return &grabpay{}
}

// concrete product
type razorpay struct {
}

func (r *razorpay) pay(rs retryStrategy) {
	err := rs.retry(func() error {
		if rand.Intn(2) == 0 { // random fail
			fmt.Println("Stripe payment failed")
			return errors.New("razorpay failed")
		}
		fmt.Println("Stripe payment successful")
		return nil
	})
	if err != nil {
		fmt.Println("Final failure after retry:", err)
	}

}

type stripe struct {
}

func (r *stripe) pay() {
	fmt.Println("paying through stripe")
}

type paytm struct {
}

func (r *paytm) pay() {
	fmt.Println("paying through paytm")
}

type grabpay struct {
}

func (r *grabpay) pay() {
	fmt.Println("paying through grabpay")
}

func main() {
	f := GetPaymentFactory("sea")
	f.createCardProcessor().pay()
	f.createWalletPayment().pay()
}

type retryStrategy interface {
	retry(action func() error) error
}

type noRetry struct {
}

func (n *noRetry) retry(action func() error) error {
	err := action()
	if err != nil {
		return err
	}
	return nil
}

type fixedRetry struct {
	attempts int
}

func (f *fixedRetry) retry(action func() error) error {
	for i := 0; i < f.attempts; i++ {
		err := action()
		if err == nil {
			return nil
		}
	}
	return errors.New("failed after retries")
}
