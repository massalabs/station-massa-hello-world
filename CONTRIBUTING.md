# Contributing

We welcome and appreciate contributions from the community. Whether it's a bug report, feature request,
code contribution, or anything else, we're grateful for your effort. Here's a brief guide to get you started.

Any contribution you make will be reflected on [CONTRIBUTORS](CONTRIBUTORS.md).

To get an overview of the project, read the [README](README.md).

## What you can Contribute

We welcome contributions of all kinds! Here are some ways you can help:

* [Report bugs](#bug-reports) and [suggest features](#feature-requests)
* Improve the documentation
* [Write code, such as bug fixes and new features](#code-contributions)
* Submit design improvements

## Code of Conduct

Help us keep this project open and inclusive and keep our community approachable and respectable.

## Bug Reports

If you find any bugs, feel free to open an issue in the
[issue tracker](https://github.com/massalabs/station-massa-hello-world/issues). Please include as much information as possible,
such as operating system, browser version, etc.

## Feature Requests

If you have an idea for a new feature, feel free to open an issue in the
[issue tracker](https://github.com/massalabs/station-massa-hello-world/issues). Please include as much information as possible,
such as what the feature would do, why it would be useful, etc.

## Code Contributions

We are always looking for developers to help us improve the codebase. If you are interested in contributing code, please
take a look at the [open issues](https://github.com/massalabs/station-massa-hello-world/issues) and follow the steps below:

1. [Fork](https://help.github.com/en/github/getting-started-with-github/fork-a-repo) the repository
2. Create a feature branch off of the `main` branch
3. Make your changes
4. Submit a pull request

When contributing:

* Ensure that all new features or changes are aligned with our product vision.
* Write clean, easy-to-understand code.
* Write tests for your code in order to ensure that it works as expected.
* Make sure all tests pass before submitting a pull request.
* Follow our code conventions and formatting standards.
* Document your code thoroughly.

## Questions

If you have any questions, feel free to open an issue in the
[issue tracker](https://github.com/massalabs/station-massa-hello-world/issues).

*Thank you for taking the time to contribute to our project! We look forward to seeing your contributions!*

## Developer guide

This section helps developer getting started.

If you want to contribute, please refer to our [CONTRIBUTING](CONTRIBUTING.md) guide.

### Install Task

Follow the installation instructions here:
[task-install](https://taskfile.dev/installation/)

### Install dependencies

```shell
task install
```

### Build

Generate the projects: go-swagger, wails, web-frontend:

```shell
task generate
```

```shell
task build
```

### Test

```shell
task test
```

### Run

For development purpose, you can run the plugin in standalone mode: it will not try to register with MassaStation.

```shell
task run
```

All in one build & run:

```shell
task build-run
```

The `STANDALONE` environment variable is to run the plugin without MassaStation.

Now navigate into <http://localhost:8080>. Note that some features will not work if
[MassaStation-server](https://github.com/massalabs/station) is not running.

**Install manually the plugin for Massa Station:**

For development purpose, you can install the plugin manually:

```shell
task install-plugin
```

This will create MassaStation plugin directories and copy the binary file created in the previous step so that
MassaStation can detect the plugin and launch it.
