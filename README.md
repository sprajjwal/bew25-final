# Image downloader tool

## Problem

Here at *Advertising Startup* we have to work with stock images all the time. The Image downloader CLI tool downloads any JPG from a given URL. Additionally it can take file names to name the downloaded image as a flag `-file`

## Usage

1. Clone the repo.
2. Call the CLI using `./getImage -url=<URL> -file=<file>`

## Example

Using only the `url` flag.
![using flag](./screenshots/flag_operator.png "Using a flag")

Image Download Confirmation.
![image confirmation](./screenshots/success.png "Confirmation")

When trying to download images that might have the same file name as existing one
![redownload](screenshots/redownload.png "Redownload")

Random file name gets assigned for existing files.
![random file](screenshots/random_name.png "random file name")

Additionally images can be downloaded with custom names like this.
![custom name flag](screenshots/get_with_custom_name.png)

This creates `image.jpg` for us.
![custon image name](screenshots/custom_name.png "custom named file")