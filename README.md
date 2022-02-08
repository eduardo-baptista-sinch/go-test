MySQL commands:

```
dc exec mysql bash
mysql -uroot -p fullcycle 
create table courses(id varchar(255), name varchar(255), description varchar(255), status varchar(255));
```

Kafka commands:

```
dc exec kafka bash
kafka-topics --bootstrap-server=localhost:9092 --topic=courses --create --partitions=3 --replication-factor=1

# Send messages to topic:

kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses
```