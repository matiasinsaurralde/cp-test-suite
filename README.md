cp-test-suite
==

# General requirements

- Web server to serve the bundle files.
- Tyk build with CP support. Global settings: `bundle_base_url`, `enable_bundle_downloader`, `public_key_path`.
- Key to sign bundles.

# My setup

My private key is in this path: `~/dev/tyk-cli/mykey.pem`, the corresponding public key is in `~/dev/tyk-cli/mykey.pub`.

My `tyk.conf` has the following settings:

```
"coprocess_options": {
  "enable_coprocess": true,
  "coprocess_grpc_server": ""
},
"bundle_base_url": "http://127.0.0.1/dev/",
"enable_bundle_downloader": true,
"public_key_path": "/Users/matias/dev/tyk-cli/mykey.pub"
```

`bundle_base_url` is pointing to a local nginx instance.

## Lua

Check [here](lua/README.md)

## Python

Check [here](python/README.md)
