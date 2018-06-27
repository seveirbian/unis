# unis-cli api

## unisctl

unisctl

Usage: unisctl COMMAND

A client to communicate with unis-apiserver

Commands:
  connect    Connect to the remote unis-apiserver
  images     List images in remote registry
  nodes      Display the status of all edge nodes
  ps         List containers
  pull       Pull an image from remote registry
  push       Push an image to remote registry
  rm         Remove on or more containers on edge nodes
  rmi        Remove one or more images in remote registry
  run        Run a container on a edge node
  signin     Sign in
  signup     Sign up a new account
  stats      Display the status of all components of unis
  stop       Stop one or more running containers
  tag        Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
  version    Show the unis-cli and unis-apiserver version information

Run 'unisctl COMMAND --help' for more information on a command.

## unisctl connect --help

Usage: unisctl connect [SERVER]

Connect to unis-apiserver

Options:
      --help             Print usage

## unisctl images --help

Usage: unisctl images [OPTIONS]

List images

Options:
      --help             Print usage

## unisctl nodes --help

Usage: unisctl nodes [OPTIONS]

List the status of all edge nodes

Options:
      --help              Print Usage

## unisctl ps --help

Usage: unisctl ps [OPTIONS]

List containers

Options:
      --help             Print Usage
  -a, --all              Show all containers (dafault shows just running)

## unisctl pull --help

Usage: unisctl pull [OPTIONS] NAME[:TAG]

Pull an image from a registry

Options:
      --help             Print usage

## unisctl push --help

Usage: unisctl push [OPTIONS] NAME[:TAG]

Push an image to a registry.

Options:
      --help             Print usage
  -f, --configure-file   Add configure file with image

## unisctl rm --help

Usage: unisctl rm [OPTIONS] CONTAINER [CONTAINER...]

Remove one or more containers

Options:
      --help              Print Usage
  -f, --force              Force the removal of a running container (uses SIGKILL)

## unisctl rmi --help

Usage: unisctl rmi [OPTIONS] IMAGE [IMAGE...]

Remove one or more images

Options:
      --help              Print usage
  -f, --force             Force removal of the image

## unisctl run --help

Usage: unisctl run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container which will be autoly deployed to edge node

Options:
      --help              Print usage

## unisctl signin --help

Usage: unisctl signin -u USERNAME -p PASSWORD

Sign in

Options:
      --help             Print usage

## unisctl signup --help

Usage: unisctl signup -u USERNAME -p PASSWORD

Sign up a new account

Options:
      --help             Print usage

## unisctl stats --help

Usage: unisctl stats [OPTIONS]

List the status of all components of unis

Options:
      --help              Print Usage

## unisctl stop --help

Usage: unisctl stop [OPTIONS] CONTAINER [CONTAINER...]

Stop one or more running containers

Options:
      --help              Print usage
  -t, --time int          Seconds to wait for stop before killing it (default 10)

## unisctl tag --help

Usage: docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE

Options:
      --help              Print usage

## unisctl version --help

Usage: unisctl version [OPTIONS]

Show unis component version

Options:
      --help              Print usage