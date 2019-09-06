# Explorer

A block explorer for Freeflow Token

## Run it yourself

### Prerequisites
* Caddyserver
* Freeflow Token daemon (`fftchaind`)


Make sure you have `fftchaind` (the Freeflow Token daemon) running with the explorer module enabled:
`fftchaind -M cgte`

Now start caddy from the `caddy` folder of this repository:
`caddy -conf Caddyfile.local`
and browse to http://localhost:2015
