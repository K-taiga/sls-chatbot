package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/sts"
)

func describeSpend() (*budgets.CalculatedSpend, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed NewSession: %s", err)
	}

	// SecurityTokenServiceで認証する
	identityInput := &sts.GetCallerIdentityInput{}
	callerIdentity, err := sts.New(sess).GetCallerIdentity(identityInput)
	if err != nil {
		return nil, fmt.Errorf("failed GetCallerIdentity: %s", err)
	}
	// AWSアカウントIDを取得
	accountId := callerIdentity.Account

	// AWSBudgetsから予算の一覧を取得
	input := &budgets.DescribeBudgetsInput{AccountId: accountId}
	output, err := budgets.New(sess).DescribeBudgets(input)
	if err != nil {
		return nil, fmt.Errorf("failed DescribeBudgets: %s", err)
	}

	for _, budget := range output.Budgets {
		// 月単位のコスト予算を取得
		if *budget.BudgetType == "COST" && *budget.TimeUnit == "MONTHLY" {
			return budget.CalculatedSpend, nil
		}
	}
	return nil, fmt.Errorf("not found budget: %s", output.Budgets)
}
