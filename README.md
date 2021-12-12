# ApplyConfiguration Types with DeepCopy functions

client-go provides [ApplyConfiguration types](https://pkg.go.dev/k8s.io/client-go/applyconfigurations) for typesafe Server-Side Apply.
However, those types don't have a DeepCopy function.
Especially it would be inconvenient when you embed an ApplyConfiguration type into your custom resource.

This repository provides the DeepCopy functions for ApplyConfiguration types generated by [deepcopy-gen](https://github.com/kubernetes/code-generator/tree/master/cmd/deepcopy-gen).
