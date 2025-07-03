package main

import "fmt"

type DB interface {
	Save(string) error
}

type Store struct{}
type DinamycStore struct {
	dinamycStuff string
}
type PenAndPaper struct {
	paper           string
	pencilRemaining int
}

func (s *Store) Save(data string) error {
	fmt.Println("saving data: ", data)
	return nil
}

func (p *PenAndPaper) Save(data string) error {
	if len(data) > p.pencilRemaining {
		return fmt.Errorf("not enough pencil remaining")
	}
	p.pencilRemaining -= len(data)
	fmt.Println("saving data: ", data, p.paper)
	return nil
}

func (d *DinamycStore) Save(data string) error {
	fmt.Println("saving data: ", data, d.dinamycStuff)
	return nil
}

func myExecuteFn(db DB) ExecuteFn {
	fmt.Println("testing func:", db)
	return func(data string) error {
		fmt.Println("before saving data: ", data)
		return db.Save(data)
	}
}

// 3rd party func i dont control
type ExecuteFn func(string) error

func Execute(fn ExecuteFn) {
	err := fn("Hi there now i have  more length")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

//******** More Fintech Related***************

type PaymentProcessor interface {
	ProccessPayment(amount float64) error
}

type AwesomePaymentProcessor struct {
}

func (a *AwesomePaymentProcessor) ProccessPayment(amount float64) error {
	fmt.Println("processing payment: ", amount)
	return nil
}

type FraudDetector struct {
	processor PaymentProcessor
}

func (f *FraudDetector) ProccessPayment(amount float64) error {
	if amount > 1000 {
		return fmt.Errorf("check this dudes wallet")
	}
	return f.processor.ProccessPayment(amount)
}

type NotificationSender struct {
	processor PaymentProcessor
}

func (n *NotificationSender) ProccessPayment(amount float64) error {
	fmt.Println("sending notification: ", amount)
	return n.processor.ProccessPayment(amount)
}

//***************************
//***************************

func main() {
	store := &Store{}
	dinamycStore := &DinamycStore{
		dinamycStuff: "dinamyc stuff",
	}
	penAndPaper := &PenAndPaper{
		paper:           "paper",
		pencilRemaining: 10,
	}

	decoratedFn := myExecuteFn(store)
	decoratedFn2 := myExecuteFn(dinamycStore)
	decoratedFn3 := myExecuteFn(penAndPaper)
	Execute(decoratedFn)
	Execute(decoratedFn2)
	Execute(decoratedFn3)
	fmt.Println("--------------------------------")
	baseProcessor := &AwesomePaymentProcessor{}
	notificationSender := &NotificationSender{processor: baseProcessor}
	fraudDetector := &FraudDetector{processor: notificationSender}

	fraudDetector.ProccessPayment(500)
}
