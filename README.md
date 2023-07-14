# Web Crawler

The Web Crawler is an application/API that crawls websites and saves the HTML content to files.

## Description

The Web Crawler allows you to specify a list of target URLs to crawl. It fetches the HTML content of each URL and saves it as a separate HTML file. This can be useful for tasks such as data scraping, content analysis, or archival purposes.

## Prerequisites

- Go (Golang) is installed on your system. You can download Go from the official website: https://golang.org/dl/

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/arifrhm/cmlabs-backend-crawler-freelance-test.git
   ```

2. Change into the project directory:

    ```bash
    cd web-crawler
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

## Configuration
You can configure the target URLs to crawl in the crawlWebsites function in main.go. Simply add or remove URLs in the targetURLs slice.

```go
targetURLs := []string{
    "https://example.com",
    "https://example2.com",
    // Add more URLs
}
```
## Usage
1. Ensure that the application is running by executing the command go run main.go.

2. The application will crawl each URL specified in the configuration.

3. The crawled HTML content will be saved as separate HTML files, with each file named after the domain name of the corresponding URL.

## Contribution
Contributions are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.