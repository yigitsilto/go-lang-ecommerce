#!/bin/bash

# .env dosyasını oluştur (eğer yoksa)
if [ ! -f .env ]; then
  if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
    cp env/.koreenv .env
  elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
    cp env/.peynircilerenv .env
  fi
fi


if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
  if [ -z "$(docker images -q gocommercekore:latest 2>/dev/null)" ]; then
    docker build -t gocommercekore --no-cache -f Dockerfile .
  fi
elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
  if [ -z "$(docker images -q gocommercepeynirciler:latest 2>/dev/null)" ]; then
    docker build -t gocommercepeynirciler --no-cache -f Dockerfile .
  fi
fi

if [ "$(git rev-parse --abbrev-ref HEAD)" = "main" ]; then
  docker stack deploy -c docker-compose-kore.yaml gokore --resolve-image=always
elif [ "$(git rev-parse --abbrev-ref HEAD)" = "peynirciler" ]; then
  docker stack deploy -c docker-compose-peynirciler.yaml gopeynirciler --resolve-image=always
fi
