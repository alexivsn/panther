package api

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/panther-labs/panther/api/lambda/alerts/models"
	"github.com/panther-labs/panther/internal/log_analysis/alerts_api/table"
)

var (
	timeInTest = time.Now()

	alertItems = []*table.AlertItem{
		{
			RuleID:       "ruleId",
			AlertID:      "alertId",
			UpdateTime:   timeInTest,
			CreationTime: timeInTest,
			Severity:     "INFO",
			DedupString:  "dedupString",
			LogTypes:     []string{"AWS.CloudTrail"},
			EventCount:   100,
		},
	}

	expectedAlertSummary = []*models.AlertSummary{
		{
			RuleID:        aws.String("ruleId"),
			AlertID:       aws.String("alertId"),
			UpdateTime:    aws.Time(timeInTest),
			CreationTime:  aws.Time(timeInTest),
			Severity:      aws.String("INFO"),
			DedupString:   aws.String("dedupString"),
			EventsMatched: aws.Int(100),
		},
	}
)

func TestListAlertsForRule(t *testing.T) {
	tableMock := &tableMock{}
	alertsDB = tableMock

	input := &models.ListAlertsInput{
		RuleID:            aws.String("ruleId"),
		PageSize:          aws.Int(10),
		ExclusiveStartKey: aws.String("startKey"),
	}

	tableMock.On("ListByRule", "ruleId", aws.String("startKey"), aws.Int(10)).
		Return(alertItems, aws.String("lastKey"), nil)
	result, err := API{}.ListAlerts(input)
	require.NoError(t, err)

	assert.Equal(t, &models.ListAlertsOutput{
		Alerts:           expectedAlertSummary,
		LastEvaluatedKey: aws.String("lastKey"),
	}, result)
}

func TestListAllAlerts(t *testing.T) {
	tableMock := &tableMock{}
	alertsDB = tableMock

	input := &models.ListAlertsInput{
		PageSize:          aws.Int(10),
		ExclusiveStartKey: aws.String("startKey"),
	}

	tableMock.On("ListAll", aws.String("startKey"), aws.Int(10)).
		Return(alertItems, aws.String("lastKey"), nil)
	result, err := API{}.ListAlerts(input)
	require.NoError(t, err)

	assert.Equal(t, &models.ListAlertsOutput{
		Alerts:           expectedAlertSummary,
		LastEvaluatedKey: aws.String("lastKey"),
	}, result)
}
