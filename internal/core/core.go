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

// [ ] Play - starts playing / начинает воспроизведение
func (pl *playlist) play() error {
	t := time.NewTimer()
	return nil
}

// TODO:
// [ ] Pause - приостанавливает воспроизведение
// [ ] AddSong - добавляет в конец плейлиста песню
// [ ] Next воспроизвести след песню
// [ ] Prev воспроизвести предыдущую песню
