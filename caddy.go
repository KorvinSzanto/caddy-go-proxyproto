package caddygoproxyproto

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func init() {
	caddy.RegisterModule(Wrapper{})
}

func (Wrapper) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.listeners.go_proxyproto",
		New: func() caddy.Module { return new(Wrapper) },
	}
}

func (w *Wrapper) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		// No same-line options are supported
		if d.NextArg() {
			return d.ArgErr()
		}

		for d.NextBlock(0) {
			switch d.Val() {
			case "timeout":
				if !d.NextArg() {
					return d.ArgErr()
				}
				dur, err := caddy.ParseDuration(d.Val())
				if err != nil {
					return d.Errf("parsing go_proxyproto timeout duration: %v", err)
				}
				w.Timeout = caddy.Duration(dur)

			default:
				return d.ArgErr()
			}
		}
	}
	return nil
}

var (
	_ caddy.Provisioner     = (*Wrapper)(nil)
	_ caddy.Module          = (*Wrapper)(nil)
	_ caddy.ListenerWrapper = (*Wrapper)(nil)
	_ caddyfile.Unmarshaler = (*Wrapper)(nil)
)
