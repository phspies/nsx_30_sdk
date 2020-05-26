# LdapProbeError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | The cause of the error. CONNECTION_TIMEOUT: The connection timed out when contating the                     LDAP server. Check the hostname/ip. CONNECTION_REFUSED: The connection was refused when contacting the                     LDAP server. Ensure that the LDAP server is                     up and that you have the correct ip/hostname. STARTTLS_FAILED: Unable to use StartTLS to upgrade the                  connection to use TLS. Ensure the LDAP server                  supports TLS and if not, use LDAP or LDAPS                  as the protocol. INVALID_CREDENTIALS: The username and/or password are incorrect. THUMBPRINT_MISMATCH: A certificate thumbprint was provided in the                      LDAP server configuration, but did not match                      the certificate presented by the LDAP server. BASE_DN_NOT_FOUND: The configured base_dn does not                    exist on the LDAP server or is not                    readable. SSL_HANDSHAKE_ERROR: An error occurred while establishing a secure                      connection with the LDAP server. GENERAL_ERROR: An undetermined error occurred.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

