package adhan

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	beep "github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/spf13/viper"
)

type Player struct {
	streamer beep.StreamSeekCloser
	format   beep.Format
	isInit   bool
}

func NewPlayer() *Player {
	return &Player{}
}

func (a *Player) Initialize() error {
	path, err := filepath.Abs(os.ExpandEnv(viper.GetString("adhan")))
	if err != nil {
		return fmt.Errorf("failed to resolve file path: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open MP3 file: %w", err)
	}

	a.streamer, a.format, err = mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode MP3: %w", err)
	}

	if !a.isInit {
		err = speaker.Init(a.format.SampleRate, a.format.SampleRate.N(time.Second/10))
		if err != nil {
			return fmt.Errorf("failed to initialize speaker: %w", err)
		}
		a.isInit = true
	}

	return nil
}

func (a *Player) Duration() time.Duration {
	return a.format.SampleRate.D(a.streamer.Len())
}

func (a *Player) Stop() {
	speaker.Clear()
}

func (a *Player) Play() error {
	speaker.Clear()

	if err := a.streamer.Seek(0); err != nil {
		return fmt.Errorf("failed to seek to adhan start: %w", err)
	}

	speaker.Play(a.streamer)
	return nil
}
