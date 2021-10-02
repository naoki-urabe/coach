#!/bin/bash
USER_HOST=`echo "$USER@$HOST"`
SYNC_PATH=`echo "$USER_HOST:/home/$USER/app/coach"`
SSH_PORT="ssh -p $HOST_PORT"
rsync -e "$SSH_PORT" -acvz  --delete ./ $SYNC_PATH
ssh $USER_HOST -p $HOST_PORT "cd app/coach;/usr/bin/docker-compose down -v; docker image prune -f;docker pull ghcr.io/naoki-urabe/coach/api:prod;docker pull ghcr.io/naoki-urabe/coach/front:prod;docker pull ghcr.io/naoki-urabe/coach/task:prod;/usr/bin/docker-compose -f docker-compose.prod.yml up -d"