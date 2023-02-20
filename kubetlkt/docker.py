import docker
import logging

logger = logging.getLogger(__name__)


def build_image(repo, name, tag="latest") -> tuple:
    client = docker.from_env()
    logger.info("Building the docker image")
    try:
        image = client.images.build(path="kubetlkt", tag=repo + "/" + name + ":" + tag)[
            0
        ]
    except docker.errors.BuildError as err:
        logger.error("Image build failed")
        logger.error(err)
    except docker.errors.APIError as err:
        logger.error("Docker API error")
        logger.error(err)
    else:
        logger.info("Docker image " + repo + ":" + tag + " built successfully")
        return image


def push_image(repo, name, tag="latest"):
    client = docker.from_env()
    logger.info("Pushing the docker image to the registry")
    resp = client.api.push(
        repo + "/" + name,
        tag="latest",
        stream=True,
        decode=True,
    )
    for line in resp:
        logger.info(line)
