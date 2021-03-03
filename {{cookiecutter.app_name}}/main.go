package main

import (
	"github.com/kitabisa/{{ cookiecutter.app_name }}/cmd"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	zlog.Logger = zlog.With().Caller().Logger()

    cmd.Execute()
}
