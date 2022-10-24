# Twilio-Video-Demo
An application built using Svelte and Go to demonstrate setting up video chats via Twilio's video API


Technical Requirements
* Go 1.19 or higher
* An up-to-date NPM installation


Installation
------------
```bash
$ git clone https://github.com/ybjozee/Twilio-Video-Demo.git
$ cd Twilio-Video-Demo
```

Usage
-----
Set up the frontend 

```bash
$ cd frontend
$ npm install
```

For the backend, make a local version of the `.env` file

```bash
$ cd backend
$ cp .env .env.local
```

Update the relevant Twilio keys in `.env.local`

``` ini
TWILIO_ACCOUNT_SID="<<TWILIO_ACCOUNT_SID>>"
TWILIO_API_KEY_SID="<<TWILIO_API_KEY_SID>>"
TWILIO_API_SECRET_KEY="<<TWILIO_API_SECRET_KEY>>"
```

Run your application
```bash
# frontend
$ npm run dev
# application will run on port 8080

# backend
$ go run main.go
# application will run on port 8000
```