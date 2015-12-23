#!/bin/bash

set -e

(cf uninstall-plugin "Whoami Plugin" || true) && go build -o whoami-plugin main.go && cf install-plugin whoami-plugin
