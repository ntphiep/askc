package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func GetDailyCost() (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("unable to load AWS config: %w", err)
	}
	ce := costexplorer.NewFromConfig(cfg)

	end := time.Now()
	start := end.AddDate(0, 0, -1)
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: ptr(start.Format("2006-01-02")),
			End:   ptr(end.Format("2006-01-02")),
		},
		Granularity: types.GranularityDaily,
		Metrics:     []string{"UnblendedCost"},
	}
	
	result, err := ce.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to get cost: %w", err)
	}

	if len(result.ResultsByTime) == 0 {
		return "No cost data found.", nil
	}

	cost := result.ResultsByTime[0].Total["UnblendedCost"].Amount
	var costStr string
	if cost != nil {
		costStr = *cost
	} else {
		costStr = "0.00"
	}
	
	return fmt.Sprintf("AWS Daily Cost for %s: $%s", start.Format("2006-01-02"), costStr), nil
}

// ptr returns a pointer to a string
func ptr(s string) *string {
	return &s
}
