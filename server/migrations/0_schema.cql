CREATE KEYSPACE IF NOT EXISTS k8s_playground
  WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE IF NOT EXISTS k8s_playground.records (
  id uuid PRIMARY KEY,
  message text,
  created_date timestamp,
) WITH comment = 'The records pushed from Kafka.';
