# TRC20 Address
[![Go Reference](https://pkg.go.dev/badge/github.com/spcily/trc20.svg)](https://pkg.go.dev/github.com/spcily/trc20)

a Golang package to generate a new TRC20 wallet address offline or get one from private key.

## Usage

```golang
import (
  "github.com/spcily/trc20/address"
)
```

### Generate a new address

```golang
address := address.Generate()
```

### Get from private key

privateKeyHex is a hex string of private key. eg. `0xf5737b0a37b87451a52863e0bbec15ab06983540f41b5b06b322ad1a1f8eef72`

```golang
address := address.GetFromPrivateKey(privateKeyHex)
```

### Return

`address.PrivateKeyHex`: private key in hex string format (begins with 0x)

`address.AddressHex`: wallet address in hex string format (begins with 0x)

`address.Base58Address`: wallet address in base58check format

# LICENSE

MIT
