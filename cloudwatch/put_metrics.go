package cloudwatch

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var (
	maxMetricsPerPut = 20
)

func PutMetrics(namespace string, values []MetricPoint) error {
	svc, err := services.NewCloudwatch()
	if err != nil {
		return fmt.Errorf("could not create cloudwatch service: %s", err)
	}

	batch := make([]*cloudwatch.MetricDatum, 0)
	for _, item := range values {
		if len(batch) >= maxMetricsPerPut {
			input := &cloudwatch.PutMetricDataInput{
				Namespace:  aws.String(namespace),
				MetricData: batch,
			}
			if _, err := svc.PutMetricData(input); err != nil {
				return err
			}
			batch = []*cloudwatch.MetricDatum{item.ToAWS()}
		} else {
			batch = append(batch, item.ToAWS())
		}
	}
	if len(batch) > 0 {
		input := &cloudwatch.PutMetricDataInput{
			Namespace:  aws.String(namespace),
			MetricData: batch,
		}
		_, err := svc.PutMetricData(input)
		return err
	}
	return nil
}
