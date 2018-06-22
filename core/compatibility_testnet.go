package core

// CompatibilityTestNet ..
type CompatibilityTestNet struct {
	transferFromContractEventRecordableHeight uint64

	acceptFuncAvailableHeight uint64

	randomAvailableHeight uint64

	dateAvailableHeight uint64

	recordCallContractResultHeight uint64

	nvmMemoryLimitWithoutInjectHeight uint64

	wsResetRecordDependencyHeight uint64

	v8JSLibVersionControlHeight uint64

	transferFromContractFailureEventRecordableHeight uint64

	newNvmExeTimeoutConsumeGasHeight uint64

	v8JSLibVersionHeightMap *V8JSLibVersionHeightMap
}

// NewCompatibilityTestNet ..
func NewCompatibilityTestNet() Compatibility {
	return &CompatibilityTestNet{
		transferFromContractEventRecordableHeight:        199666,
		acceptFuncAvailableHeight:                        199666,
		randomAvailableHeight:                            199666,
		dateAvailableHeight:                              199666,
		recordCallContractResultHeight:                   199666,
		nvmMemoryLimitWithoutInjectHeight:                281600,
		wsResetRecordDependencyHeight:                    281600,
		v8JSLibVersionControlHeight:                      460000,
		transferFromContractFailureEventRecordableHeight: 460000,
		newNvmExeTimeoutConsumeGasHeight:                 460000,
		v8JSLibVersionHeightMap: &V8JSLibVersionHeightMap{
			Data: map[string]uint64{
				"1.0.5": 460000, // v8JSLibVersionControlHeight
				"1.1.0": 500000,
			},
			DescKeys: []string{"1.1.0", "1.0.5"},
		},
	}
}

// TransferFromContractEventRecordableHeight ..
func (c *CompatibilityTestNet) TransferFromContractEventRecordableHeight() uint64 {
	return c.transferFromContractEventRecordableHeight
}

// AcceptFuncAvailableHeight ..
func (c *CompatibilityTestNet) AcceptFuncAvailableHeight() uint64 {
	return c.acceptFuncAvailableHeight
}

// RandomAvailableHeight ..
func (c *CompatibilityTestNet) RandomAvailableHeight() uint64 {
	return c.randomAvailableHeight
}

// DateAvailableHeight ..
func (c *CompatibilityTestNet) DateAvailableHeight() uint64 {
	return c.dateAvailableHeight
}

// RecordCallContractResultHeight ..
func (c *CompatibilityTestNet) RecordCallContractResultHeight() uint64 {
	return c.recordCallContractResultHeight
}

// NvmMemoryLimitWithoutInjectHeight ..
func (c *CompatibilityTestNet) NvmMemoryLimitWithoutInjectHeight() uint64 {
	return c.nvmMemoryLimitWithoutInjectHeight
}

// WsResetRecordDependencyHeight ..
func (c *CompatibilityTestNet) WsResetRecordDependencyHeight() uint64 {
	return c.wsResetRecordDependencyHeight
}

// V8JSLibVersionControlHeight ..
func (c *CompatibilityTestNet) V8JSLibVersionControlHeight() uint64 {
	return c.v8JSLibVersionControlHeight
}

// TransferFromContractFailureEventRecordableHeight ..
func (c *CompatibilityTestNet) TransferFromContractFailureEventRecordableHeight() uint64 {
	return c.transferFromContractFailureEventRecordableHeight
}

// NewNvmExeTimeoutConsumeGasHeight ..
func (c *CompatibilityTestNet) NewNvmExeTimeoutConsumeGasHeight() uint64 {
	return c.newNvmExeTimeoutConsumeGasHeight
}

// V8JSLibVersionHeightMap ..
func (c *CompatibilityTestNet) V8JSLibVersionHeightMap() *V8JSLibVersionHeightMap {
	return c.v8JSLibVersionHeightMap
}
