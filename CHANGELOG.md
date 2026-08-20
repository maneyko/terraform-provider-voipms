# Changelog

All notable changes to this provider will be documented in this file.

## Unreleased

### Fixed

- `GetSubAccount` / import by numeric id: VoIP.ms `getSubAccounts?account=<id>` returns `no_subaccount`, so the client now falls back to listing all sub-accounts and matching by id.

### Added

- `voipms_did` accepts `pop_hostname` (`newyork7.voip.ms` or `New York 7`) so configs do not have to use a numeric POP id. `voipms_server` can look up a POP by `hostname` or `name` as well as `pop`.
- Named lookups: voicemail by display `name`, forwarding/callback by `description`, phonebook group/entry by `name`, sub-account by `username`. DID `voicemail_name` and phonebook `group_name` resolve the same way as `pop_hostname`.
- Resources and data sources for sub-accounts, DIDs (routing + SMS), forwarding, voicemail, callbacks, caller-ID filters, phonebook entries/groups, and POP servers.
- Single and list data sources for each of those objects.
- Resources to create/update/delete them (DID configure-only: no order/cancel).
- Registry-style docs generated with tfplugindocs.
- Provider credentials also accept `voip_ms_username` / `voip_ms_api_key`.
- Unit tests for inventory client methods and provider registration.
- `make install` / `make install-plugin` for local use from another Terraform repo.
- GitHub repository `vetal-ca-org/terraform-provider-voipms` (HashiCorp provider naming).
