// Code generated by scripts/dbgen/main.go. DO NOT EDIT.
package database

// ForeignKeyConstraint represents a named foreign key constraint on a table.
type ForeignKeyConstraint string

// ForeignKeyConstraint enums.
const (
	ForeignKeyAPIKeysUserIDUUID                             ForeignKeyConstraint = "api_keys_user_id_uuid_fkey"                               // ALTER TABLE ONLY api_keys ADD CONSTRAINT api_keys_user_id_uuid_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyFrobulatorsOrgID                              ForeignKeyConstraint = "frobulators_org_id_fkey"                                  // ALTER TABLE ONLY frobulators ADD CONSTRAINT frobulators_org_id_fkey FOREIGN KEY (org_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyFrobulatorsUserID                             ForeignKeyConstraint = "frobulators_user_id_fkey"                                 // ALTER TABLE ONLY frobulators ADD CONSTRAINT frobulators_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyGitAuthLinksOauthAccessTokenKeyID             ForeignKeyConstraint = "git_auth_links_oauth_access_token_key_id_fkey"            // ALTER TABLE ONLY external_auth_links ADD CONSTRAINT git_auth_links_oauth_access_token_key_id_fkey FOREIGN KEY (oauth_access_token_key_id) REFERENCES dbcrypt_keys(active_key_digest);
	ForeignKeyGitAuthLinksOauthRefreshTokenKeyID            ForeignKeyConstraint = "git_auth_links_oauth_refresh_token_key_id_fkey"           // ALTER TABLE ONLY external_auth_links ADD CONSTRAINT git_auth_links_oauth_refresh_token_key_id_fkey FOREIGN KEY (oauth_refresh_token_key_id) REFERENCES dbcrypt_keys(active_key_digest);
	ForeignKeyGitSSHKeysUserID                              ForeignKeyConstraint = "gitsshkeys_user_id_fkey"                                  // ALTER TABLE ONLY gitsshkeys ADD CONSTRAINT gitsshkeys_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);
	ForeignKeyGroupMembersGroupID                           ForeignKeyConstraint = "group_members_group_id_fkey"                              // ALTER TABLE ONLY group_members ADD CONSTRAINT group_members_group_id_fkey FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE;
	ForeignKeyGroupMembersUserID                            ForeignKeyConstraint = "group_members_user_id_fkey"                               // ALTER TABLE ONLY group_members ADD CONSTRAINT group_members_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyGroupsOrganizationID                          ForeignKeyConstraint = "groups_organization_id_fkey"                              // ALTER TABLE ONLY groups ADD CONSTRAINT groups_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyJfrogXrayScansAgentID                         ForeignKeyConstraint = "jfrog_xray_scans_agent_id_fkey"                           // ALTER TABLE ONLY jfrog_xray_scans ADD CONSTRAINT jfrog_xray_scans_agent_id_fkey FOREIGN KEY (agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyJfrogXrayScansWorkspaceID                     ForeignKeyConstraint = "jfrog_xray_scans_workspace_id_fkey"                       // ALTER TABLE ONLY jfrog_xray_scans ADD CONSTRAINT jfrog_xray_scans_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
	ForeignKeyNotificationMessagesNotificationTemplateID    ForeignKeyConstraint = "notification_messages_notification_template_id_fkey"      // ALTER TABLE ONLY notification_messages ADD CONSTRAINT notification_messages_notification_template_id_fkey FOREIGN KEY (notification_template_id) REFERENCES notification_templates(id) ON DELETE CASCADE;
	ForeignKeyNotificationMessagesUserID                    ForeignKeyConstraint = "notification_messages_user_id_fkey"                       // ALTER TABLE ONLY notification_messages ADD CONSTRAINT notification_messages_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyNotificationPreferencesNotificationTemplateID ForeignKeyConstraint = "notification_preferences_notification_template_id_fkey"   // ALTER TABLE ONLY notification_preferences ADD CONSTRAINT notification_preferences_notification_template_id_fkey FOREIGN KEY (notification_template_id) REFERENCES notification_templates(id) ON DELETE CASCADE;
	ForeignKeyNotificationPreferencesUserID                 ForeignKeyConstraint = "notification_preferences_user_id_fkey"                    // ALTER TABLE ONLY notification_preferences ADD CONSTRAINT notification_preferences_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyOauth2ProviderAppCodesAppID                   ForeignKeyConstraint = "oauth2_provider_app_codes_app_id_fkey"                    // ALTER TABLE ONLY oauth2_provider_app_codes ADD CONSTRAINT oauth2_provider_app_codes_app_id_fkey FOREIGN KEY (app_id) REFERENCES oauth2_provider_apps(id) ON DELETE CASCADE;
	ForeignKeyOauth2ProviderAppCodesUserID                  ForeignKeyConstraint = "oauth2_provider_app_codes_user_id_fkey"                   // ALTER TABLE ONLY oauth2_provider_app_codes ADD CONSTRAINT oauth2_provider_app_codes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyOauth2ProviderAppSecretsAppID                 ForeignKeyConstraint = "oauth2_provider_app_secrets_app_id_fkey"                  // ALTER TABLE ONLY oauth2_provider_app_secrets ADD CONSTRAINT oauth2_provider_app_secrets_app_id_fkey FOREIGN KEY (app_id) REFERENCES oauth2_provider_apps(id) ON DELETE CASCADE;
	ForeignKeyOauth2ProviderAppTokensAPIKeyID               ForeignKeyConstraint = "oauth2_provider_app_tokens_api_key_id_fkey"               // ALTER TABLE ONLY oauth2_provider_app_tokens ADD CONSTRAINT oauth2_provider_app_tokens_api_key_id_fkey FOREIGN KEY (api_key_id) REFERENCES api_keys(id) ON DELETE CASCADE;
	ForeignKeyOauth2ProviderAppTokensAppSecretID            ForeignKeyConstraint = "oauth2_provider_app_tokens_app_secret_id_fkey"            // ALTER TABLE ONLY oauth2_provider_app_tokens ADD CONSTRAINT oauth2_provider_app_tokens_app_secret_id_fkey FOREIGN KEY (app_secret_id) REFERENCES oauth2_provider_app_secrets(id) ON DELETE CASCADE;
	ForeignKeyOrganizationMembersOrganizationIDUUID         ForeignKeyConstraint = "organization_members_organization_id_uuid_fkey"           // ALTER TABLE ONLY organization_members ADD CONSTRAINT organization_members_organization_id_uuid_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyOrganizationMembersUserIDUUID                 ForeignKeyConstraint = "organization_members_user_id_uuid_fkey"                   // ALTER TABLE ONLY organization_members ADD CONSTRAINT organization_members_user_id_uuid_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyParameterSchemasJobID                         ForeignKeyConstraint = "parameter_schemas_job_id_fkey"                            // ALTER TABLE ONLY parameter_schemas ADD CONSTRAINT parameter_schemas_job_id_fkey FOREIGN KEY (job_id) REFERENCES provisioner_jobs(id) ON DELETE CASCADE;
	ForeignKeyProvisionerDaemonsOrganizationID              ForeignKeyConstraint = "provisioner_daemons_organization_id_fkey"                 // ALTER TABLE ONLY provisioner_daemons ADD CONSTRAINT provisioner_daemons_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyProvisionerJobLogsJobID                       ForeignKeyConstraint = "provisioner_job_logs_job_id_fkey"                         // ALTER TABLE ONLY provisioner_job_logs ADD CONSTRAINT provisioner_job_logs_job_id_fkey FOREIGN KEY (job_id) REFERENCES provisioner_jobs(id) ON DELETE CASCADE;
	ForeignKeyProvisionerJobsOrganizationID                 ForeignKeyConstraint = "provisioner_jobs_organization_id_fkey"                    // ALTER TABLE ONLY provisioner_jobs ADD CONSTRAINT provisioner_jobs_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyProvisionerKeysOrganizationID                 ForeignKeyConstraint = "provisioner_keys_organization_id_fkey"                    // ALTER TABLE ONLY provisioner_keys ADD CONSTRAINT provisioner_keys_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyTailnetAgentsCoordinatorID                    ForeignKeyConstraint = "tailnet_agents_coordinator_id_fkey"                       // ALTER TABLE ONLY tailnet_agents ADD CONSTRAINT tailnet_agents_coordinator_id_fkey FOREIGN KEY (coordinator_id) REFERENCES tailnet_coordinators(id) ON DELETE CASCADE;
	ForeignKeyTailnetClientSubscriptionsCoordinatorID       ForeignKeyConstraint = "tailnet_client_subscriptions_coordinator_id_fkey"         // ALTER TABLE ONLY tailnet_client_subscriptions ADD CONSTRAINT tailnet_client_subscriptions_coordinator_id_fkey FOREIGN KEY (coordinator_id) REFERENCES tailnet_coordinators(id) ON DELETE CASCADE;
	ForeignKeyTailnetClientsCoordinatorID                   ForeignKeyConstraint = "tailnet_clients_coordinator_id_fkey"                      // ALTER TABLE ONLY tailnet_clients ADD CONSTRAINT tailnet_clients_coordinator_id_fkey FOREIGN KEY (coordinator_id) REFERENCES tailnet_coordinators(id) ON DELETE CASCADE;
	ForeignKeyTailnetPeersCoordinatorID                     ForeignKeyConstraint = "tailnet_peers_coordinator_id_fkey"                        // ALTER TABLE ONLY tailnet_peers ADD CONSTRAINT tailnet_peers_coordinator_id_fkey FOREIGN KEY (coordinator_id) REFERENCES tailnet_coordinators(id) ON DELETE CASCADE;
	ForeignKeyTailnetTunnelsCoordinatorID                   ForeignKeyConstraint = "tailnet_tunnels_coordinator_id_fkey"                      // ALTER TABLE ONLY tailnet_tunnels ADD CONSTRAINT tailnet_tunnels_coordinator_id_fkey FOREIGN KEY (coordinator_id) REFERENCES tailnet_coordinators(id) ON DELETE CASCADE;
	ForeignKeyTemplateVersionParametersTemplateVersionID    ForeignKeyConstraint = "template_version_parameters_template_version_id_fkey"     // ALTER TABLE ONLY template_version_parameters ADD CONSTRAINT template_version_parameters_template_version_id_fkey FOREIGN KEY (template_version_id) REFERENCES template_versions(id) ON DELETE CASCADE;
	ForeignKeyTemplateVersionVariablesTemplateVersionID     ForeignKeyConstraint = "template_version_variables_template_version_id_fkey"      // ALTER TABLE ONLY template_version_variables ADD CONSTRAINT template_version_variables_template_version_id_fkey FOREIGN KEY (template_version_id) REFERENCES template_versions(id) ON DELETE CASCADE;
	ForeignKeyTemplateVersionWorkspaceTagsTemplateVersionID ForeignKeyConstraint = "template_version_workspace_tags_template_version_id_fkey" // ALTER TABLE ONLY template_version_workspace_tags ADD CONSTRAINT template_version_workspace_tags_template_version_id_fkey FOREIGN KEY (template_version_id) REFERENCES template_versions(id) ON DELETE CASCADE;
	ForeignKeyTemplateVersionsCreatedBy                     ForeignKeyConstraint = "template_versions_created_by_fkey"                        // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;
	ForeignKeyTemplateVersionsOrganizationID                ForeignKeyConstraint = "template_versions_organization_id_fkey"                   // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyTemplateVersionsTemplateID                    ForeignKeyConstraint = "template_versions_template_id_fkey"                       // ALTER TABLE ONLY template_versions ADD CONSTRAINT template_versions_template_id_fkey FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE;
	ForeignKeyTemplatesCreatedBy                            ForeignKeyConstraint = "templates_created_by_fkey"                                // ALTER TABLE ONLY templates ADD CONSTRAINT templates_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT;
	ForeignKeyTemplatesOrganizationID                       ForeignKeyConstraint = "templates_organization_id_fkey"                           // ALTER TABLE ONLY templates ADD CONSTRAINT templates_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE;
	ForeignKeyUserLinksOauthAccessTokenKeyID                ForeignKeyConstraint = "user_links_oauth_access_token_key_id_fkey"                // ALTER TABLE ONLY user_links ADD CONSTRAINT user_links_oauth_access_token_key_id_fkey FOREIGN KEY (oauth_access_token_key_id) REFERENCES dbcrypt_keys(active_key_digest);
	ForeignKeyUserLinksOauthRefreshTokenKeyID               ForeignKeyConstraint = "user_links_oauth_refresh_token_key_id_fkey"               // ALTER TABLE ONLY user_links ADD CONSTRAINT user_links_oauth_refresh_token_key_id_fkey FOREIGN KEY (oauth_refresh_token_key_id) REFERENCES dbcrypt_keys(active_key_digest);
	ForeignKeyUserLinksUserID                               ForeignKeyConstraint = "user_links_user_id_fkey"                                  // ALTER TABLE ONLY user_links ADD CONSTRAINT user_links_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentLogSourcesWorkspaceAgentID      ForeignKeyConstraint = "workspace_agent_log_sources_workspace_agent_id_fkey"      // ALTER TABLE ONLY workspace_agent_log_sources ADD CONSTRAINT workspace_agent_log_sources_workspace_agent_id_fkey FOREIGN KEY (workspace_agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentMetadataWorkspaceAgentID        ForeignKeyConstraint = "workspace_agent_metadata_workspace_agent_id_fkey"         // ALTER TABLE ONLY workspace_agent_metadata ADD CONSTRAINT workspace_agent_metadata_workspace_agent_id_fkey FOREIGN KEY (workspace_agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentPortShareWorkspaceID            ForeignKeyConstraint = "workspace_agent_port_share_workspace_id_fkey"             // ALTER TABLE ONLY workspace_agent_port_share ADD CONSTRAINT workspace_agent_port_share_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentScriptsWorkspaceAgentID         ForeignKeyConstraint = "workspace_agent_scripts_workspace_agent_id_fkey"          // ALTER TABLE ONLY workspace_agent_scripts ADD CONSTRAINT workspace_agent_scripts_workspace_agent_id_fkey FOREIGN KEY (workspace_agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentStartupLogsAgentID              ForeignKeyConstraint = "workspace_agent_startup_logs_agent_id_fkey"               // ALTER TABLE ONLY workspace_agent_logs ADD CONSTRAINT workspace_agent_startup_logs_agent_id_fkey FOREIGN KEY (agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAgentsResourceID                     ForeignKeyConstraint = "workspace_agents_resource_id_fkey"                        // ALTER TABLE ONLY workspace_agents ADD CONSTRAINT workspace_agents_resource_id_fkey FOREIGN KEY (resource_id) REFERENCES workspace_resources(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceAppStatsAgentID                      ForeignKeyConstraint = "workspace_app_stats_agent_id_fkey"                        // ALTER TABLE ONLY workspace_app_stats ADD CONSTRAINT workspace_app_stats_agent_id_fkey FOREIGN KEY (agent_id) REFERENCES workspace_agents(id);
	ForeignKeyWorkspaceAppStatsUserID                       ForeignKeyConstraint = "workspace_app_stats_user_id_fkey"                         // ALTER TABLE ONLY workspace_app_stats ADD CONSTRAINT workspace_app_stats_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);
	ForeignKeyWorkspaceAppStatsWorkspaceID                  ForeignKeyConstraint = "workspace_app_stats_workspace_id_fkey"                    // ALTER TABLE ONLY workspace_app_stats ADD CONSTRAINT workspace_app_stats_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES workspaces(id);
	ForeignKeyWorkspaceAppsAgentID                          ForeignKeyConstraint = "workspace_apps_agent_id_fkey"                             // ALTER TABLE ONLY workspace_apps ADD CONSTRAINT workspace_apps_agent_id_fkey FOREIGN KEY (agent_id) REFERENCES workspace_agents(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceBuildParametersWorkspaceBuildID      ForeignKeyConstraint = "workspace_build_parameters_workspace_build_id_fkey"       // ALTER TABLE ONLY workspace_build_parameters ADD CONSTRAINT workspace_build_parameters_workspace_build_id_fkey FOREIGN KEY (workspace_build_id) REFERENCES workspace_builds(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceBuildsJobID                          ForeignKeyConstraint = "workspace_builds_job_id_fkey"                             // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_job_id_fkey FOREIGN KEY (job_id) REFERENCES provisioner_jobs(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceBuildsTemplateVersionID              ForeignKeyConstraint = "workspace_builds_template_version_id_fkey"                // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_template_version_id_fkey FOREIGN KEY (template_version_id) REFERENCES template_versions(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceBuildsWorkspaceID                    ForeignKeyConstraint = "workspace_builds_workspace_id_fkey"                       // ALTER TABLE ONLY workspace_builds ADD CONSTRAINT workspace_builds_workspace_id_fkey FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceResourceMetadataWorkspaceResourceID  ForeignKeyConstraint = "workspace_resource_metadata_workspace_resource_id_fkey"   // ALTER TABLE ONLY workspace_resource_metadata ADD CONSTRAINT workspace_resource_metadata_workspace_resource_id_fkey FOREIGN KEY (workspace_resource_id) REFERENCES workspace_resources(id) ON DELETE CASCADE;
	ForeignKeyWorkspaceResourcesJobID                       ForeignKeyConstraint = "workspace_resources_job_id_fkey"                          // ALTER TABLE ONLY workspace_resources ADD CONSTRAINT workspace_resources_job_id_fkey FOREIGN KEY (job_id) REFERENCES provisioner_jobs(id) ON DELETE CASCADE;
	ForeignKeyWorkspacesOrganizationID                      ForeignKeyConstraint = "workspaces_organization_id_fkey"                          // ALTER TABLE ONLY workspaces ADD CONSTRAINT workspaces_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE RESTRICT;
	ForeignKeyWorkspacesOwnerID                             ForeignKeyConstraint = "workspaces_owner_id_fkey"                                 // ALTER TABLE ONLY workspaces ADD CONSTRAINT workspaces_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE RESTRICT;
	ForeignKeyWorkspacesTemplateID                          ForeignKeyConstraint = "workspaces_template_id_fkey"                              // ALTER TABLE ONLY workspaces ADD CONSTRAINT workspaces_template_id_fkey FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE RESTRICT;
)
