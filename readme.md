```json
"streamServer": {
  "type": "SrtLiveServer",
  "statsUrl": "http://localhost:8080/stats.json",
  "publisher": "publish/live/srt"
},
```

```json
"software": {
  "type": "Obs", // NOALBS supports OBS WebSocket v4 and v5. To still use v4 use type ObsOld.
  "host": "host.docker.internal", // Don't change this
  "password": "example", // Password to the OBS Websockets.
  "port": 4455, // Port to the OBS Websockets.
  "collections": {
    "twitch": {
      "profile": "twitch_profile",
      "collection": "twitch_scenes"
    },
    "kick": {
      "profile": "kick_profile",
      "collection": "kick_scenes"
    }
  }
},
```

This runs SRTLA, SRT, Stats server, and NOALBS all in one using the config supplied.
It will send SRT back to OBS using the host.docker.internal:4455 srt address.

OBS Config

Create a MEDIA source
Set the source to: srt://0.0.0.0:5001?mode=listener

Update config in noalbs/config.json and .env to set the correct SRT external port, Twitch user config and noalbs srt config.
As per the above, please keep the streamserver and software settings similar, specifically the stats URL and the host.docker.internal on the software.
You acn change everything else as needed.