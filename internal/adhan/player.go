//go:generate mockgen -package adhan_test -destination mock_player_test.go . Player
package adhan

import "time"

type Player interface {
	Initialize() error
	Duration() time.Duration
	Play() error
	Stop() error
}

func NewPlayer() Player {
	return &Mp3Player{}
}
