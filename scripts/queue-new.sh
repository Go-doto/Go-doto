#!/bin/bash

if [ "$1" = "" ]
then
  echo "Pleaser set topic"
  exit
fi

docker exec kafka_doto /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic $1 --from-beginning
