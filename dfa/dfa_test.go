package dfa

import (
    "fmt"
    "testing"
)

func bin(n int) string {
    return fmt.Sprintf("%b", n)
}

func TestDfaShouldCorrectlyDecideIntegers(t *testing.T) {
    for modulus := 1; modulus < 200; modulus++ {
        dfa, err := DfaDivisibleBy(modulus)
        if err != nil {
            t.Fail()
        }

        for n := 0; n < 1000; n++ {
            expected := n % modulus == 0
            result := dfa.Eval(bin(n))

            if result != expected {
                t.Fail()
            }
        }
    }
}
