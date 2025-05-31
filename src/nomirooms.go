package main

import (
    NomiKin "github.com/cydonia1978/NomiKinGo"
)

type Room struct {
    Name    string
    Note    string
    Uuid    string
    Backchanneling bool
    Nomis   []NomiKin.Nomi
    RandomResponseChance int
}

