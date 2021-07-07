package actions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xdbfoundation/go/protocols/frontier"
	"github.com/xdbfoundation/go/services/frontier/internal/context"
	"github.com/xdbfoundation/go/services/frontier/internal/db2"
	"github.com/xdbfoundation/go/services/frontier/internal/db2/history"
	"github.com/xdbfoundation/go/services/frontier/internal/ledger"
	"github.com/xdbfoundation/go/services/frontier/internal/resourceadapter"
	"github.com/xdbfoundation/go/support/errors"
	"github.com/xdbfoundation/go/support/render/hal"
	"github.com/xdbfoundation/go/support/render/problem"
	"github.com/xdbfoundation/go/xdr"
)

// AssetStatsHandler is the action handler for the /asset endpoint
type AssetStatsHandler struct {
	LedgerState *ledger.State
}

func (handler AssetStatsHandler) validateAssetParams(code, issuer string, pq db2.PageQuery) error {
	if code != "" {
		if !xdr.ValidAssetCode.MatchString(code) {
			return problem.MakeInvalidFieldProblem(
				"asset_code",
				fmt.Errorf("%s is not a valid asset code", code),
			)
		}
	}

	if issuer != "" {
		if _, err := xdr.AddressToAccountId(issuer); err != nil {
			return problem.MakeInvalidFieldProblem(
				"asset_issuer",
				fmt.Errorf("%s is not a valid asset issuer", issuer),
			)
		}
	}

	if pq.Cursor != "" {
		parts := strings.SplitN(pq.Cursor, "_", 3)
		if len(parts) != 3 {
			return problem.MakeInvalidFieldProblem(
				"cursor",
				errors.New("cursor must contain exactly one colon"),
			)
		}

		cursorCode, cursorIssuer, assetType := parts[0], parts[1], parts[2]
		if !xdr.ValidAssetCode.MatchString(cursorCode) {
			return problem.MakeInvalidFieldProblem(
				"cursor",
				fmt.Errorf("%s is not a valid asset code", cursorCode),
			)
		}

		if _, err := xdr.AddressToAccountId(cursorIssuer); err != nil {
			return problem.MakeInvalidFieldProblem(
				"cursor",
				fmt.Errorf("%s is not a valid asset issuer", cursorIssuer),
			)
		}

		if _, ok := xdr.StringToAssetType[assetType]; !ok {
			return problem.MakeInvalidFieldProblem(
				"cursor",
				fmt.Errorf("%s is not a valid asset type", assetType),
			)
		}

	}

	return nil
}

func (handler AssetStatsHandler) findIssuersForAssets(
	historyQ *history.Q,
	assetStats []history.ExpAssetStat,
) (map[string]history.AccountEntry, error) {
	issuerSet := map[string]bool{}
	issuers := []string{}
	for _, assetStat := range assetStats {
		if issuerSet[assetStat.AssetIssuer] {
			continue
		}
		issuerSet[assetStat.AssetIssuer] = true
		issuers = append(issuers, assetStat.AssetIssuer)
	}

	accountsByID := map[string]history.AccountEntry{}
	accounts, err := historyQ.GetAccountsByIDs(issuers)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		accountsByID[account.AccountID] = account
		delete(issuerSet, account.AccountID)
	}

	// Note it's possible that no accounts can be found for certain issuers.
	// That can occur because an account can be removed when there are only empty trustlines
	// pointing to it. We still continue to serve asset stats for such issuers.

	return accountsByID, nil
}

// GetResourcePage returns a page of offers.
func (handler AssetStatsHandler) GetResourcePage(
	w HeaderWriter,
	r *http.Request,
) ([]hal.Pageable, error) {
	ctx := r.Context()

	code, err := getString(r, "asset_code")
	if err != nil {
		return nil, err
	}

	issuer, err := getString(r, "asset_issuer")
	if err != nil {
		return nil, err
	}

	pq, err := GetPageQuery(handler.LedgerState, r, DisableCursorValidation)
	if err != nil {
		return nil, err
	}

	if err = handler.validateAssetParams(code, issuer, pq); err != nil {
		return nil, err
	}

	historyQ, err := context.HistoryQFromRequest(r)
	if err != nil {
		return nil, err
	}

	assetStats, err := historyQ.GetAssetStats(code, issuer, pq)
	if err != nil {
		return nil, err
	}

	issuerAccounts, err := handler.findIssuersForAssets(historyQ, assetStats)
	if err != nil {
		return nil, err
	}

	var response []hal.Pageable
	for _, record := range assetStats {
		var assetStatResponse frontier.AssetStat

		resourceadapter.PopulateAssetStat(
			ctx,
			&assetStatResponse,
			record,
			issuerAccounts[record.AssetIssuer],
		)
		response = append(response, assetStatResponse)
	}

	return response, nil
}
