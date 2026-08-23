data "voipms_voicemail" "johns_voicemail" {
  name = "John"
}

data "voipms_forwarding" "mobile" {
  description = "Mobile"
}

data "voipms_subaccount" "gateway" {
  username = "gateway"
}

data "voipms_server" "nyc" {
  hostname = "newyork7.voip.ms"
}

resource "voipms_did" "home" {
  did                  = "5550001001"
  note                 = "Home line"
  routing              = data.voipms_subaccount.gateway.route
  pop_hostname         = data.voipms_server.nyc.hostname
  dialtime             = 30
  voicemail            = data.voipms_voicemail.johns_voicemail.id
  failover_busy        = data.voipms_voicemail.johns_voicemail.route
  failover_noanswer    = data.voipms_voicemail.johns_voicemail.route
  failover_unreachable = data.voipms_forwarding.mobile.route

  sms_enabled     = true
  webhook         = "https://example.lambda-url.us-east-1.on.aws/"
  webhook_enabled = true
}
