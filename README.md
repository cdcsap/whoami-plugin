# whoami-plugin

[![Build Status](https://travis-ci.org/jtuchscherer/whoami-plugin.svg?branch=master)](https://travis-ci.org/jtuchscherer/whoami-plugin)

## Installation

```bash
 cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
 cf install-plugin "Whoami Plugin" -r CF-Community

```

## Usage
`cf whoami`

## Uninstall

```bash
cf uninstall "Whoami Plugin"
```

## Testing

Run tests
```bash
./scripts/test.sh
```

If you want to install the plugin locally and test it manually
```bash
./scripts/install.sh
```

## Releasing

In order to create a new release, follow these steps

1. Create local tag and binaries
  ```
  ./scripts/build-all.sh release VERSION_NUMBER #(e.g. 0.7.0)
  ```
1. Copy the output of the previous command from the first line (should be '- name: Firehose Plugin' to the last checksum line (should be something like checksum: fde5fd52c40ea4c34330426c09c143a76a77a8db)
1. Push the tag `git push --follow-tags`
1. On github, create new release based on new tag [here](https://github.com/cloudfoundry/firehose-plugin/releases/new)
1. Upload the three binaries from the ./bin folders to the release (Linux, OSX and Win64)
1. Fork [this repo](https://github.com/cloudfoundry-incubator/cli-plugin-repo) and clone it locally
1. Edit the repo-index.yml
  ```
  vi repo-index.yml
  ```
  to override the existing section about the firehose plugin with the text previously copied in Step 2.
1. Push the change to your fork
1. Create a PR against the [original repo](https://github.com/cloudfoundry-incubator/cli-plugin-repo/compare)

```
