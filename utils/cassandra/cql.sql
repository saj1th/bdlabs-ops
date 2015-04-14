CREATE KEYSPACE forecaster
           WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

USE forecaster;

CREATE TABLE predictions (
 sku varchar,
 date timestamp,
 volume double,
 sale  double,
 PRIMARY KEY (sku, date));


CREATE TABLE forecast(
                      sku varchar,
                      date timestamp,
                      volume double,
                      sale  double,
                      PRIMARY KEY (sku, date));

INSERT INTO predictions (sku, date, volume, sale)  VALUES ('62c36092', '2007-01-02', 1000, 20000);

CREATE TABLE admin_users (
 id int,
 email varchar,
 name varchar,
 password varchar,
 status  varchar,
 PRIMARY KEY (email));

CREATE INDEX id ON admin_users (id);




CREATE TABLE forecaster_jobs (
 userId varchar,
 id varchar,
 filepath varchar,
 stage varchar,
 PRIMARY KEY (id));

CREATE INDEX userId ON forecaster_jobs (userId);

