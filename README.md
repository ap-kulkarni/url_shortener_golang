# URL Shortener 
This repository contains solution for the URL Shortener.

The program listens on port 8080. The shortened urls are hard coded with localhost:8080 as the host part. 

## API Description
### Shorten URL
|                |              |
|----------------|--------------|
|  **Endpoint**  | POST /shorten|
|**Request Body**| {"url": "url-to-shorten"}|
|  **Response**  | {"short_url": "shortened-url"}|
|  **Sample CURL req** | curl -i -X POST http://localhost:8080/shorten -d \"{ \\\"url\\\" : \\\"https://y45r.com/search?q=fu45\\\"} \" |
|  **Sample Response**| {"short_url": "http://localhost:8080/h10U93l"} |

### Get long URL
|                |              |
|----------------|--------------|
|  **Endpoint**  | GET /(short-segment)|
|**Request Body**| -                    |
|  **Response**  | Long URL in **location** header with **301** status code|
| **Sample CURL req**| curl -i http://localhost:8080/h10U93l |

### Metrics
While generating metric, we consider complete host part of the url. This is because e.g. go.dev and go.com will be different domains.
However, considering only one word domains will make them same.

|                |              |
|----------------|--------------|
|  **Endpoint**  | GET /metric|
|**Request Body**| -                    |
|  **Response**  | Metric data in json with domain as field name and domain count as field value|
| **Sample CURL req**| curl -i http://localhost:8080/metric |
| **Sample Response**|{"y45r.com": 1,"yahoo.com": 1}        |


## Running instructions


### Running locally
To run locally, checkout the project and run with go run
```
git clone https://github.com/ap-kulkarni/url_shortener_golang.git url_shortener
cd url_shortener
go run main.go
```

### Running with Docker
#### With image from docker hub
```
docker run -d -p 8080:8080 ameyk2409/url_shortener
```

#### With docker file from the repository
```
git clone https://github.com/ap-kulkarni/url_shortener_golang.git url_shortener
cd url_shortener
docker build -t url_shortener .
docker run -d -p 8080:8080 url_shortener
```
