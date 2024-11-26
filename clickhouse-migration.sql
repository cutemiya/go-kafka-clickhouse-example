CREATE TABLE offer_queue
(
    pnr       String,
    arrival   String,
    departure String,
    price     Float32,
    tripId    Int
)
    ENGINE = Kafka
        SETTINGS kafka_broker_list = 'kafka:9092',
            kafka_topic_list = 'insert',
            kafka_group_name = 'mutation-click',
            kafka_format = 'JSONEachRow',
            kafka_row_delimiter = '\n';

CREATE TABLE offer
(
    pnr       String,
    arrival   String,
    departure String,
    price     Float32,
    tripId    Int
)
    ENGINE = MergeTree
        ORDER BY (tripId, pnr);

CREATE MATERIALIZED VIEW queue_to_offer
    TO offer AS
SELECT pnr,
       arrival,
       departure,
       price,
       tripId
FROM offer_queue;
