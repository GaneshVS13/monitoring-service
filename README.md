# Monitoring Service
This is the external URL monitoring service which monitors the response status of external URLs.
URLs to be monitored are configurable. These URLs are loaded in the application on application startup.

# Prometheus metrics
Below metrics of each monitored URL are published using Prometheus metrics monitoring APIs.
1. Response status - If the service is up then metrics with value 1 is published else the value is set to 0.
2. Response duration - The response time of the URL in miliseconds is published.
Both the metrics are of type Guage vector with the URL as label is published.

## Examples 
###### HELP sample_external_url_response_ms Duration of HTTP request in miliseconds
###### TYPE sample_external_url_response_ms gauge
sample_external_url_response_ms{url="https://httpstat.us/200"} 693
sample_external_url_response_ms{url="https://httpstat.us/503"} 694

###### HELP sample_external_url_up Status of HTTP response
###### TYPE sample_external_url_up gauge
sample_external_url_up{url="https://httpstat.us/200"} 1
sample_external_url_up{url="https://httpstat.us/503"} 0

To visualize these metrics on Prometheus dashboard, go to http://localhost:9090
Add the sample_external_url_up or sample_external_url_response_ms in search and execute it.

# Grafana dashboard
To see the above metrics on Grafana, login to http://localhost:3000 with username and password as 'admin'.
Add the Prometheus datasource and then add the new dashboard with queries as 'sample_external_url_response_ms' and 'sample_external_url_up'.

# Usage
To run the application, go to /src folder path and run command 'make up-docker-env'.
This command will start the docker containers for service, prometheus and grafana.
