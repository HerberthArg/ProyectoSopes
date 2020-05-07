FROM python:3.8-alpine
RUN pip install pika

ENV PYTHONUNBUFFERED=1
WORKDIR /usr/src/app
COPY consumer.py .


CMD ["python", "/usr/src/app/consumer.py"]