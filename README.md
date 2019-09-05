# fftchain Blockchain

The fftchain Blockchain repository.

## How to use fftchainc

### devnet

You can run the chain easily on your computer to play with the software already.

First step is to launch a daemon from your console:
```
fftchaind --network devnet --no-bootstrap -Mgctwbe
```

the above launches a fftchain daemon on devnet, using no bootstrap
(meaning it doesn't try to connect to bootstrap nodes or wait for such nodes to know if you're sync or not),
enabling also the explorer module.

Once you have that you can recover the genesis wallet so you can start creating blocks and have money to spend:

```
fftchainc wallet recover --plain \
    --seed "carbon boss inject cover mountain fetch fiber fit tornado cloth wing dinosaur proof joy intact fabric thumb rebel borrow poet chair network expire else"
```

As this wallet is recovered as a plain wallet it does not have to be unlocked and is ready for use:

```
$ fftchainc wallet
Wallet status:
Encrypted, Unlocked
Confirmed Balance:   100006530 FFT
Locked Balance:      0 FFT
Unconfirmed Delta:   + 0 FFT
BlockStakes:         3000 BS
```

Please consult the `--help` menus of the `fftchainc` command and all its subcommands for more information on how to use the CLI.

### Using multiple wallets on the same machine

A single `fftchaind` daemon doesn't allow multiple wallets for the time being.
In order to have multiple wallets running on the same machine you therefore need
to run multiple `fftchaind` daemons, with each daemon:
  - using a unique persistent directory (either by starting each daemon from a different directory or
    by explicitly setting it using the `--persistent-dir` flag);
  - exposing itself using a unique port.
These different can manually be connected to one another using the `fftchainc gateway connect localhost:22312` command.
