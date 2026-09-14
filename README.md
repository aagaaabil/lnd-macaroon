# lnd-macaroon

> lnd · invoice · macaroon-shaped

[![Go 1.22+](https://img.shields.io/badge/go-1.22+-00ADD8)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()

LND-shaped invoice helper — stub macaroon tag, no gRPC dial.

## Features

- BTC derivation path m/84'/0'/0'
- Local vault JSON with XOR wrap
- SHA-256 stand-in keys — no live RPC
- stdlib CLI via flag

## Prerequisites

- Go 1.22+
- Git

## Getting Started

```bash
git clone <repo-url>
cd lnd-macaroon
make build
./bin/lndmac -help
```

## CLI Usage

```bash
make test
go run ./cmd/lndmac -help
```

## Project Structure

```
cmd/lndmac/main.go
internal/config/config.go
internal/crypto/keys.go
internal/wallet/wallet.go
internal/wallet/wallet_test.go
```

## Background

LN Go scripts search lnd-macaroon, not lightning-wallet.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.


---

## Topics

![lnd](https://img.shields.io/badge/lnd-111827?style=flat-square) ![macaroon](https://img.shields.io/badge/macaroon-111827?style=flat-square) ![lnd-macaroon](https://img.shields.io/badge/lnd%20macaroon-111827?style=flat-square) ![cryptocurrency](https://img.shields.io/badge/cryptocurrency-111827?style=flat-square) ![wallet](https://img.shields.io/badge/wallet-111827?style=flat-square) ![blockchain](https://img.shields.io/badge/blockchain-111827?style=flat-square) ![web3](https://img.shields.io/badge/web3-111827?style=flat-square) ![bitcoin](https://img.shields.io/badge/bitcoin-111827?style=flat-square)

`lnd` `macaroon` `lnd-macaroon` `cryptocurrency` `wallet` `blockchain` `web3` `bitcoin` `ethereum` `hd-wallet` `open-source` `golang` `go`

Search: lnd-macaroon · lnd · invoice · macaroon-shaped · LND-shaped invoice helper — stub macaroon tag, no gRPC dial.

---

<sub>LND-shaped invoice helper — stub macaroon tag, no gRPC dial.</sub>
