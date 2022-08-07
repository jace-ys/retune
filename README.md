# Retune

Retune is a stable port of [go-micro.dev/v4/config](https://github.com/asim/go-micro/tree/v4.0.0/config), providing a pluggable library for working with static and dynamic configuration in [12-Factor applications](https://12factor.net/) written in Go.

## Features

- **Dynamic Loading** - Load configuration from multiple source as and when needed. Go Config manages watching config sources 
in the background and automatically merges and updates an in memory view. 

- **Pluggable Sources** - Choose from any number of sources to load and merge config. The backend source is abstracted away into 
a standard format consumed internally and decoded via encoders. Sources can be env vars, flags, file, etcd, k8s configmap, etc.

- **Mergeable Config** - If you specify multiple sources of config, regardless of format, they will be merged and presented in 
a single view. This massively simplifies priority order loading and changes based on environment.

- **Observe Changes** - Optionally watch the config for changes to specific values. Hot reload your app using Go Config's watcher. 
You don't have to handle ad-hoc hup reloading or whatever else, just keep reading the config and watch for changes if you need 
to be notified.

- **Sane Defaults** - In case config loads badly or is completely wiped away for some unknown reason, you can specify fallback 
values when accessing any config values directly. This ensures you'll always be reading some sane default in the event of a problem.
