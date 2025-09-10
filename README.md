# Query AWS daily cost and send report to Slack at 8pm

This app uses Go to query AWS Cost Explorer and send a daily report to Slack.

## Setup

1. Set up AWS credentials (via environment or config file).
2. Set the Slack webhook URL as an environment variable:
   - `SLACK_WEBHOOK_URL`
3. Install dependencies:
   - Run `go mod tidy`
4. Build and run:
   - `go run main.go`

## Scheduling

To run daily at 8pm, use Windows Task Scheduler or a cron job (on Linux).

Example (Windows Task Scheduler):
- Trigger: Daily at 8:00 PM
- Action: Start a program
- Program/script: `pwsh.exe`
- Add arguments: `-Command "cd 'c:\Users\Hiep\Desktop\askc'; go run main.go"`

## Customization
- You can modify the report format or Slack message in `main.go`.
