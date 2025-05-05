# Smart Guardian

Smart Guardian is a DeFi Risk Monitoring & Alerting Platform written in Go. It continuously tracks on-chain LTV, utilization, APR, and slippage across lending, trading, and vault protocols, automatically notifying you of potential risks in your positions.

---

## 🚀 Why Smart Guardian?

- **Proactive Risk Management**: Get real-time alerts on liquidation risks, utilization spikes, and unusual volatility before they impact your portfolio.
- **Modular & Extensible**: Designed with a clear package layout—easily plug in new protocols, add GraphQL/gRPC endpoints, or integrate Kafka for streaming analytics.
- **Built for Developers**: Leverage Go’s strong typing, goroutines, and channels to build reliable, high-performance monitoring services.
- **Full DevOps Support**: Includes configuration loading (Viper), structured logging (Zap), automated database migrations, and CI/CD workflows for seamless deployments.

---

## ⚙️ Key Features

- **On-Chain Data Integration**: Fetch metrics via JSON-RPC or TheGraph in parallel worker pools.
- **Customizable Alert Rules**: Define thresholds on TVL, APR, LTV, and more, with email/webhook notifications.
- **CLI & HTTP API**: Command-line tools for migrations/backfills, RESTful endpoints for metrics and alerts, plus Server-Sent Events/Webhooks for live updates.
- **Robust Configuration**: Centralized settings with Viper, environment-based overrides, and sensible defaults.
- **Structured Logging & Error Handling**: Zap-powered logs and wrapped errors for clear diagnostics.
- **Testing & CI**: Unit and integration tests, Golang best-practice linting, formatting, vetting, and GitHub Actions pipeline.

---

## 🛠️ Getting Started

### Prerequisites

- Go 1.20+ installed
- PostgreSQL database
- (Optional) Redis for caching

### Installation

```bash
git clone https://github.com/yourusername/smart-guardian.git
cd smart-guardian
go mod tidy
cp .env.example .env
