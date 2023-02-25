import platformdirs
import logging
import importlib.resources
import os.path
import configparser

logger = logging.getLogger(__name__)


def parse_config(name: str, config_name: str) -> object:
    config = configparser.ConfigParser()
    user_config_full_path = user_config_file(config_name)
    package_config_full_path = str(importlib.resources.path(name, config_name))

    if os.path.isfile(user_config_full_path):
        try:
            logger.info("Reading " + user_config_full_path + " configuration")
            config.read(user_config_full_path)
        except FileNotFoundError as e:
            logger.error(user_config_full_path + " user config file not found")
    else:
        try:
            logger.info("Reading " + package_config_full_path + " configuration")
            config.read(package_config_full_path)
        except FileNotFoundError as e:
            logger.error(package_config_full_path + " default config file not found")

    return config


def write_user_config(package_name: str, name: str, repo: str):
    user_config_dir = platformdirs.user_config_dir()
    user_config_full_path = os.path.join(user_config_dir, name)

    try:
        config_file = importlib.resources.read_text(package_name, name)

        with open(user_config_full_path, "w") as dest:
            dest.writelines(config_file)
    except FileNotFoundError as e:
        logger.error("Default config file not found")

    config = configparser.ConfigParser()

    try:
        logger.info("Writing " + user_config_full_path + " configuration")
        config["DEFAULT"] = {"repo": repo}
        with open(user_config_full_path, "w") as configfile:
            config.write(configfile)
    except FileNotFoundError as e:
        logger.error(user_config_full_path + " user config file not found")


def user_config_file(config_name: str) -> str:
    user_config_dir = platformdirs.user_config_dir()
    full_path = os.path.join(user_config_dir, config_name)

    return full_path
