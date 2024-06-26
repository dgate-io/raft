<p align="center">
  <img alt="raft" src="assets/raft_logo.png">
</p>


![build](https://github.com/dgate-io/raft/actions/workflows/build.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/dgate-io/raft)](https://pkg.go.dev/github.com/dgate-io/raft)
![GitHub](https://img.shields.io/github/license/dgate-io/raft)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dgate-io/raft)

# Raft

This package provides a simple, easy-to-understand, and reliable implementation of [Raft](https://en.wikipedia.org/wiki/Raft_(algorithm)) using Go. Raft is a [consensus](https://en.wikipedia.org/wiki/Consensus_(computer_science)) protocol designed to manage replicated logs in a distributed system. Its purpose is to ensure fault-tolerant coordination and consistency among a group of nodes, making it suitable for building reliable systems. Potential use cases include distributed file systems, consistent key-value stores, and service discovery.

# Features

The following features are currently supported:
- Automated Snapshots
- Concurrent Snapshot Transfer from Leader to Followers
- Dynamic Membership Changes
- Linearizable and Lease-Based Read-Only Operations
- Prevote and Leader Stickyness

# Protocol Overview

Raft is based on a leader-follower model, where one node is elected as the leader and coordinates the replication process. Time is divided into terms, and the leader is elected for each term through a leader election process. The leader receives client requests, which are then replicated to other nodes called followers. The followers maintain a log of all state changes, and the leader's responsibility is to ensure that all followers have consistent logs by sending them entries to append. Safety is guaranteed by requiring a majority of nodes to agree on the state changes, ensuring that no conflicting states are committed. 

# Installation and Usage

First, make sure you have Go `1.20` or a higher version installed on your system.
You can download and install Go from the official [Go website](https://golang.org/). 

Then, install the raft package by running

```shell
go get -u github.com/dgate-io/raft
```

Once you have the package installed, you may refer to the [raft package reference page](https://pkg.go.dev/github.com/dgate-io/raft) for basic usage.
An example of how to use raft as well as a set of [Jepsen](https://github.com/jepsen-io/jepsen) tests can be found [here](https://github.com/dgate-io/raft-example).

# Contributing

Other developers are encouraged to contribute to this project and pull requests are welcome. Please read
these [guidelines](.github/CONTRIBUTING.md) if you are interested in contributing.
