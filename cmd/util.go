package main

import "math/rand"

func randSelect[S []E, E any](s S) E {
	return s[rand.Intn(len(s))]
}

func pop[S []E, E any](s S) (E, S) {
	popped := s[0]
	return popped, s[1:]
}
