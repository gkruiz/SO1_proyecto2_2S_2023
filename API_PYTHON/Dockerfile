FROM python:3.7-alpine

WORKDIR /

COPY api.py /
COPY ips.sh /

    
RUN pip3 install -U flask-cors && \
pip install mysql-connector-python && \
   pip install redis  && \
   pip install redis[hiredis] && \
   pip install waitress

ENTRYPOINT ["python3", "-u", "api.py"]
