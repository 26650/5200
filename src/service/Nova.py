import boto3
import json
import os
from datetime import datetime
from model import Chat
from dotenv import load_dotenv

load_dotenv(os.path.join(os.path.dirname(__file__), '..', '.env'))

# AWS credentials from environment variables
aws_access_key_id = os.getenv('AWS_ACCESS_KEY_ID')
aws_secret_access_key = os.getenv('AWS_SECRET_ACCESS_KEY')
region_name = 'us-west-2'  # Replace with your AWS region if necessary

# Create a boto3 client for Bedrock Runtime
client = boto3.client(
    "bedrock-runtime",
    region_name=region_name,
    aws_access_key_id=aws_access_key_id,
    aws_secret_access_key=aws_secret_access_key
)

async def nova_number(input_message : Chat.Message) -> str:

    LITE_MODEL_ID = "us.amazon.nova-pro-v1:0"

    system_list = [
        {
            "text": "int 형식만 답변 가능, 입력된 단어의 글자 수 출력, 답변 예시: 3"
        }
    ]

    inf_params = {
        "max_new_tokens": 20,
        "top_p": 0,
        "top_k": 20,
        "temperature": 0
    }

    input_text = input_message.content
    message_list = [{"role": "user", "content": [{"text": input_text}]}]

    request_body = {
        "schemaVersion": "messages-v1",
        "messages": message_list,
        "system": system_list,
        "inferenceConfig": inf_params,
    }

    # Invoke the model without response stream
    response = client.invoke_model(
        modelId=LITE_MODEL_ID,
        body=json.dumps(request_body)
    )

    # Process the response
    response_body = json.loads(response["body"].read())
    print(response_body)
    result = ""
    for content_item in response_body.get("output", {}).get("message", {}).get("content", []):
        result += content_item.get("text", "")
    print(result)

    return result