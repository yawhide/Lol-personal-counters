# Lol-personal-counters
Submission for the riot april 2016 api challenge

## Setting up the technology
1. Install golang by following these steps exaclty on the [golang website](https://golang.org/doc/install)
2. Ensure you have $GOPATH environment variable and that you can run `go` and `go get` from the command line.
3. Run `mkdir -p $GOPATH/src/github.com/yawhide` to setup the workspace
4. clone the git repo by running `git clone https://github.com/yawhide/Lol-personal-counters.git %GOPATH/src/github.com/yawhide/Lol-personal-counters` (must be inside the `%GOPATH/src/github.com/yawhide/Lol-personal-counters` directory) and ensure you are on the `Riot2016Submission` branch
5. Install postgres 9.5
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
6. cd to `%GOPATH/src/github.com/yawhide/Lol-personal-counters` and run `go get github.com/tools/godep`
7. run `$GOPATH\bin\godep restore`
8. If you get an error talking about "error downloading dep (launchpad.net/go-xdg): exec "brz" blah blah blah 
    - run `go get -u -insecure github.com/yawhide/go-lol`
9. run `go run main.go api.go analytics.go` from the github root folder
10. if all goes well, the application starts with no errors, you should see `Server started` and you can navigate to `localhost:8080`
11. hit `ctrl-c` or `cmd-c` to stop the server
12. navigate to `scripts/` inside the github folder and run `go run data.go api.go analytics.go`
    - if you see a bunch of urls hitting champion.gg, that is good. we are getting the latest win rates for the current patch. wait until its done (it should exit)
13. now we are ready to go! run `cd ..` to go back to the root directory and run `go install .`
14. now finally run `$GOPATH/bin/Lol-personal-counters` and navigate to `localhost:8080`. Please play around and enjoy the app!
15. if you experience any issues, please email me at `ert.mcscrad@gmail.com` or write down an app note on my application. Thanks!

### BRZ command
If you are getting an error talking about some `brz` command, try installing [bazaar](http://wiki.bazaar.canonical.com/Download) and adding it to your path

### Frameworks used
You may need to manually `go get` these to make the app work, please run `go get` then enter the github link below
- github.com/tools/godep
- github.com/spf13/viper
- gopkg.in/pg.v4
- github.com/yawhide/go-lol


### start postgres on windows
`"pg_ctl" -D datafolderpath -l log\logfile start`

### access postgres on ubuntu
`sudo -i -u postgres`
