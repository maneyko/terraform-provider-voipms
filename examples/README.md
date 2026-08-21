These examples are used by documentation generation (`make generate`).

The household-style configuration in `examples/complete/main.tf` is the example on the Terraform Registry provider page.

Copy `examples/provider/provider.tf` into a directory with a data-source or
resource example, set `VOIPMS_USERNAME` and `VOIPMS_PASSWORD`, and use a
Terraform CLI config with `dev_overrides` if you are running a locally built
binary.

Resource examples live under `examples/resources/<type>/` (`resource.tf` plus
`import.sh`). Data source examples live under `examples/data-sources/<type>/`.
