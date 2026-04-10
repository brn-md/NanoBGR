# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-04-04

### Added
- **Infrastructure**: Initial `docker-compose.yml` orchestrating Go, Python, Redis, and MinIO.
- **Go API**: High-performance image upload engine with Fiber.
- **Python Worker**: Background image processor utilizing `rembg` (U2-Net) and OpenCV.
- **Queue/State**: Async task management using Redis Lists and Keys.
- **Storage**: S3-compatible storage integration with MinIO.
- **Architecture**: Implemented High-Res Background Removal proxy-mask technique.

## [0.2.0] - 2026-04-06

### Added
- **Frontend**: Full React/Vite dashboard with drag-and-drop support and automatic polling.
- **Docker**: Added `frontend` service to `docker-compose.yml` using Nginx for production serving.
- **API**: Added CORS middleware to Go Fiber to allow cross-origin requests from the new frontend.
- **Docs**: Created a comprehensive line-by-line project walkthrough for developers.

### Fixed
- **IA Processing**: Enabled `alpha_matting` in the Python worker to significantly improve edge detection (fingers, hair) and prevent accidental removal of fine details.
- **Docker Infrastructure**: Fixed frontend indentation in `docker-compose.yml` to prevent YAML parsing errors.
- **Structure**: Moved `web` folder to the project root for better service separation.

## [0.1.0] - 2026-04-04
