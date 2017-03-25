# vlt

Just a simple client to read secrets from HashiCorp's Vault.

## Configuration

Default configuration is read from ```~/.vault.json```:

```
{
  "server": "some-vault-host.somewhere.com",
  "port": 443,
  "token": "some-secret-token",
  "tls": true,
  "insecure": false
}
```

Configration is also possible via CLI parameters.

## Usage

To read a secret, do something like this:

```vlt --secret "path/to/secret" --key "some-key"```

## Build

Just call ```make```.