package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

type (
	// Rubyist represents the P&P role of the Rubyist
	Rubyist struct {
		Character
	}
)

// NewRubyist ...
func NewRubyist() *Rubyist {
	return &Rubyist{Character: Character{H: 100, Name: "Rubyist"}}
}

// Art renders the player's ascii art with the player state
func (r Rubyist) Art() string {
	return fmt.Sprintf(rubyistArt, r.H, r.X)
}

// Skills returns the list of abilities the Rubyist has
// If XP is larger than 100 [DuckTyping, Module, DarkMagic]
// If XP is larger than 10 [DuckTyping, Module, ]
// Anything else [DuckTyping]
func (r Rubyist) Skills() []pnp.Skill {
	if r.XP() > 100 {
		return []pnp.Skill{pnp.DuckTyping, pnp.Module, pnp.DarkMagic}
	}
	if r.XP() > 10 {
		return []pnp.Skill{pnp.DuckTyping, pnp.Module}
	}
	return []pnp.Skill{pnp.DuckTyping}
}

func (r Rubyist) String() string {
	return r.Name
}