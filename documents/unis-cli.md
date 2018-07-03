# unis-cli

## unisctl

unisctl

Usage: unisctl COMMAND

A client to communicate with unis-apiserver

Commands:  
  connect    Connect to the remote unis-apiserver  
  images     List images in remote registry  
  nodes      Display the status of all edge nodes  
  ps         List containers  
  pull       Pull an image from registry  
  push       Push an image to registry  
  rmi        Remove one or more images in remote registry  
  run        Run a container on a edge node  
  signin     Sign in  
  stats      Display the status of all components of unis  
  stop       Stop one or more running containers  
  tag        Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE  
  version    Show the unis-cli and unis-apiserver version information  

Run 'unisctl COMMAND --help' for more information on a command.

## unisctl connect --help

Usage: unisctl connect SERVER

Connect to unis-apiserver

Options:
      --help             Print usage

## unisctl images --help

Usage: unisctl images [OPTIONS]

List images

Options:
  -a, --all              Show all images (default private images)
      --help             Print usage

## unisctl nodes --help

Usage: unisctl nodes [OPTIONS]

List the status of all edge nodes

Options:
  -a, --all               Show all nodes (default private nodes)
      --help              Print Usage

## unisctl ps --help

Usage: unisctl ps [OPTIONS]

List instances

Options:
  -a, --all              Show all instances (dafault private instances)
      --help             Print Usage

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

## unisctl rmi --help

Usage: unisctl rmi [OPTIONS] IMAGE [IMAGE...]

Remove one or more images

Options:
      --help              Print usage

## unisctl run --help

Usage: unisctl run [OPTIONS] IMAGE

Run a container on an edge node

Options:
      --help              Print usage

## unisctl signin --help

Usage: unisctl signin [OPTIONS]

Sign in

Options:
      --help            help for run
  -p, --password        password
  -u, --username        username

## unisctl stats --help

Usage: unisctl stats [OPTIONS]

List the status of all components of unis

Options:
      --help              Print Usage

## unisctl stop --help

Usage: unisctl stop [OPTIONS] INSTANCE [INSTANCE...]

Stop one or more running instances

Options:
      --help              Print usage

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
