FROM docker.elastic.co/beats/filebeat:7.4.2
COPY filebeat.yml /usr/share/filebeat/filebeat.yml
COPY data.log /usr/share/filebeat/logs/data.log
USER root
RUN chown root:filebeat /usr/share/filebeat/filebeat.yml
USER filebeat