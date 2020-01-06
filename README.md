
# faas stonks

This uses serverless technology to get a stock quote and 30 day sparkline from Yahoo Finance.

# Deployment

- Nimbella account
- Namespace with object storage enabled
- Deploy command: `nim project deploy stonks --remote-build`

> Note: the `--remote-build` is required to build this Go function into a binary before deploying the action. 
