import boto3
import json
from datetime import datetime

# Cria uma sessão com o perfil da AWS (opcional, se já configurado)
ec2_client = boto3.client('ec2')

instances = ec2_client.describe_instances()
volumes = ec2_client.describe_volumes()

def datetime_converter(o):
    if isinstance(o, datetime):
        return o.isoformat()  # Converte o datetime para formato ISO (string)

# Chama describe_instances() para obter as informações
instances = ec2_client.describe_instances()

# Salva a resposta em um arquivo JSON, usando a função de conversão
with open('instances.json', 'w') as json_file:
    json.dump(instances, json_file, default=datetime_converter, indent=4)

with open('volumes.json', 'w') as json_file:
    json.dump(volumes, json_file, default=datetime_converter, indent=4)


print ("Dados salvos em instances.json")
