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
- **Docker API**: Updated `api/Dockerfile` from Go 1.22 to 1.23-alpine for better compatibility with modern dependencies.
- **Docker Worker**: Replaced obsolete `libgl1-mesa-glx` with `libgl1` in the Python-slim base image to resolve build errors.
- **Dependencies**: Downgraded `minio-go/v7` to `v7.0.70` in `go.mod` to ensure compatibility with Go 1.23 environment.
- **Git**: Added robust `.gitignore` file to prevent accidental commits of binary and database data.
