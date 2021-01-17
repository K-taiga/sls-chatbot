package main

// キーがstringでバリューもstringの階層型の連想配列をマッピング
// ["IAM Access Analyzer":["CRITICAL", "HIGH", "MEDIUM"]]
var filters = map[string][]string{
	"IAM Access Analyzer": []string{"CRITICAL", "HIGH", "MEDIUM"},
	"GuardDuty":           []string{"CRITICAL", "HIGH"},
	"Security Hub":        []string{"CRITICAL"},
}

func filter(findings *[]Finding) *[]Finding {
	result := []Finding{}
	for _, finding := range *findings {
		// findingの中にfiletersのキーのプロダクトがある
		if severities, ok := filters[finding.ProductName]; ok {
			// filtersのセキュリティの重要度と合致するものがあれば返却
			// 重要度が入った配列と実際のfidingのセキュリティレベルの文字列を渡す
			if contains(severities, finding.SeverityLabel) {
				result = append(result, finding)
			}
		}
	}
	return &result
}

func contains(items []string, str string) bool {
	for _, item := range items {
		if item == str {
			return true
		}
	}
	return false
}
