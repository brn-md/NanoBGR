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

### Fixed
- **Docker API**: Updated `api/Dockerfile` from Go 1.23 to 1.24-alpine to support modern dependencies (minio-go/v7 v7.0.100+).
- **Go Compilation**: Removed unused `mime/multipart` import in `minio.go` to comply with strict Go compiler rules.
- **Python Naming**: Renamed `app/queue.py` to `app/task_queue.py` to avoid shadowing the Python standard library `queue` module.
- **Docker Infrastructure**: Set `PYTHONPATH=/app` and removed unnecessary volume mounting for the Go API container to prevent binary overwrites.
- **Git**: Added robust `.gitignore` file and documentation on CRLF/LF line ending handlings on Windows.
- **Docker Worker**: Fixed `libgl1-mesa-glx` (obsolete) with `libgl1` in Python base image.
