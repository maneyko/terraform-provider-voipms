terraform {
  required_providers {
    voipms = {
      source  = "vetal-ca-org/voipms"
      version = "~> 0.2"
    }
  }
}

provider "voipms" {
  # Set VOIPMS_USERNAME and VOIPMS_PASSWORD in the environment.
  # Allow-list this machine's public IP under SOAP & REST/JSON API.
}

variable "gateway_sip_password" {
  type        = string
  sensitive   = true
  description = "SIP password for the FreeSWITCH / PBX sub-account."
}

variable "voicemail_pin" {
  type        = string
  sensitive   = true
  description = "Shared mailbox PIN for the example voicemail boxes."
}

# Named DID routes. VoIP.ms routing is type:target (account / fwd / vm / sys / none).
# account: is the full SIP login; fwd: and vm: match description or mailbox name.
locals {
  route_gateway     = "account:${voipms_subaccount.gateway.account}"
  route_vm_alex     = "vm:${voipms_voicemail.alex.name}"
  route_alex_mobile = "fwd:${voipms_forwarding.alex_mobile.description}"
}

# POP by hostname (also accepts display name, e.g. "New York 7").
data "voipms_server" "nyc" {
  hostname = "newyork7.voip.ms"
}

# SIP trunk for a local PBX. username is the suffix only (max 12 chars);
# the full login is account (e.g. 100001_gateway).
resource "voipms_subaccount" "gateway" {
  username              = "gateway"
  password              = var.gateway_sip_password
  description           = "Home FreeSWITCH gateway"
  device_type           = "1"
  allowed_codecs        = "ulaw;g722"
  nat                   = "no"
  encrypted_sip_traffic = true
  canada_routing        = "premium"
}

resource "voipms_voicemail" "alex" {
  mailbox        = "101"
  name           = "Alex"
  password       = var.voicemail_pin
  skip_password  = true
  email          = "alex@example.com"
  attach_message = true
  timezone       = "America/New_York"
}

resource "voipms_voicemail" "jordan" {
  mailbox        = "102"
  name           = "Jordan"
  password       = var.voicemail_pin
  skip_password  = true
  email          = "jordan@example.com"
  attach_message = true
  timezone       = "America/New_York"
}

resource "voipms_forwarding" "alex_mobile" {
  phone_number = "5550002001"
  description  = "Alex mobile"
}

resource "voipms_callback" "alex_mobile" {
  description     = "Alex mobile"
  number          = "15550002001"
  callerid_number = "5550001001"
}

resource "voipms_phonebook_group" "family" {
  name = "Family"
}

resource "voipms_phonebook_entry" "alex_mobile" {
  name       = "Alex mobile"
  number     = "5550002001"
  group_name = voipms_phonebook_group.family.name
  speed_dial = "21"
}

resource "voipms_phonebook_entry" "jordan_mobile" {
  name       = "Jordan mobile"
  number     = "5550002002"
  group_name = voipms_phonebook_group.family.name
  speed_dial = "22"
}

# DID must already be on the account. Terraform configures routing; it does
# not order or cancel the number. Destroy only drops state.
resource "voipms_did" "home" {
  did            = "5550001001"
  note           = "Home line"
  routing        = local.route_gateway
  pop_hostname   = data.voipms_server.nyc.hostname
  dialtime       = 30
  cnam           = true
  voicemail_name = voipms_voicemail.alex.name

  failover_busy        = local.route_vm_alex
  failover_noanswer    = local.route_vm_alex
  failover_unreachable = local.route_alex_mobile

  sms_enabled     = true
  webhook         = "https://example.com/sms"
  webhook_enabled = true
}

resource "voipms_caller_id_filter" "blocked_prefix" {
  callerid = "999XXXXXXXX"
  did      = "all"
  routing  = "sys:hangup"
  note     = "Blocked caller ID prefix"
}
