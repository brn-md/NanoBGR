import traceback
import time
import sys
from app.config import QUEUE_NAME
from app.task_queue import RedisClient
from app.storage import MinioClient
from app.processor import process_image

def main():
    print("Initializing Worker...", flush=True)
    try:
        redis_client = RedisClient()
        minio_client = MinioClient()
        print("Connected to Redis and Minio", flush=True)
    except Exception as e:
        print(f"Failed to initialize clients: {e}", flush=True)
        sys.exit(1)

    print("Listening for tasks...", flush=True)
    while True:
        try:
            task = redis_client.get_task(QUEUE_NAME)
            if not task:
                continue
            
            task_id = task.get("id")
            raw_image_path = task.get("raw_image")
            print(f"Processing task {task_id}: {raw_image_path}", flush=True)
            
            try:
                # Update status
                redis_client.set_status(task_id, "processing")
                
                # Download
                image_bytes = minio_client.download_image(raw_image_path)
                
                # Process
                processed_bytes = process_image(image_bytes)
                
                # Upload
                result_path = f"processed/{task_id}.png"
                minio_client.upload_image(result_path, processed_bytes)
                
                # Update status to done
                redis_client.set_status(task_id, "done")
                print(f"Task {task_id} completed successfully.", flush=True)
                
            except Exception as e:
                print(f"Error processing {task_id}: {traceback.format_exc()}", sys.stderr, flush=True)
                redis_client.set_status(task_id, "error")

        except Exception as e:
            print(f"Queue error: {e}", sys.stderr, flush=True)
            time.sleep(1)

if __name__ == "__main__":
    main()
