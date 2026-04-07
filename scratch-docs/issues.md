# Prod version

Prod docker setup
- remove dev tools (like air)
- remove database from compose
- have a base config, an override config for dev and a prod config for prod

Database need to live outside of the api container (RDS for example)
- Implement the config to point to the prod DB
- When the connection fail in the prod env, DATABASE_URL should not fallback to localhost, but throw an informative error
- Migrations should live in a deploy pipeline
- Review security and SSL connection

Security
- Store prod secrets safely (e.g. AWS Secrets Manager)

