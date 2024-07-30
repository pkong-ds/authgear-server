package config

var _ = Schema.Add("BotProtectionConfig", `
{
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"enabled": { "type": "boolean" },
    "provider": { "$ref": "#/$defs/BotProtectionProvider" },
		"builtin_flow_requirement": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirement" }
	}
}
`)

var _ = Schema.Add("BotProtectionProvider", `
{
	"type": "object",
	"additionalProperties": false,
	"required": ["type"],
	"properties": {
		"type": { "type": "string", "enum": ["cloudflare", "recaptchav2"] },
		"site_key": { "type": "string" }
	},
	"allOf": [
		{
			"if": {
				"properties": {
					"type": {
						"enum": ["cloudflare", "recaptchav2"]
					}
				},
				"required": ["type"]
			},
			"then": {
				"required": ["site_key"]
			}
		}
	]
}
`)

var _ = Schema.Add("BotProtectionBuiltinFlowRequirement", `
{
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"signup_or_login": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirementObject" },
		"account_recovery": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirementObject" },
		"password": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirementObject" },
		"oob_otp_email": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirementObject" },
		"oob_otp_sms": { "$ref": "#/$defs/BotProtectionBuiltinFlowRequirementObject" }
	}
}
`)

var _ = Schema.Add("BotProtectionBuiltinFlowRequirementObject", `
{
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"mode": { "$ref": "#/$defs/BotProtectionRiskMode" }
	},
	"required": ["mode"]
}
`)

var _ = Schema.Add("BotProtectionRiskMode", `
{
	"type": "string",
	"enum": ["never", "always"]
}
`)

type BotProtectionConfig struct {
	Enabled                bool                                 `json:"enabled,omitempty"`
	Provider               *BotProtectionProvider               `json:"provider,omitempty" nullable:"true"`
	BuiltinFlowRequirement *BotProtectionBuiltinFlowRequirement `json:"builtin_flow_requirement,omitempty" nullable:"true"`
}

type BotProtectionProvider struct {
	Type    BotProtectionProviderType `json:"type,omitempty"`
	SiteKey string                    `json:"site_key,omitempty"` // only for cloudflare, recaptchav2
}

type BotProtectionProviderType string

const (
	BotProtectionProviderTypeCloudflare  BotProtectionProviderType = "cloudflare"
	BotProtectionProviderTypeRecaptchaV2 BotProtectionProviderType = "recaptchav2"
)

type BotProtectionBuiltinFlowRequirement struct {
	SignupOrLogin   *BotProtectionBuiltinFlowRequirementObject `json:"signup_or_login,omitempty"`
	AccountRecovery *BotProtectionBuiltinFlowRequirementObject `json:"account_recovery,omitempty"`
	Password        *BotProtectionBuiltinFlowRequirementObject `json:"password,omitempty"`
	OOBOTPEmail     *BotProtectionBuiltinFlowRequirementObject `json:"oob_otp_email,omitempty"`
	OOBOTPSms       *BotProtectionBuiltinFlowRequirementObject `json:"oob_otp_sms,omitempty"`
}

type BotProtectionBuiltinFlowRequirementObject struct {
	Mode BotProtectionRiskMode `json:"mode,omitempty"`
}

type BotProtectionRiskMode string

const (
	BotProtectionRiskModeNever  BotProtectionRiskMode = "never"
	BotProtectionRiskModeAlways BotProtectionRiskMode = "always"
)
