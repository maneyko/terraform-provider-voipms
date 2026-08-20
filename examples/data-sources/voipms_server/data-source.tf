data "voipms_server" "newyork7" {
  hostname = "newyork7.voip.ms"
}

output "pop_id" {
  value = data.voipms_server.newyork7.pop
}

output "label" {
  value = "${data.voipms_server.newyork7.name} (${data.voipms_server.newyork7.hostname})"
}
