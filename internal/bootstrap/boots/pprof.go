package boots

import (
	"everything-template/internal/vars"
	"everything-template/pkg/logger"
	"net/http"
	_ "net/http/pprof"
)

func InitPprof() {
	if vars.Config.Pprof.Enabled {
		go func() {
			addr := vars.Config.Pprof.Host + ":" + vars.Config.Pprof.Port
			logger.Infow("[pprof] start on " + addr)
			err := http.ListenAndServe(addr, nil)
			if err != nil {
				panic(err)
			}
		}()
	}

}
