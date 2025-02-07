package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
)

// Digit is the number struct containing all the information about the number
type Digit struct {
	Number     int      `json:"number"`
	Is_prime   bool     `json:"is_prime"`
	Is_perfect bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	Digit_sum  int      `json:"digit_sum"`
	Fun_fact   string   `json:"fun_fact"`
	Count      int      `json:"-"`
}

// Primer checks if the field Number value is a prime number.
func (d *Digit) Primer() {
	if d.Number <= 1 {
		d.Is_prime = false
		return
	}
	for i := 2; i <= d.Number/2; i++ {
		if d.Number%i == 0 {
			d.Is_prime = false
			return
		}
	}
	d.Is_prime = true
}

// Perfecter checks if the field "Number" is a perfect number.
func (d *Digit) Perfecter() {
	var sum int

	for i := 1; i <= d.Number/2; i++ {
		if d.Number%i == 0 {
			sum += i
		}
	}

	d.Is_perfect = d.Number == sum
}

// DigiCounter is a method of Digit update field Count with the number of values in the provided number.
func (d *Digit) DigiCounter() {
	if d.Number == 0 {
		d.Count = 1
		return
	}

	var sum int
	rem := d.Number

	for rem > 0 {

		sum += rem % 10

		rem /= 10
		d.Count++
	}
}

// Armstronger checks if the number is a Armstrong number and updates the Digit properities slice
func (d *Digit) Armstonger() {
	var sum int
	var armstrongValue float64
	rem := d.Number

	for i := 0; i < d.Count; i++ {
		val := rem % 10

		armstrongValue += math.Pow(float64(val), float64(d.Count))
		sum += val

		rem /= 10
	}

	d.Digit_sum = sum

	if int(armstrongValue) == d.Number {
		d.Properties = append(d.Properties, "Armstrong")
	}
}

// Parity checks if the number is an odd or prime number
func (d *Digit) Parityer() {
	str := "odd"
	if d.Number%2 == 0 {
		str = "even"
	}

	d.Properties = append(d.Properties, str)
}

// Facter updates the Digit struct Fun_fact field with a fact about the number.
func (d *Digit) Facter() error {
	link := fmt.Sprintf("http://numbersapi.com/%v/math", d.Number)
	res, err := http.Get(link)
	if err != nil {
		d.Fun_fact = "No facts here!"
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status %v and value %v", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	d.Fun_fact = string(body)
	return nil
}

// Setup function populates the Digit struct
func (d *Digit) Setup() {
	d.DigiCounter()
	d.Primer()
	d.Perfecter()
	d.Parityer()
	d.Armstonger()
	if err := d.Facter(); err != nil {
		d.Fun_fact = fmt.Sprintln("Nothing to see here :)")
	}

}

// Setup the digit
func SetupNum(num int) (Digit, error) {
	numero := math.Abs(float64(num))

	digit := Digit{
		Number: int(numero),
	}

	digit.Setup()

	return digit, nil

}
