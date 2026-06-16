package main

type SleepService struct{}

type AppInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type DashboardSummary struct {
	LastNightUsage     string  `json:"lastNightUsage"`
	AHI                float64 `json:"ahi"`
	Leak95             float64 `json:"leak95"`
	Pressure95         float64 `json:"pressure95"`
	ImportedNightCount int     `json:"importedNightCount"`
}

func (s *SleepService) AppInfo() AppInfo {
	return AppInfo{
		Name:        "Go Sleep",
		Version:     "0.1.0",
		Description: "A local-first CPAP and sleep data dashboard.",
	}
}

func (s *SleepService) GetDashboardSummary() DashboardSummary {
	return DashboardSummary{
		LastNightUsage:     "6h 42m",
		AHI:                3.8,
		Leak95:             18.2,
		Pressure95:         16.4,
		ImportedNightCount: 0,
	}
}
