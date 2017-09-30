#!/bin/bash

nohup gotty -w tmux new-session -s $1 ssh manager@139.59.88.117 &

