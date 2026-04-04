import redis
import json
from app.config import REDIS_URL

class RedisClient:
    def __init__(self):
        self.client = redis.from_url(REDIS_URL, decode_responses=True)
        self.client.ping()

    def get_task(self, queue_name, timeout=5):
        task = self.client.brpop(queue_name, timeout=timeout)
        if task:
            _, payload = task
            return json.loads(payload)
        return None

    def set_status(self, task_id, status):
        key = f"status:{task_id}"
        self.client.setex(key, 86400, status)
