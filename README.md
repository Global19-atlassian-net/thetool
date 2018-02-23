# thetool
Build the Gloo universe

## Installing
### Prerequisite
`thetool` uses Docker to run the build process and build Docker images. To install Docker, please
visit [Install Docker](https://docs.docker.com/install/)

You need [Helm](https://helm.sh/) to deploy Gloo to Kubernetes. For information, please visit
[Install Helm](https://docs.helm.sh/using_helm/#installing-helm)

### Downloading and Installing
Download the latest release from https://github.com/solo-io/thetool/releases/latest/

If you prefer to compile your own binary or work on the development `thetool` please use the following command:

```
go get github.com/solo-io/thetool
``` 

## Getting Started
### Initialize
Create a working directory for `thetool`

```
mkdir gloo
cd gloo
```

Initialize `thetool` with default set of Gloo features. Optionally, you specify a default user id for Docker
using the `-u` flag.

```
thetool init -u solo-io
```

You can look at the default set of features by using the `list` command.

```
thetool list

Repository:       https://github.com/solo-io/gloo-plugins.git
Name:             aws_lambda
Gloo Directory:   aws
Envoy Directory:  aws/envoy
Enabled:          true

```

### Select the Gloo features
You can enable or disable any of the features calling `enable` or `disable` command with the name of the feature.

```
thetool disable -n aws_lambda
```

If you list again, you will see the `echo` feature is disabled.

```
thetool list

Repository:       https://github.com/solo-io/gloo-plugins.git
Name:             aws_lambda
Gloo Directory:   aws
Envoy Directory:  aws/envoy
Enabled:          false

```

### Build
Once you have selected the features you want to include, you can build Gloo and its components using the `build` command.

```
thetool build all
```

You can also choose to build individual components of Gloo by specifying the name of the component like `envoy` or `gloo`.

The build command builds the appropriate binaries and their corresponding Docker images. It then publishes these images to Docker registry. If you do not want to publish, you can pass a flag to `thetool`

```
thetool build all --publish=false
```

Note: In order to deploy Gloo to Kubernetes, you need to publish the Docker images.


### Deploy

You can use the `deploy` command to deploy Gloo and its components to different environments.

Here, we are looking at deploying Gloo to Kubernetes. `thetool` uses Helm to deploy Gloo
and its components.

Note: If you used custom Docker tags when building Gloo and its components, you must provide
the same tag to `deploy` command to deploy those images.

```
thetool deploy k8s
```

If you want to generate the Helm chart values without deploying please pass the `--dry-run` flag.

The Helm chart used by Gloo is available at [gloo-chart](https://github.com/solo-io/gloo-chart)

## Adding Your Own Feature

`thetool` can build Gloo with your custom Gloo features by adding your own feature repository to the list.

You can add or remove your feature repository using `add` and `delete` commands.

```
thetool add -r https://github.com/axhixh/gloo-magic.git -c 37a53fefe0a267fe3f4704c35e3721a4b6032f2a
```

You can verify by looking at the repository list with `list-repo` command.

```
thetool list-repo

Repository:  https://github.com/solo-io/gloo-plugins.git
Commit:      7bff2ff6c6ee707d8c09100de0bb7f869bd7488d

Repository:  https://github.com/axhixh/gloo-magic.git
Commit:      7bff2ff6c6ee707d8c09100de0bb7f869bd7488d
```

When you add a Gloo feature repository, it loads the file `features.json` in the root folder to
find what features are available. It uses the file to identify the Gloo plugin folder and envoy
filter folder for the feature.

To learn more about writing your own Gloo feature, please read the Gloo documentation.
