CQL="CREATE KEYSPACE doto WITH REPLICATION = { 'class': 'NetworkTopologyStrategy', 'datacenter1': 1 };"

until echo $CQL | cqlsh; do
  echo "cqlsh: Cassandra is unavailable to initialize - will retry later"
  sleep 2
done &

exec /docker-entrypoint.sh "$@"