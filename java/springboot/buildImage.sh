#!/bin/bash 

mvn clean package

_NAME=$(xml sel -N ns="http://maven.apache.org/POM/4.0.0" -t -v "//ns:project/ns:artifactId/text()" ./pom.xml)
_VERSION=$(xml sel -N ns="http://maven.apache.org/POM/4.0.0" -t -v "//ns:project/ns:version/text()" ./pom.xml)
echo $_NAME $_VERSION
docker build -t ${_NAME}:${_VERSION} -f Dockerfile .
