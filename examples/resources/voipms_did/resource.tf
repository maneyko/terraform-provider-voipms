locals {
  my_pops = {
    ny7 = "newyork7.voip.ms"
  }
}

data "voipms_voicemail" "johns_voicemail" {
  name = "John"
}

data "voipms_forwarding" "mobile" {
  description = "Mobile"
}

data "voipms_subaccount" "gateway" {
  username = "gateway"
}

resource "voipms_did" "home" {
  did                  = "5550001001"
  note                 = "Home line"
  routing              = "account:${data.voipms_subaccount.gateway.account}"
  pop_hostname         = local.my_pops.ny7
  dialtime             = 30
  voicemail            = data.voipms_voicemail.johns_voicemail.mailbox
  failover_busy        = "vm:${data.voipms_voicemail.johns_voicemail.mailbox}"
  failover_noanswer    = "vm:${data.voipms_voicemail.johns_voicemail.mailbox}"
  failover_unreachable = "vm:${data.voipms_voicemail.johns_voicemail.mailbox}"

  sms_enabled     = true
  webhook         = "https://example.lambda-url.us-east-1.on.aws/"
  webhook_enabled = true
}
