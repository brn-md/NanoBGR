# NanoBGR: High-Resolution Background Remover SaaS

NanoBGR is a production-ready, microservices-based MVP for high-resolution background removal. It decouples high-performance AI inference from image composition to preserve 100% of the original image quality.

## 🚀 Key Features
- **High-Res Preservation**: Uses a proxy-masking technique (Alpha Matting) to process high-resolution images without memory exhaustion.
- **Asynchronous Architecture**: Go-based API handles traffic while Python-based Workers deal with heavy AI processing.
- **S3-Compatible Storage**: Integrated with MinIO for scalable asset management.
- **Microservices Orchestration**: Fully Dockerized for easy local development and cloud deployment.

## 🏗️ Architecture Stack
- **Orchestrator (API)**: Golang (Fiber)
- **AI Worker**: Python 3.10 (OpenCV, rembg/U2-Net, PIL)
- **Message Broker**: Redis (Task Queuing & Status Cache)
- **Object Storage**: MinIO (S3-compatible)
- **Infrastructure**: Docker Compose

## 🛠️ High-Res Technique (The "Secret Sauce")
Traditional AI background removers often downscale the image to fit into VRAM, losing quality. NanoBGR solves this by:
1. **Fetching** the original high-resolution image.
2. **Creating** a low-resolution proxy (1024px) for fast AI inference.
3. **Generating** a mask (Alpha Matte) on the proxy.
4. **Upscaling** the mask back to original dimensions using Bilinear/Linear filters.
5. **Merging** the upscaled mask with the *untouched* original RGB pixels.

## 🚦 How to Run

### Prerequisites
- Docker & Docker Compose
- (Optional) WSL2 for Windows users for better performance

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/brnmd96/NanoBGR.git
   cd NanoBGR
   ```
2. Start the infrastructure:
   ```bash
   docker-compose up --build
   ```

### 📡 API Endpoints
- **POST `/upload`**: Upload a multipart/form image. Returns a `UUID`.
- **GET `/status/:id`**: Check processing progress. Returns URLs for original and processed images.

## ⚙️ Environment Variables
Check `.env.example` for details, but by default:
- `API_PORT`: 3000
- `MINIO_ROOT_USER`: admin
- `MINIO_ROOT_PASSWORD`: supersecret123
