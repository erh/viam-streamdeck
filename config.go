package viamstreamdeck

import (
	"fmt"
	"slices"
)

type KeyConfig struct {
	Key int

	Text      string
	TextColor string `json:"text_color"`

	Color string
	Image string

	Component string
	Method    string
	Args      []interface{}
}

func (kc *KeyConfig) Validate() error {
	if kc.Component == "" {
		return fmt.Errorf("need a component")
	}
	if kc.Method == "" {
		return fmt.Errorf("need a component")
	}
	return nil
}

type Config struct {
	Brightness int
	Keys       []KeyConfig
}

func (c *Config) Validate(p string) ([]string, []string, error) {
	ret := []string{}

	for _, k := range c.Keys {
		err := k.Validate()
		if err != nil {
			return nil, nil, err
		}

		if !slices.Contains(ret, k.Component) {
			ret = append(ret, k.Component)
		}

	}

	return nil, ret, nil
}
