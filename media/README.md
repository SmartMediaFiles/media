# Media Type, File Type, and File Extension Architecture

## Overview

This project implements a flexible and extensible system for managing media types, file types, and file extensions. The architecture is designed to provide a clear hierarchy and relationship between these entities, allowing for easy categorization and manipulation of various media files.

## Core Components

### 1. Media Type

Media Type represents the highest level of categorization for media files. It is defined as a custom type in Go:

The `MediaType` is used to group related file types under a common category, such as "image", "video", "raw", etc.

### 2. File Type

File Type represents a specific format within a media type. It is defined as:

The `FileType` is used to further categorize media files within a specific media type, such as "JPEG", "PNG", "AVC", etc.


### 3. File Extension

File Extension represents the actual file extension associated with a file type. It is defined as:

The `FileExtension` is used to identify the file type and extension of a media file, such as ".jpg", ".png", ".mp4", etc.

## Relationship Between Components

The relationship between these components is hierarchical:

- A `MediaType` can have multiple `FileTypes`.
- A `FileType` can have multiple `FileExtensions`.
