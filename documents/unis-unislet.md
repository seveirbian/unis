# unis-unislet

## unislet

unislet

Usage: unislet COMMAND

An edge node to execute apiserver cmd

Commands:  
  connect    Connect to the remote unis-apiserver  
  run        Run as unis component
  signin     Sign in
  stats      Display the status of this node    
  version    Show the unis-cli and unis-apiserver version information  

Run 'unislet COMMAND --help' for more information on a command.

## unislet connect --help

Usage: unislet connect SERVER

Connect to unis-apiserver

Options:
      --help             Print usage

## unislet run --help

Usage: unislet run [OPTIONS] 

Run as unis component

Options:
  -e, --environment       Set environment(Docker or Unikernel)
    --help              Print usage
  -p, --password          Set password
  -t, --NodeType          Set node type
  -u, --username          Set user name

## unislet signin --help

Usage: unislet signin [OPTIONS]

Sign in

Options:
      --help            help for run
  -p, --password        password
  -u, --username        username

## unislet stats --help

Usage: unislet stats [OPTIONS]

List the status of this node

Options:
      --help              Print Usage

## unislet version --help

Usage: unislet version [OPTIONS]

Show unislet version

Options:
      --help              Print usage
