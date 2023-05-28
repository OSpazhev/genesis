# BTC application

## How to run
Run the next commands to start the app:
- `docker build -t btc-application .`
- `docker run -p 8080:8080 btc-application`

App is available on http://localhost:8080

### Volume
If you want to remember subscribers from previous run, please run docker with parameter `-v` like:
- `docker build -t btc-application . -v <local path>:/data`

### Email 
To send emails, you should set your `SENDER_EMAIL` and `SENDER_PASSWORD` from mailtrap.io  
If you want to use any other platform, you can also set `SMTP_HOST` and `SMTP_PORT`

#### P.s.: 
I know that key storing in Dockerfile (and in git) is not secure, but it's just for reviewer


#### Author
Spazhev Oleksandr