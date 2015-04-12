SPARK_VERSION='spark-1.2.0'
HADOOP_VERSION='hadoop2.4'
SPARK_BINARY="${SPARK_VERSION}-bin-${HADOOP_VERSION}"
SPARK_TAR="${SPARK_BINARY}.tgz"
FORECASTER_VERSION='0.2'
FORECASTER_JAR="forecaster-${FORECASTER_VERSION}.jar"
RUN_DIR='/home/jclouds/forecaster/bin'

[ -d $RUN_DIR ] || mkdir -p $RUN_DIR
cd $RUN_DIR
hdfs dfs -copyToLocal /opt/"$SPARK_TAR"
hdfs dfs -copyToLocal /projects/forecaster/bin/$FORECASTER_JAR
tar zxvf ${SPARK_TAR}
rm -f ${SPARK_TAR}


$SPARK_BINARY/bin/spark-submit \
 	--class org.bitbucket.saj1th.forecaster.Forecast \
 	--master mesos://x.x.x.x:5050 \
  	$RUN_DIR/$FORECASTER_JAR \
  	--master mesos://x.x.x.x:5050 \
  	--sparkexecutor hdfs://x.x.x.x/opt/$SPARK_TAR \
  	--cassandrahost x.x.x.x  \
  	--modelspath  hdfs://x.x.x.x:8020/projects/forecaster/models \
  	hdfs://x.x.x.x:8020/projects/forecaster/data/traindata.csv

rm -rf $RUN_DIR