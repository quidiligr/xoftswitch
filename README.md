# XoftSwitch

**XoftSwitch** is an open-source VoIP switching & management layer that sits on top of **Asterisk + FreePBX** to make **WebRTC extensions** simple, repeatable, and enterprise-ready. It’s the companion server to the **XoftPhone** clients, enabling single sign-on, near-zero-config onboarding, and optional AI-assisted administration.

> Free and extensible • WebRTC-first • AI-assisted ops • PMS integration • Booking & Drive-thru modules

---

## Key Features

- **WebRTC-first extensions**  
  Opinionated defaults for DTLS/SRTP, ICE, transport-wss, presence, and AOR/contact tuning.

- **FreePBX automation**  
  Bulk Handler CSV generation + import, `fwconsole` integration, and safe idempotent updates.

- **XoftPhone SSO**  
  Designed to pair with XoftPhone apps for seamless sign-in and multi-device presence.

- **AI integration**  
  Optional assistants to help generate secure secrets, review configs, and explain call-flows.

- **PMS & Booking**  
  Integrations for hospitality (PMS), booking flows, and a **drive-thru** workflow for quick-serve.

- **Batteries included**  
  Systemd unit, HTML templates, and CLI tooling to bootstrap a new deployment fast.

---

## Architecture (High-Level)

```
XoftPhone (iOS/Android/Web/Desktop)
        │  WebRTC (WSS/DTLS/SRTP)
        ▼
   Asterisk <— FreePBX (BMO/Bulk Handler)
        ▲
        │  CLI + CSV + Templates
        │
  XoftSwitch (this repo)
   • Go services / CLI
   • HTML templates (admin)
   • Bulk import/export helpers
   • Optional AI helpers
```

---

## Requirements

- **Asterisk/FreePBX** 16–18 (tested), `fwconsole` in PATH
- **Go** 1.21+ (builds with modules)
- **PHP** (for FreePBX bootstrap calls)
- Linux x86_64 (primary), systemd

---

## Quick Start

```bash
# 1) Build
go mod tidy
go build -o ./bin/xoftswitch ./xoftswitch

# 2) (Optional) Install service files
sudo cp etc/systemd/system/xoftswitch.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable xoftswitch
sudo systemctl start xoftswitch

# 3) Generate & import extensions via FreePBX Bulk Handler
./bin/xoftswitch gen --template ./etc/xoftswitch/templates/addextension.html       --out /tmp/extensions.csv --start 2001 --end 2005       --display "XO {ext}" --email "{ext}@xoftphone.com"

# Import to FreePBX
fwconsole bulkimport --type=extensions /tmp/extensions.csv
fwconsole reload
```

> **Note:** WebRTC presets assume `transport-wss`, `media_encryption=dtls`, `rtp_symmetric=yes`, `rewrite_contact=yes`, `force_rport=yes`, and DTLS cert auto-gen (tune per site policy).

---

## Configuration

XoftSwitch reads sensible defaults; most flags can be provided via CLI or env vars:

- `--php-path` (default `php`)
- `--freepbx-conf` (default `/etc/freepbx.conf`)
- `--fwconsole-path` (default `fwconsole`)
- `--parallel` (default `1`) — be conservative with BMO concurrency
- `--per-call-timeout` (default `15s`)
- `--template` — required for CSV/HTML generation
- `--out` — output CSV/HTML
- `--start/--end` — extension range (supports single like `2001-2001`)
- `--webrtc yes|no`, `--media-encryption dtls|no`, etc.

---

## Security Notes

- **Never** commit real secrets or certificates. Use environment variables or your secret manager.
- Validate WSS/TLS certs and CAs for WebRTC endpoints.
- Limit `fwconsole`/FreePBX access to admins and CI runners you trust.

---

## Roadmap

- [ ] First-class DEB/RPM build recipes in CI
- [ ] Sample Kubernetes manifests
- [ ] Auto-provision XoftPhone clients
- [ ] More PMS providers + hospitality playbooks
- [ ] Drive-thru UX reference flows

---

## Contributing

PRs and issues welcome! Please file reproducible bug reports and include:
- Asterisk/FreePBX versions
- Go version
- Steps to reproduce + logs (minus secrets)

---

## License

[MIT](LICENSE)
