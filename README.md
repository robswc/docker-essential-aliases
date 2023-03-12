![docker-essential-aliases-logo](https://user-images.githubusercontent.com/38849824/222996047-9c454b3c-c8f5-4199-9f6d-0e24f5043ea6.png)

# Docker Essential Aliases
Docker Essential Aliases (dea) is a set of aliases and utility functions for Docker.
The main goal of dea is to reduce the amount of typing and time spent on common Docker commands and tasks.

Please note, this was a project mainly for personal use, but I love open source and I hope others find it useful!
Just be aware I don't have 100% coverage (yet!)

* [Usage](#usage)
* [Features](#features)
* [Installation](#installation)
  * [Linux](#linux)
  * [Building from source](#building-from-source)

# Usage

```bash
dea [command] [options]
````

# Features

### Remembers last container
```bash
# Enter a container
dea enter [container] [options] # same as docker exec -it [container] /bin/bash
dea enter hello-world

# _If_ no container is specified, dea will use the last container you interacted with (logs, enter, etc)

dea enter # enters hello-world
dea logs # shows logs for hello-world

docker rm -f hello-world # remove hello-world
docker run hello-world # create new hello-world container
dea enter # first tries w/container ID but falls back on name, so it will still enter hello-world
```



### Easier ps
```bash
dea ps --columns [columns,sep,by,comma] # same as docker ps --format "table {{columns}}"
dea ps --set-columns [columns,sep,by,comma] # sets the default columns for dea ps
dea ps --reset-columns # resets the default columns for dea ps
```

### More to come!

If you have any suggestions, please [open an issue](https://github.com/robswc/docker-essential-aliases/issues/new)! And
label it "Feature Request."


# Installation

### Linux

Currently, only Linux is supported. If you would like to see support for other platforms, 
please [open an issue](https://github.com/robswc/docker-essential-aliases/issues/new).  I'm 99% sure it will work on Mac tho!

```bash
# Download the latest release via github api
curl -s https://api.github.com/repos/robswc/docker-essential-aliases/releases/latest | grep "browser_download_url.*dea" | cut -d : -f 2,3 | tr -d \" | wget -qi -

# Make executable
chmod +x dea

# Move to /usr/local/bin
sudo mv dea /usr/local/bin/dea
```

### Building from source

To build from source, you can clone the repo and use make!

```bash
# Clone the repo
git clone https://github.com/robswc/docker-essential-aliases.git
# Move into the directory
cd docker-essential-aliases
# Build
make

# Move to /usr/local/bin
sudo mv dea /usr/local/bin/dea
```