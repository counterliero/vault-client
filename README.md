# vlt

Just a simple client to read secrets from HashiCorp's Vault.

## Configuration

Configuration is read from ```~/.vault.json```:

```json
{
  "server": "some-vault-host.somewhere.com",
  "port": 443,
  "token": "some-secret-token",
  "tls": true,
  "insecure": false
}
```
Configration is also possible via CLI parameters.

If there is no configuration file and nothing is passed via CLI, the default settings will be used:

```json
{
  "server": "localhost",
  "port": 8200,
  "tls": true,
  "insecure": false
}
```


## Usage

To read a secret, do something like this:

```bash
vlt --secret "path/to/secret" --key "some-key"
```

If you don't use a configuration file, you can pass all available settings as CLI parameter:

```bash
vlt --server "some-vault.secure.tld" \
    --port 443 \
    --token "some-token" \
    --secret "some/secret/to/read" \
    --key "password"
```

## Build

Just call ```make```.
