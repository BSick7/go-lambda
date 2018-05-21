package cloudwatch

import (
	"fmt"

	"github.com/BSick7/go-lambda/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type MetricDatum interface {
	ToAWS() *cloudwatch.MetricDatum
}

func PutMetrics(namespace string, values []MetricDatum) error {
	svc, err := services.NewCloudwatch()
	if err != nil {
		return fmt.Errorf("could not create cloudwatch service: %s", err)
	}

	data := make([]*cloudwatch.MetricDatum, 0)
	for _, item := range values {
		data = append(data, item.ToAWS())
	}
	input := &cloudwatch.PutMetricDataInput{
		Namespace:  aws.String(namespace),
		MetricData: data,
	}
	_, err = svc.PutMetricData(input)
	return err
}
