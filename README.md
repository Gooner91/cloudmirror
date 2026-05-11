# cloudmirror

> **Work in progress.** Core config management is done; Google Drive sync is under active development.

A lightweight Go CLI that periodically backs up local files and folders to cloud storage. You define source-to-destination mappings once, and cloudmirror handles the rest. Google Drive is the first supported backend, with more planned.

---

## Table of contents

- [Overview](#overview)
- [Project status](#project-status)
- [Installation](#installation)
- [Google Drive setup](#google-drive-setup)
- [Usage](#usage)
  - [Managing backup mappings](#managing-backup-mappings)
- [Config file](#config-file)
- [Development](#development)
- [Roadmap](#roadmap)

---

## Overview

cloudmirror works around the concept of **mappings** — a source glob pattern pointing to files or directories on your machine, paired with a destination path in your cloud storage of choice.

```
/home/user/documents/**/*.pdf  →  gdrive:/Backups/PDFs
/home/user/projects/notes/     →  gdrive:/Backups/Notes
```

Once mappings are registered, cloudmirror will resolve the globs, walk the matching paths, and upload everything to the specified destination.

---

## Project status

| Area                        | Status        |
|-----------------------------|---------------|
| CLI scaffolding (Cobra)     | Done          |
| Config store (add/delete)   | Done          |
| Config validation           | Done          |
| Google Drive OAuth2 auth    | Done          |
| Google Drive file listing   | Done (WIP)    |
| Google Drive file upload    | In progress   |
| Sync command                | Not started   |
| Periodic / scheduled sync   | Not started   |
| Additional cloud backends   | Not started   |

---

## Installation

Requires Go 1.24+.

```bash
git clone https://github.com/Gooner91/cloudmirror.git
cd cloudmirror
go build -o cloudmirror .
```

Optionally move the binary somewhere on your `$PATH`:

```bash
mv cloudmirror /usr/local/bin/
```

---

## Google Drive setup

cloudmirror uses OAuth2 to authenticate with Google Drive. You need a `credentials.json` file from a Google Cloud project before running any Drive-related commands.

1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Create a project (or select an existing one).
3. Enable the **Google Drive API** for that project.
4. Under **APIs & Services → Credentials**, create an **OAuth 2.0 Client ID** (Desktop app type).
5. Download the credentials and save the file as `credentials.json` in the directory where you run `cloudmirror`.

On first use, cloudmirror will open an authorization URL in your terminal. Visit the URL, grant access, paste the authorization code back into the terminal, and a `token.json` will be saved locally for future runs.

> `credentials.json` and `token.json` are excluded from version control via `.gitignore`. Never commit these files.

---

## Usage

```
cloudmirror [command]

Commands:
  config      Manage backup mappings
  help        Help about any command
```

### Managing backup mappings

**Add a mapping**

Register a source glob and a destination path in the config:

```bash
cloudmirror config add --srcGlob "/home/user/docs/*.pdf" --dest "Backups/PDFs"
```

| Flag        | Required | Description                                              |
|-------------|----------|----------------------------------------------------------|
| `--srcGlob` | Yes      | Glob pattern matching the local files/dirs to back up    |
| `--dest`    | Yes      | Destination folder path in the target cloud storage      |

**Delete a mapping**

Remove a previously registered mapping (both `--srcGlob` and `--dest` must match exactly):

```bash
cloudmirror config delete --srcGlob "/home/user/docs/*.pdf" --dest "Backups/PDFs"
```

---

## Config file

Mappings are persisted as JSON at:

| Platform        | Path                                          |
|-----------------|-----------------------------------------------|
| Linux / macOS   | `$XDG_CONFIG_HOME/cloudmirror/config.json`    |
| Fallback        | `/etc/cloudmirror/config.json`                |

On most Linux systems `$XDG_CONFIG_HOME` defaults to `~/.config`, so the effective path is `~/.config/cloudmirror/config.json`.

Example file after adding two mappings:

```json
[
  {
    "SrcGlob": "/home/user/documents/*.pdf",
    "Dest": "Backups/PDFs"
  },
  {
    "SrcGlob": "/home/user/projects/notes",
    "Dest": "Backups/Notes"
  }
]
```

The config file is created automatically on first `config add`. Its permissions are set to `0600` (owner read/write only).

**Validation rules**

- `SrcGlob` must be non-empty and a valid Go filepath glob pattern.
- `Dest` must be non-empty.
- Duplicate mappings (same `SrcGlob` + `Dest` pair) are rejected.

---

## Development

```bash
# Run tests
go test ./...

# Build with live-reload (requires air)
air
```

**Project layout**

```
cloudmirror/
├── cmd/                        # Cobra CLI commands
│   ├── root.go                 # Root command
│   ├── config.go               # `config` subcommand group
│   ├── add.go                  # `config add`
│   └── delete.go               # `config delete`
├── internal/
│   ├── config/
│   │   ├── types.go            # Config and ConfigList types
│   │   ├── store.go            # Save, Delete, validate, persist
│   │   └── store_test.go       # Unit tests for config store
│   └── google_drive/
│       ├── auth.go             # OAuth2 client setup
│       ├── service.go          # Drive API service constructor
│       └── client.go           # Drive API operations (ListFiles, ...)
└── main.go
```

---

## Roadmap

- [ ] `cloudmirror sync` — resolve mappings and upload matching files to Google Drive
- [ ] Progress output and dry-run mode
- [ ] Periodic sync via cron or a built-in scheduler
- [ ] `config list` command to display current mappings
- [ ] Support for additional backends (S3, Dropbox, etc.)
- [ ] Incremental sync (skip unchanged files)
