package procedure

const (
	IntoPath = "info"
)

type InfoInput struct {
}

type InfoOutput struct {
	Code                    int64           `json:"code"`
	Descr                   string          `json:"descr"`
	Account                 float64         `json:"account"`
	OrganizationDescription string          `json:"organizationDescr"`
	OrganizationCode        string          `json:"ocode"`
	Sources                 []string        `json:"source"`
	ReceiveNumbers          []string        `json:"receive_numbers,omitempty"`
	UserAccess              map[string]bool `json:"userAccess"`
}

type InfoUserAccessOutput struct {
	CanExportContacts            bool `json:"canExportContacts,omitempty"`
	СanDeleteGroups              bool `json:"canDeleteGroups,omitempty"`
	СanDeleteContacts            bool `json:"canDeleteContacts,omitempty"`
	CanConfig                    bool `json:"canConfig,omitempty"`
	CanViewReportNums            bool `json:"canViewReportNums,omitempty"`
	CanViewNumbersInPhonebook    bool `json:"canViewNumbersInPhonebook,omitempty"`
	CanEditPhonebook             bool `json:"canEditPhonebook,omitempty"`
	CanEditContacts              bool `json:"canEditContacts,omitempty"`
	CanEditStopList              bool `json:"canEditStopList,omitempty"`
	CanUseAllGroups              bool `json:"canUseAllGroups,omitempty"`
	CanModifyContent             bool `json:"canModifyContent,omitempty"`
	CanDealerConfig              bool `json:"canDealerConfig,omitempty"`
	CanTarifConfig               bool `json:"canTarifConfig,omitempty"`
	CanManageAllOrgs             bool `json:"canManageAllOrgs,omitempty"`
	CanReportsCenter             bool `json:"canReportsCenter,omitempty"`
	CanViewAllReports            bool `json:"canViewAllReports,omitempty"`
	CanTarifInfo                 bool `json:"canTarifInfo,omitempty"`
	CanChangeTarif               bool `json:"canChangeTarif,omitempty"`
	CanManageInputNumbers        bool `json:"canManageInputNumbers,omitempty"`
	CanManageClientLoginPassword bool `json:"canManageClientLoginPassword,omitempty"`
	CanSendSMS                   bool `json:"canSendSMS,omitempty"`
	CanSendEmail                 bool `json:"canSendEmail,omitempty"`
	CanDopModules                bool `json:"canDopModules,omitempty"`
	CanOnline                    bool `json:"canOnline,omitempty"`
	CanMakePayment               bool `json:"canMakePayment,omitempty"`
	CanChangeUseAlfaSourceAll    bool `json:"canChangeUseAlfaSourceAll,omitempty"`
	CanPm                        bool `json:"canPm,omitempty"`
	CanOrgList                   bool `json:"canOrgList,omitempty"`
	CanOrgListExport             bool `json:"canOrgListExport,omitempty"`
	CanDutyAndSkipped            bool `json:"canDutyAndSkipped,omitempty"`
	CanUseAllSmsTemplates        bool `json:"canUseAllSmsTemplates,omitempty"`
	CanPager                     bool `json:"canPager,omitempty"`
	CanMyFinReports              bool `json:"canMyFinReports,omitempty"`
	CanDoc                       bool `json:"canDoc,omitempty"`
	CanMyEvents                  bool `json:"canMyEvents,omitempty"`
	CanViewNews                  bool `json:"canViewNews,omitempty"`
	CanConfigOnline              bool `json:"canConfigOnline,omitempty"`
	CanPcount                    bool `json:"canPcount,omitempty"`
	CanChangePassword            bool `json:"canChangePassword,omitempty"`
	CanReportsSMS                bool `json:"canReportsSMS,omitempty"`
	CanReportsEmail              bool `json:"canReportsEmail,omitempty"`
	CanDealerLeftMenu            bool `json:"canDealerLeftMenu,omitempty"`
	CanAgent                     bool `json:"canAgent,omitempty"`
	CanLcabDocuments             bool `json:"canLcabDocuments,omitempty"`
}
