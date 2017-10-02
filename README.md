# Gobeet
Gobeet is a simple website health-checker written in go

Gobeet takes a configuration file and will consistently ping your website, sending you an email if any route is down for too long.

## Arguments
Gobeet takes two command line arguments:

* `--config=config.json`: Points to the configuration file to use.
* `--password=******`: The email password for the configured email address.

## Config File
Gobeet configuration files are composed of three main sections
* [Email Configuration](#email-configuration)
* [Websites](#sites)
* [Pages](#pages)

### Email Configuration
This is where email is configured. If you don't want to get emails, don't add this section.

A standard email configuration looks like this:

```
"email" : {
  "recipients": [
    "recipient1@recipient.com",
    "recipient2@recipient.com"
  ],
  "template": "sampletemplate.tmpl"
  "address": "test@sender.com",
  "server": "smtp.recipient.com",
  "port": "587",
}
```
* `recipients`: Who gets an alert when your site is down.
* `template`: Path to a valid [Golang template](#email-template) to use for the email.
* `address`: The email address to send out the alerts.
* `server`: SMTP email server.
* `port`: STMP email server port. Usually 587.

### Sites
A list of all domains to test.

A typical sites list will look like so:

```
"sites": {
  "url": "www.github.com",
  "interval": 60,
  "pages": [
    ...
  ]
}
```
* `url`: The URL to test.
* `interval`: How many seconds to wait before testing all pages in the site again
* `pages`: A list of pages

### Pages
A page is simply a part of your site that gets tested at a specified interval.

```
{
  "url": "https://github.com/danman113/",
  "method": "GET",
  "timeout": 5000,
  "status": 200,
  "duration": 10
}
```

* `url`: The page to test.
* `method`: The HTTP method to test the page with.
* `timeout`: The number of milliseconds to wait before deciding if the page is down.
* `status`: The expected HTTP status code the page is supposed to return. Gobeet will send an error if this result is different than this.
* `duration`: How many seconds to wait after a page is down before sending an email.

## Email Template
Gobeet uses standard [Golang HTML Templates](https://golang.org/pkg/html/template/). The two exposed variables are:
* `Site`: The current page that's down.
* `Error`: The error object.

```
<div class="header">
  <h1>Your website, <a href="{{.Site.Url}}">{{.Site.Url}}</a> has been down for {{.Site.Duration}} seconds!</h1>
</div>
<div class="content">
  <h3>The error is: </h3>
  <div class="pre">
    {{.Error}}
  </div>
</div>
```
