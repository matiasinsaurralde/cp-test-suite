Python
==

# Setup

To build this bundle and sign it, run:

```
tyk-cli bundle build -key ~/dev/tyk-cli/mykey.pem
```

This will generate a `bundle.zip`, in the current working directory. After this, I copy the file to my `~/dev` directory, which is visible through the web server:

```
cp bundle.zip ~/dev/test-bundle-python
```

We need to specify the `custom_middleware_bundle` in our test API spec.

```
"custom_middleware_bundle": "test-bundle-python"
```

This name will be merged with `bundle_base_url`, resulting in something like (in my config):
```
http://127.0.0.1/dev/test-bundle-python
```

(This is a ZIP file, with no extension)

If everything goes fine, we should see something like this in the log:

```
[Oct 24 17:49:41]  INFO main: --> Loading API: Tyk Test API
[Oct 24 17:49:41]  INFO main: ----> Tracking: (no host)
[Oct 24 17:49:41]  INFO main: ----> Fetching Bundle: test-bundle
[Oct 24 17:49:41]  INFO main: ----> Loading bundle: test-bundle
[Oct 24 17:49:41]  INFO main: ----> Verifying bundle: test-bundle
[Oct 24 17:49:41]  INFO main: ----> Bundle is valid, adding to spec: test-bundle
[Oct 24 17:49:41]  INFO main: ----> Checking security policy: Open
```

At this point the custom middleware block of this API should be loaded, and a test request should display the injected header:

```
% curl http://localhost:8080/test-api/headers
{
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "Mypyheader": "mypyvalue",
    "User-Agent": "Go-http-client/1.1"
  }
}
```
