## go-redpanda

Demo project that orchestrates [redpanda/connect](https://github.com/redpanda-data/connect) data stream
from a [spicedb watch](https://docs.redpanda.com/redpanda-connect/components/inputs/spicedb_watch/) input to a [kafka](https://docs.redpanda.com/redpanda-connect/components/outputs/kafka/) output.

---

### environment

- Docker
- Make
- Golang

---

### setting up kafka

- compose up

```bash
docker compose up -d
```

- create a session

```bash
docker exec --workdir /opt/kafka/bin -it broker bash
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
> --bootstrap-server localhost:9092 \
> --from-beginning
```

---

### setting up redpanda connect

```bash
make rpk-connect-run
```
