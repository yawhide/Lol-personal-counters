# Lol-personal-counters
Submission for the riot april 2016 api challenge

## Setting up the technology
1. Install golang by following these steps exaclty on the [golang website](https://golang.org/doc/install)
2. Ensure you have $GOPATH environment variable and that you can run `go` and `go get` from the command line
3. clone the git repo by running `git clone https://github.com/yawhide/Lol-personal-counters.git %GOPATH/src/github.com/yawhide/Lol-personal-counters` (must be inside the `%GOPATH/src/github.com/yawhide/Lol-personal-counters` directory) and ensure you are on the `Riot2016Submission` branch
4. Install postgres 9.5
    - Please leave the port on 5432
    - On windows you can go easy mode by downloading this [installer](http://www.enterprisedb.com/products-services-training/pgdownload#windows), ensure you choose v9.5
    - On ubuntu, you can follow this guide to setup postgres9.5: [guide](https://www.howtoforge.com/tutorial/how-to-install-postgresql-95-on-ubuntu-12_04-15_10/)
    - On mac, you can download [Postgres.app](http://postgresapp.com/)
    - Ensure postgres is running
    - run `psql` on the command line and it should start the command line version of postgres
    - run `\du` and find out which role has superuser
    - copy the config.json.sample to config.json and edit the file
    - add the superuser username/password into the config.json under the `postgres` key
    - mind as well add the riot api key and the champion.gg api key (the champion.gg api key will be provided in the notes on the application)
5. cd to `%GOPATH/src/github.com/yawhide/Lol-personal-counters` and run `go get github.com/tools/godep`
6. run `$GOPATH\bin\godep restore`
7. run `go install .` from the github root folder
8. if all goes well, the application starts with no errors, you should see `Server started` and you can navigate to `localhost:8080`

### Frameworks used
- https://github.com/codegangsta/gin
- https://github.com/spf13/viper
- https://gopkg.in/pg.v4
- https://github.com/tools/godep
- https://github.com/yawhide/go-lol


### start postgres on windows
`"pg_ctl" -D datafolderpath -l log\logfile start`

### access postgres
`sudo -i -u postgres`
