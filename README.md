# kubectl watch-diff

A simple kubectl plugin that watches the changes in the k8s resource and
prints the diff.

## Installation

```bash
go install github.com/alexmt/kubectl-watch-diff@latest
```

## Usage

```bash
kubectl watch diff cm argocd-cm
```

[![asciicast](https://asciinema.org/a/670880.svg)](https://asciinema.org/a/670880)