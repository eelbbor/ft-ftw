Feature Toggles For The Win
==========

This is the base code for a Feature Toggle service.  The goal is to enable a RESTful service for creating and using Feature Toggles during development.  If your working in a SOA it can be a bit of a pain to implement Feature Toggles independently in every service when adding a cross-cutting feature.  A singular source for toggles can be tremendously useful, which is the goal of this work.

Project structure is built to allow for multiple go projects rather than just the singular directory definition.  You will need to add the following to your environment (as this project evovles I will add more clarity to this):

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/projects/go:$HOME/projects/ft-ftw
export PATH=$PATH:${GOPATH//://bin:}/bin
```
