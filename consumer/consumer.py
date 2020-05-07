import pika
import time

connection = pika.BlockingConnection(
    pika.ConnectionParameters(host='localhost')
)

channel = connection.channel()

channel.queue_declare(queue='cola',durable=True)
print('[*] Esperando por mensajes')

def callback(ch,method,properties,body):
    print("[x] received %r" % body)
    print("[x] Done")
    ch.basic_ack(delivery_tag=delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue="cola", on_message_callback=callback)

channel.start_consuming()