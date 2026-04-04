# NanoBGR: High-Resolution Background Remover SaaS

[🇧🇷 Português](#português) | [🇺🇸 English](#english)

---

<a name="português"></a>
## 🇧🇷 Português

NanoBGR é um MVP pronto para produção de um SaaS de remoção de fundo de imagem. Ele utiliza uma arquitetura de microserviços para processar imagens de alta resolução sem perder a qualidade original.

### 🚀 Principais Funcionalidades
- **Preservação de Alta Resolução**: Técnica de Alpha Matting por Proxy para processar fotos pesadas sem estourar a memória RAM/VRAM.
- **Arquitetura Assíncrona**: API em Go para o tráfego e Worker em Python para o processamento pesado de IA.
- **Armazenamento S3**: Integrado com MinIO para gerenciamento escalável de arquivos.

### 🏗️ Como Rodar
1.  Inicie a infraestrutura:
    ```bash
    docker-compose up --build
    ```

### 🧪 Testando o Workflow
**Passo 1: Upload da imagem** (Windows Powershell):
```powershell
curl.exe -F "image=@C:\caminho\para\foto.jpg" http://localhost:3000/upload
```
**Passo 2: Consultar Status**:
```bash
curl http://localhost:3000/status/ID-DO-UPLOAD
```

---

<a name="english"></a>
## 🇺🇸 English

NanoBGR is a production-ready microservices-based SaaS for background removal. It decouples AI inference from image composition to preserve 100% of the original quality.

### 🚀 Key Features
- **High-Res Preservation**: Proxy-mask technique (Alpha Matting) to handle high-resolution files without memory exhaustion.
- **Asynchronous Flow**: Go-based API handles requests while Python Workers deal with heavy AI inference.
- **S3-Compatible Storage**: Built-in MinIO integration.

### 🏗️ How to Run
1.  Start the infrastructure:
    ```bash
    docker-compose up --build
    ```

### 🧪 Testing the Workflow
**Step 1: Upload an image** (Windows Powershell):
```powershell
curl.exe -F "image=@C:\path\to\photo.jpg" http://localhost:3000/upload
```
**Step 2: Check Processing Status**:
```bash
curl http://localhost:3000/status/UPLOAD-ID
```

## ⚙️ Environment Variables
Check `.env.example` for details. Default ports:
- API: 3000
- MinIO: 9001
