# Lol-personal-counters
Submission for the Riot April 2016 api challenge, the `Riot2016Submission` branch is for you while the master branch will be our on going effect to improve this website.

## About
This is our submission for the April 2016 api challenge! We wanted to make an app that helps the community, no matter a player’s rank. We love playing around counters during champion select, so we created a tool to help us determine the best champion in our arsenal to pick.

Lol personal counters is a website that determines both popular and personal win rates. With this information, it helps you figure out which champion you already know (and do well with), that you should pick to counter your lane opponent. The website is simple to use and understand but provides huge insight! There are other sites that determine counter matchups, but none of these sites restrict their lists to champions we actively play. For example, if you don't own Karthus, it’s not helpful for you to find out that he is the best counter to Katarina. With our site however, a champion you do own, like Orianna, is a logical counter to show.

The biggest hurdle for this api challenge, from a technical standpoint, was to combine personal win rates of champion matchups (and getting enough data for it!) with third party website win rates. After creating the website, we were a little disappointed about champion diversity (there are a few champions who are played most often). We decided to keep personal win rates and third party win rates separate because we didn’t have a good algorithm to properly weigh both win rates. If we continue this project, we would rate personal win rates higher than third party ones. Third party win rates are biased because they are so heavily weighted by platinum+ player performance, which not everyone in the community can live up to.

In conclusion, we would like to thank Riot for giving us the opportunity to show off our skills and give back to the community we love. Thanks!

## Set backs
- Proper loading page when users first check a matchup. Please be patient!! We are getting your rank stats! (little hack is to just refresh the page after a few second and the win rates will be determined in the background)
- Proper algorithm to combine personal win rates and third party win rates

## Setting up the technology
1. Install golang by following these steps exaclty on the [golang website](https://golang.org/doc/install)
2. Ensure you have `$GOPATH` environment variable and that you can run `go` and `go get` from the command line.
3. Run `mkdir -p $GOPATH/src/github.com/yawhide` to setup the workspace
4. Clone the git repo by running `git clone https://github.com/yawhide/Lol-personal-counters.git %GOPATH/src/github.com/yawhide/Lol-personal-counters` (must be inside the `%GOPATH/src/github.com/yawhide/Lol-personal-counters` directory) and ensure you are on the `Riot2016Submission` branch
5. Install postgres 9.5
    - Please leave the port on 5432
    - On windows you can download this [installer](http://www.enterprisedb.com/products-services-training/pgdownload#windows), ensure you choose v9.5
    - On ubuntu, you can follow this guide to setup postgres9.5: [guide](https://www.howtoforge.com/tutorial/how-to-install-postgresql-95-on-ubuntu-12_04-15_10/)
    - On mac, you can download [Postgres.app](http://postgresapp.com/)
    - Ensure postgres is running
    - Run `psql` on the command line and it should start the command line version of postgres
    - Run `\du` and find out which role has superuser
    - Copy the config.json.sample to config.json and edit the file
    - Add the superuser username/password into the config.json under the `postgres` key
    - Add the riot api key and the champion.gg api key (the champion.gg api key will be provided in the notes on the application)
6. cd to `%GOPATH/src/github.com/yawhide/Lol-personal-counters` and run `go get github.com/tools/godep`
7. Run `$GOPATH\bin\godep restore`
8. If you get an error talking about "error downloading dep (launchpad.net/go-xdg): exec "brz" blah blah blah 
    - run `go get -u -insecure github.com/yawhide/go-lol`
9. Run `go run main.go api.go analytics.go` from the github root folder
10. If all goes well, the application starts with no errors, you should see `Server started` and you can navigate to `localhost:8080`
11. Hit `ctrl-c` or `cmd-c` to stop the server
12. Navigate to `scripts/` inside the github folder and run `go run data.go api.go analytics.go`
    - if you see a bunch of urls hitting champion.gg, that is good. we are getting the latest win rates for the current patch. wait until its done (it should exit)
13. Now we are ready to go! run `cd ..` to go back to the root directory and run `go install .`
14. Now finally run `$GOPATH/bin/Lol-personal-counters` and navigate to `localhost:8080`. Please play around and enjoy the app!
15. If you experience any issues, please email me us at `ert.mcscrad@gmail.com` or write down an app note on my application. Thanks!

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
