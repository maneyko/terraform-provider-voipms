---
page_title: "Releasing a new version"
subcategory: ""
description: |-
  How to publish a new version of the VoIP.ms provider to the Terraform Registry.
---

# Releasing a new version

The public listing is [registry.terraform.io/providers/vetal-ca-org/voipms](https://registry.terraform.io/providers/vetal-ca-org/voipms). Source address: `vetal-ca-org/voipms`.

GitHub Actions plus GoReleaser build the binaries, write checksums, and GPG-sign the checksums. After the first publish, a new GitHub Release is enough — do not click Publish again and do not re-upload the GPG key.

## Prerequisites (one-time)

These are already in place for this repository:

- Public GitHub repo `vetal-ca-org/terraform-provider-voipms`
- Namespace `vetal-ca-org` claimed in HCP Terraform, with signing key `F3ADF9C3A8C694B3`
- Actions secrets `GPG_PRIVATE_KEY` and `PASSPHRASE`
- Workflow [`.github/workflows/release.yml`](https://github.com/vetal-ca-org/terraform-provider-voipms/blob/master/.github/workflows/release.yml) and [`.goreleaser.yml`](https://github.com/vetal-ca-org/terraform-provider-voipms/blob/master/.goreleaser.yml)

## Each release

1. Land the change on `master` through a pull request. CI must be green (`make test` / Tests workflow).
2. From `master`, tag the **next** semantic version. There must not be a branch with the same name as the tag.

   | Change | Tag example |
   | --- | --- |
   | Bug fix or docs | `v0.1.1` |
   | New resource / data source, compatible | `v0.2.0` |
   | Breaking schema or behavior | `v1.0.0` |

   ```shell
   git checkout master
   git pull
   git tag v0.1.1
   git push origin v0.1.1
   ```

3. Confirm **Actions → Release** succeeds. The GitHub Release must include platform zips, `terraform-provider-voipms_<version>_manifest.json`, `_SHA256SUMS`, and `_SHA256SUMS.sig`.
4. Wait for the Terraform Registry to ingest the release (usually a few minutes). It appears as a new version of `vetal-ca-org/voipms`.
5. Consumers upgrade with `terraform init -upgrade` (or a tighter `version` constraint).

## Do not

- Reuse or move an existing tag (`v0.1.0` stays `v0.1.0` forever).
- Replace zip files or checksums on a published GitHub Release.
- Remove the GPG public key from the `vetal-ca-org` namespace. To rotate, **add** a new key, update the Actions secrets, then tag; leave the old key so older versions still verify.
- Put secrets, live phone numbers, or SIP passwords in git. Docs and examples use fictional `555` numbers.

## Registry docs

Provider pages on the Registry come from `docs/` in the tagged commit. After you change examples or templates, regenerate before merging:

```shell
make generate
```

The household-style configuration on the provider page is [`examples/complete/main.tf`](https://github.com/vetal-ca-org/terraform-provider-voipms/blob/master/examples/complete/main.tf).
