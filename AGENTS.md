# Agent Notes

This repository holds latency-app, the small sample web service the AstroPulse examples build and deploy, and the `.astropulse/` files that deploy it. These notes apply to every person and coding agent that changes it.

## This repository is public

Every commit, branch name, pull request, issue, comment and file here is world-readable and permanent. Write all of it for an outside reader who uses the examples.

Never include:

- **Internal references.** No internal project, repository, issue or pull request names or numbers. No release plans, release checklists, internal milestones or statements about what production does or accepts.
- **Identifiers.** No real AWS account IDs, ARNs, project IDs, subscription IDs, organization IDs, hostnames, domains, IP addresses, cluster names, or details of local or test environments (kind, a laptop, a test workspace).
- **Secrets.** No credentials, tokens, keys, kubeconfigs or registry passwords, not even expired ones.
- **Tool attribution.** No `Co-Authored-By` lines naming a tool, no "Generated with" footers and no session links.

Use placeholders a reader replaces:

| Kind | Placeholder |
|------|-------------|
| AWS account ID | `123456789012` |
| Azure subscription ID | `00000000-0000-0000-0000-000000000000` |
| GCP project | `your-project` |
| Domain | `your-domain.com` |
| Cluster | `your-cluster` |
| Container registry | `ghcr.io/your-org` |

## Pull requests and commits

- A description says what changed in the examples and why, in terms a reader of this repository understands.
- If a change depends on a future AstroPulse release, keep that dependency out of the repository. The person who merges it tracks the timing.
- Commit as the human you work for, with a conventional subject under 50 characters.
- Before you open or merge a pull request, scan the whole diff and its description for everything listed above.

## The sample

- `.astropulse/` holds the files a reader applies to deploy this service with AstroPulse. They must apply as written against the current AstroPulse API, with placeholders as the only changes a reader makes.
- Use the API's own spellings for enum values (for example `Image`, `OnDemand`).
- Keep the service small and dependency-light: it exists to be built, deployed, updated and rolled back in the examples.
