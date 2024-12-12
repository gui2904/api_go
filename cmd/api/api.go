package main

type application struct {
        config config
}

type config struct {
		    addr string
}

func (app *application) run() error {
        mux := http.NewServeMux()

		    srv := &http.Server {
				        Addr: app.config.addr,
                Handler: mux,
		    }

		    return srv.ListenAndServe()
}
