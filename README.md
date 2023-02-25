<p align="center">
<img src="assets/logo.png" width="500" height="150" alt="KubeToolkit logo"/>
</p>
<hr>

Kubernetes Toolkit(<b>kubetlkt</b>) seemlesly launches a deployment from a container image containing a set of tools to investigate and troubleshoot a Kubernetes cluster. It does not make use of the ephimeral containers future for compatibility reasons with clusters running Kubernetes version prior to 1.25

# Current version: 0.3.0
## Installation

```
pip install .
```
Note: On Ubuntu 22.04 you might need to export an env variable `DEB_PYTHON_INSTALL_LAYOUT=deb_system` because of this [issue](https://github.com/pypa/setuptools/issues/3269#issuecomment-1254507377)

## Usage

```
kubetlkt --help
```

### Commands

`create` creates the Kubernetes deployment in the `default` namespace

`clean` removes the deployment, created with the `create` command

`image` build and push your on image to Docker Hub. In order to use, you need to overwrite the default repository with the `config` command first

`config` overwrite the default configuration. Currently only the Docker Hub repository namespace

## Autocompletion

Add autocompletion to your shell of choice
### Bash

Save the autocompletion script

```
_KUBETLKT_COMPLETE=bash_source kubetlkt > ~/.kubetlkt-complete
```
Source the file in `~/.bashrc`
```
. ~/.kubetlkt-complete
```


### Zsh

```
_KUBETLKT_COMPLETE=zsh_source kubetlkt > ~/.kubetlkt-complete
```
Source the file in `~/.zshrc`
```
. ~/.kubetlkt-complete
```

## Tools

`curl`<br />
`wget`<br />
`jq`<br />
`unzip`<br />
`dnsutils`<br />
`traceroute`<br />
`telnet`<br />
`netcat`<br />
`net-tools`<br />

## FAQ

<b>Whoa, wrapping this thing in Python is a huge overkill. This can be achieved with a few commands and a yaml file</b>

Yeah, but where is fun in that? However i do plan to develop this tool further into a more capable debugging tool, so i've decided to lay out the foundation from the start and see where it leads