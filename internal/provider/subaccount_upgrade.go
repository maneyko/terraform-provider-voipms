package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/vetal-ca-org/terraform-provider-voipms/internal/client"
)

type subaccountModelV0 struct {
	ID                   types.String `tfsdk:"id"`
	Account              types.String `tfsdk:"account"`
	Username             types.String `tfsdk:"username"`
	Description          types.String `tfsdk:"description"`
	Protocol             types.String `tfsdk:"protocol"`
	AuthType             types.String `tfsdk:"auth_type"`
	Password             types.String `tfsdk:"password"`
	IP                   types.String `tfsdk:"ip"`
	DeviceType           types.String `tfsdk:"device_type"`
	CallerIDNumber       types.String `tfsdk:"callerid_number"`
	CanadaRouting        types.String `tfsdk:"canada_routing"`
	LockInternational    types.String `tfsdk:"lock_international"`
	InternationalRoute   types.String `tfsdk:"international_route"`
	MusicOnHold          types.String `tfsdk:"music_on_hold"`
	Language             types.String `tfsdk:"language"`
	AllowedCodecs        types.String `tfsdk:"allowed_codecs"`
	DTMFMode             types.String `tfsdk:"dtmf_mode"`
	NAT                  types.String `tfsdk:"nat"`
	SIPTraffic           types.Bool   `tfsdk:"sip_traffic"`
	MaxExpiry            types.Int64  `tfsdk:"max_expiry"`
	RTPTimeout           types.Int64  `tfsdk:"rtp_timeout"`
	RTPHoldTimeout       types.Int64  `tfsdk:"rtp_hold_timeout"`
	IPRestriction        types.String `tfsdk:"ip_restriction"`
	EnableIPRestriction  types.Bool   `tfsdk:"enable_ip_restriction"`
	POPRestriction       types.String `tfsdk:"pop_restriction"`
	EnablePOPRestriction types.Bool   `tfsdk:"enable_pop_restriction"`
	RecordCalls          types.Bool   `tfsdk:"record_calls"`
	Allow225             types.Bool   `tfsdk:"allow225"`
	InternalExtension    types.String `tfsdk:"internal_extension"`
	InternalVoicemail    types.String `tfsdk:"internal_voicemail"`
	InternalDialtime     types.String `tfsdk:"internal_dialtime"`
	EnableInternalCNAM   types.Bool   `tfsdk:"enable_internal_cnam"`
	DialingMode          types.String `tfsdk:"dialing_mode"`
	DefaultE911          types.String `tfsdk:"default_e911"`
	CallPickupBehavior   types.String `tfsdk:"call_pickup_behavior"`
}

type subaccountModelV1 struct {
	ID                   types.String `tfsdk:"id"`
	Account              types.String `tfsdk:"account"`
	Username             types.String `tfsdk:"username"`
	Description          types.String `tfsdk:"description"`
	Protocol             types.String `tfsdk:"protocol"`
	AuthType             types.String `tfsdk:"auth_type"`
	Password             types.String `tfsdk:"password"`
	IP                   types.String `tfsdk:"ip"`
	DeviceType           types.String `tfsdk:"device_type"`
	CallerIDNumber       types.String `tfsdk:"callerid_number"`
	CanadaRouting        types.String `tfsdk:"canada_routing"`
	LockInternational    types.String `tfsdk:"lock_international"`
	InternationalRoute   types.String `tfsdk:"international_route"`
	MusicOnHold          types.String `tfsdk:"music_on_hold"`
	Language             types.String `tfsdk:"language"`
	AllowedCodecs        types.String `tfsdk:"allowed_codecs"`
	DTMFMode             types.String `tfsdk:"dtmf_mode"`
	NAT                  types.String `tfsdk:"nat"`
	SIPTraffic           types.Bool   `tfsdk:"sip_traffic"`
	MaxExpiry            types.Int64  `tfsdk:"max_expiry"`
	RTPTimeout           types.Int64  `tfsdk:"rtp_timeout"`
	RTPHoldTimeout       types.Int64  `tfsdk:"rtp_hold_timeout"`
	IPRestriction        types.String `tfsdk:"ip_restriction"`
	EnableIPRestriction  types.Bool   `tfsdk:"enable_ip_restriction"`
	POPRestriction       types.String `tfsdk:"pop_restriction"`
	EnablePOPRestriction types.Bool   `tfsdk:"enable_pop_restriction"`
	RecordCalls          types.Bool   `tfsdk:"record_calls"`
	Allow225Balance      types.Bool   `tfsdk:"allow_225_balance"`
	InternalExtension    types.String `tfsdk:"internal_extension"`
	InternalVoicemail    types.String `tfsdk:"internal_voicemail"`
	InternalDialtime     types.String `tfsdk:"internal_dialtime"`
	EnableInternalCNAM   types.Bool   `tfsdk:"enable_internal_cnam"`
	DialingMode          types.String `tfsdk:"dialing_mode"`
	DefaultE911          types.String `tfsdk:"default_e911"`
	CallPickupBehavior   types.String `tfsdk:"call_pickup_behavior"`
}

func (r *subaccountResource) UpgradeState(_ context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: subaccountResourceAttributesV0(),
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior subaccountModelV0
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, subaccountModel{
					ID:                   prior.ID,
					Route:                accountRouteValue(prior.Account),
					Account:              prior.Account,
					Username:             prior.Username,
					Description:          prior.Description,
					Protocol:             prior.Protocol,
					AuthType:             prior.AuthType,
					Password:             prior.Password,
					IP:                   prior.IP,
					DeviceType:           prior.DeviceType,
					CallerIDNumber:       prior.CallerIDNumber,
					CanadaRouting:        prior.CanadaRouting,
					LockInternational:    prior.LockInternational,
					InternationalRoute:   prior.InternationalRoute,
					MusicOnHold:          prior.MusicOnHold,
					Language:             prior.Language,
					AllowedCodecs:        prior.AllowedCodecs,
					DTMFMode:             prior.DTMFMode,
					NAT:                  prior.NAT,
					SIPTraffic:           prior.SIPTraffic,
					MaxExpiry:            prior.MaxExpiry,
					RTPTimeout:           prior.RTPTimeout,
					RTPHoldTimeout:       prior.RTPHoldTimeout,
					IPRestriction:        prior.IPRestriction,
					EnableIPRestriction:  prior.EnableIPRestriction,
					POPRestriction:       prior.POPRestriction,
					EnablePOPRestriction: prior.EnablePOPRestriction,
					RecordCalls:          prior.RecordCalls,
					Allow225Balance:      prior.Allow225,
					InternalExtension:    prior.InternalExtension,
					InternalVoicemail:    prior.InternalVoicemail,
					InternalDialtime:     prior.InternalDialtime,
					EnableInternalCNAM:   prior.EnableInternalCNAM,
					DialingMode:          prior.DialingMode,
					DefaultE911:          prior.DefaultE911,
					CallPickupBehavior:   prior.CallPickupBehavior,
				})...)
			},
		},
		1: {
			PriorSchema: &schema.Schema{
				Attributes: subaccountResourceAttributesV1(),
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior subaccountModelV1
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, subaccountModel{
					ID:                   prior.ID,
					Route:                accountRouteValue(prior.Account),
					Account:              prior.Account,
					Username:             prior.Username,
					Description:          prior.Description,
					Protocol:             prior.Protocol,
					AuthType:             prior.AuthType,
					Password:             prior.Password,
					IP:                   prior.IP,
					DeviceType:           prior.DeviceType,
					CallerIDNumber:       prior.CallerIDNumber,
					CanadaRouting:        prior.CanadaRouting,
					LockInternational:    prior.LockInternational,
					InternationalRoute:   prior.InternationalRoute,
					MusicOnHold:          prior.MusicOnHold,
					Language:             prior.Language,
					AllowedCodecs:        prior.AllowedCodecs,
					DTMFMode:             prior.DTMFMode,
					NAT:                  prior.NAT,
					SIPTraffic:           prior.SIPTraffic,
					MaxExpiry:            prior.MaxExpiry,
					RTPTimeout:           prior.RTPTimeout,
					RTPHoldTimeout:       prior.RTPHoldTimeout,
					IPRestriction:        prior.IPRestriction,
					EnableIPRestriction:  prior.EnableIPRestriction,
					POPRestriction:       prior.POPRestriction,
					EnablePOPRestriction: prior.EnablePOPRestriction,
					RecordCalls:          prior.RecordCalls,
					Allow225Balance:      prior.Allow225Balance,
					InternalExtension:    prior.InternalExtension,
					InternalVoicemail:    prior.InternalVoicemail,
					InternalDialtime:     prior.InternalDialtime,
					EnableInternalCNAM:   prior.EnableInternalCNAM,
					DialingMode:          prior.DialingMode,
					DefaultE911:          prior.DefaultE911,
					CallPickupBehavior:   prior.CallPickupBehavior,
				})...)
			},
		},
	}
}

func priorOptBool() schema.BoolAttribute {
	return schema.BoolAttribute{
		Optional:      true,
		Computed:      true,
		PlanModifiers: []planmodifier.Bool{optBool()},
	}
}

func subaccountResourceAttributesV0() map[string]schema.Attribute {
	attrs := copySchemaAttrs(subaccountResourceAttributes())
	delete(attrs, "allow_225_balance")
	delete(attrs, "encrypted_sip_traffic")
	delete(attrs, "route")
	attrs["allow225"] = priorOptBool()
	attrs["sip_traffic"] = priorOptBool()
	return attrs
}

func subaccountResourceAttributesV1() map[string]schema.Attribute {
	attrs := copySchemaAttrs(subaccountResourceAttributes())
	delete(attrs, "encrypted_sip_traffic")
	delete(attrs, "route")
	attrs["sip_traffic"] = priorOptBool()
	return attrs
}

func accountRouteValue(account types.String) types.String {
	if account.IsNull() || account.IsUnknown() || account.ValueString() == "" {
		return types.StringNull()
	}
	return types.StringValue(client.AccountRoute(account.ValueString()))
}

func copySchemaAttrs(src map[string]schema.Attribute) map[string]schema.Attribute {
	attrs := make(map[string]schema.Attribute, len(src)+1)
	for k, v := range src {
		attrs[k] = v
	}
	return attrs
}
