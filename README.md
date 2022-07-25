# TRC20 Address

a Golang package to generate a new TRC20 wallet address offline or get one from private key.

## Usage

```golang
import (
  "github.com/spcily/trc20"
)
```

### Generate a new address

```golang
privateKey, hexAddress, base58Address := trc20.Generate()
```

### Get from private key

privateKeyHex is a hex string of private key. eg. `0xf5737b0a37b87451a52863e0bbec15ab06983540f41b5b06b322ad1a1f8eef72`.

```golang
privateKey, hexAddress, base58Address := trc20.GetFromPrivateKey(privateKeyHex)
```

### Returns slice

[0] = privateKeyHex: private key in hex string format

[1] = hexAddress: wallet address in hex string format

[2] = base58Address: wallet address in base58check format

# LICENSE

MIT
