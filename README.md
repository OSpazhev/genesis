# BTC application

## How to run
Run the next commands to start the app:
- `docker build -t btc-application .`
- `docker run -p 8080:8080 btc-application`

App is available on http://localhost:8080

If you want to remember subscribers from previous run, please run docker with parameter `-v` like:
- `docker build -t btc-application . -v <local path>:/data`






Author: Spazhev Oleksandr