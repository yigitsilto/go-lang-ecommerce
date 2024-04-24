#!/bin/bash

# .env dosyasını oluştur (eğer yoksa)
if [ ! -f .env ]; then
  if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
    cp env/.koreenv .env
  elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
    cp env/.peynircilerenv .env
  fi
 elif [ "$(git rev-parse --abbrev-ref HEAD)" = "proshop" ]; then
     cp env/.proenv .env
   fi
fi


if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
    docker stack rm gokore
    docker build -t gocommercekore --no-cache -f Dockerfile .

elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
    docker stack rm gopeynirciler
    docker build -t gocommercepeynirciler --no-cache -f Dockerfile .
elif [ "$(git rev-parse --abbrev-ref HEAD)" = "proshop" ]; then
    docker stack rm gopro
    docker build -t gocommercepro --no-cache -f Dockerfile .
fi

if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
  docker stack deploy -c docker-compose-kore.yaml gokore --resolve-image=always
elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
  docker stack deploy -c docker-compose-peynirciler.yaml gopeynirciler --resolve-image=always
elif [ "$(git rev-parse --abbrev-ref HEAD)" = "proshop" ]; then
  docker stack deploy -c docker-compose-pro.yaml gopro --resolve-image=always
fi
