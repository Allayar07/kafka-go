# Step-1
## clone the repo: 
```
git clone git@github.com:Allayar07/kafka-go.git
```

* And run ```go mod tidy```

# Step-2
## Up the docker containers
```
docker compose up
```
# Step-3
## Create topics in the kafka:
* First topic
```
docker exec kafka_local bash -c '/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic topic_1 --bootstrap-server localhost:9092
```
* Second topic:
```
docker exec kafka_local bash -c '/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic topic_2 --bootstrap-server localhost:9092
```

# Step-4
## run the app ```go run main.go```
# Step-5
## Write something to topic_1:
```
docker exec kafka_local bash
```
And then : ```cd /opt/bitnami/kafka/bin```

Then write message to topic_1:
```
kafka-console-producer.sh --topic topic_1 --bootstrap-server localhost:9092
```
write something and press enter

# Step-5
## See messages from topic_2:
```
docker exec kafka_local bash
```
And then : ```cd /opt/bitnami/kafka/bin```

Then see messages in the topic_2:
```
kafka-console-consumer.sh --topic topic_2 --from-beginning --bootstrap-server localhost:9092
```

* Or you can get it from here: ```go run consumer/consume.go```

Documentation of library : https://github.com/segmentio/kafka-go