package adhan

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/spf13/viper"
)

type Mp3Player struct {
	buffer      beep.Buffer
	format      beep.Format
	initialized bool
}

func (a *Mp3Player) Initialize() error {
	path, err := filepath.Abs(os.ExpandEnv(viper.GetString("adhan")))
	if err != nil {
		return fmt.Errorf("failed to resolve file path: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open MP3 file: %w", err)
	}

	var streamer beep.StreamSeekCloser
	streamer, a.format, err = mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode MP3: %w", err)
	}
	defer streamer.Close()

	a.buffer = *beep.NewBuffer(a.format)
	a.buffer.Append(streamer)

	return nil
}

func (a *Mp3Player) Play() error {
	if !a.initialized {
		err := speaker.Init(a.format.SampleRate, a.format.SampleRate.N(time.Second/10))
		if err != nil {
			return fmt.Errorf("failed to initialize speaker: %w", err)
		}
		a.initialized = true
	}

	speaker.Clear()
	streamer := a.buffer.Streamer(0, a.buffer.Len())
	speaker.Play(streamer)
	return nil
}

func (a *Mp3Player) Duration() time.Duration {
	return a.format.SampleRate.D(a.buffer.Len())
}

func (a *Mp3Player) Stop() error {
	speaker.Clear()
	return nil
}
