import click
import coloredlogs, logging
import kubetoolkit.docker as docker
import kubetoolkit.kubernetes as kube
import kubetoolkit.config as config


# Set global name used in Docker and Kubernetes
NAME = "kubetlkt"

# Configure logging
logger = logging.getLogger(__name__)
coloredlogs.install(level="INFO", fmt="%(asctime)s %(levelname)-2s %(message)s")


@click.command()
@click.option(
    "--action",
    help="start: creates deployment \
                                cleanup: removes deployment",
    required=True,
)
@click.option(
    "--repo",
    default="evsoroka",
    show_default=True,
    help="Docker repository name to push the image. If not specified, the public one will be used",
)
def cli(action, repo, name=NAME):
    logger.info("Firing up")
    if action == "start":
        if repo != "evsoroka":
            docker.build_image(repo, name)
            docker.push_image(repo, name)

        # Create kubernetes object
        kube.kube_action(repo, name, action)
    elif action == "cleanup":
        logger.error("I'm here")
        config.get_config_dir()
        kube.kube_action(repo, name, action)
    else:
        logger.error(
            "Provide a desired action with the --action option: start, cleanup"
        )
