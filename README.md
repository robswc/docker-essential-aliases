![docker-essential-aliases-logo](https://user-images.githubusercontent.com/38849824/222996047-9c454b3c-c8f5-4199-9f6d-0e24f5043ea6.png)

# CURRENTLY IN TESTING

# Installation

## Linux

```bash
# Download the latest release via github api
curl -s https://api.github.com/repos/robswc/docker-essential-aliases/releases/latest | grep "browser_download_url.*dea" | cut -d : -f 2,3 | tr -d \" | wget -qi -

# Make executable
chmod +x dea

# Move to /usr/local/bin
sudo mv dea /usr/local/bin/dea
```