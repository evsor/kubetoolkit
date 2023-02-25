import click
import coloredlogs, logging
import kubetoolkit.docker as docker
import kubetoolkit.kubernetes as kube
import kubetoolkit.config as conf


# Set globals
NAME = "kubetlkt"
PACKAGE_NAME = "kubetoolkit"
CONFIG_NAME = ".kubetlktconfig"
CONFIG = None


# Configure logging
logger = logging.getLogger(__name__)
coloredlogs.install(level="INFO", fmt="%(asctime)s %(levelname)-2s %(message)s")


@click.group()
def cli():
    """Kubernetes debugging toolkit"""
    logger.info("Firing up")
    global CONFIG
    CONFIG = conf.parse_config(PACKAGE_NAME, CONFIG_NAME)


@cli.command(help="Set user configuration")
@click.option("--repo", help="Set the DockerHub repository", required=True)
def config(repo):
    conf.write_user_config(PACKAGE_NAME, CONFIG_NAME, repo)


@cli.command(help="Build and push image")
def image():
    repo = CONFIG["DEFAULT"]["repo"]

    docker.build_image(repo, NAME, PACKAGE_NAME)
    docker.push_image(repo, NAME)


@cli.command(help="Create debug pod")
def create():
    repo = CONFIG["DEFAULT"]["repo"]
    kube.kube_action(repo, NAME, "start")


@cli.command(help="Remove debug pod")
def clean():
    repo = CONFIG["DEFAULT"]["repo"]
    kube.kube_action(repo, NAME, "cleanup")
