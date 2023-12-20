# GoDom - Domain and SSL Expiration Checker

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

- Check domain validity and SSL certificate expiration for a list of domains.
- Send email notifications with domain and SSL expiration details.
- Supports reading a list of domains from a file.
- Configurable via command-line flags for flexibility and ease of use.

## Requirements

- Go 1.20 or higher recommended (for build).
- Access to an SMTP server for sending emails.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/KeepSec-Technologies/GoDom
    ```

2. Navigate to the cloned directory:

    ```bash
    cd GoDom
    ```

3. Build the tool:

    ```bash
    go build -o godom
    ```

## Building from Source

1. Ensure you have Go installed on your system. You can download Go from [here](https://golang.org/dl/).
2. Clone the repository:

    ```bash
    git clone https://github.com/KeepSec-Technologies/GoDom
    ```

3. Navigate to the cloned directory:

    ```bash
    cd GoDom
    ```

4. Build the tool:

    ```bash
    go build -o godom
    ```

## Usage

Run the GoDom tool with the required flags:

```bash
./godom --smtp-server=<smtp_server> --smtp-port=<smtp_port> \
        --smtp-username=<username> --smtp-password=<password> \
        --from-email=<from_email> --to-email=<to_email> \
        --domains-file=<path_to_domains_file>
```

### Flags

```text
-s, --smtp-server: SMTP server for sending emails.
-p, --smtp-port: SMTP server port (default: 587).
-u, --smtp-username: Username for SMTP authentication.
-w, --smtp-password: Password for SMTP authentication.
-f, --from-email: Email address to send notifications from.
-t, --to-email: Email address to send notifications to.
-d, --domains-file: Path to the file containing domain names.
```

### Example

```bash
./godom --smtp-server=smtp.example.com --smtp-port=587 \
        --smtp-username=user@example.com --smtp-password=pass \
        --from-email=<noreply@example.com> --to-email=<admin@example.com> \
        --domains-file=domains.txt
```

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues to improve the functionality or fix problems with GoDom.

## License

This project is licensed under MIT - see the LICENSE file for details.
