# NanoBGR — High-Resolution Background Removal Project

[🇧🇷 Português](#português) | [🇺🇸 English](#english)

---

<a name="português"></a>
## 🇧🇷 Português

NanoBGR é um **projeto de remoção de fundo de imagens** com foco em processamento de imagens em alta resolução sem perda de qualidade.

O projeto explora:
- Arquitetura de microserviços
- Processamento assíncrono
- Separação entre inferência de IA e composição final da imagem

### Funcionalidades
- Processamento em alta resolução usando abordagem de Alpha Matting por proxy
- API em Go responsável por upload, status e orquestração
- Worker em Python para processamento pesado de IA
- Armazenamento S3-compatível via MinIO

### Arquitetura (Visão Geral)
- **API (Go)**: Recebe imagens, cria jobs e expõe status
- **Worker (Python)**: Executa inferência e gera máscara
- **Composição**: Aplica a máscara na imagem original em alta resolução
- **Storage**: MinIO para entrada e saída de arquivos

### Como Rodar
docker-compose up --build

### Testando o Workflow

**Upload da imagem (Windows PowerShell):**
curl.exe -F "image=@C:\caminho\para\foto.jpg" http://localhost:3000/upload

**Consultar status:**
curl http://localhost:3000/status/ID-DO-UPLOAD

---

<a name="english"></a>
## 🇺🇸 English

NanoBGR is a **background removal project** focused on handling high-resolution images without quality loss.

The project explores:
- Microservice architecture
- Asynchronous processing
- Decoupling AI inference from final image composition

### Features
- High-resolution processing using a proxy-based alpha matting approach
- Go API for uploads, job orchestration, and status tracking
- Python worker for AI-heavy processing
- S3-compatible storage via MinIO

### Architecture Overview
- **API (Go)**: Handles uploads, job creation, and status
- **Worker (Python)**: Runs inference and generates masks
- **Composition**: Applies masks to the original high-resolution image
- **Storage**: MinIO for input/output files

### How to Run
docker-compose up --build

### Testing the Workflow

**Upload an image (Windows PowerShell):**
curl.exe -F "image=@C:\path\to\photo.jpg" http://localhost:3000/upload

**Check status:**
curl http://localhost:3000/status/UPLOAD-ID

### Environment Variables
See `.env.example`.

Default ports:
- API: 3000
- MinIO: 9001
