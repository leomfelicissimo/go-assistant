package server

import (
	"belazap.com/assistant/rest"
)

func Initialize() {
	rest.StartHttp()
}
