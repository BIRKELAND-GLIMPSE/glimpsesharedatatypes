package glimpsesharedatatypes

// GpmTopic defines the topic structure for the Glimpse platform.
type GpmTopic struct {
	ID                        int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupTopicID              string  `gorm:"not null;index" json:"group_topic_id default ''"`
	Title                     string  `gorm:"type:varchar(255);not null;unique" json:"title"`
	Slug                      string  `gorm:"type:varchar(255);not null" json:"slug"`
	Description               string  `gorm:"type:text" json:"description"`
	DateCreated               int64   `gorm:"not null" json:"date_created"`
	IsActive                  bool    `gorm:"not null;default:false;index" json:"is_active"`
	CreatorID                 string  `gorm:"type:varchar(255);not null" json:"creator_id"`
	ResolutionInfo            string  `gorm:"type:text" json:"resolution_info"`
	IsResolved                bool    `gorm:"not null;default:false" json:"is_resolved"`
	EndTimeUTC                int64   `gorm:"index" json:"end_time_utc"`
	ResolvedAtUTC             int64   `gorm:"" json:"resolved_at_utc"`
	EndedAtUTC                int64   `gorm:"" json:"ended_at_utc"`
	ResolvedOptionID          *int64  `gorm:"default:null" json:"resolved_option_id"`
	ResolvedOptionDescription string  `gorm:"type:text" json:"resolved_option_description"`
	Results                   string  `gorm:"type:text" json:"results"`
	IsPrimaryTopic            bool    `gorm:"not null;default:false" json:"is_primary_topic"`
	Category                  string  `gorm:"type:varchar(255);index" json:"category"`
	Alpha                     float64 `gorm:"type:numeric;default:0.05" json:"alpha"`
	TopicImageUrl             string  `gorm:"type:text;default:''" json:"topic_image_url"`
	SubsidyAmount             int64   `gorm:"type:numeric;default:0" json:"subsidy_amount"`
	SubsidisedUserId          string  `gorm:"type:varchar(255);default:''" json:"subsidised_user_id"`
	MarketSubsidised          bool    `gorm:"not null;default:false;index" json:"market_subsidised"`
	TopicType                 string  `gorm:"type:varchar(255);default:''" json:"topic_type"`
	Recurrence                string  `gorm:"type:varchar(255);default:'one_time'" json:"recurrence"`
	ImageUrl                  string  `gorm:"type:text;default:''" json:"image_url"`
	EstimatedResolveTime      int64   `gorm:"type:bigint;default:0" json:"estimated_resolve_time"`
	TotalAmountInMarket       float64 `gorm:"type:numeric;default:0" json:"total_amount_in_market"`
	ValueOnCreation           int64   `gorm:"type:numeric;default:0" json:"value_on_creation"`
	GeneratorId               uint    `gorm:"type:numeric;default:0" json:"generator_id"`
	DateStringToResolve       string  `gorm:"column:date_string_to_resolve;type:varchar(255);default:null"`
	DateUTCString             string  `gorm:"type:varchar(32);default:null" json:"date_utc_string"`
	SymbolOne                 string  `gorm:"type:varchar(255);default:''" json:"symbol_one"`
	SymbolTwo                 string  `gorm:"type:varchar(255);default:''" json:"symbol_two"`
	DurationCategory          string  `gorm:"type:varchar(255);default:''" json:"duration_category"`
	TradingViewSymbol         string  `gorm:"type:varchar(255);default:''" json:"trading_view_symbol"`

	IsReconciled              bool    `gorm:"not null;default:false;index" json:"is_reconciled"`
	ReconciledAt              *int64  `gorm:"type:bigint;default:null" json:"reconciled_at"`
	ReconciliationTxHash      *string `gorm:"type:varchar(64);default:null;index" json:"reconciliation_tx_hash"` // Bitcoin on-chain transaction ID (txid): 64 hex characters (32 bytes)
	FinalAMMPnLMilliSats      int64   `gorm:"type:bigint;default:0" json:"final_amm_pnl_milli_sats"`
	ReconciledAmountMilliSats int64   `gorm:"type:bigint;default:0" json:"reconciled_amount_milli_sats"`
	ReconciledBy              *string `gorm:"type:varchar(255);default:null" json:"reconciled_by"`
}

// LmsrTopicOptions defines the options for a GpmTopic in the Glimpse platform.
type LmsrTopicOptions struct {
	ID               int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	TopicID          int64   `gorm:"not null;index" json:"topic_id"`
	Description      string  `gorm:"type:varchar(255);not null;index" json:"description"`
	Odds             float64 `gorm:"type:numeric;default:0" json:"odds"`
	Shares           float64 `gorm:"type:numeric;default:0" json:"shares"`
	SubsidisedShares float64 `gorm:"type:numeric;default:0" json:"subsidised_shares"`
	ImageUrl         string  `gorm:"type:text;default:''" json:"image_url"`
	SeedProbability  float64 `gorm:"type:numeric;default:0" json:"seed_probability"`
}
