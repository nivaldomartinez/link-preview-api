# Link-Preview Backend

[link-prevue](https://github.com/nivaldomartinez/link-prevue)

[link-prevue Demo](https://link-prevue.herokuapp.com/)

[link-prevue Demo Repo](https://github.com/nivaldomartinez/link-prevue-demo)

## Install GoDep

In the root project folder run the following command

```sh
$ go get github.com/tools/godep
```

and run

```sh
$ godep save
```

## Deploy on Heroku

heroku login

```sh
$ heroku login
Enter your Heroku credentials.
Email: you@email.com
Password (typing will be hidden):
Logged in as you@email.com
```

add code and dependencies, like

```sh
$ git add -A .
$ git commit -m "code and deps"
```

create a heroku app

```sh
$ heroku create appname
```

now deploy your app

```sh
$ git push heroku master
```

visit the app deployed

```sh
$ heroku open
```

:heart:
