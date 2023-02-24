<p align="center">
<img src="assets/logo.png" width="500" height="150" alt="KubeToolkit logo"/>
</p>
<hr>

Kubernetes Toolkit(<b>kubetlkt</b>) seemlesly launches a deployment from a container image containing a set of tools to investigate and troubleshoot a Kubernetes cluster. It does not make use of the ephimeral containers future for compatibility reasons with clusters running Kubernetes version prior to 1.25

# Current version: 0.2.1
## Installation

```
pip install .
```
Note: On Ubuntu 22.04 you will need to export an env variable `DEB_PYTHON_INSTALL_LAYOUT=deb_system` because of this [issue](https://github.com/pypa/setuptools/issues/3269#issuecomment-1254507377)

## Usage

```
kubetlkt --action start
```

`--action` flag is a required argument that takes 2 values: start and cleanup

`start` action will create the Kubernetes deployment in the `default` namespace

`cleanup` will remove it

There is also an optional `--repo` argument. It should be a valid DockerHub repository. If passed, the script will build the image and push it to the DockerHub registry. It gives the user control over what image to run. If you use this flag, it should be passes together with the `--action` command on every action

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