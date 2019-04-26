# Portunus

Portunus was the ancient Roman god of keys and doors. However, this repo does not
contain the god. It contains Portunus, a small and self-contained user/group
management and authentication service.

TODO explain more

## Environment variables

| Variable | Default | Explanation |
| -------- | ------- | ----------- |
| `PORTUNUS_LDAP_SUFFIX` | *(required)* | The DN of the topmost entry in your LDAP directory. Must currently be a sequence of `dc=xxx` RDNs. (This requirement may be lifted in future versions.) See [*LDAP directory structure*](#ldap-directory-structure) for details. |
| `PORTUNUS_SLAPD_BINARY` | `slapd` | Where to find the binary of slapd (the OpenLDAP server). Semantics match those of `execvp(3)`: If the supplied value is not a path containing slashes, `$PATH` will be searched for it. |
| `PORTUNUS_SLAPD_SCHEMA_DIR` | `/etc/openldap/schema` | Where to find OpenLDAP's schema definitions. |
| `XDG_RUNTIME_DIR` | *(required)* | Portunus will store all runtime state in the directory `$XDG_RUNTIME_DIR/portunus`. If `XDG_RUNTIME_DIR` is not set by the login manager, set it to a directory on a tmpfs that is writable by the current user. When starting Portunus with root privileges, `XDG_RUNTIME_DIR=/run` is a sensible choice. |
