#!/bin/bash
USER_HOST=`echo "$USER@$HOST"`
SYNC_PATH=`echo "$USER_HOST:/home/$USER/app/coach"`
SSH_PORT="ssh -p $HOST_PORT"
rsync -e "$SSH_PORT" -acvz  --delete ./ $SYNC_PATH
ssh $USER_HOST -p $HOST_PORT "cd app/coach;/usr/bin/docker-compose down -v; docker pull ghcr.io/naoki-urabe/coach/api:prod;docker pull ghcr.io/naoki-urabe/coach/front:prod;docker pull ghcr.io/naoki-urabe/coach/task:prod;docker system prune -f;/usr/bin/docker-compose -f prod.yml up -d"
