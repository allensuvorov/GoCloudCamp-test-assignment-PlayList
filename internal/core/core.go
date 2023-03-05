package core

import "time"

type song struct {
	name     string
	duration time.Duration
}

type track struct {
	song
	next *track
	prev *track
}

type playlist struct {
	name      string
	firstSong *track
}
