import { Context, FormattedMessage } from "@oursky/react-messageformat";
import React, { useCallback, useContext, useState } from "react";
import cn from "classnames";
import ScreenContent from "../../ScreenContent";
import ScreenTitle from "../../ScreenTitle";
import styles from "./BotProtectionConfigurationScreen.module.css";
import {
  BotProtectionProviderType,
  PortalAPIAppConfig,
  PortalAPISecretConfig,
  PortalAPISecretConfigUpdateInstruction,
} from "../../types";
import {
  AppSecretConfigFormModel,
  useAppSecretConfigForm,
} from "../../hook/useAppSecretConfigForm";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { useAppSecretVisitToken } from "./mutations/generateAppSecretVisitTokenMutation";
import { AppSecretKey } from "./globalTypes.generated";
import ShowError from "../../ShowError";
import ShowLoading from "../../ShowLoading";
import { useLocationEffect } from "../../hook/useLocationEffect";
import { produce } from "immer";
import { clearEmptyObject } from "../../util/misc";
import FormContainer from "../../FormContainer";
import ScreenDescription from "../../ScreenDescription";
import Toggle from "../../Toggle";
import WidgetTitle from "../../WidgetTitle";
import { DefaultEffects, IButtonProps, Image, Label } from "@fluentui/react";
import { useSystemConfig } from "../../context/SystemConfigContext";
import recaptchaV2LogoURL from "../../images/recaptchav2_logo.svg";
import cloudflareLogoURL from "../../images/cloudflare_logo.svg";
import WidgetDescription from "../../WidgetDescription";
import FormTextField from "../../FormTextField";
import PrimaryButton from "../../PrimaryButton";
import { startReauthentication } from "./Authenticated";

const MASKED_SECRET = "***************";

const SECRET_KEY_FORM_FIELD_ID = "secret-key-form-field";

interface LocationState {
  isOAuthRedirect: boolean;
}
function isLocationState(raw: unknown): raw is LocationState {
  return (
    raw != null &&
    typeof raw === "object" &&
    (raw as Partial<LocationState>).isOAuthRedirect != null
  );
}

interface FormState {
  enabled: boolean;
  providerType: BotProtectionProviderType;
  siteKey: string;
  secretKey: string | null;
}

function constructFormState(
  config: PortalAPIAppConfig,
  secrets: PortalAPISecretConfig
): FormState {
  const enabled = config.bot_protection?.enabled ?? false;
  const providerType = config.bot_protection?.provider?.type ?? "recaptchav2";
  const siteKey = config.bot_protection?.provider?.site_key ?? "";
  const secretKey = secrets.botProtectionProviderSecret?.secretKey ?? null;

  return {
    enabled,
    providerType,
    siteKey,
    secretKey,
  };
}

function constructConfig(
  config: PortalAPIAppConfig,
  secrets: PortalAPISecretConfig,
  _initialState: FormState,
  currentState: FormState,
  _effectiveConfig: PortalAPIAppConfig
): [PortalAPIAppConfig, PortalAPISecretConfig] {
  return produce([config, secrets], ([config, secrets]) => {
    config.bot_protection ??= {};
    config.bot_protection.provider ??= {};
    config.bot_protection = {
      enabled: currentState.enabled,
      provider: {
        type: currentState.providerType,
        site_key: currentState.siteKey,
      },
    };

    if (currentState.secretKey != null) {
      secrets.botProtectionProviderSecret = {
        secretKey: currentState.secretKey,
        type: currentState.providerType,
      };
    }
    clearEmptyObject(config);
  });
}

function constructSecretUpdateInstruction(
  _config: PortalAPIAppConfig,
  _secrets: PortalAPISecretConfig,
  currentState: FormState
): PortalAPISecretConfigUpdateInstruction | undefined {
  if (currentState.secretKey == null) {
    return undefined;
  }
  return {
    botProtectionProviderSecret: {
      action: "set",
      data: {
        secretKey: currentState.secretKey,
        type: currentState.providerType,
      },
    },
  };
}

interface ProviderCardProps {
  className?: string;
  logoSrc?: any;
  logoWidth?: number;
  logoHeight?: number;
  children?: React.ReactNode;
  onClick?: IButtonProps["onClick"];
  isSelected?: boolean;
  disabled?: boolean;
}

function ProviderCard(props: ProviderCardProps) {
  const {
    className,
    disabled,
    isSelected,
    children,
    onClick,
    logoSrc,
    logoHeight = 48,
    logoWidth = 48,
  } = props;

  const {
    themes: {
      main: {
        palette: { themePrimary },
        semanticColors: { disabledBackground: backgroundColor },
      },
    },
  } = useSystemConfig();

  return (
    <div
      style={{
        boxShadow: disabled ? undefined : DefaultEffects.elevation4,
        borderColor: isSelected ? themePrimary : "transparent",
        backgroundColor: disabled ? backgroundColor : undefined,
        cursor: disabled ? "not-allowed" : undefined,
      }}
      className={cn(className, styles.providerCard)}
      onClick={disabled ? undefined : onClick}
    >
      {logoSrc != null ? (
        <Image src={logoSrc} width={logoWidth} height={logoHeight} />
      ) : null}
      <Label className={styles.providerCardLabel}>{children}</Label>
    </div>
  );
}

export interface BotProtectionConfigurationContentProps {
  form: AppSecretConfigFormModel<FormState>;
}

const BotProtectionConfigurationContent: React.VFC<BotProtectionConfigurationContentProps> =
  function BotProtectionConfigurationContent(props) {
    const { form } = props;
    const { state, setState } = form;

    const { renderToString } = useContext(Context);

    const locationState = useLocationEffect((state: LocationState) => {
      if (state.isOAuthRedirect) {
        window.location.hash = "";
        window.location.hash = "#" + SECRET_KEY_FORM_FIELD_ID;
      }
    });

    const [revealed, setRevealed] = useState(
      locationState?.isOAuthRedirect ?? false
    );

    const onChangeEnabled = useCallback(
      (_event, checked?: boolean) => {
        if (checked != null) {
          setState((state) => {
            return {
              ...state,
              enabled: checked,
            };
          });
        }
      },
      [setState]
    );

    const onClickProviderRecaptchaV2 = useCallback(
      (e: React.MouseEvent<unknown>) => {
        e.preventDefault();
        e.stopPropagation();
        if (state.providerType === "recaptchav2") {
          return;
        }
        setState((state) => {
          return {
            ...state,
            providerType: "recaptchav2",
          };
        });
      },
      [setState, state.providerType]
    );

    const onClickProviderCloudflare = useCallback(
      (e: React.MouseEvent<unknown>) => {
        e.preventDefault();
        e.stopPropagation();
        if (state.providerType === "cloudflare") {
          return;
        }
        setState((state) => {
          return {
            ...state,
            providerType: "cloudflare",
          };
        });
      },
      [setState, state.providerType]
    );

    const onChangeSiteKey = useCallback(
      (_, value?: string) => {
        if (value != null) {
          setState((state) => {
            return {
              ...state,
              siteKey: value,
            };
          });
        }
      },
      [setState]
    );

    const onChangeSecretKey = useCallback(
      (_, value?: string) => {
        if (value != null) {
          setState((state) => {
            return {
              ...state,
              secretKey: value,
            };
          });
        }
      },
      [setState]
    );

    const navigate = useNavigate();
    const onClickReveal = useCallback(
      (e: React.MouseEvent<unknown>) => {
        e.preventDefault();
        e.stopPropagation();

        if (state.secretKey != null) {
          setRevealed(true);
          return;
        }

        const locationState: LocationState = {
          isOAuthRedirect: true,
        };

        startReauthentication(navigate, locationState).catch((e) => {
          // Normally there should not be any error.
          console.error(e);
        });
      },
      [navigate, state.secretKey]
    );

    return (
      <ScreenContent>
        <ScreenTitle className={styles.widget}>
          <FormattedMessage id="BotProtectionConfigurationScreen.title" />
        </ScreenTitle>
        <ScreenDescription className={styles.widget}>
          <FormattedMessage id="BotProtectionConfigurationScreen.description" />
        </ScreenDescription>
        <div className={styles.content}>
          <Toggle
            // TODO: figure out 4px gap between toggle and label
            checked={state.enabled}
            onChange={onChangeEnabled}
            label={renderToString(
              "BotProtectionConfigurationScreen.enable.label"
            )}
            inlineLabel={false}
          />
          {state.enabled ? (
            <div className={styles.enabledContent}>
              <WidgetTitle>
                <FormattedMessage id="BotProtectionConfigurationScreen.challengeProvider.title" />
              </WidgetTitle>
              <div className={styles.providerCardContainer}>
                <ProviderCard
                  className={styles.columnLeft}
                  onClick={onClickProviderRecaptchaV2}
                  isSelected={state.providerType === "recaptchav2"}
                  logoSrc={recaptchaV2LogoURL}
                >
                  <FormattedMessage id="BotProtectionConfigurationScreen.provider.recaptchaV2.label" />
                </ProviderCard>
                <ProviderCard
                  className={styles.columnRight}
                  onClick={onClickProviderCloudflare}
                  isSelected={state.providerType === "cloudflare"}
                  logoSrc={cloudflareLogoURL}
                >
                  <FormattedMessage id="BotProtectionConfigurationScreen.provider.cloudflare.label" />
                </ProviderCard>
              </div>
              {state.providerType === "recaptchav2" ? (
                <WidgetDescription>
                  <FormattedMessage id="BotProtectionConfigurationScreen.provider.recaptchaV2.description" />
                </WidgetDescription>
              ) : (
                <WidgetDescription>
                  <FormattedMessage id="BotProtectionConfigurationScreen.provider.cloudflare.description" />
                </WidgetDescription>
              )}
              <FormTextField
                type="text"
                label={renderToString(
                  "BotProtectionConfigurationScreen.provider.siteKey.label"
                )}
                value={state.siteKey}
                required={true}
                onChange={onChangeSiteKey}
                parentJSONPointer=""
                fieldName="siteKey"
              />
              <div className={styles.secretKeyInputContainer}>
                <FormTextField
                  className={
                    revealed
                      ? styles.secretKeyInputWithoutReveal
                      : styles.secretKeyInputWithReveal
                  }
                  id={SECRET_KEY_FORM_FIELD_ID}
                  type="text"
                  label={renderToString(
                    "BotProtectionConfigurationScreen.provider.secretKey.label"
                  )}
                  value={revealed ? state.secretKey ?? "" : MASKED_SECRET}
                  required={true}
                  onChange={onChangeSecretKey}
                  parentJSONPointer=""
                  fieldName="secretKey"
                  readOnly={!revealed}
                />
                {!revealed ? (
                  <PrimaryButton
                    className={styles.secretKeyRevealButton}
                    onClick={onClickReveal}
                    text={<FormattedMessage id="reveal" />}
                  />
                ) : null}
              </div>
            </div>
          ) : null}
        </div>
      </ScreenContent>
    );
  };

const BotProtectionConfigurationScreen1: React.VFC<{
  appID: string;
  secretToken: string | null;
}> = function BotProtectionConfigurationScreen1({ appID, secretToken }) {
  const form = useAppSecretConfigForm({
    appID,
    secretVisitToken: secretToken,
    constructFormState,
    constructConfig,
    constructSecretUpdateInstruction,
  });

  if (form.isLoading) {
    return <ShowLoading />;
  }

  if (form.loadError) {
    return <ShowError error={form.loadError} onRetry={form.reload} />;
  }

  return (
    <FormContainer form={form}>
      <BotProtectionConfigurationContent form={form} />
    </FormContainer>
  );
};

const SECRETS = [AppSecretKey.BotProtectionProviderSecret];

const BotProtectionConfigurationScreen: React.VFC =
  function BotProtectionConfigurationScreen() {
    const { appID } = useParams() as { appID: string };
    const location = useLocation();
    const [shouldRefreshToken] = useState<boolean>(() => {
      const { state } = location;
      if (isLocationState(state) && state.isOAuthRedirect) {
        return true;
      }
      return false;
    });
    const { token, error, retry } = useAppSecretVisitToken(
      appID,
      SECRETS,
      shouldRefreshToken
    );
    if (error) {
      return <ShowError error={error} onRetry={retry} />;
    }

    if (token === undefined) {
      return <ShowLoading />;
    }

    return (
      <BotProtectionConfigurationScreen1 appID={appID} secretToken={token} />
    );
  };

export default BotProtectionConfigurationScreen;