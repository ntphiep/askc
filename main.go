package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    ce "github.com/aws/aws-sdk-go-v2/service/costexplorer"
    cetypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func main() {
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    client := ce.NewFromConfig(cfg)

    input := &ce.GetCostAndUsageWithResourcesInput{
        TimePeriod: &cetypes.DateInterval{
            Start: aws.String("2025-08-30"), // inclusive
            End:   aws.String("2025-09-09"), // exclusive
        },
        Granularity: cetypes.GranularityDaily,
        Metrics:     []string{"UnblendedCost"},
        GroupBy: []cetypes.GroupDefinition{
            {
                Type: cetypes.GroupDefinitionTypeDimension,
                Key:  aws.String("SERVICE"), // cái key này khác so với SERVICE
            },
        },
		Filter: &cetypes.Expression{
			Dimensions: &cetypes.DimensionValues{
				Key:    cetypes.DimensionLinkedAccount,
				Values: []string{"014498663963"},
			},
		},
	}

    resp, err := client.GetCostAndUsageWithResources(context.Background(), input)
    if err != nil {
        log.Fatal(err)
    }

    for _, day := range resp.ResultsByTime {
        fmt.Println("Period:", *day.TimePeriod.Start, "-", *day.TimePeriod.End)
        for _, group := range day.Groups {
            fmt.Printf("Resource: %v | Cost: %v\n", group.Keys, group.Metrics["UnblendedCost"].Amount)
        }
    }
}
