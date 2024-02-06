<p align="center">
 <img src="https://github.com/KeepSec-Technologies/GoDom/assets/108779415/430a7dd3-be24-4f32-8c12-818bc2ad26bb"
</p>

# GoDom - Domain and SSL Expiration Checker
![License](https://img.shields.io/github/license/KeepSec-Technologies/GoDom)
![GitHub issues](https://img.shields.io/github/issues-raw/KeepSec-Technologies/GoDom)
![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/KeepSec-Technologies/GoDom/main)

GoDom is a command-line tool written in Go, designed to check the validity of domain names and SSL certificates expirations, and then send notifications via email. It's an efficient way to stay ahead of domain and SSL renewals, ensuring your websites and services remain uninterrupted.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Building from Source](#building-from-source)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Check domain validity and SSL certificate expiration from a list of domains.
- Send email notifications with domain validity and SSL expiration details.
- Supports reading a list of domains from a file.
- Configurable via command-line arguments for flexibility and ease of use.

## Requirements

- Go 1.20 or higher recommended (for build).
- Access to an SMTP server for sending emails.
- (Optional) Making a cronjob out of this is the intended way to use it.

## Installation

1. Download the binary with wget:

    ```shell
    wget https://github.com/KeepSec-Technologies/GoDom/releases/download/1.1/godom_linux_amd64_1.1.tar.gz
    ```

2. Unpack it with tar

    ```shell
    tar -xf godom_linux_amd64_1.1.tar.gz
    ```

3. Move it to your /usr/local/bin/ (Optional):

    ```shell
    sudo mv godom /usr/local/bin/godom
    ```

## Building from Source

1. Ensure you have Go installed on your system. You can download Go from [here](https://golang.org/dl/).
2. Clone the repository:

    ```shell
    git clone https://github.com/KeepSec-Technologies/GoDom
    ```

3. Navigate to the cloned directory:

    ```shell
    cd GoDom
    ```

4. Build the tool:

    ```shell
    CGO_ENABLED=0 go build -a -installsuffix cgo -o godom .
    ```

## Usage

Put your domains in a text file, eg: domains.txt:
```text
example.com
example.org
example.ca
```

Run the GoDom tool with the required flags:

```shell
./godom --smtp-server <smtp_server> --smtp-port <smtp_port> \
        --smtp-username <username> --smtp-password <password> \
        --from-email <from_email> --to-email <to_email> \
        --domains-file <path_to_domains_file>
```

Flags:

```text
  -s, --smtp-server         SMTP server for sending emails
  -p, --smtp-port           SMTP server port
  -u, --smtp-username       Username for SMTP authentication
  -w, --smtp-password       Password for SMTP authentication
  -f, --from-email          Email address to send notifications from
  -c, --config              Path to the SMTP json config file which replaces the above arguments
  -t, --to-email            Email address to send notifications to
  -d, --domains-file        Path to the file containing domain names
```

Examples:

```shell
./godom -s smtp.example.com -p 587 -u user@example.com -w password123 -f godom@example.com -t admin@example.com -d domains.txt
```

or

```shell
./godom -c path/to/config.json -t admin@example.com -d domains.txt
```

Example of json config file to pass to GoDom:

```json
{
    "smtp_server": "mail.example.com",
    "smtp_port": 587,
    "smtp_username": "user@example.com",
    "smtp_password": "password123",
    "from_email": "mail2go@example.com"
}
```

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues to improve the functionality or fix problems with GoDom.

## License

This project is licensed under MIT - see the LICENSE file for details.
