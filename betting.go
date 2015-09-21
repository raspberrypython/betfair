// Copyright 2013 Alessandro De Donno

// "Betfair API-NG Golang Library" is dual-licensed: for free software projects
// please refer to GPLv3 (see declaration above), for commercial software
// please contact the author.
// If you are a contributor and need any clarification, please contact the
// author.

// For free software projects:

// This file is part of "Betfair API-NG Golang Library".

// "Betfair API-NG Golang Library" is free software: you can redistribute it
// and/or modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// "Betfair API-NG Golang Library" is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with "Betfair API-NG Golang Library".  If not, see
// <http://www.gnu.org/licenses/>.

package betfair

import (
	"encoding/json"
	"strings"
	"time"
	// "log"
)

type TimeRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type MarketFilter struct {
	TextQuery       string     `json:"textQuery,omitempty"`
	ExchangeIds     []string   `json:"exchangeIds,omitempty"`
	EventIds        []string   `json:"eventIds,omitempty"`
	EventTypeIds    []string   `json:"eventTypeIds,omitempty"`
	MarketCountries []string   `json:"marketCountries,omitempty"`
	MarketIds       []string   `json:"marketIds,omitempty"`
	CompetitionIds  []string   `json:"competitionIds,omitempty"`
	MarketTypeCodes []string   `json:"marketTypeCodes,omitempty"`
	MarketStartTime *TimeRange `json:"marketStartTime,omitempty"`
}

type PriceProjection struct {
	PriceData      []string `json:"priceData,omitempty"`
	Virtualise     bool
	RolloverStakes bool
}

type Params struct {
	MarketFilter     *MarketFilter    `json:"filter,omitempty"`
	MarketIds        []string         `json:"marketIds,omitempty"`
	PriceProjection  *PriceProjection `json:"priceProjection,omitempty"`
	MaxResults       int              `json:"maxResults,omitempty"`
	Locale           string           `json:"locale,omitempty"`
	MarketProjection []string         `json:"marketProjection,omitempty"`
	OrderProjection  string           `json:"orderProjection,omitempty"`
	MatchProjection  string           `json:"matchProjection,omitempty"`
}

type EventType struct {
	Id   string
	Name string
}

type EventTypeResult struct {
	EventType   *EventType
	MarketCount int
}

type Competition struct {
	Id   string
	Name string
}

type CompetitionResult struct {
	Competition       *Competition
	MarketCount       int
	CompetitionRegion string
}

type CountryCodeResult struct {
	CountryCode string
	MarketCount int
}

type Event struct {
	Id          string
	Name        string
	CountryCode string
	Timezone    string
	Venue       string
	OpenDate    time.Time
}

type EventResult struct {
	Event       *Event
	MarketCount int
}

type MarketBook struct {
	MarketId              string
	IsMarketDataDelayed   bool
	Status                string
	BetDelay              int
	BspReconciled         bool
	Complete              bool
	Inplay                bool
	NumberOfWinners       int
	NumberOfRunners       int
	NumberOfActiveRunners int
	LastMatchTime         time.Time
	TotalMatched          float32
	TotalAvailable        float32
	CrossMatching         bool
	RunnersVoidable       bool
	Version               int
	Runners               []Runner
}

type RunnerStatus string

var (
	RunnserStatusActive RunnerStatus = "ACTIVE"
)

type Runner struct {
	SelectionId     int
	Handicap        float32
	Status          string // ACTIVE и тд.
	LastPriceTraded float32
	TotalMatched    float32
	Ex              *ExchangePrices
}

type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize
}

type PriceSize struct {
	Price, Size float32
}

// Information about the Runners (selections) in a market.
type RunnerCatalog struct {
	SelectionId  int
	RunnerName   string
	Handicap     float32
	SortPriority int
	Metadata     map[string]string
}

// Information about a market.
type MarketCatalogue struct {
	MarketId        string
	MarketName      string
	MarketStartTime *time.Time
	Description     *MarketDescription
	Runners         []RunnerCatalog
	EventType       *EventType
	Competition     *Competition
	Event           *Event
}

// Market definition.
type MarketDescription struct {
	PersistenceEnabled bool
	BspMarket          bool
	MarketTime         time.Time
	SuspendTime        time.Time
	SettleTime         time.Time
	BettingType        string
	TurnInPlayEnabled  bool
	MarketType         string
	Regulator          string
	MarketBaseRate     float32
	DiscountAllowed    bool
	Wallet             string
	Rules              string
	RulesHasDate       bool
	Clarifications     string
}

// MarketType Result.
type MarketTypeResult struct {
	MarketType  string
	MarketCount int
}

// Returns a list of Competitions (i.e., World Cup 2013) associated with the
// markets selected by the MarketFilter.
func (s *Session) ListCompetitions(filter *MarketFilter) ([]CompetitionResult, error) {
	var results []CompetitionResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listCompetitions", params, &results)
	return results, err
}

// Returns a list of Countries associated with the markets selected by the
// MarketFilter.
func (s *Session) ListCountries(filter *MarketFilter) ([]CountryCodeResult, error) {
	var results []CountryCodeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listCountries", params, &results)
	return results, err
}

// Returns a list of Events (i.e, Reading vs. Man United) associated with the
// markets selected by the MarketFilter.
func (s *Session) ListEvents(filter *MarketFilter) ([]EventResult, error) {
	var results []EventResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listEvents", params, &results)
	return results, err
}

// Returns a list of Event Types (i.e. Sports) associated with the markets
// selected by the MarketFilter.
func (s *Session) ListEventTypes(filter *MarketFilter) ([]EventTypeResult, error) {
	var results []EventTypeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listEventTypes", params, &results)
	return results, err
}

// Returns a list of dynamic data about markets. Dynamic data includes prices,
// the status of the market, the status of selections, the traded volume, and
// the status of any orders you have placed in the market.
func (s *Session) ListMarketBook(marketIds []string) ([]MarketBook, error) {
	var results []MarketBook
	params := new(Params)
	params.MarketIds = marketIds
	// rf.
	params.PriceProjection = &PriceProjection{PriceData: []string{"EX_BEST_OFFERS"}}
	params.MatchProjection = "NO_ROLLUP"
	err := doBettingRequest(s, "listMarketBook", params, &results)
	return results, err
}

// Returns a list of information about markets that does not change (or
// changes very rarely). You use listMarketCatalogue to retrieve the name
// of the market, the names of selections and other information about markets.
// Market Data Request Limits apply to requests made to listMarketCatalogue.
func (s *Session) ListMarketCatalogue(filter *MarketFilter, maxResults int) ([]MarketCatalogue, error) {
	var results []MarketCatalogue
	params := new(Params)
	params.MarketFilter = filter
	params.MaxResults = maxResults

	// rf.
	params.MarketProjection = []string{"RUNNER_METADATA", "EVENT", "MARKET_START_TIME", "MARKET_DESCRIPTION"}
	err := doBettingRequest(s, "listMarketCatalogue", params, &results)
	return results, err
}

// Returns a list of market types (i.e. MATCH_ODDS, NEXT_GOAL) associated
// with the markets selected by the MarketFilter. The market types are always
// the same, regardless of locale.
func (s *Session) ListMarketTypes(filter *MarketFilter) ([]MarketTypeResult, error) {
	var results []MarketTypeResult
	params := new(Params)
	params.MarketFilter = filter
	err := doBettingRequest(s, "listMarketTypes", params, &results)
	return results, err
}

func doBettingRequest(s *Session, method string, params *Params, v interface{}) error {

	params.Locale = s.config.Locale

	bytes, err := json.Marshal(params)

	if err != nil {
		return err
	}
	body := strings.NewReader(string(bytes))

	data, err := doRequest(s, "betting", method, body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
