# 🛡️ Security Policy

## Supported Versions

We maintain security support for the latest **major version** of OptiTerra. Previous versions may receive patches at the discretion of the maintainers.

| Version        | Supported          |
|----------------|--------------------|
| `v1.x`         | ✅ Active support   |
| `< v1.0.0`     | ❌ No longer supported |

---

## 🔔 Reporting a Vulnerability

If you discover a security vulnerability **in the code, configuration, deployment, or documentation** of this project, please **report it privately**.

### 📩 Contact

- **Primary:** `security@optiterra.org`
- Backup (if unreachable): open a GitHub issue with `[SECURITY]` in the title and no sensitive info

We will respond within **72 hours** and aim to resolve verified issues in **7 working days**.

---

## 🧪 Disclosure Policy

- All security issues must be **reported privately first**.
- We will acknowledge the report and assign a severity level.
- We may ask for clarification or a proof of concept (PoC).
- Once resolved, we will publish:
  - A changelog entry
  - A GitHub Security Advisory (if applicable)
  - Credit to the reporter (optional)
  - A CVE ID (if warranted)

---

## 🔐 Security Practices

We aim to follow industry best practices, including:

- Linting, `go vet`, and race condition checks
- `gosec` (Go Security Linter) in CI/CD
- JWT auth & HTTPS-ready deployment
- Regular dependency updates (`dependabot`, `go mod tidy`)
- Minimal external dependencies
- Clear governance and CLA for contributors

---

## 🧩 Contributing Secure Code

When contributing new code or features, please:

- Avoid introducing unnecessary third-party packages.
- Use interfaces for dependency injection when testing.
- Follow the architecture guidelines in `docs/architecture.md`.
- Add tests for input validation, edge cases, and abuse scenarios.

---

## 📚 References

- [OWASP Top 10 for APIs](https://owasp.org/www-project-api-security/)
- [The Go Blog - Secure Coding Guidelines](https://blog.golang.org/security)
- [CNCF Security Best Practices](https://github.com/cncf/tag-security)

---

> Thank you for helping us protect the integrity of this project, the open-source community, and the future of sustainable technology.

