import logging
import sys

from kubernetes import client, config
from kubernetes.client.rest import ApiException

logger = logging.getLogger(__name__)


def create_deployment_object(repo, name):
    # Container section
    container = client.V1Container(
        name=name,
        image=repo + "/" + name + ":" + "latest",
        command=["tail", "-f", "/dev/null"],
    )

    # Spec section
    spec = client.V1PodTemplateSpec(
        metadata=client.V1ObjectMeta(labels={"app": name}),
        spec=client.V1PodSpec(containers=[container]),
    )

    # Deployment spec section
    deployment_spec = client.V1DeploymentSpec(
        replicas=1, template=spec, selector={"matchLabels": {"app": name}}
    )

    # Instantiate the deployment object
    deployment = client.V1Deployment(
        api_version="apps/v1",
        kind="Deployment",
        metadata=client.V1ObjectMeta(name=name + "-deployment"),
        spec=deployment_spec,
    )

    return deployment


def create_deployment(api, deployment, namespace="default"):
    # Create deployment
    try:
        response = api.create_namespaced_deployment(
            body=deployment, namespace=namespace
        )
    except ApiException as err:
        logger.error("Kubernetes API Error")
        logger.error(err)
    else:
        logger.info(
            "Deployment `"
            + response.metadata.name
            + "` created successfully in `"
            + response.metadata.namespace
            + "` namespace"
        )


def delete_deployment(api, name, namespace="default"):
    # Delete deployment
    try:
        response = api.delete_namespaced_deployment(
            name=name + "-deployment",
            namespace=namespace,
            body=client.V1DeleteOptions(
                propagation_policy="Foreground", grace_period_seconds=5
            ),
        )
    except ApiException as err:
        logger.error("Kubernetes API Error")
        logger.error(err)
    else:
        logger.info(
            "Deployment `"
            + name
            + "-deployment` removed successfully from `"
            + namespace
            + "` namespace"
        )


def kube_action(repo, name, action):
    try:
        config.load_kube_config()
    except:
        logger.error("Error loading Kubernetes config")
        sys.exit("Exiting")

    apps_v1 = client.AppsV1Api()

    if action == "start":
        logger.info("Creating " + name + " deployment")
        deployment = create_deployment_object(repo, name)
        create_deployment(apps_v1, deployment)
    elif action == "cleanup":
        logger.info("Removing " + name + " deployment")
        delete_deployment(apps_v1, name)
