---
type: feature
priority: high
created: 2026-01-07T20:45:00Z
status: implemented
tags: [auth, stremio, proxy, caddy]
keywords: [STREAMX_AUTH_TOKEN, manifest, link generation, stremio://]
patterns: [os.Getenv, c.BaseURL, manifest.json]
---

# FEATURE: Support Static Auth Token for Reverse Proxy Environments

## Description
Modify the StreamX container to support a static authentication token via the `STREAMX_AUTH_TOKEN` environment variable. This token must be appended to internal links generated for Stremio (manifests, resources, and install links) to ensure authenticated clickthroughs when running behind a reverse proxy (like Caddy) that requires token-based auth for non-browser clients (like Stremio) that do not support cookies.

## Context
The user runs StreamX behind a reverse proxy requiring `?token=sk-<token>`. While browser-based configuration works via cookies, Stremio clients do not persist these cookies and thus fail to authenticate subsequent requests. Appending the token to generated links within the Stremio ecosystem solves this.

## Requirements

### Functional Requirements
- Load `STREAMX_AUTH_TOKEN` from environment variables.
- Append `token=<value>` as a query parameter to:
    - The `stremio://` install link generated in the configuration UI.
    - The `Logo` URL in the manifest JSON.
    - Resource URLs (Catalog, Meta, Stream) inside the manifest if they point back to the addon.
- Ensure existing query parameters are preserved (append with `&`).
- The feature is optional; if the env var is not set, behavior remains unchanged.
- Token should NOT be visible in the configuration form fields (server-side injection only).

### Non-Functional Requirements
- Minimal performance impact on manifest generation.
- No leakage of the token to external services (e.g., Real Debrid, Cinemeta).

## Current State
- Links are generated using `c.BaseURL()` or hardcoded relative paths without auth parameters.
- `HandleGetManifest` and `HandleGetStreams` use `c.BaseURL()` but don't account for external auth tokens.
- `configure.html` uses JavaScript to build the `stremio://` link.

## Desired State
- A centralized helper or middleware identifies "internal" generated links and appends the token if `STREAMX_AUTH_TOKEN` is present.
- Manifest responses contain authenticated URLs for Stremio's internal consumption.

## Research Context

### Keywords to Search
- `STREAMX_AUTH_TOKEN` - New environment variable name.
- `HandleGetManifest` - Main entry point for Stremio integration.
- `c.BaseURL()` - Used throughout `addon.go` to build absolute links.
- `stremio://` - Pattern in `configure.html` for the install button.

### Patterns to Investigate
- `internal/addon/addon.go`: `HandleGetManifest` (lines 125-163) and `HandleGetStreams`.
- `internal/static/configure.html`: JS logic for `installLink`.
- `cmd/server/main.go`: Config struct parsing for the new env var.

### Key Decisions Made
- **Token Name**: `STREAMX_AUTH_TOKEN`.
- **Scope**: Limited to Stremio manifests and install links; external services excluded.
- **Implementation**: Server-side injection preferred for manifest URLs; JS/Server-side coordination for the install link.

## Success Criteria

### Automated Verification
- [ ] Test that `manifest.json` contains the token in the Logo URL when env var is set.
- [ ] Test that `manifest.json` does NOT contain the token when env var is unset.

### Manual Verification
- [ ] Verify the "Install" link on the `/configure` page includes the token.
- [ ] Verify Stremio can fetch the logo and manifest resources through the proxy using the token.
