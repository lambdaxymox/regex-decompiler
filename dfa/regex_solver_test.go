package dfa

import (
    "fmt"
    "testing"
    "regexp"
)

func bin(n int) string {
    return fmt.Sprintf("%b", n)
}

func TestRegexExpressionShouldDecideDivN(t *testing.T) {
    for modulus := 0; modulus < 20; modulus++ {
        re := make_solver(0, modulus).solve()
        matcher, err := regexp.Compile(re)
        if err != nil {
            t.Fail()
        }

        for n := 0; n < 1000; n++ {
            expected := bin(n)
            if !matcher.MatchString(expected) {
                t.Fail()
            }
        }
    }
}
