CREATE KEYSPACE forecaster
           WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

USE forecaster;

CREATE TABLE predictions (
 sku varchar,
 date timestamp,
 volume double,
 sale  double,
 PRIMARY KEY (sku, date));


INSERT INTO predictions (sku, date, volume, sale)  VALUES ('62c36092', '2007-01-02', 1000, 20000);