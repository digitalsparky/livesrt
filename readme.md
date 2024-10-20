### OBS Config

Create a MEDIA source
Unselect "Local file"
Set the source to: srt://127.0.0.1:7654?streamid=play/live/%YOURKEY%
Set the Image Format to mpegts
Select "Use hardware encoding where available"

### WINDOWS
In order to run on windows, install docker desktop and setup your system with WSL2
Instructions available here: https://docs.docker.com/desktop/install/windows-install/

### Setup

Update config in noalbs/config.json and .env to set the correct SRT external port, Twitch user config and noalbs srt config.
The example configuration is already setup for use with the provided NOALBS and SRT server addresses, so you can ignore the stream config:

```json
"streamServers": [
      {
        "streamServer": {
          "type": "SrtLiveServer",
          "statsUrl": "http://srt:8181/stats",
          "publisher": "publish/live/%YOURKEY%"
        },
        "name": "SRT",
        "priority": 0,
        "overrideScenes": null,
        "dependsOn": null,
        "enabled": true
      }
    ]
  },
```
*Important*: change 'feed1' to a your own stream key (see security section)

All setup, you can use the following docker-compose.yml

```yaml
services:
    srt:
        image: b3ckontwitch/sls-b3ck-edit:latest
        volumes:
            - ./srt/sls.conf:/etc/sls/sls.conf
        ports:
            - 7654:30000/udp
            - 8181:8181
    noalbs:
        build: noalbs-docker/
        image: digitalsparky/noalbs:latest
        restart: unless-stopped
        volumes:
            - ./noalbs:/opt/noalbs
```

- Download this archive and extract to a folder of your choice
- Open Docker Desktop then open the terminal (bottom right corner of the screen to bring it up)
- Type cd then drag and drop the folder you just extracted into the terminal, this should auto-complete the path, if it doesn't you may need to find the path manually.
- Copy noalbs/.env.example to noalbs/.env and configure
  - Set the TWITCH_BOT_USERNAME to either your twitch username, or your bot username
  - Generate an oauth key for that user (login to twitch as that user, then visit:  https://twitchapps.com/tmi/)
  - Set the value of TWITCH_BOT_OAUTH to the oauth key generated
- Copy noalbs/config.json.example to noalbs/config.json and configure
  - Full configuration directives are available here: https://github.com/Edward-Wu/srt-live-server/wiki/Directives
  - Update the lines with comments and REMOVE THE COMMENTS (or it will not start!)

- Run: docker compose up -d


### Security

%YOURKEY% is a key placeholder, generate a random string and copy it in to each of the locations you find %YOURKEY% in this guide.

On your mobile app AND in noalbs/config.json set publisher:

```
publish/live/%YOURKEY%
```

change your OBS media source to:

```
srt://127.0.0.1:7654?streamid=play/live/%YOURKEY%
```

Set your mobile SRT stream to:

```
srt://%YOURIP%:7654?streamid=publish/live/%YOURKEY%
```

(where %YOURIP% is your IP address)

Firewall:

- Forward port 7654 protocol UDP to your local system via your router/firewall


# Like my stuff?

Would you like to buy me a coffee or send me a tip?
While it's not expected, I would really appreciate it.

[![Paypal](https://www.paypalobjects.com/webstatic/mktg/Logo/pp-logo-100px.png)](https://paypal.me/MattSpurrier) <a href="https://www.buymeacoffee.com/digitalsparky" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/white_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
