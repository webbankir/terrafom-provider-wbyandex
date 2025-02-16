// Code generated by protoc-gen-goext. DO NOT EDIT.

package billing

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type Budget_BudgetSpec = isBudget_BudgetSpec

func (m *Budget) SetBudgetSpec(v Budget_BudgetSpec) {
	m.BudgetSpec = v
}

func (m *Budget) SetId(v string) {
	m.Id = v
}

func (m *Budget) SetName(v string) {
	m.Name = v
}

func (m *Budget) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Budget) SetBillingAccountId(v string) {
	m.BillingAccountId = v
}

func (m *Budget) SetStatus(v BudgetStatus) {
	m.Status = v
}

func (m *Budget) SetCostBudget(v *CostBudgetSpec) {
	m.BudgetSpec = &Budget_CostBudget{
		CostBudget: v,
	}
}

func (m *Budget) SetExpenseBudget(v *ExpenseBudgetSpec) {
	m.BudgetSpec = &Budget_ExpenseBudget{
		ExpenseBudget: v,
	}
}

func (m *Budget) SetBalanceBudget(v *BalanceBudgetSpec) {
	m.BudgetSpec = &Budget_BalanceBudget{
		BalanceBudget: v,
	}
}

type CostBudgetSpec_StartType = isCostBudgetSpec_StartType

func (m *CostBudgetSpec) SetStartType(v CostBudgetSpec_StartType) {
	m.StartType = v
}

func (m *CostBudgetSpec) SetAmount(v string) {
	m.Amount = v
}

func (m *CostBudgetSpec) SetNotificationUserAccountIds(v []string) {
	m.NotificationUserAccountIds = v
}

func (m *CostBudgetSpec) SetThresholdRules(v []*ThresholdRule) {
	m.ThresholdRules = v
}

func (m *CostBudgetSpec) SetFilter(v *ConsumptionFilter) {
	m.Filter = v
}

func (m *CostBudgetSpec) SetResetPeriod(v ResetPeriodType) {
	m.StartType = &CostBudgetSpec_ResetPeriod{
		ResetPeriod: v,
	}
}

func (m *CostBudgetSpec) SetStartDate(v string) {
	m.StartType = &CostBudgetSpec_StartDate{
		StartDate: v,
	}
}

func (m *CostBudgetSpec) SetEndDate(v string) {
	m.EndDate = v
}

type ExpenseBudgetSpec_StartType = isExpenseBudgetSpec_StartType

func (m *ExpenseBudgetSpec) SetStartType(v ExpenseBudgetSpec_StartType) {
	m.StartType = v
}

func (m *ExpenseBudgetSpec) SetAmount(v string) {
	m.Amount = v
}

func (m *ExpenseBudgetSpec) SetNotificationUserAccountIds(v []string) {
	m.NotificationUserAccountIds = v
}

func (m *ExpenseBudgetSpec) SetThresholdRules(v []*ThresholdRule) {
	m.ThresholdRules = v
}

func (m *ExpenseBudgetSpec) SetFilter(v *ConsumptionFilter) {
	m.Filter = v
}

func (m *ExpenseBudgetSpec) SetResetPeriod(v ResetPeriodType) {
	m.StartType = &ExpenseBudgetSpec_ResetPeriod{
		ResetPeriod: v,
	}
}

func (m *ExpenseBudgetSpec) SetStartDate(v string) {
	m.StartType = &ExpenseBudgetSpec_StartDate{
		StartDate: v,
	}
}

func (m *ExpenseBudgetSpec) SetEndDate(v string) {
	m.EndDate = v
}

func (m *BalanceBudgetSpec) SetAmount(v string) {
	m.Amount = v
}

func (m *BalanceBudgetSpec) SetNotificationUserAccountIds(v []string) {
	m.NotificationUserAccountIds = v
}

func (m *BalanceBudgetSpec) SetThresholdRules(v []*ThresholdRule) {
	m.ThresholdRules = v
}

func (m *BalanceBudgetSpec) SetStartDate(v string) {
	m.StartDate = v
}

func (m *BalanceBudgetSpec) SetEndDate(v string) {
	m.EndDate = v
}

func (m *ConsumptionFilter) SetServiceIds(v []string) {
	m.ServiceIds = v
}

func (m *ConsumptionFilter) SetCloudFoldersFilters(v []*CloudFoldersConsumptionFilter) {
	m.CloudFoldersFilters = v
}

func (m *CloudFoldersConsumptionFilter) SetCloudId(v string) {
	m.CloudId = v
}

func (m *CloudFoldersConsumptionFilter) SetFolderIds(v []string) {
	m.FolderIds = v
}

func (m *ThresholdRule) SetType(v ThresholdType) {
	m.Type = v
}

func (m *ThresholdRule) SetAmount(v string) {
	m.Amount = v
}

func (m *ThresholdRule) SetNotificationUserAccountIds(v []string) {
	m.NotificationUserAccountIds = v
}
