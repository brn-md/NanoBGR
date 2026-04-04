import boto3
from botocore.client import Config
from app.config import MINIO_ENDPOINT, MINIO_ACCESS_KEY, MINIO_SECRET_KEY, MINIO_USE_SSL, MINIO_BUCKET

class MinioClient:
    def __init__(self):
        protocol = "https" if MINIO_USE_SSL else "http"
        self.s3 = boto3.client(
            's3',
            endpoint_url=f"{protocol}://{MINIO_ENDPOINT}",
            aws_access_key_id=MINIO_ACCESS_KEY,
            aws_secret_access_key=MINIO_SECRET_KEY,
            config=Config(signature_version='s3v4'),
            region_name='us-east-1'
        )
        self.bucket = MINIO_BUCKET

    def download_image(self, object_name):
        obj = self.s3.get_object(Bucket=self.bucket, Key=object_name)
        return obj['Body'].read()

    def upload_image(self, object_name, image_bytes, content_type='image/png'):
        self.s3.put_object(
            Bucket=self.bucket,
            Key=object_name,
            Body=image_bytes,
            ContentType=content_type
        )
