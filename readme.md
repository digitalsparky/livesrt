

This runs SRTLA, SRT, Stats server, and NOALBS all in one using the config supplied.
It will send SRT back to OBS using the host.docker.internal:4455 srt address.

OBS Config

Create a MEDIA source
Unselect "Local file"
Set the source to: srt://127.0.0.1:7654?streamid=play/live/feed1
Set the Image Format to mpegts
Select "Use hardware encoding where available"

Update config in noalbs/config.json and .env to set the correct SRT external port, Twitch user config and noalbs srt config.
The example configuration is already setup for use with the provided NOALBS and SRT server addresses, so you can ignore the stream config:

```json
"streamServers": [
      {
        "streamServer": {
          "type": "SrtLiveServer",
          "statsUrl": "http://srt:8181/stats",
          "publisher": "publish/live/feed1"
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

=== WINDOWS
In order to run on windows, install docker desktop and setup your system with WSL2
Instructions available here: https://docs.docker.com/desktop/install/windows-install/

=== Setup

- Download this archive and extract to a folder of your choice
- Open Docker Desktop then open the terminal (bottom right corner of the screen to bring it up)
- Type cd then drag and drop the folder you just extracted into the terminal, this should auto-complete the path, if it doesn't you may need to find the path manually.
- Copy noalbs/.env.example to noalbs/.env and configure
-- Set the TWITCH_BOT_USERNAME to either your twitch username, or your bot username
-- Generate an oauth key for that user (login to twitch as that user, then visit:  https://twitchapps.com/tmi/)
-- Set the value of TWITCH_BOT_OAUTH to the oauth key generated
- Copy noalbs/config.json.example to noalbs/config.json and configure
-- Full configuration directives are available here: https://github.com/Edward-Wu/srt-live-server/wiki/Directives
-- Update the lines with comments and REMOVE THE COMMENTS (or it will not start!)

- Run: docker compose up -d