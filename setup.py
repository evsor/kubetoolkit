from setuptools import setup

setup(
    name="kubetlkt",
    version="0.1.0",
    py_modules=["kubetlkt"],
    install_requires=[
        "click==8.1.3",
        "coloredlogs==15.0.1",
        "kubernetes==26.1.0",
        "docker==6.0.1",
    ],
    entry_points={
        "console_scripts": [
            "kubetlkt = kubetlkt:cli",
        ],
    },
)
