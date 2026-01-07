# Work Plan: Proxy Token Persistence Support

This plan implements support for `STREAMX_AUTH_TOKEN` environment variable to ensure authenticated clickthroughs for Stremio behind reverse proxies.

## Phase 1: Core Configuration
- [x] **Modify `cmd/server/main.go`**:
    - Add `AuthToken string `env:"STREAMX_AUTH_TOKEN"`` to the `config` struct.
    - Pass this token to the `addon.New` constructor using a new functional option.
- [x] **Modify `internal/addon/addon.go`**:
    - Add `authToken string` field to the `Addon` struct.
- [x] **Modify `internal/addon/options.go`**:
    - Add `WithAuthToken(token string) Option` to populate the new field.

## Phase 2: URL Helper Implementation
- [x] **Create/Update `internal/addon/util.go`**:
    - Implement `appendAuthToken(urlStr string, token string) string` helper.
    - This function should parse the URL, check if a token is provided, and append `token=<token>` as a query parameter (handling existing parameters with `&`).

## Phase 3: Manifest & Resource Updates
- [x] **Modify `HandleGetManifest` in `internal/addon/addon.go`**:
    - Use the helper to append the token to the `Logo` URL.
    - Ensure any future `Catalog` or `Background` URLs in the manifest also use this helper.
- [x] **Modify `HandleGetStreams` in `internal/addon/addon.go`**:
    - Review internal stream/download URL construction (e.g., lines 312) to ensure the token is preserved if needed for internal redirects.

## Phase 4: Configuration UI Integration
- [x] **Modify `HandleConfigure` in `internal/static/static.go`**:
    - Instead of just serving the raw `configure.html`, use `strings.Replace` to inject the `STREAMX_AUTH_TOKEN` into a safe location (e.g., a hidden input or a data attribute).
- [x] **Modify `internal/static/configure.html`**:
    - Update the `updateLink` JavaScript function to read the injected token and append it to the `stremio://` and manifest URLs.

## Phase 5: Verification
- [x] **Automated Tests**:
    - [x] Add unit tests for `appendAuthToken` helper.
    - [x] Mock environment variables and verify manifest JSON output.
- [x] **Manual Verification**:
    - [x] Set `STREAMX_AUTH_TOKEN=test-token`.
    - [x] Access `/configure` and verify the "Install" link.
    - [x] Access `/manifest.json` and verify the Logo URL.

