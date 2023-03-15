//go:build !windows

package sing_tun

import (
	tun "github.com/metacubex/sing-tun"
)

func tunOpen(options tun.Options) (tun.Tun, error) {
	return tun.New(options)
}
