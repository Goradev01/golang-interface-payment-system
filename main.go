package main

import (
	"fmt"
)

type PaymentGateway interface {
	Pay(amount float64) error
}

type PayStack struct {
}

type FlutterWave struct {
}

func (service PayStack) Pay(amount float64) error{
 fmt.Println("Paystack Pay",amount );
 return nil
}
func (service FlutterWave) Pay(amount float64) error{
 fmt.Println("Flutterwave Pay",amount );
 return nil
}

type CheckoutGateway struct{
gateway	PaymentGateway ;
}

func (checkout CheckoutGateway) Checkout(amt float64){
err := checkout.gateway.Pay(amt)

if err != nil{
	fmt.Println("Payment fail")
}
fmt.Println("Payment Successfull")
}

func main(){
	paymentMethod := PayStack{}

	check := CheckoutGateway{
		gateway: paymentMethod,
	}

    check.Checkout(60)
}