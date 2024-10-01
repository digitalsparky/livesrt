#!/bin/bash

# /opt/livesrt/bin/srtla_rec 7654 127.0.0.1 5002 &
/opt/livesrt/bin/srt-live-transmit -st:yes -a -tm:1 -t 120 -statsout:/tmp/srtstats.json -pf:json "srt://:5002?mode=listener" "srt://host.docker.internal:5001?mode=caller"