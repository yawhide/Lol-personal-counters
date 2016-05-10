# Lol-personal-counters
Submission for the Riot April 2016 api challenge

## About
This is our submission for the April 2016 api challenge! We wanted to make an app that would help the community out there no matter how good they are (non-ranked up to challenger). We love playing around counters during champion select and we thought making a tool that could help us determine the best champion to play in our arsenal would be amazing for us and the community. 

Lol personal counters is a website that gets both popular win rates and your own personal win rates. It shows you in a very sexy manner which champions you should play given your enemy and lane. its simple to use and understand but provides huge insight! We like some other sites that tell us counter matchups, but no of them curate their lists strickly to the champions we actively play. If we don't own Kathus, I don't want to see how he is the BEST champion to play into Katarina midlane. But we do own Orianna so its only logical to show her instead.

The biggest hurdles for this api challenge, from a technical standpoint, was to combine personal win rates of champion matchups (and getting enough data for it!) with third party website win rates. After creating the website, we were a little disappointed on the champion usage diversity (some champions were playing against a lot more then others). We decided to keep the personal win rates and the third party win rates separate because we couldnt think of a good mathematical squation to properly weight these two win rates. Going forward, we would definitely rate personal win rates higher then third party as the third party win rates are biased to platinum+ players (which MOST of the community isnt there yet).

In conclusion, we would like to thank YOU riot for giving up the oppotunity to show off our skills and give back to the community we love. Thanks!

## Set backs
- Proper loading page when users first check a matchup. Please be patient!! We are getting your rank stats!
- Proper algorithm to combine personal win rates and third party win rates

## Setting up the technology
1. Install golang by following these steps exaclty on the [golang website](https://golang.org/doc/install)
2. Ensure you have `$GOPATH` environment variable and that you can run `go` and `go get` from the command line.
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
15. if you experience any issues, please email me us at `ert.mcscrad@gmail.com` or write down an app note on my application. Thanks!

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
