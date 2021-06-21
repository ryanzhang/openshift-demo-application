FROM quay.io/rzhang/py-greeting:csg
MAINTAINER Ryan Zhang <rzhang@redhat.com>
ADD ./webapp /opt/webapp/
WORKDIR /opt/webapp
EXPOSE 5000
CMD ["python", "app.py"]

