package player

import (
	"github.com/olshevskiy87/goradiooooo/params"
)

func (r *RadioooooPlayer) SetParams(p *params.Params) {
	switch p.Mode {
	case params.MODE_RANDOM:
		r.url = URL_PLAY_RANDOM
	case params.MODE_EXPLORE:
		fallthrough
	case params.MODE_TAXI:
		r.url = URL_PLAY
	}
	r.params = p
}
