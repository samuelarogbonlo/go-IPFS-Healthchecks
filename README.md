# GO-IPFS-Healthcheck

A plugin for [kubo](https://github.com/ipfs/kubo) that serves a healthcheck endpoint which returns the status of the IPFS node.

# Plugin Installation

This is a preloaded plugin built in-tree into go-ipfs when it's compiled.

```sh
git clone https://github.com/ipfs/kubo

cd kubo

# Pull in the plugin (you can specify a version other than "latest" if you'd like)
go get github.com/ceramicnetwork/go-ipfs-healthcheck/plugin@latest

# Add the plugin to the [preload list](https://github.com/ipfs/go-ipfs/blob/master/docs/plugins.md#preloaded-plugins)
echo "\nhealthcheck github.com/ceramicnetwork/go-ipfs-healthcheck/plugin 0" >> plugin/loader/preload_list

# Download dependencies
go mod download

# Build go-ipfs with the plugin
make build

# If an error occurs, try
go mod tidy
make build
```

# Usage

Run Kubo IPFS node and check its status.
```sh
./cmd/ipfs/ipfs daemon
curl -X GET http://localhost:8011
```

# Healthcheck Endpoint Statistics Review

## Steps
- Install and run IPFS as a daemon. Follow the docs [here](https://docs.ipfs.tech/) for more information.
- Use the CID value from [here](https://github.com/ipfs/kubo/pull/8429/files) or any other CID value from an IPFS setup.
- Run `run cmd/main.go` at the root directory
- Check [http://localhost:8011/healthcheck](http://localhost:8011/healthcheck) on your browser and ensure the result is `"status": "ok"`.

At this point, great to see that everything works! Your output shows a successful health check response with the IPFS DAG statistics.

# Resources

[Kubo Plugins](https://github.com/ipfs/kubo/blob/master/docs/plugins.md)

# Maintainers

[@samuelarogbonlo](https://github.com/samuelarogbonlo)

# License

Fully open source and dual-licensed under MIT and Apache 2.
