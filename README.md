## redpanda-playground

Demo project that orchestrates [redpanda/connect](https://github.com/redpanda-data/connect) data streams, with [redpanda/console](https://github.com/redpanda-data/console) as the UI.

**examples**
- from a [spicedb watch](https://docs.redpanda.com/redpanda-connect/components/inputs/spicedb_watch/) input to a [kafka](https://docs.redpanda.com/redpanda-connect/components/outputs/kafka/) output.

---

### environment

- Docker
- Make
- Golang

---

### setup

- compose up

```bash
docker compose up -d
```

---

### spicedb watch to kafka

#### setting up kafka

- create a session

```bash
docker exec --workdir /opt/kafka/bin -it connect-output bash
```

- create a topic

```bash
./kafka-topics.sh --create --topic spicypanda \
  --bootstrap-server localhost:9092 \
  --partitions 1 \
  --replication-factor 1
```

- list topics

```bash
./kafka-topics.sh --list --bootstrap-server localhost:9092
```

- read messages from a topic

```bash
./kafka-console-consumer.sh --topic spicypanda \
 --bootstrap-server localhost:9092 \
 --from-beginning
```

---

###  run redpanda connect

```bash
./cli.sh connect_run
```

---

### setting up spicedb 

- write schema

```bash
./cli.sh spicedb write-schema
```

- setup data

```bash
./cli.sh spicedb setup-data
