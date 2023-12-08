# 检查操作系统类型
ifeq ($(OS),Windows_NT)
  IGNORE_OUTPUT := >NUL 2>&1
	RM := powershell.exe -Command "Remove-Item -Force -Recurse"
else
  IGNORE_OUTPUT := >/dev/null 2>&1
	RM := rm -rf
endif

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

EXE_NAME = video-thumb
VERSION := 1.0.0
EXTENSION ?=

ifeq ($(GOOS),windows)
	EXTENSION = .exe
endif

# 输出文件名
OUTPUT_NAME := $(EXE_NAME)-$(VERSION)-$(GOARCH)_$(GOOS)$(EXTENSION)

# 编译目标文件夹
BUILD_DIR := build

# 编译命令
BUILD_CMD_RELEASE := go build -a -v -trimpath -ldflags="-s -w" -tags="release" -o $(BUILD_DIR)/$(OUTPUT_NAME) github.com/kmou424/go-video-thumb/cmd
BUILD_CMD_DEBUG := go build -a -v -trimpath -ldflags="-s -w" -o $(BUILD_DIR)/$(OUTPUT_NAME) github.com/kmou424/go-video-thumb/cmd

.PHONY: all build clean

all: clean build

build:
	@echo "Building $(OUTPUT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(BUILD_CMD_RELEASE)
	@if [ -n "$$UPX_ENABLED" ]; then \
		echo "Compressing executable using UPX..."; \
		upx -9 $(BUILD_DIR)/$(OUTPUT_NAME) $(IGNORE_OUTPUT); \
	fi
	@echo "Build completed."

debug:
	@echo "Building debugging $(OUTPUT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(BUILD_CMD_RELEASE)
	@if [ -n "$$UPX_ENABLED" ]; then \
		echo "Compressing executable using UPX..."; \
		upx -9 $(BUILD_DIR)/$(OUTPUT_NAME) $(IGNORE_OUTPUT); \
	fi
	@echo "Build completed."

clean:
	@echo "Cleaning..."
	@$(RM) $(BUILD_DIR)
	@echo "Clean completed."
