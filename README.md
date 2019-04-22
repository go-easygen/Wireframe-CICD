
# Wireframe-CICD - Wireframe CI/CD Demo

Wireframe CI/CD Demo

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-easygen/Wireframe-CICD?status.svg)](http://godoc.org/github.com/go-easygen/Wireframe-CICD)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-easygen/Wireframe-CICD)](https://goreportcard.com/report/github.com/go-easygen/Wireframe-CICD)
[![travis Status](https://travis-ci.org/go-easygen/Wireframe-CICD.svg?branch=master)](https://travis-ci.org/go-easygen/Wireframe-CICD)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

## TOC
- [SYNOPSIS](#synopsis)
- [OPTIONS](#options)
- [DESCRIPTION](#description)
- [Download/Install](#downloadinstall)
  - [Using `apt`](#using-`apt`)
  - [Download binaries](#download-binaries)
  - [Debian package](#debian-package)
  - [Install Source](#install-source)
- [EXAMPLES](#examples)
- [SEE ALSO](#see-also)
- [AUTHORS](#authors)

The following are just borrowed from [`ffcvt`](https://github.com/suntong/ffcvt).
Only to show the TOC feature.

## SYNOPSIS

    ffcvt [flags] [video_filename]


## OPTIONS

For `flags`, run `ffcvt` without parameters to see the comprehensive list and explanation. 


## DESCRIPTION

Command `ffcvt` is an easy to use video converter that takes the burden of such daunting task from normal Joe, to harness the fantastic high efficiency audio/video codec/encoding capability long been available in `ffmpeg`, by simplifying the convoluted `ffmpeg` command line parameters.

## Download/Install

### Using `apt`

The `ffcvt` is now officially in Debian repository, so the installation is now as simple as a `apt install`/`apt-get install`:

    apt install ffcvt

### Download binaries

- The latest binary executables are available under  
https://bintray.com/suntong/bin/Wireframe-CICD/  
as the result of the Continuous-Integration process.
- I.e., they are built right from the source code during _every_ git commit _automatically_ by [travis-ci](https://travis-ci.org/).
- Pick & choose the latest binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `Wireframe-CICD-linux-amd64` file (if your OS and its architecture is not available in the download list, please let me know and I'll add it).
- Download the executable and put it somewhere in the PATH
- You may want to rename it to a shorter name instead, e.g., `Wireframe-CICD`, after downloading it.


### Debian package

Debian package _repo_ is available at https://dl.bintray.com/suntong/deb.  
The _browse-able_ repo view is at https://bintray.com/suntong/deb.

```
echo "deb [trusted=yes] https://dl.bintray.com/suntong/deb all main" | sudo tee /etc/apt/sources.list.d/suntong-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/suntong-debs.list
apt-cache policy Wireframe-CICD

sudo apt-get install -y Wireframe-CICD
```

### Install Source

If you prefer to compile and install `Wireframe-CICD` from source, although a manual process, it's pretty straightforward and simple.

1. Get the source via `git clone` or [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies).
1. Do `cd Wireframe-CICD`, then issue [`go build`](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) without any other parameters.
1. Copy the generated executable somewhere in the PATH

That's it, it's ready to roll. 


## EXAMPLES

`ffcvt` comes with comprehensive online documents. For further details,
please go to:

https://github.com/suntong/ffcvt#encoding-help

https://github.com/suntong/ffcvt/wiki


## SEE ALSO

`ffmpeg(1)`


## AUTHORS

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")
