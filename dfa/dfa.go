package dfa

import (
    "fmt"
    "strings"
)

type AcceptingStates []int

func (acceptingStates *AcceptingStates) contains(state int) bool {
    for _, acceptingState := range *acceptingStates {
        if  state == acceptingState {
            return true
        }
    }

    return false
}

type ErrModulusTooSmall int

func (e ErrModulusTooSmall) Error() string {
    return fmt.Sprintf("Modulus too small: %v; modulus must be at least 1.", int(e))
}

type Dfa struct {
    alphabet        string
    startingState   int
    acceptingStates AcceptingStates
    nextState       map[int]map[rune]int
}

func (dfa *Dfa) String() string {
    str := fmt.Sprintf("ALPHABET: \"%v\"\n", dfa.alphabet)
    str += fmt.Sprintf("STARTING STATE: %v\n", dfa.startingState)
    str += fmt.Sprintf("ACCEPTING STATES: %v\n", dfa.acceptingStates)
    str += fmt.Sprintf("STATE TABLE: \n")
    for q, mq := range dfa.nextState {
        for ch, next := range mq {
            str += fmt.Sprintf("nextState(%v, \"%v\") = %v\n", q, string(ch), next)
        }
    }
    
    return str
}

func (dfa *Dfa) Eval(str string) bool {
    if str == "" {
        return false
    }

    q := dfa.startingState
    for _, ch := range str {
        if strings.ContainsRune(dfa.alphabet, ch) {
            q = dfa.nextState[q][ch]
        } else {
            return false
        }
    }
    
    return dfa.acceptingStates.contains(q)
}

func DfaDivisibleBy(modulus int) (Dfa, error) {
    if modulus < 1 {
        return Dfa{}, ErrModulusTooSmall(modulus)
    }

    alphabet := "01"
    startingState := 0
    acceptingStates := []int{0}
    nextState := make(map[int]map[rune]int)
    for q := 0; q < modulus; q++ {
        nextState[q] = make(map[rune]int)
        nextState[q]['0'] = (2*q + 0) % modulus
        nextState[q]['1'] = (2*q + 1) % modulus
    }

    dfa := Dfa{alphabet, startingState, acceptingStates, nextState}

    return dfa, nil
}
