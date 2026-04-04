FROM python:3.10-slim

WORKDIR /app

# Install system dependencies for OpenCV and rembg
RUN apt-get update && apt-get install -y --no-install-recommends \
    libglib2.0-0 \
    libsm6 \
    libxext6 \
    libxrender-dev \
    libgl1 \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt .

# Upgrade pip and install requirements
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt

# Cache U2Net model during build
ENV U2NET_HOME=/home/.u2net
RUN mkdir -p $U2NET_HOME && \
    python -c "import rembg; rembg.new_session('u2net')"

COPY . .

CMD ["python", "-u", "app/main.py"]
