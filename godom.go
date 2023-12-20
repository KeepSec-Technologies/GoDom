package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/domainr/whois"
)

var (
	// Long-form flags
	smtpServer  string
	smtpPort    string
	username    string
	password    string
	fromEmail   string
	toEmail     string
	domainsFile string

	// Short-form flags
	smtpServerShort  string
	smtpPortShort    string
	usernameShort    string
	passwordShort    string
	fromEmailShort   string
	toEmailShort     string
	domainsFileShort string
)

func init() {
	// Long-form flags
	flag.StringVar(&smtpServer, "smtp-server", "", "SMTP server for sending emails")
	flag.StringVar(&smtpPort, "smtp-port", "587", "SMTP server port")
	flag.StringVar(&username, "smtp-username", "", "Username for SMTP authentication")
	flag.StringVar(&password, "smtp-password", "", "Password for SMTP authentication")
	flag.StringVar(&fromEmail, "from-email", "", "Email address to send notifications from")
	flag.StringVar(&toEmail, "to-email", "", "Email address to send notifications to")
	flag.StringVar(&domainsFile, "domains-file", "", "Path to the file containing domain names")

	// Short-form flags
	flag.StringVar(&smtpServerShort, "s", "", "SMTP server for sending emails (short)")
	flag.StringVar(&smtpPortShort, "p", "587", "SMTP server port (short)")
	flag.StringVar(&usernameShort, "u", "", "Username for SMTP authentication (short)")
	flag.StringVar(&passwordShort, "w", "", "Password for SMTP authentication (short)")
	flag.StringVar(&fromEmailShort, "f", "", "Email address to send notifications from (short)")
	flag.StringVar(&toEmailShort, "t", "", "Email address to send notifications to (short)")
	flag.StringVar(&domainsFileShort, "d", "", "Path to the file containing domain names (short)")
}

func main() {
	flag.Parse()

	// Override long-form flags with short-form flags if set
	if smtpServerShort != "" {
		smtpServer = smtpServerShort
	}
	if smtpPortShort != "" {
		smtpPort = smtpPortShort
	}
	if usernameShort != "" {
		username = usernameShort
	}
	if passwordShort != "" {
		password = passwordShort
	}
	if fromEmailShort != "" {
		fromEmail = fromEmailShort
	}
	if toEmailShort != "" {
		toEmail = toEmailShort
	}
	if domainsFileShort != "" {
		domainsFile = domainsFileShort
	}

	// Check if required flags are missing
	if smtpServer == "" || username == "" || password == "" || fromEmail == "" || toEmail == "" || domainsFile == "" {
		usage()
	}

	domains, err := getDomains(domainsFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range domains {
		sslExpDate := checkSSLExpiration(domain)
		domainExpDate := checkDomainExpiration(domain)

		message := fmt.Sprintf("Domain: %s\nSSL Expiration: %s\nDomain Expiration: %s\n\n", domain, sslExpDate, domainExpDate)
		sendEmail(fromEmail, toEmail, "Domain and SSL Check Results", message)
	}

	fmt.Println("Execution complete")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  -s, --smtp-server         SMTP server for sending emails")
	fmt.Fprintln(os.Stderr, "  -p, --smtp-port           SMTP server port")
	fmt.Fprintln(os.Stderr, "  -u, --smtp-username       Username for SMTP authentication")
	fmt.Fprintln(os.Stderr, "  -w, --smtp-password       Password for SMTP authentication")
	fmt.Fprintln(os.Stderr, "  -f, --from-email          Email address to send notifications from")
	fmt.Fprintln(os.Stderr, "  -t, --to-email            Email address to send notifications to")
	fmt.Fprintln(os.Stderr, "  -d, --domains-file        Path to the file containing domain names")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "  Ensure all required flags are provided.")
	os.Exit(1)
}

// getDomains retrieves a list of domains from the specified file
func getDomains(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}

// checkSSLExpiration checks the SSL certificate expiration of a domain
func checkSSLExpiration(domain string) string {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		log.Println("Error connecting:", err)
		return "Error"
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	return cert.NotAfter.Format(time.RFC1123)
}

// checkDomainExpiration checks the domain expiration date
func checkDomainExpiration(domain string) string {
	// Perform a WHOIS query
	response, err := whois.Fetch(domain)
	if err != nil {
		log.Printf("Error performing WHOIS query for %s: %v\n", domain, err)
		return "Error"
	}

	return parseWhoisOutput(response.String())
}
func parseWhoisOutput(output string) string {
	// Common patterns in 'whois' output for expiration date
	var expirationPatterns = []string{
		"Expiry Date:",          // Common pattern
		"Registry Expiry Date:", // Another common pattern
		"Expires On:",           // Another variation
	}

	for _, pattern := range expirationPatterns {
		if strings.Contains(output, pattern) {
			// Extract the line containing the expiration date
			lines := strings.Split(output, "\n")
			for _, line := range lines {
				if strings.Contains(line, pattern) {
					// Extract and return the date portion of the line
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						dateStr := strings.TrimSpace(parts[1])
						// Optionally, further parse dateStr if needed to standardize the date format
						return dateStr
					}
				}
			}
		}
	}

	return "Unknown" // Return "Unknown" or handle it differently if no pattern matches
}

// sendEmail sends an email with the given details
func sendEmail(from, to, subject, body string) {
	// Set up authentication information
	auth := smtp.PlainAuth("", username, password, smtpServer)

	// Email headers and body
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}
