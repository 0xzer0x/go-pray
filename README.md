<a id="readme-top"></a>

## Overview ⚡

Powerful and user-friendly command-line interface (CLI) application that helps Muslims stay on top of their daily prayers. With features like prayer time notifications, prayer calendars, and flexible configuration options. 🕌

### Features

- **Prayer Notifications**: Start the daemon to receive desktop notifications at prayer times. 🔔
- **Next Prayer**: Quickly find the next upcoming prayer. 🕒
- **Prayer Calendar**: View prayer times for today or any date. 📅
- **Flexible Output**: Supports multiple output formats (e.g., JSON, plain text). 📜

## Installation 📥

> [!WARNING]
> Currently, only the following platforms are supported:
>
> - Linux (x86_64)

You can run the following command to install the latest release of `go-pray`:

```sh
curl -sLo - https://github.com/0xzer0x/go-pray/raw/refs/heads/main/install.sh | bash
```

By default, `go-pray` is installed to `./bin`. To customize the install directory, run the following:

```sh
curl -sLo - https://github.com/0xzer0x/go-pray/raw/refs/heads/main/install.sh | env INSTALL_DIR=$HOME/.local/bin bash
```

To install a specific version, use the following:

```sh
curl -sLo - https://github.com/0xzer0x/go-pray/raw/refs/heads/main/install.sh | env INSTALL_VERSION=0.1.3 bash
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Configuration ⚙️

`go-pray` reads from a `config.yml` file to manage settings. You can specify a custom configuration file using the `--config` flag. By default, the application searches for `config.yml` in the following paths:

1. `$XDG_CONFIG_HOME/go-pray`
2. `$HOME/.config/go-pray`
3. `$HOME/.go-pray`

An example [config.yml](./config.example.yml) is available. You can modify it to suit your location and preferences.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Building From Source 🔨

### Steps

1. Clone the repository:

```bash
git clone https://github.com/0xzer0x/go-pray.git
cd go-pray
```

2. Install dependencies for [oto](https://github.com/ebitengine/oto?tab=readme-ov-file#prerequisite):

```bash
apt install libasound2-dev
```

3. Build the application:

```bash
go build -o go-pray
```

3. Move the binary to your `$PATH`:

```bash
mv go-pray /usr/local/bin/
```

## Contributing 👥

Contributions are welcome! To get started:

1. Fork the repository
2. Create a branch for your feature (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing-feature'`)
4. Push the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Acknowledgments ✨

- [AdhanGo](https://github.com/mnadev/adhango/) library for calculation of prayer times.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License 📜

Distributed under the GPL v3 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
